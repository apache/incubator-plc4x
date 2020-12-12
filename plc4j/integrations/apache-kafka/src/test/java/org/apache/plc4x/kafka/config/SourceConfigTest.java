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
package org.apache.plc4x.kafka.config;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

import java.io.StringReader;
import java.util.HashMap;
import java.util.Map;
import java.util.Properties;

public class SourceConfigTest {

    @Test
    public void parseConfig() throws Exception {
        Properties properties = new Properties();
        properties.load(new StringReader("name=plc-source-test\n" +
            "connector.class=org.apache.plc4x.kafka.Plc4xSourceConnector\n" +
            "\n" +
            "defaults.topic=some/default\n" +
            "\n" +
            "sources=machineA,machineB,machineC\n" +
            "sources.machineA.connectionString=s7://1.2.3.4/1/1\n" +
            "sources.machineA.jobReferences=s7-dashboard,s7-heartbeat\n" +
            "sources.machineA.jobReferences.s7-heartbeat.topic=heartbeat\n" +
            "\n" +
            "sources.machineB.connectionString=s7://10.20.30.40/1/1\n" +
            "sources.machineB.topic=heartbeat\n" +
            "sources.machineB.jobReferences=s7-heartbeat\n" +
            "\n" +
            "sources.machineC.connectionString=ads://1.2.3.4.5.6\n" +
            "sources.machineC.topic=heartbeat\n" +
            "sources.machineC.jobReferences=ads-heartbeat\n" +
            "\n" +
            "jobs=s7-dashboard,s7-heartbeat,ads-heartbeat\n" +
            "jobs.s7-dashboard.interval=500\n" +
            "jobs.s7-dashboard.fields=inputPreasure,outputPreasure,temperature\n" +
            "jobs.s7-dashboard.fields.inputPreasure=%DB.DB1.4:INT\n" +
            "jobs.s7-dashboard.fields.outputPreasure=%Q1:BYT\n" +
            "jobs.s7-dashboard.fields.temperature=%I3:INT\n" +
            "\n" +
            "jobs.s7-heartbeat.interval=1000\n" +
            "jobs.s7-heartbeat.fields=active\n" +
            "jobs.s7-heartbeat.fields.active=%I0.2:BOOL\n" +
            "\n" +
            "jobs.ads-heartbeat.interval=1000\n" +
            "jobs.ads-heartbeat.fields=active\n" +
            "jobs.ads-heartbeat.fields.active=Main.running\n"));
        SourceConfig sourceConfig = SourceConfig.fromPropertyMap(toStringMap(properties));

        assertNotNull(sourceConfig);
        assertEquals(3, sourceConfig.getSources().size(), "Expected 3 sources");
        assertEquals(3, sourceConfig.getJobs().size(), "Expected 3 jobs");
    }

    private static Map<String, String> toStringMap(Properties properties) {
        Map<String, String> map = new HashMap<>();
        for (String stringPropertyName : properties.stringPropertyNames()) {
            map.put(stringPropertyName, properties.getProperty(stringPropertyName));
        }
        return map;
    }

}
