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
package org.apache.plc4x.java.amsads.connection;

import io.netty.channel.Channel;
import io.netty.channel.ChannelHandler;
import io.netty.channel.ChannelInitializer;
import io.netty.channel.ChannelPipeline;
import org.apache.plc4x.java.amsads.protocol.Ads2PayloadProtocol;
import org.apache.plc4x.java.amsads.protocol.Payload2SerialProtocol;
import org.apache.plc4x.java.amsads.protocol.Plc4x2AdsProtocol;
import org.apache.plc4x.java.amsads.readwrite.AmsNetId;
import org.apache.plc4x.java.base.protocol.SingleItemToSingleRequestProtocol;
import org.apache.plc4x.java.serial.connection.connection.SerialChannelFactory;

import java.util.concurrent.CompletableFuture;

public class AdsSerialPlcConnection extends AdsAbstractPlcConnection {

    private AdsSerialPlcConnection(String serialPort, AmsNetId targetAmsNetId, int targetAmsPort) {
        this(serialPort, targetAmsNetId, targetAmsPort, generateAmsNetId(), generateAMSPort());
    }

    private AdsSerialPlcConnection(String serialPort, AmsNetId targetAmsNetId, int targetAmsPort, AmsNetId sourceAmsNetId, int sourceAmsPort) {
        super(new SerialChannelFactory(serialPort), targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort);
    }

    public static AdsSerialPlcConnection of(String serialPort, AmsNetId targetAmsNetId, int targetAmsPort) {
        return new AdsSerialPlcConnection(serialPort, targetAmsNetId, targetAmsPort);
    }

    public static AdsSerialPlcConnection of(String serialPort, AmsNetId targetAmsNetId, int targetAmsPort, AmsNetId sourceAmsNetId, int sourceAmsPort) {
        return new AdsSerialPlcConnection(serialPort, targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort);
    }

    @Override
    protected ChannelHandler getChannelHandler(CompletableFuture<Void> sessionSetupCompleteFuture) {
        return new ChannelInitializer() {
            @Override
            protected void initChannel(Channel channel) {
                // Build the protocol stack for communicating with the ads protocol.
                ChannelPipeline pipeline = channel.pipeline();
                pipeline.addLast(new Payload2SerialProtocol());
                pipeline.addLast(new Ads2PayloadProtocol());
                pipeline.addLast(new Plc4x2AdsProtocol(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, fieldMapping));
                pipeline.addLast(new SingleItemToSingleRequestProtocol(AdsSerialPlcConnection.this, AdsSerialPlcConnection.this, null, timer));
            }
        };
    }

}
