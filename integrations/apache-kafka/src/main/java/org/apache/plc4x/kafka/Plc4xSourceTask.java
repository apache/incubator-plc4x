/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.kafka;

import org.apache.kafka.common.config.AbstractConfig;
import org.apache.kafka.connect.data.Schema;
import org.apache.kafka.connect.data.SchemaBuilder;
import org.apache.kafka.connect.data.Struct;
import org.apache.kafka.connect.errors.ConnectException;
import org.apache.kafka.connect.source.SourceRecord;
import org.apache.kafka.connect.source.SourceTask;
import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.connection.PlcConnection;
import org.apache.plc4x.java.api.connection.PlcReader;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.exceptions.UncheckedPlcInvalidFieldException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.kafka.util.VersionUtil;

import java.util.*;
import java.util.concurrent.*;

/**
 * Source Connector Task polling the data source at a given rate.
 * A timer thread is scheduled which sets the fetch flag to true every rate milliseconds.
 * When poll() is invoked, the calling thread waits until the fetch flag is set for WAIT_LIMIT_MILLIS.
 * If the flag does not become true, the method returns null, otherwise a fetch is performed.
 */
public class Plc4xSourceTask extends SourceTask {
    private final static long WAIT_LIMIT_MILLIS = 100;
    private final static long TIMEOUT_LIMIT_MILLIS = 5000;

    private final static String URL_FIELD = "url";
    private final static String QUERY_FIELD = "query";

    private final static Schema KEY_SCHEMA =
        new SchemaBuilder(Schema.Type.STRUCT)
            .field(URL_FIELD, Schema.STRING_SCHEMA)
            .field(QUERY_FIELD, Schema.STRING_SCHEMA)
            .build();

    private String topic;
    private String url;
    private List<String> queries;

    private PlcConnection plcConnection;
    private PlcReader plcReader;
    private PlcReadRequest plcRequest;



    // TODO: should we use shared (static) thread pool for this?
    private ScheduledExecutorService scheduler;
    private ScheduledFuture<?> timer;
    private boolean fetch = true;

    @Override
    public String version() {
        return VersionUtil.getVersion();
    }

    @Override
    public void start(Map<String, String> props) {
        AbstractConfig config = new AbstractConfig(Plc4xSourceConnector.CONFIG_DEF, props);
        topic = config.getString(Plc4xSourceConnector.TOPIC_CONFIG);
        url = config.getString(Plc4xSourceConnector.URL_CONFIG);
        queries = config.getList(Plc4xSourceConnector.QUERIES_CONFIG);

        openConnection();

        plcReader = plcConnection.getReader()
            .orElseThrow(() -> new ConnectException("PlcReader not available for this type of connection"));


        PlcReadRequest.Builder builder = plcReader.readRequestBuilder();
        for (String query : queries) {
            builder.addItem(query, query);
        }
        try {
            plcRequest = builder.build();
        } catch (PlcInvalidFieldException e) {
            // TODO how should this be handled?
            throw new UncheckedPlcInvalidFieldException(e);
        }

        int rate = Integer.valueOf(props.get(Plc4xSourceConnector.RATE_CONFIG));
        scheduler = Executors.newScheduledThreadPool(1);
        timer = scheduler.scheduleAtFixedRate(Plc4xSourceTask.this::scheduleFetch, rate, rate, TimeUnit.MILLISECONDS);
    }

    @Override
    public void stop() {
        timer.cancel(true);
        scheduler.shutdown();
        closeConnection();
    }

    @Override
    public List<SourceRecord> poll() throws InterruptedException {
        return awaitFetch(WAIT_LIMIT_MILLIS) ? doFetch() : null;
    }

    private void openConnection() {
        try {
            plcConnection = new PlcDriverManager().getConnection(url);
            plcConnection.connect();
        } catch (PlcConnectionException e) {
            throw new ConnectException("Could not establish a PLC connection", e);
        }
    }

    private void closeConnection() {
        if (plcConnection != null) {
            try {
                plcConnection.close();
            } catch (Exception e) {
                throw new RuntimeException("Caught exception while closing connection to PLC", e);
            }
        }
    }

    /**
     * Schedule next fetch operation.
     */
    private synchronized void scheduleFetch() {
        fetch = true;
        notify();
    }

    /**
     * Wait for next scheduled fetch operation.
     * @param milliseconds maximum time to wait
     * @throws InterruptedException if the thread is interrupted
     * @return true if a fetch should be performed, false otherwise
     */
    private synchronized boolean awaitFetch(long milliseconds) throws InterruptedException {
        if (!fetch) {
            wait(milliseconds);
        }
        try {
            return fetch;
        } finally {
            fetch = false;
        }
    }

    private List<SourceRecord> doFetch() throws InterruptedException {
        final CompletableFuture<PlcReadResponse<?>> response = plcReader.read(plcRequest);
        try {
            final PlcReadResponse<?> received = response.get(TIMEOUT_LIMIT_MILLIS, TimeUnit.MILLISECONDS);
            return extractValues(received);
        } catch (ExecutionException e) {
            throw new ConnectException("Could not fetch data from source", e);
        } catch (TimeoutException e) {
            throw new ConnectException("Timed out waiting for data from source", e);
        }
    }

    private List<SourceRecord> extractValues(PlcReadResponse<?> response) {
        final List<SourceRecord> result = new LinkedList<>();
        for (String query : queries) {
            final PlcResponseCode rc = response.getResponseCode(query);
            if (!rc.equals(PlcResponseCode.OK))  {
                continue;
            }

            Struct key = new Struct(KEY_SCHEMA)
                .put(URL_FIELD, url)
                .put(QUERY_FIELD, query);

            Object value = response.getObject(query);
            Schema valueSchema = getSchema(value);
            Long timestamp = System.currentTimeMillis();
            Map<String, String> sourcePartition = new HashMap<>();
            sourcePartition.put("url", url);
            sourcePartition.put("query", query);
            Map<String, Long> sourceOffset = Collections.singletonMap("offset", timestamp);

            SourceRecord record =
                new SourceRecord(
                    sourcePartition,
                    sourceOffset,
                    topic,
                    KEY_SCHEMA,
                    key,
                    valueSchema,
                    value
                );

            result.add(record);
        }

        return result;
    }

    private Schema getSchema(Object value) {
        Objects.requireNonNull(value);

        if (value instanceof Byte)
            return Schema.INT8_SCHEMA;

        if (value instanceof Short)
            return Schema.INT16_SCHEMA;

        if (value instanceof Integer)
            return Schema.INT32_SCHEMA;

        if (value instanceof Long)
            return Schema.INT64_SCHEMA;

        if (value instanceof Float)
            return Schema.FLOAT32_SCHEMA;

        if (value instanceof Double)
            return Schema.FLOAT64_SCHEMA;

        if (value instanceof Boolean)
            return Schema.BOOLEAN_SCHEMA;

        if (value instanceof String)
            return Schema.STRING_SCHEMA;

        if (value instanceof byte[])
            return Schema.BYTES_SCHEMA;

        // TODO: add support for collective and complex types
        throw new ConnectException(String.format("Unsupported data type %s", value.getClass().getName()));
    }

}