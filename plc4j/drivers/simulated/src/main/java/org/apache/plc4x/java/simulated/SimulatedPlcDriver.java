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
package org.apache.plc4x.java.simulated;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.simulated.connection.SimulatedPlcConnection;
import org.apache.plc4x.java.simulated.connection.TestDevice;
import org.apache.plc4x.java.spi.PlcDriver;
import org.osgi.service.component.annotations.Component;

/**
 * Test driver holding its state in the client process.
 * The URL schema is {@code test:<device_name>}.
 * Devices are created each time a connection is established and should not be reused.
 * Every device contains a random value generator accessible by address {@code random}.
 * Any value can be stored into test devices, however the state will be gone when connection is closed.
 */
@Component(service = PlcDriver.class, immediate = true)
public class SimulatedPlcDriver implements PlcDriver {

    @Override
    public String getProtocolCode() {
        return "test";
    }

    @Override
    public String getProtocolName() {
        return "PLC4X Test Protocol";
    }

    @Override
    public PlcConnection connect(String url) throws PlcConnectionException {
        // TODO: perform further checks
        String deviceName = url.substring(5);
        if (deviceName.isEmpty()) {
            throw new PlcConnectionException("Invalid URL: no device name given.");
        }
        TestDevice device = new TestDevice(deviceName);
        return new SimulatedPlcConnection(device);
    }

    @Override
    public PlcConnection connect(String url, PlcAuthentication authentication) throws PlcConnectionException {
        throw new PlcConnectionException("Test driver does not support authentication.");
    }

}
