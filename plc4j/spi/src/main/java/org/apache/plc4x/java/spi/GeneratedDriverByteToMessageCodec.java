/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi;

import io.netty.buffer.ByteBuf;
import io.netty.channel.ChannelHandlerContext;
import io.netty.handler.codec.ByteToMessageCodec;
import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.spi.generation.Message;
import org.apache.plc4x.java.spi.generation.MessageIO;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;

public abstract class GeneratedDriverByteToMessageCodec<T extends Message> extends ByteToMessageCodec<T> {

    private static final Logger logger = LoggerFactory.getLogger(GeneratedDriverByteToMessageCodec.class);

    private final boolean bigEndian;
    private final Object[] parserArgs;
    private final MessageIO<T, T> io;

    public GeneratedDriverByteToMessageCodec(MessageIO<T, T> io, Class<T> clazz, boolean bigEndian, Object[] parserArgs) {
        super(clazz);
        this.io = io;
        this.bigEndian = bigEndian;
        this.parserArgs = parserArgs;
    }

    @Override
    protected void encode(ChannelHandlerContext ctx, T packet, ByteBuf byteBuf) throws Exception {
        WriteBuffer buffer = new WriteBuffer(packet.getLengthInBytes(), !bigEndian);
        io.serialize(buffer, packet);
        byteBuf.writeBytes(buffer.getData());
        logger.debug("Sending bytes to PLC for message {} as data {}", packet, Hex.encodeHexString(buffer.getData()));
    }

    @Override
    protected void decode(ChannelHandlerContext ctx, ByteBuf byteBuf, List<Object> out) throws Exception {
        logger.trace("Receiving bytes, trying to decode Message...");
        // As long as there is data available, continue checking the content.
        while(byteBuf.readableBytes() > 0) {
            // Check if enough data is present to process the entire package.
            int packetSize = getPacketSize(byteBuf);
            if(packetSize == -1 || packetSize > byteBuf.readableBytes()) {
                return;
            }

            // Read the packet data into a new ReadBuffer
            byte[] bytes = new byte[packetSize];
            byteBuf.readBytes(bytes);
            System.out.println();
            ReadBuffer readBuffer = new ReadBuffer(bytes, !bigEndian);

            try {
                // Parse the packet.
                T packet = io.parse(readBuffer, parserArgs);

                // Pass the packet to the pipeline.
                out.add(packet);

                // It seems that one batch of 16 messages is the maximum, so we have to give up
                // and process the rest next time.
                if(out.size() >= 16) {
                    return;
                }
            } catch (Exception e) {
                logger.warn("Error decoding package with content [" + Hex.encodeHexString(bytes) + "]: "
                    + e.getMessage(), e);
                // Just remove any trailing junk ... if there is any.
                removeRestOfCorruptPackage(byteBuf);
            }
        }
    }

    abstract protected int getPacketSize(ByteBuf byteBuf);

    abstract protected void removeRestOfCorruptPackage(ByteBuf byteBuf);

}
