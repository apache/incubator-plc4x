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

public class Constants {

    public static final String DEFAULT_TOPIC_CONFIG = "default.topic";
    public static final String DEFAULT_TOPIC_DOC = "Default topic to be used, if not otherwise configured.";

    public static final String SOURCES_CONFIG = "sources";
    public static final String SOURCES_DOC = "List of source names that will be configured.";
    public static final String SOURCES_DEFAULT = "";

    public static final String JOBS_CONFIG = "jobs";
    public static final String JOBS_DOC = "List of job names that will be configured.";
    public static final String JOBS_DEFAULT = "";

    public static final String TOPIC_CONFIG = "topic";
    public static final String INTERVAL_CONFIG = "interval";
    public static final String FIELDS_CONFIG = "fields";

    public static final String CONNECTION_STRING_CONFIG = "connectionString";
    public static final String CONNECTION_STRING_DOC = "PLC4X Connection String";

    public static final String JOB_REFERENCES_CONFIG = "jobReferences";

    public static final String KAFKA_POLL_RETURN_CONFIG = "pollReturnInterval";
    public static final String KAFKA_POLL_RETURN_DOC = "Default poll return interval to be used, if not otherwise configured.";
    public static final Integer KAFKA_POLL_RETURN_DEFAULT = 5000;

    public static final String BUFFER_SIZE_CONFIG = "bufferSize";
    public static final String BUFFER_SIZE_DOC = "Default buffer size to be used, if not otherwise configured.";
    public static final Integer BUFFER_SIZE_DEFAULT = 1000;

    /*
     * Config of the task.
     */
    public static final String CONNECTION_NAME_CONFIG = "connection-name";
    public static final String CONNECTION_NAME_STRING_DOC = "Connection Name";

    // Syntax for the queries: {job-name}:{topic}:{rate}:{field-alias}#{field-address}:{field-alias}#{field-address}...,{topic}:{rate}:....
    public static final String QUERIES_CONFIG = "queries";
    public static final String QUERIES_DOC = "Field queries to be sent to the PLC";

    /*
     * Configuration of the output.
     */
    public static final String SOURCE_NAME_FIELD = "sourceName";
    public static final String JOB_NAME_FIELD = "jobName";

}
