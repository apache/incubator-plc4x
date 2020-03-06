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
package org.apache.plc4x.camel;

import org.apache.camel.*;
import org.apache.camel.spi.ExceptionHandler;
import org.apache.camel.support.AsyncProcessorConverterHelper;
import org.apache.camel.support.LoggingExceptionHandler;
import org.apache.camel.support.service.ServiceSupport;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.ArrayList;
import java.util.Collection;
import java.util.List;
import java.util.concurrent.*;

public class Plc4XConsumer extends ServiceSupport implements Consumer {
    private static final Logger LOGGER = LoggerFactory.getLogger(Plc4XConsumer.class);

    private Plc4XEndpoint endpoint;
    private AsyncProcessor processor;
    private ExceptionHandler exceptionHandler;
    private PlcConnection plcConnection;
    private List<String> fieldQuery;
    private Class<?> dataType;
    private PlcSubscriptionResponse subscriptionResponse;


    private ScheduledExecutorService executorService = Executors.newSingleThreadScheduledExecutor();
    private ScheduledFuture<?> future;

    public Plc4XConsumer(Plc4XEndpoint endpoint, Processor processor) throws PlcException {
        this.endpoint = endpoint;
        this.dataType = endpoint.getDataType();
        this.processor = AsyncProcessorConverterHelper.convert(processor);
        this.exceptionHandler = new LoggingExceptionHandler(endpoint.getCamelContext(), getClass());
        this.plcConnection = endpoint.getConnection();
        this.fieldQuery = endpoint.getAddress();
    }

    @Override
    public String toString() {
        return "Plc4XConsumer[" + endpoint + "]";
    }

    @Override
    public Endpoint getEndpoint() {
        return endpoint;
    }

    public ExceptionHandler getExceptionHandler() {
        return exceptionHandler;
    }

    public void setExceptionHandler(ExceptionHandler exceptionHandler) {
        this.exceptionHandler = exceptionHandler;
    }

    @Override
    protected void doStart() throws InterruptedException, ExecutionException {
        int nb=0;
        PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
        if (fieldQuery.size()>1){
            int i=0;
            for(String query : fieldQuery){
                builder.addItem(String.valueOf(i++),query);
            }
        }
        else{
            builder.addItem("default",fieldQuery.get(0));
        }
        PlcReadRequest request = builder.build();
        future = executorService.schedule(() -> {
            request.execute().thenAccept(response -> {
                    try {
                        Exchange exchange = endpoint.createExchange();
                        if (fieldQuery.size()>1){
                            int i=0;
                            List<Object> values = new ArrayList<>();
                            for(String query : fieldQuery){
                                values.add(response.getObject(String.valueOf(i++)));
                            }
                            exchange.getIn().setBody(values);
                        }
                        else {
                            exchange.getIn().setBody(unwrapIfSingle(response.getAllObjects("default")));
                        }
                        processor.process(exchange);
                    } catch (Exception e) {
                        exceptionHandler.handleException(e);
                    }
                });
        }, 500, TimeUnit.MILLISECONDS);
    }

    @Override
    protected void doStop() throws InterruptedException, ExecutionException, TimeoutException {
        // First stop the polling process
        if (future != null) {
            future.cancel(true);
        }
    }

    private Object unwrapIfSingle(Collection collection) {
        if (collection.isEmpty()) {
            return null;
        }
        if (collection.size() == 1) {
            return collection.iterator().next();
        }
        return collection;
    }

    @Override
    public Processor getProcessor() {
        return this.processor;
    }
}