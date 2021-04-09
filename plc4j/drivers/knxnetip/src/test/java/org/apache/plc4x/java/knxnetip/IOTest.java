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

package org.apache.plc4x.java.knxnetip;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.dataformat.xml.XmlMapper;
import org.apache.commons.codec.binary.Hex;
import org.apache.plc4x.java.knxnetip.readwrite.KnxNetIpMessage;
import org.apache.plc4x.java.knxnetip.readwrite.io.KnxNetIpMessageIO;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.junit.jupiter.api.Test;

import java.util.Arrays;

public class IOTest {

    @Test
    public void testXml() throws Exception {
        byte[] rData = Hex.decodeHex("061004200019047f02002900bce0110509d405008044a68000");
        ObjectMapper mapper = new XmlMapper().enableDefaultTyping();
        ReadBuffer rBuf = new ReadBuffer(rData);
        KnxNetIpMessage packet = new KnxNetIpMessageIO().parse(rBuf);
        String xml = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(packet);
        System.out.println(xml);
        KnxNetIpMessage pack2 = mapper.readValue(xml, KnxNetIpMessage.class);
        System.out.println(pack2);
    }

    @Test
    public void testJson() throws Exception {
        byte[] rData = Hex.decodeHex("0610020500180801c0a82a46c4090801c0a82a46c40a0203");
        ObjectMapper mapper = new ObjectMapper().enableDefaultTyping();
        ReadBuffer rBuf = new ReadBuffer(rData);
        KnxNetIpMessage packet = new KnxNetIpMessageIO().parse(rBuf);
        String json = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(packet);
        System.out.println(json);
        KnxNetIpMessage pack2 = mapper.readValue(json, KnxNetIpMessage.class);
        System.out.println(pack2);
    }

    @Test
    public void testParser() throws Exception {
        byte[] rData = Hex.decodeHex("0610020500180801c0a82a46c4090801c0a82a46c40a0203");
        long start = System.currentTimeMillis();
        int numRunsParse = 20000;

        KnxNetIpMessageIO knxNetIPMessageIO = new KnxNetIpMessageIO();

        // Benchmark the parsing code
        KnxNetIpMessage packet = null;
        for(int i = 0; i < numRunsParse; i++) {
            ReadBuffer rBuf = new ReadBuffer(rData);
            packet = KnxNetIpMessageIO.staticParse(rBuf);
        }
        long endParsing = System.currentTimeMillis();

        System.out.println("Parsed " + numRunsParse + " packets in " + (endParsing - start) + "ms");
        System.out.println("That's " + ((float) (endParsing - start) / numRunsParse) + "ms per packet");

        // Benchmark the serializing code
        int numRunsSerialize = 20000;
        byte[] oData = null;
        for(int i = 0; i < numRunsSerialize; i++) {
            WriteBuffer wBuf = new WriteBuffer(packet.getLengthInBytes());
            KnxNetIpMessageIO.staticSerialize(wBuf, packet);
            oData = wBuf.getData();
        }
        long endSerializing = System.currentTimeMillis();

        System.out.println("Serialized " + numRunsSerialize + " packets in " + (endSerializing - endParsing) + "ms");
        System.out.println("That's " + ((float) (endSerializing - endParsing) / numRunsSerialize) + "ms per packet");
        if(!Arrays.equals(rData, oData)) {
            for(int i = 0; i < rData.length; i++) {
                if(rData[i] != oData[i]) {
                    System.out.println("Difference in byte " + i);
                }
            }
            System.out.println("Not equals");
        } else {
            System.out.println("Bytes equal");
        }
    }

}
