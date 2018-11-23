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
package org.apache.plc4x.java.s7.connection;

import io.netty.channel.embedded.EmbeddedChannel;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.s7.types.S7ControllerType;
import org.junit.After;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.Timeout;

import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;

import static org.hamcrest.core.IsEqual.equalTo;
import static org.hamcrest.core.IsNull.notNullValue;
import static org.junit.Assert.assertThat;

public class S7PlcConnectionIT {

    @Rule
    public Timeout globalTimeout = Timeout.seconds(4); // 4 seconds max per method tested

    private S7PlcTestConnection SUT;

    @Before
    public void setUp() {
        SUT = new S7PlcTestConnection(1, 2,
            "pdu-size=1&max-amq-caller=2&max-amq-callee=3&unknown=parameter&unknown-flag", S7ControllerType.S7_1500);
    }

    @After
    public void tearDown() throws PlcConnectionException{
        if(SUT.isConnected()) {
            SUT.close();
        }
        SUT = null;
    }

    @Test
    public void connectAndClose() throws Exception {
        SUT.connect();
        SUT.close();
    }

    @Test
    public void read() throws Exception {
        SUT.connect();
        EmbeddedChannel channel = (EmbeddedChannel) SUT.getChannel();
        assertThat("No outbound messages should exist.", channel.outboundMessages().size(), equalTo(0));

        PlcReadRequest request = SUT.readRequestBuilder().addItem("test", "%Q0.4:BOOL").build();
        CompletableFuture<PlcReadResponse> responseFuture = SUT.read(request);
        // Check that one message has been sent.
        assertThat("Exactly one outbound message should exist after sending.",
            channel.outboundMessages().size(), equalTo(1));
        SUT.verifyPcapFile("org/apache/plc4x/java/s7/connection/s7-read-var-request.pcapng");

        // Manually feed a packet response into the channel.
        SUT.sendPcapFile("org/apache/plc4x/java/s7/connection/s7-read-var-response.pcapng");

        // Now get the response as it was processed by the connection.
        PlcReadResponse response = responseFuture.get(200, TimeUnit.MILLISECONDS);

        assertThat(response, notNullValue());

        SUT.close();
    }

    @Test
    public void write() throws Exception {
        SUT.connect();
        EmbeddedChannel channel = (EmbeddedChannel) SUT.getChannel();
        assertThat("No outbound messages should exist.", channel.outboundMessages().size(), equalTo(0));

        PlcWriteRequest request = SUT.writeRequestBuilder().addItem("test", "%Q0.4:BOOL", true).build();
        CompletableFuture<PlcWriteResponse> responseFuture = SUT.write(request);
        // Check that one message has been sent.
        assertThat("Exactly one outbound message should exist after sending.",
            channel.outboundMessages().size(), equalTo(1));
        SUT.verifyPcapFile("org/apache/plc4x/java/s7/connection/s7-write-var-request.pcapng");

        // Manually feed a packet response into the channel.
        SUT.sendPcapFile("org/apache/plc4x/java/s7/connection/s7-write-var-response.pcapng");

        // Now get the response as it was processed by the connection.
        PlcWriteResponse response = responseFuture.get(200, TimeUnit.MILLISECONDS);

        assertThat(response, notNullValue());

        SUT.close();
    }

}
