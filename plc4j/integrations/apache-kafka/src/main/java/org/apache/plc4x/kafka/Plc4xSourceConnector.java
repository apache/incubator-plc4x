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

import org.apache.commons.lang3.math.NumberUtils;
import org.apache.kafka.common.config.Config;
import org.apache.kafka.common.config.ConfigDef;
import org.apache.kafka.common.config.ConfigValue;
import org.apache.kafka.connect.connector.Task;
import org.apache.kafka.connect.source.SourceConnector;
import org.apache.plc4x.kafka.config.Job;
import org.apache.plc4x.kafka.config.JobReference;
import org.apache.plc4x.kafka.config.Source;
import org.apache.plc4x.kafka.config.SourceConfig;
import org.apache.plc4x.kafka.util.VersionUtil;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.*;

public class Plc4xSourceConnector extends SourceConnector {

    private static final Logger log = LoggerFactory.getLogger(Plc4xSourceConnector.class);

    public static final String DEFAULT_TOPIC_CONFIG = "default-topic";
    private static final String DEFAULT_TOPIC_DOC = "Default topic to be used, if not otherwise configured.";

    public static final String SOURCES_CONFIG = "sources";
    private static final String SOURCES_DOC = "List of source names that will be configured.";

    public static final String JOBS_CONFIG = "jobs";
    private static final String JOBS_DOC = "List of job names that will be configured.";

    private static final String CONNECTION_STRING_CONFIG = "connectionString";
    private static final String JOB_REFERENCES_CONFIG = "jobReferences";
    private static final String TOPIC_CONFIG = "topic";
    private static final String INTERVAL_CONFIG = "interval";
    private static final String FIELDS_CONFIG = "fields";

    private SourceConfig sourceConfig;

    @Override
    public void start(Map<String, String> props) {
        sourceConfig = SourceConfig.fromPropertyMap(props);
    }

    @Override
    public void stop() {
        sourceConfig = null;
    }

    @Override
    public Class<? extends Task> taskClass() {
        return Plc4xSourceTask.class;
    }

    @Override
    public List<Map<String, String>> taskConfigs(int maxTasks) {
        // Initially we planned to have the simple assumption that one task maps to one PLC connection.
        // But we could easily say that one scraper instance maps to a task and one scraper task can
        // process multiple PLC connections. But I guess this would be an optimization as we have to
        // balance the load manually.
        if(sourceConfig.getJobs().size() > maxTasks) {
            // not enough tasks
            log.warn("NOT ENOUGH TASKS!");
            return Collections.emptyList();
        }

        // For each configured source we'll start a dedicated scraper instance collecting
        // all the scraper jobs enabled for this source.
        List<Map<String, String>> configs = new LinkedList<>();
        for (Source source : sourceConfig.getSources()) {
            // Build a list of job configurations only containing the ones referenced from
            // the current source.
            StringBuilder query = new StringBuilder();
            for (JobReference jobReference : source.getJobReferences()) {
                Job job = sourceConfig.getJob(jobReference.getName());
                if(job == null) {
                    log.warn(String.format("Couldn't find referenced job '%s'", jobReference.getName()));
                } else {
                    query.append(",").append(jobReference.getName()).append("|").append(jobReference.getTopic());
                    query.append("|").append(job.getInterval());
                    for (Map.Entry<String, String> field : job.getFields().entrySet()) {
                        String fieldName = field.getKey();
                        String fieldAddress = field.getValue();
                        query.append("|").append(fieldName).append("#").append(fieldAddress);
                    }
                }
            }

            // Create a new task configuration.
            Map<String, String> taskConfig = new HashMap<>();
            taskConfig.put(Plc4xSourceTask.CONNECTION_NAME_CONFIG, source.getName());
            taskConfig.put(Plc4xSourceTask.PLC4X_CONNECTION_STRING_CONFIG, source.getConnectionString());
            taskConfig.put(Plc4xSourceTask.QUERIES_CONFIG, query.toString().substring(1));
            configs.add(taskConfig);
        }
        return configs;
    }

