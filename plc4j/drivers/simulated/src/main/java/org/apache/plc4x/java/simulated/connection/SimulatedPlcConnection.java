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
package org.apache.plc4x.java.simulated.connection;

import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.base.connection.AbstractPlcConnection;
import org.apache.plc4x.java.base.messages.*;
import org.apache.plc4x.java.base.messages.items.BaseDefaultFieldItem;

import java.util.HashMap;
import java.util.Map;
import java.util.Optional;
import java.util.concurrent.CompletableFuture;

/**
 * Connection to a test device.
 * This class is not thread-safe.
 */
public class SimulatedPlcConnection extends AbstractPlcConnection implements PlcReader, PlcWriter {
    private final TestDevice device;
    private boolean connected = false;

    public SimulatedPlcConnection(TestDevice device) {
        this.device = device;
    }

    @Override
    public void connect() {
        connected = true;
    }

    @Override
    public boolean isConnected() {
        return connected;
    }

    @Override
    public void close() {
        connected = false;
    }

    @Override
    public boolean canRead() {
        return true;
    }

    @Override
    public boolean canWrite() {
        return true;
    }

    @Override
    public PlcReadRequest.Builder readRequestBuilder() {
        return new DefaultPlcReadRequest.Builder(this, new TestFieldHandler());
    }

    @Override
    public PlcWriteRequest.Builder writeRequestBuilder() {
        return new DefaultPlcWriteRequest.Builder(this, new TestFieldHandler());
    }

    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        if(!(readRequest instanceof InternalPlcReadRequest)) {
            throw new IllegalArgumentException("Read request doesn't implement InternalPlcReadRequest");
        }
        InternalPlcReadRequest request = (InternalPlcReadRequest) readRequest;
        Map<String, Pair<PlcResponseCode, BaseDefaultFieldItem>> fields = new HashMap<>();
        for (String fieldName : request.getFieldNames()) {
            TestField field = (TestField) request.getField(fieldName);
            Optional<BaseDefaultFieldItem> fieldItemOptional = device.get(field);
            ImmutablePair<PlcResponseCode, BaseDefaultFieldItem> fieldPair;
            boolean present = fieldItemOptional.isPresent();
            fieldPair = present
                ? new ImmutablePair<>(PlcResponseCode.OK, fieldItemOptional.get())
                : new ImmutablePair<>(PlcResponseCode.NOT_FOUND, null);
            fields.put(fieldName, fieldPair);
        }
        PlcReadResponse response = new DefaultPlcReadResponse(request, fields);
        return CompletableFuture.completedFuture(response);
    }

    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        if(!(writeRequest instanceof InternalPlcWriteRequest)) {
            throw new IllegalArgumentException("Read request doesn't implement InternalPlcWriteRequest");
        }
        InternalPlcWriteRequest request = (InternalPlcWriteRequest) writeRequest;
        Map<String, PlcResponseCode> fields = new HashMap<>();
        for (String fieldName : request.getFieldNames()) {
            TestField field = (TestField) request.getField(fieldName);
            BaseDefaultFieldItem fieldItem = request.getFieldItem(fieldName);
            device.set(field, fieldItem);
            fields.put(fieldName, PlcResponseCode.OK);
        }
        PlcWriteResponse response = new DefaultPlcWriteResponse(request, fields);
        return CompletableFuture.completedFuture(response);
    }

    @Override
    public String toString() {
        return String.format("test:%s", device);
    }

}