    @Override
    @SuppressWarnings("unchecked")
    public Config validate(Map<String, String> connectorConfigs) {
        ////////////////////////////////////////////////////
        // Get the static part of the config
        Config config = super.validate(connectorConfigs);

        ////////////////////////////////////////////////////
        // Add the dynamic parts of the config

        // Find the important config elements
        String defaultTopic = null;
        ConfigValue sources = null;
        ConfigValue jobs = null;
        for (ConfigValue configValue : config.configValues()) {
            switch (configValue.name()) {
                case DEFAULT_TOPIC_CONFIG:
                    defaultTopic = (String) configValue.value();
                    break;
                case JOBS_CONFIG:
                    jobs = configValue;
                    break;
                case SOURCES_CONFIG:
                    sources = configValue;
                    break;
                default:
                    // Just ignore the others.
            }
        }

        // Configure the jobs first (As we reference them from the sources)
        List<Object> foundJobs = new LinkedList<>();
        if(jobs != null) {
            final List<String> jobNames = (List<String>) jobs.value();
            for (String jobName : jobNames) {
                String jobIntervalConfig = JOBS_CONFIG + "." + jobName + "." + INTERVAL_CONFIG;
                ConfigValue jobIntervalConfigValue = new ConfigValue(jobIntervalConfig);
                config.configValues().add(jobIntervalConfigValue);
                String jobIntervalString = connectorConfigs.get(jobIntervalConfig);
                if (jobIntervalString == null) {
                    jobIntervalConfigValue.value(null);
                    jobIntervalConfigValue.addErrorMessage(jobIntervalConfig + " is mandatory");
                } else if (NumberUtils.isParsable(jobIntervalString)) {
                    int jobInterval = NumberUtils.toInt(jobIntervalString);
                    if (jobInterval > 0) {
                        jobIntervalConfigValue.value(jobInterval);
                    } else {
                        jobIntervalConfigValue.value(null);
                        jobIntervalConfigValue.addErrorMessage(jobIntervalConfig + " must be greater than 0");
                    }
                } else {
                    jobIntervalConfigValue.value(null);
                    jobIntervalConfigValue.addErrorMessage(jobIntervalConfig + " must be a numeric value greater than 0");
                }

                String jobFieldsConfig = JOBS_CONFIG + "." + jobName + "." + FIELDS_CONFIG;
                final ConfigValue jobFieldsConfigValue = new ConfigValue(jobFieldsConfig);
                if (!connectorConfigs.containsKey(jobFieldsConfig)) {
                    jobFieldsConfigValue.value(null);
                    jobFieldsConfigValue.addErrorMessage(jobFieldsConfig + " is mandatory");
                } else {
                    String[] jobFieldNames = connectorConfigs.getOrDefault(jobFieldsConfig, "").split(",");
                    jobFieldsConfigValue.value(jobFieldNames);
                    if (jobFieldNames.length == 0) {
                        jobFieldsConfigValue.addErrorMessage(jobFieldsConfig + " at least has to contain one field name");
                    } else {
                        for (String jobFieldName : jobFieldNames) {
                            String jobFieldAddressConfig =
                                JOBS_CONFIG + "." + jobName + "." + FIELDS_CONFIG + "." + jobFieldName;
                            final ConfigValue jobFieldAddressConfigValue = new ConfigValue(jobFieldAddressConfig);
                            String jobFieldAddress = connectorConfigs.get(jobFieldAddressConfig);
                            jobFieldAddressConfigValue.value(jobFieldAddress);
                            if ((jobFieldAddress == null) || jobFieldAddress.isEmpty()) {
                                jobFieldAddressConfigValue.addErrorMessage(jobFieldAddressConfig + " is mandatory");
                            }
                            // TODO: Validate the address ...
                        }
                    }
                }

                foundJobs.add(jobName);
            }
        }

        // Configure the sources
        if(sources != null) {
            final List<String> sourceNames = (List<String>) sources.value();
            for (String sourceName : sourceNames) {
                String connectionStringConfig = SOURCES_CONFIG + "." + sourceName + "." + CONNECTION_STRING_CONFIG;
                final ConfigValue sourceConnectionStringConfigValue = new ConfigValue(connectionStringConfig);
                config.configValues().add(sourceConnectionStringConfigValue);
                String connectionString = connectorConfigs.get(connectionStringConfig);
                sourceConnectionStringConfigValue.value();
                if (connectionString == null) {
                    sourceConnectionStringConfigValue.addErrorMessage(connectionStringConfig + " is mandatory");
                } else {
                    // TODO: Check if the connection string is valid.

                    String sourceTopicConfig = SOURCES_CONFIG + "." + sourceName + "." + TOPIC_CONFIG;
                    final ConfigValue sourceTopicConfigValue = new ConfigValue(sourceTopicConfig);
                    config.configValues().add(sourceTopicConfigValue);
                    String sourceTopic = connectorConfigs.get(sourceTopicConfig);
                    sourceTopicConfigValue.value(sourceTopic);

                    String jobReferenceNamesConfig = SOURCES_CONFIG + "." + sourceName + "." + JOB_REFERENCES_CONFIG;
                    final ConfigValue jobReferenceNamesConfigValue = new ConfigValue(jobReferenceNamesConfig);
                    jobReferenceNamesConfigValue.recommendedValues(foundJobs);
                    config.configValues().add(jobReferenceNamesConfigValue);
                    if(!connectorConfigs.containsKey(jobReferenceNamesConfig)) {
                        jobReferenceNamesConfigValue.value(null);
                        jobReferenceNamesConfigValue.addErrorMessage(jobReferenceNamesConfig + " is mandatory");
                    } else {
                        String[] jobReferenceNames = connectorConfigs.getOrDefault(jobReferenceNamesConfig, "").split(",");
                        jobReferenceNamesConfigValue.value(jobReferenceNames);
                        // Check at least one job is referenced
                        if (jobReferenceNames.length == 0) {
                            jobReferenceNamesConfigValue.addErrorMessage(jobReferenceNamesConfig + " is mandatory");
                        }
                        for (String jobReferenceName : jobReferenceNames) {
                            // Check the references reference configured jobs
                            if (!foundJobs.contains(jobReferenceName)) {
                                jobReferenceNamesConfigValue.addErrorMessage(jobReferenceNamesConfig +
                                    " references non-existent job " + jobReferenceName);
                            }
                            // Check if a topic is specified at some level
                            else {
                                String jobReferenceTopicNameConfig = SOURCES_CONFIG + "." + sourceName + "." +
                                    JOB_REFERENCES_CONFIG + "." + jobReferenceName + TOPIC_CONFIG;
                                final ConfigValue jobReferenceTopicNameConfigValue = new ConfigValue(jobReferenceTopicNameConfig);
                                config.configValues().add(jobReferenceTopicNameConfigValue);
                                String jobReferenceTopic = connectorConfigs.get(jobReferenceTopicNameConfig);
                                jobReferenceTopicNameConfigValue.value(jobReferenceTopic);
                                if ((jobReferenceTopic == null) && (sourceTopic == null) && (defaultTopic == null)) {
                                    jobReferenceTopicNameConfigValue.addErrorMessage(
                                        "No topic definition found at any level for " + jobReferenceTopicNameConfig);
                                }
                            }
                        }
                    }
                }
            }
        }

        return config;
    }

    @Override
    public ConfigDef config() {
        return new ConfigDef()
            .define(DEFAULT_TOPIC_CONFIG, ConfigDef.Type.STRING, ConfigDef.Importance.LOW, DEFAULT_TOPIC_DOC)
            .define(SOURCES_CONFIG, ConfigDef.Type.LIST, ConfigDef.Importance.HIGH, SOURCES_DOC)
            .define(JOBS_CONFIG, ConfigDef.Type.LIST, ConfigDef.Importance.HIGH, JOBS_DOC);
    }

    @Override
    public String version() {
        return VersionUtil.getVersion();
    }

}
