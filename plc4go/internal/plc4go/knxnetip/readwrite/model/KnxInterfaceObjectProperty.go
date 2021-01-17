//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package model

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

type KnxInterfaceObjectProperty uint32

type IKnxInterfaceObjectProperty interface {
	PropertyDataType() KnxPropertyDataType
	Name() string
	PropertyId() uint8
	ObjectType() KnxInterfaceObjectType
	Serialize(io utils.WriteBuffer) error
}

const (
	KnxInterfaceObjectProperty_PID_UNKNOWN                                                          KnxInterfaceObjectProperty = 0
	KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_TYPE                                              KnxInterfaceObjectProperty = 1
	KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_NAME                                              KnxInterfaceObjectProperty = 2
	KnxInterfaceObjectProperty_PID_GENERAL_SEMAPHOR                                                 KnxInterfaceObjectProperty = 3
	KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_REFERENCE                                   KnxInterfaceObjectProperty = 4
	KnxInterfaceObjectProperty_PID_GENERAL_LOAD_STATE_CONTROL                                       KnxInterfaceObjectProperty = 5
	KnxInterfaceObjectProperty_PID_GENERAL_RUN_STATE_CONTROL                                        KnxInterfaceObjectProperty = 6
	KnxInterfaceObjectProperty_PID_GENERAL_TABLE_REFERENCE                                          KnxInterfaceObjectProperty = 7
	KnxInterfaceObjectProperty_PID_GENERAL_SERVICE_CONTROL                                          KnxInterfaceObjectProperty = 8
	KnxInterfaceObjectProperty_PID_GENERAL_FIRMWARE_REVISION                                        KnxInterfaceObjectProperty = 9
	KnxInterfaceObjectProperty_PID_GENERAL_SERVICES_SUPPORTED                                       KnxInterfaceObjectProperty = 10
	KnxInterfaceObjectProperty_PID_GENERAL_SERIAL_NUMBER                                            KnxInterfaceObjectProperty = 11
	KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_ID                                          KnxInterfaceObjectProperty = 12
	KnxInterfaceObjectProperty_PID_GENERAL_PROGRAM_VERSION                                          KnxInterfaceObjectProperty = 13
	KnxInterfaceObjectProperty_PID_GENERAL_DEVICE_CONTROL                                           KnxInterfaceObjectProperty = 14
	KnxInterfaceObjectProperty_PID_GENERAL_ORDER_INFO                                               KnxInterfaceObjectProperty = 15
	KnxInterfaceObjectProperty_PID_GENERAL_PEI_TYPE                                                 KnxInterfaceObjectProperty = 16
	KnxInterfaceObjectProperty_PID_GENERAL_PORT_CONFIGURATION                                       KnxInterfaceObjectProperty = 17
	KnxInterfaceObjectProperty_PID_GENERAL_POLL_GROUP_SETTINGS                                      KnxInterfaceObjectProperty = 18
	KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_DATA                                        KnxInterfaceObjectProperty = 19
	KnxInterfaceObjectProperty_PID_GENERAL_ENABLE                                                   KnxInterfaceObjectProperty = 20
	KnxInterfaceObjectProperty_PID_GENERAL_DESCRIPTION                                              KnxInterfaceObjectProperty = 21
	KnxInterfaceObjectProperty_PID_GENERAL_FILE                                                     KnxInterfaceObjectProperty = 22
	KnxInterfaceObjectProperty_PID_GENERAL_TABLE                                                    KnxInterfaceObjectProperty = 23
	KnxInterfaceObjectProperty_PID_GENERAL_ENROL                                                    KnxInterfaceObjectProperty = 24
	KnxInterfaceObjectProperty_PID_GENERAL_VERSION                                                  KnxInterfaceObjectProperty = 25
	KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_LINK                                        KnxInterfaceObjectProperty = 26
	KnxInterfaceObjectProperty_PID_GENERAL_MCB_TABLE                                                KnxInterfaceObjectProperty = 27
	KnxInterfaceObjectProperty_PID_GENERAL_ERROR_CODE                                               KnxInterfaceObjectProperty = 28
	KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_INDEX                                             KnxInterfaceObjectProperty = 29
	KnxInterfaceObjectProperty_PID_GENERAL_DOWNLOAD_COUNTER                                         KnxInterfaceObjectProperty = 30
	KnxInterfaceObjectProperty_PID_DEVICE_ROUTING_COUNT                                             KnxInterfaceObjectProperty = 31
	KnxInterfaceObjectProperty_PID_DEVICE_MAX_RETRY_COUNT                                           KnxInterfaceObjectProperty = 32
	KnxInterfaceObjectProperty_PID_DEVICE_ERROR_FLAGS                                               KnxInterfaceObjectProperty = 33
	KnxInterfaceObjectProperty_PID_DEVICE_PROGMODE                                                  KnxInterfaceObjectProperty = 34
	KnxInterfaceObjectProperty_PID_DEVICE_PRODUCT_ID                                                KnxInterfaceObjectProperty = 35
	KnxInterfaceObjectProperty_PID_DEVICE_MAX_APDULENGTH                                            KnxInterfaceObjectProperty = 36
	KnxInterfaceObjectProperty_PID_DEVICE_SUBNET_ADDR                                               KnxInterfaceObjectProperty = 37
	KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_ADDR                                               KnxInterfaceObjectProperty = 38
	KnxInterfaceObjectProperty_PID_DEVICE_PB_CONFIG                                                 KnxInterfaceObjectProperty = 39
	KnxInterfaceObjectProperty_PID_DEVICE_ADDR_REPORT                                               KnxInterfaceObjectProperty = 40
	KnxInterfaceObjectProperty_PID_DEVICE_ADDR_CHECK                                                KnxInterfaceObjectProperty = 41
	KnxInterfaceObjectProperty_PID_DEVICE_OBJECT_VALUE                                              KnxInterfaceObjectProperty = 42
	KnxInterfaceObjectProperty_PID_DEVICE_OBJECTLINK                                                KnxInterfaceObjectProperty = 43
	KnxInterfaceObjectProperty_PID_DEVICE_APPLICATION                                               KnxInterfaceObjectProperty = 44
	KnxInterfaceObjectProperty_PID_DEVICE_PARAMETER                                                 KnxInterfaceObjectProperty = 45
	KnxInterfaceObjectProperty_PID_DEVICE_OBJECTADDRESS                                             KnxInterfaceObjectProperty = 46
	KnxInterfaceObjectProperty_PID_DEVICE_PSU_TYPE                                                  KnxInterfaceObjectProperty = 47
	KnxInterfaceObjectProperty_PID_DEVICE_PSU_STATUS                                                KnxInterfaceObjectProperty = 48
	KnxInterfaceObjectProperty_PID_DEVICE_PSU_ENABLE                                                KnxInterfaceObjectProperty = 49
	KnxInterfaceObjectProperty_PID_DEVICE_DOMAIN_ADDRESS                                            KnxInterfaceObjectProperty = 50
	KnxInterfaceObjectProperty_PID_DEVICE_IO_LIST                                                   KnxInterfaceObjectProperty = 51
	KnxInterfaceObjectProperty_PID_DEVICE_MGT_DESCRIPTOR_01                                         KnxInterfaceObjectProperty = 52
	KnxInterfaceObjectProperty_PID_DEVICE_PL110_PARAM                                               KnxInterfaceObjectProperty = 53
	KnxInterfaceObjectProperty_PID_DEVICE_RF_REPEAT_COUNTER                                         KnxInterfaceObjectProperty = 54
	KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_TABLE                                       KnxInterfaceObjectProperty = 55
	KnxInterfaceObjectProperty_PID_DEVICE_RANDOM_PAUSE_TABLE                                        KnxInterfaceObjectProperty = 56
	KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_NR                                          KnxInterfaceObjectProperty = 57
	KnxInterfaceObjectProperty_PID_DEVICE_HARDWARE_TYPE                                             KnxInterfaceObjectProperty = 58
	KnxInterfaceObjectProperty_PID_DEVICE_RETRANSMITTER_NUMBER                                      KnxInterfaceObjectProperty = 59
	KnxInterfaceObjectProperty_PID_DEVICE_SERIAL_NR_TABLE                                           KnxInterfaceObjectProperty = 60
	KnxInterfaceObjectProperty_PID_DEVICE_BIBATMASTER_ADDRESS                                       KnxInterfaceObjectProperty = 61
	KnxInterfaceObjectProperty_PID_DEVICE_RF_DOMAIN_ADDRESS                                         KnxInterfaceObjectProperty = 62
	KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_DESCRIPTOR                                         KnxInterfaceObjectProperty = 63
	KnxInterfaceObjectProperty_PID_DEVICE_METERING_FILTER_TABLE                                     KnxInterfaceObjectProperty = 64
	KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE                         KnxInterfaceObjectProperty = 65
	KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR                      KnxInterfaceObjectProperty = 66
	KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_GRPOBJTABLE                                   KnxInterfaceObjectProperty = 67
	KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE                           KnxInterfaceObjectProperty = 68
	KnxInterfaceObjectProperty_PID_ROUTER_LINE_STATUS                                               KnxInterfaceObjectProperty = 69
	KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCCONFIG                                             KnxInterfaceObjectProperty = 70
	KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCCONFIG                                              KnxInterfaceObjectProperty = 71
	KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCGRPCONFIG                                          KnxInterfaceObjectProperty = 72
	KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCGRPCONFIG                                           KnxInterfaceObjectProperty = 73
	KnxInterfaceObjectProperty_PID_ROUTER_ROUTETABLE_CONTROL                                        KnxInterfaceObjectProperty = 74
	KnxInterfaceObjectProperty_PID_ROUTER_COUPL_SERV_CONTROL                                        KnxInterfaceObjectProperty = 75
	KnxInterfaceObjectProperty_PID_ROUTER_MAX_ROUTER_APDU_LENGTH                                    KnxInterfaceObjectProperty = 76
	KnxInterfaceObjectProperty_PID_ROUTER_MEDIUM                                                    KnxInterfaceObjectProperty = 77
	KnxInterfaceObjectProperty_PID_ROUTER_FILTER_TABLE_USE                                          KnxInterfaceObjectProperty = 78
	KnxInterfaceObjectProperty_PID_ROUTER_RF_ENABLE_SBC                                             KnxInterfaceObjectProperty = 79
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID                          KnxInterfaceObjectProperty = 80
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS                           KnxInterfaceObjectProperty = 81
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES                  KnxInterfaceObjectProperty = 82
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD                     KnxInterfaceObjectProperty = 83
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD                             KnxInterfaceObjectProperty = 84
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_CAPABILITIES                                  KnxInterfaceObjectProperty = 85
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS                               KnxInterfaceObjectProperty = 86
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK                              KnxInterfaceObjectProperty = 87
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY                          KnxInterfaceObjectProperty = 88
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ADDRESS                                       KnxInterfaceObjectProperty = 89
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SUBNET_MASK                                      KnxInterfaceObjectProperty = 90
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEFAULT_GATEWAY                                  KnxInterfaceObjectProperty = 91
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER                                KnxInterfaceObjectProperty = 92
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MAC_ADDRESS                                      KnxInterfaceObjectProperty = 93
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS                   KnxInterfaceObjectProperty = 94
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS                        KnxInterfaceObjectProperty = 95
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TTL                                              KnxInterfaceObjectProperty = 96
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES                     KnxInterfaceObjectProperty = 97
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE                            KnxInterfaceObjectProperty = 98
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES                    KnxInterfaceObjectProperty = 99
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED                            KnxInterfaceObjectProperty = 100
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP                             KnxInterfaceObjectProperty = 101
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX                            KnxInterfaceObjectProperty = 102
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP                               KnxInterfaceObjectProperty = 103
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX                              KnxInterfaceObjectProperty = 104
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_FRIENDLY_NAME                                    KnxInterfaceObjectProperty = 105
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_BACKBONE_KEY                                     KnxInterfaceObjectProperty = 106
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE                       KnxInterfaceObjectProperty = 107
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PASSWORD_HASHES                                  KnxInterfaceObjectProperty = 108
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES                         KnxInterfaceObjectProperty = 109
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE                      KnxInterfaceObjectProperty = 110
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION                            KnxInterfaceObjectProperty = 111
	KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TUNNELLING_USERS                                 KnxInterfaceObjectProperty = 112
	KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_MODE                                           KnxInterfaceObjectProperty = 113
	KnxInterfaceObjectProperty_PID_SECURITY_P2P_KEY_TABLE                                           KnxInterfaceObjectProperty = 114
	KnxInterfaceObjectProperty_PID_SECURITY_GRP_KEY_TABLE                                           KnxInterfaceObjectProperty = 115
	KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE                       KnxInterfaceObjectProperty = 116
	KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_FAILURES_LOG                                   KnxInterfaceObjectProperty = 117
	KnxInterfaceObjectProperty_PID_SECURITY_SKI_TOOL                                                KnxInterfaceObjectProperty = 118
	KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT                                         KnxInterfaceObjectProperty = 119
	KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT_CONTROL                                 KnxInterfaceObjectProperty = 120
	KnxInterfaceObjectProperty_PID_SECURITY_SEQUENCE_NUMBER_SENDING                                 KnxInterfaceObjectProperty = 121
	KnxInterfaceObjectProperty_PID_SECURITY_ZONE_KEYS_TABLE                                         KnxInterfaceObjectProperty = 122
	KnxInterfaceObjectProperty_PID_SECURITY_GO_SECURITY_FLAGS                                       KnxInterfaceObjectProperty = 123
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_MULTI_TYPE                                          KnxInterfaceObjectProperty = 124
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DOMAIN_ADDRESS                                      KnxInterfaceObjectProperty = 125
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_RETRANSMITTER                                       KnxInterfaceObjectProperty = 126
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_SECURITY_REPORT_CONTROL                                KnxInterfaceObjectProperty = 127
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT                               KnxInterfaceObjectProperty = 128
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_BIDIR_TIMEOUT                                       KnxInterfaceObjectProperty = 129
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE                                KnxInterfaceObjectProperty = 130
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE                                  KnxInterfaceObjectProperty = 131
	KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_PROBE                                          KnxInterfaceObjectProperty = 132
	KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE                         KnxInterfaceObjectProperty = 133
	KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME                         KnxInterfaceObjectProperty = 134
	KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE                          KnxInterfaceObjectProperty = 135
	KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME                          KnxInterfaceObjectProperty = 136
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY                          KnxInterfaceObjectProperty = 137
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY                         KnxInterfaceObjectProperty = 138
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION                 KnxInterfaceObjectProperty = 139
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION               KnxInterfaceObjectProperty = 140
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME           KnxInterfaceObjectProperty = 141
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY        KnxInterfaceObjectProperty = 142
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING              KnxInterfaceObjectProperty = 143
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING            KnxInterfaceObjectProperty = 144
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP            KnxInterfaceObjectProperty = 145
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN          KnxInterfaceObjectProperty = 146
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE               KnxInterfaceObjectProperty = 147
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION       KnxInterfaceObjectProperty = 148
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE                 KnxInterfaceObjectProperty = 149
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE                KnxInterfaceObjectProperty = 150
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE                        KnxInterfaceObjectProperty = 151
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE                      KnxInterfaceObjectProperty = 152
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER            KnxInterfaceObjectProperty = 153
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE        KnxInterfaceObjectProperty = 154
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE                KnxInterfaceObjectProperty = 155
	KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2          KnxInterfaceObjectProperty = 156
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_ON_DELAY                                  KnxInterfaceObjectProperty = 157
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY                                 KnxInterfaceObjectProperty = 158
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME          KnxInterfaceObjectProperty = 159
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION                         KnxInterfaceObjectProperty = 160
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION                       KnxInterfaceObjectProperty = 161
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME                   KnxInterfaceObjectProperty = 162
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY                KnxInterfaceObjectProperty = 163
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED                             KnxInterfaceObjectProperty = 164
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME                         KnxInterfaceObjectProperty = 165
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE     KnxInterfaceObjectProperty = 166
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF              KnxInterfaceObjectProperty = 167
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE KnxInterfaceObjectProperty = 168
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF          KnxInterfaceObjectProperty = 169
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS                    KnxInterfaceObjectProperty = 170
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE                         KnxInterfaceObjectProperty = 171
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE                         KnxInterfaceObjectProperty = 172
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE                       KnxInterfaceObjectProperty = 173
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION                       KnxInterfaceObjectProperty = 174
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE                        KnxInterfaceObjectProperty = 175
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION                           KnxInterfaceObjectProperty = 176
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION               KnxInterfaceObjectProperty = 177
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE                         KnxInterfaceObjectProperty = 178
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE                        KnxInterfaceObjectProperty = 179
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING                      KnxInterfaceObjectProperty = 180
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING                    KnxInterfaceObjectProperty = 181
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE                             KnxInterfaceObjectProperty = 182
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE                           KnxInterfaceObjectProperty = 183
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE                       KnxInterfaceObjectProperty = 184
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE                KnxInterfaceObjectProperty = 185
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE                       KnxInterfaceObjectProperty = 186
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP                    KnxInterfaceObjectProperty = 187
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE          KnxInterfaceObjectProperty = 188
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN                  KnxInterfaceObjectProperty = 189
	KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE                  KnxInterfaceObjectProperty = 190
	KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION                               KnxInterfaceObjectProperty = 191
	KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE                          KnxInterfaceObjectProperty = 192
	KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE                           KnxInterfaceObjectProperty = 193
	KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION                             KnxInterfaceObjectProperty = 194
	KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE                        KnxInterfaceObjectProperty = 195
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME                     KnxInterfaceObjectProperty = 196
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME                        KnxInterfaceObjectProperty = 197
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME                           KnxInterfaceObjectProperty = 198
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME                KnxInterfaceObjectProperty = 199
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT       KnxInterfaceObjectProperty = 200
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH           KnxInterfaceObjectProperty = 201
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT             KnxInterfaceObjectProperty = 202
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE               KnxInterfaceObjectProperty = 203
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM                      KnxInterfaceObjectProperty = 204
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM                     KnxInterfaceObjectProperty = 205
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM                   KnxInterfaceObjectProperty = 206
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM                     KnxInterfaceObjectProperty = 207
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM                     KnxInterfaceObjectProperty = 208
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM                    KnxInterfaceObjectProperty = 209
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME                       KnxInterfaceObjectProperty = 210
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE                       KnxInterfaceObjectProperty = 211
	KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE              KnxInterfaceObjectProperty = 212
	KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE                         KnxInterfaceObjectProperty = 213
	KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION                             KnxInterfaceObjectProperty = 214
	KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE                         KnxInterfaceObjectProperty = 215
)

func (e KnxInterfaceObjectProperty) PropertyDataType() KnxPropertyDataType {
	switch e {
	case 0:
		{ /* '0' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 1:
		{ /* '1' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 10:
		{ /* '10' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 100:
		{ /* '100' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 101:
		{ /* '101' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 102:
		{ /* '102' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 103:
		{ /* '103' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 104:
		{ /* '104' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 105:
		{ /* '105' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 106:
		{ /* '106' */
			return KnxPropertyDataType_PDT_GENERIC_16
		}
	case 107:
		{ /* '107' */
			return KnxPropertyDataType_PDT_GENERIC_16
		}
	case 108:
		{ /* '108' */
			return KnxPropertyDataType_PDT_GENERIC_16
		}
	case 109:
		{ /* '109' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 11:
		{ /* '11' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 110:
		{ /* '110' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 111:
		{ /* '111' */
			return KnxPropertyDataType_PDT_SCALING
		}
	case 112:
		{ /* '112' */
			return KnxPropertyDataType_PDT_GENERIC_02
		}
	case 113:
		{ /* '113' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 114:
		{ /* '114' */
			return KnxPropertyDataType_PDT_GENERIC_18
		}
	case 115:
		{ /* '115' */
			return KnxPropertyDataType_PDT_GENERIC_18
		}
	case 116:
		{ /* '116' */
			return KnxPropertyDataType_PDT_GENERIC_08
		}
	case 117:
		{ /* '117' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 118:
		{ /* '118' */
			return KnxPropertyDataType_PDT_GENERIC_16
		}
	case 119:
		{ /* '119' */
			return KnxPropertyDataType_PDT_BITSET8
		}
	case 12:
		{ /* '12' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 120:
		{ /* '120' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 121:
		{ /* '121' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 122:
		{ /* '122' */
			return KnxPropertyDataType_PDT_GENERIC_19
		}
	case 123:
		{ /* '123' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 124:
		{ /* '124' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 125:
		{ /* '125' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 126:
		{ /* '126' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 127:
		{ /* '127' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 128:
		{ /* '128' */
			return KnxPropertyDataType_PDT_BITSET8
		}
	case 129:
		{ /* '129' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 13:
		{ /* '13' */
			return KnxPropertyDataType_PDT_GENERIC_05
		}
	case 130:
		{ /* '130' */
			return KnxPropertyDataType_PDT_GENERIC_03
		}
	case 131:
		{ /* '131' */
			return KnxPropertyDataType_PDT_GENERIC_04
		}
	case 132:
		{ /* '132' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 133:
		{ /* '133' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 134:
		{ /* '134' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 135:
		{ /* '135' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 136:
		{ /* '136' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 137:
		{ /* '137' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 138:
		{ /* '138' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 139:
		{ /* '139' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 14:
		{ /* '14' */
			return KnxPropertyDataType_PDT_BITSET8
		}
	case 140:
		{ /* '140' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 141:
		{ /* '141' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 142:
		{ /* '142' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 143:
		{ /* '143' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 144:
		{ /* '144' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 145:
		{ /* '145' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 146:
		{ /* '146' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 147:
		{ /* '147' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 148:
		{ /* '148' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 149:
		{ /* '149' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 15:
		{ /* '15' */
			return KnxPropertyDataType_PDT_GENERIC_10
		}
	case 150:
		{ /* '150' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 151:
		{ /* '151' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 152:
		{ /* '152' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 153:
		{ /* '153' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 154:
		{ /* '154' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 155:
		{ /* '155' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 156:
		{ /* '156' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 157:
		{ /* '157' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 158:
		{ /* '158' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 159:
		{ /* '159' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 16:
		{ /* '16' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 160:
		{ /* '160' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 161:
		{ /* '161' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 162:
		{ /* '162' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 163:
		{ /* '163' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 164:
		{ /* '164' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 165:
		{ /* '165' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 166:
		{ /* '166' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 167:
		{ /* '167' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 168:
		{ /* '168' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 169:
		{ /* '169' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 17:
		{ /* '17' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 170:
		{ /* '170' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 171:
		{ /* '171' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 172:
		{ /* '172' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 173:
		{ /* '173' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 174:
		{ /* '174' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 175:
		{ /* '175' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 176:
		{ /* '176' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 177:
		{ /* '177' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 178:
		{ /* '178' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 179:
		{ /* '179' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 18:
		{ /* '18' */
			return KnxPropertyDataType_PDT_POLL_GROUP_SETTINGS
		}
	case 180:
		{ /* '180' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 181:
		{ /* '181' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 182:
		{ /* '182' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 183:
		{ /* '183' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 184:
		{ /* '184' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 185:
		{ /* '185' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 186:
		{ /* '186' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 187:
		{ /* '187' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 188:
		{ /* '188' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 189:
		{ /* '189' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 19:
		{ /* '19' */
			return KnxPropertyDataType_PDT_GENERIC_04
		}
	case 190:
		{ /* '190' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 191:
		{ /* '191' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 192:
		{ /* '192' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 193:
		{ /* '193' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 194:
		{ /* '194' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 195:
		{ /* '195' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 196:
		{ /* '196' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 197:
		{ /* '197' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 198:
		{ /* '198' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 199:
		{ /* '199' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 2:
		{ /* '2' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 20:
		{ /* '20' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 200:
		{ /* '200' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 201:
		{ /* '201' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 202:
		{ /* '202' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 203:
		{ /* '203' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 204:
		{ /* '204' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 205:
		{ /* '205' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 206:
		{ /* '206' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 207:
		{ /* '207' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 208:
		{ /* '208' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 209:
		{ /* '209' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 21:
		{ /* '21' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 210:
		{ /* '210' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 211:
		{ /* '211' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 212:
		{ /* '212' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 213:
		{ /* '213' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 214:
		{ /* '214' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 215:
		{ /* '215' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 22:
		{ /* '22' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 23:
		{ /* '23' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 24:
		{ /* '24' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 25:
		{ /* '25' */
			return KnxPropertyDataType_PDT_VERSION
		}
	case 26:
		{ /* '26' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 27:
		{ /* '27' */
			return KnxPropertyDataType_PDT_GENERIC_08
		}
	case 28:
		{ /* '28' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 29:
		{ /* '29' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 3:
		{ /* '3' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 30:
		{ /* '30' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 31:
		{ /* '31' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 32:
		{ /* '32' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 33:
		{ /* '33' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 34:
		{ /* '34' */
			return KnxPropertyDataType_PDT_BITSET8
		}
	case 35:
		{ /* '35' */
			return KnxPropertyDataType_PDT_GENERIC_10
		}
	case 36:
		{ /* '36' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 37:
		{ /* '37' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 38:
		{ /* '38' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 39:
		{ /* '39' */
			return KnxPropertyDataType_PDT_GENERIC_04
		}
	case 4:
		{ /* '4' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 40:
		{ /* '40' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 41:
		{ /* '41' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 42:
		{ /* '42' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 43:
		{ /* '43' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 44:
		{ /* '44' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 45:
		{ /* '45' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 46:
		{ /* '46' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 47:
		{ /* '47' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 48:
		{ /* '48' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 49:
		{ /* '49' */
			return KnxPropertyDataType_PDT_ENUM8
		}
	case 5:
		{ /* '5' */
			return KnxPropertyDataType_PDT_CONTROL
		}
	case 50:
		{ /* '50' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 51:
		{ /* '51' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 52:
		{ /* '52' */
			return KnxPropertyDataType_PDT_GENERIC_10
		}
	case 53:
		{ /* '53' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 54:
		{ /* '54' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 55:
		{ /* '55' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 56:
		{ /* '56' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 57:
		{ /* '57' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 58:
		{ /* '58' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 59:
		{ /* '59' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 6:
		{ /* '6' */
			return KnxPropertyDataType_PDT_CONTROL
		}
	case 60:
		{ /* '60' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 61:
		{ /* '61' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 62:
		{ /* '62' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 63:
		{ /* '63' */
			return KnxPropertyDataType_PDT_GENERIC_02
		}
	case 64:
		{ /* '64' */
			return KnxPropertyDataType_PDT_GENERIC_08
		}
	case 65:
		{ /* '65' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 66:
		{ /* '66' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 67:
		{ /* '67' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 68:
		{ /* '68' */
			return KnxPropertyDataType_PDT_GENERIC_08
		}
	case 69:
		{ /* '69' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 7:
		{ /* '7' */
			return KnxPropertyDataType_PDT_UNKNOWN
		}
	case 70:
		{ /* '70' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 71:
		{ /* '71' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 72:
		{ /* '72' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 73:
		{ /* '73' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 74:
		{ /* '74' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 75:
		{ /* '75' */
			return KnxPropertyDataType_PDT_GENERIC_01
		}
	case 76:
		{ /* '76' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 77:
		{ /* '77' */
			return KnxPropertyDataType_PDT_ENUM8
		}
	case 78:
		{ /* '78' */
			return KnxPropertyDataType_PDT_BINARY_INFORMATION
		}
	case 79:
		{ /* '79' */
			return KnxPropertyDataType_PDT_FUNCTION
		}
	case 8:
		{ /* '8' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 80:
		{ /* '80' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 81:
		{ /* '81' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 82:
		{ /* '82' */
			return KnxPropertyDataType_PDT_UNSIGNED_INT
		}
	case 83:
		{ /* '83' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 84:
		{ /* '84' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 85:
		{ /* '85' */
			return KnxPropertyDataType_PDT_BITSET8
		}
	case 86:
		{ /* '86' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 87:
		{ /* '87' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 88:
		{ /* '88' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 89:
		{ /* '89' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 9:
		{ /* '9' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 90:
		{ /* '90' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 91:
		{ /* '91' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 92:
		{ /* '92' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 93:
		{ /* '93' */
			return KnxPropertyDataType_PDT_GENERIC_06
		}
	case 94:
		{ /* '94' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 95:
		{ /* '95' */
			return KnxPropertyDataType_PDT_UNSIGNED_LONG
		}
	case 96:
		{ /* '96' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 97:
		{ /* '97' */
			return KnxPropertyDataType_PDT_BITSET16
		}
	case 98:
		{ /* '98' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	case 99:
		{ /* '99' */
			return KnxPropertyDataType_PDT_UNSIGNED_CHAR
		}
	default:
		{
			return 0
		}
	}
}

func (e KnxInterfaceObjectProperty) Name() string {
	switch e {
	case 0:
		{ /* '0' */
			return "Unknown Interface Object Property"
		}
	case 1:
		{ /* '1' */
			return "Interface Object Type"
		}
	case 10:
		{ /* '10' */
			return "Services Supported"
		}
	case 100:
		{ /* '100' */
			return ""
		}
	case 101:
		{ /* '101' */
			return ""
		}
	case 102:
		{ /* '102' */
			return ""
		}
	case 103:
		{ /* '103' */
			return ""
		}
	case 104:
		{ /* '104' */
			return ""
		}
	case 105:
		{ /* '105' */
			return ""
		}
	case 106:
		{ /* '106' */
			return ""
		}
	case 107:
		{ /* '107' */
			return ""
		}
	case 108:
		{ /* '108' */
			return ""
		}
	case 109:
		{ /* '109' */
			return ""
		}
	case 11:
		{ /* '11' */
			return "KNX Serial Number"
		}
	case 110:
		{ /* '110' */
			return ""
		}
	case 111:
		{ /* '111' */
			return ""
		}
	case 112:
		{ /* '112' */
			return ""
		}
	case 113:
		{ /* '113' */
			return ""
		}
	case 114:
		{ /* '114' */
			return ""
		}
	case 115:
		{ /* '115' */
			return ""
		}
	case 116:
		{ /* '116' */
			return ""
		}
	case 117:
		{ /* '117' */
			return ""
		}
	case 118:
		{ /* '118' */
			return ""
		}
	case 119:
		{ /* '119' */
			return ""
		}
	case 12:
		{ /* '12' */
			return "Manufacturer Identifier"
		}
	case 120:
		{ /* '120' */
			return ""
		}
	case 121:
		{ /* '121' */
			return ""
		}
	case 122:
		{ /* '122' */
			return ""
		}
	case 123:
		{ /* '123' */
			return ""
		}
	case 124:
		{ /* '124' */
			return ""
		}
	case 125:
		{ /* '125' */
			return ""
		}
	case 126:
		{ /* '126' */
			return ""
		}
	case 127:
		{ /* '127' */
			return ""
		}
	case 128:
		{ /* '128' */
			return ""
		}
	case 129:
		{ /* '129' */
			return ""
		}
	case 13:
		{ /* '13' */
			return "Application Version"
		}
	case 130:
		{ /* '130' */
			return ""
		}
	case 131:
		{ /* '131' */
			return ""
		}
	case 132:
		{ /* '132' */
			return ""
		}
	case 133:
		{ /* '133' */
			return ""
		}
	case 134:
		{ /* '134' */
			return ""
		}
	case 135:
		{ /* '135' */
			return ""
		}
	case 136:
		{ /* '136' */
			return ""
		}
	case 137:
		{ /* '137' */
			return ""
		}
	case 138:
		{ /* '138' */
			return ""
		}
	case 139:
		{ /* '139' */
			return ""
		}
	case 14:
		{ /* '14' */
			return "Device Control"
		}
	case 140:
		{ /* '140' */
			return ""
		}
	case 141:
		{ /* '141' */
			return ""
		}
	case 142:
		{ /* '142' */
			return ""
		}
	case 143:
		{ /* '143' */
			return ""
		}
	case 144:
		{ /* '144' */
			return ""
		}
	case 145:
		{ /* '145' */
			return ""
		}
	case 146:
		{ /* '146' */
			return ""
		}
	case 147:
		{ /* '147' */
			return ""
		}
	case 148:
		{ /* '148' */
			return ""
		}
	case 149:
		{ /* '149' */
			return ""
		}
	case 15:
		{ /* '15' */
			return "Order Info"
		}
	case 150:
		{ /* '150' */
			return ""
		}
	case 151:
		{ /* '151' */
			return ""
		}
	case 152:
		{ /* '152' */
			return ""
		}
	case 153:
		{ /* '153' */
			return ""
		}
	case 154:
		{ /* '154' */
			return ""
		}
	case 155:
		{ /* '155' */
			return ""
		}
	case 156:
		{ /* '156' */
			return ""
		}
	case 157:
		{ /* '157' */
			return ""
		}
	case 158:
		{ /* '158' */
			return ""
		}
	case 159:
		{ /* '159' */
			return ""
		}
	case 16:
		{ /* '16' */
			return "PEI Type"
		}
	case 160:
		{ /* '160' */
			return ""
		}
	case 161:
		{ /* '161' */
			return ""
		}
	case 162:
		{ /* '162' */
			return ""
		}
	case 163:
		{ /* '163' */
			return ""
		}
	case 164:
		{ /* '164' */
			return ""
		}
	case 165:
		{ /* '165' */
			return ""
		}
	case 166:
		{ /* '166' */
			return ""
		}
	case 167:
		{ /* '167' */
			return ""
		}
	case 168:
		{ /* '168' */
			return ""
		}
	case 169:
		{ /* '169' */
			return ""
		}
	case 17:
		{ /* '17' */
			return "PortADDR"
		}
	case 170:
		{ /* '170' */
			return ""
		}
	case 171:
		{ /* '171' */
			return ""
		}
	case 172:
		{ /* '172' */
			return ""
		}
	case 173:
		{ /* '173' */
			return ""
		}
	case 174:
		{ /* '174' */
			return ""
		}
	case 175:
		{ /* '175' */
			return ""
		}
	case 176:
		{ /* '176' */
			return ""
		}
	case 177:
		{ /* '177' */
			return ""
		}
	case 178:
		{ /* '178' */
			return ""
		}
	case 179:
		{ /* '179' */
			return ""
		}
	case 18:
		{ /* '18' */
			return "Polling Group Settings"
		}
	case 180:
		{ /* '180' */
			return ""
		}
	case 181:
		{ /* '181' */
			return ""
		}
	case 182:
		{ /* '182' */
			return ""
		}
	case 183:
		{ /* '183' */
			return ""
		}
	case 184:
		{ /* '184' */
			return ""
		}
	case 185:
		{ /* '185' */
			return ""
		}
	case 186:
		{ /* '186' */
			return ""
		}
	case 187:
		{ /* '187' */
			return ""
		}
	case 188:
		{ /* '188' */
			return ""
		}
	case 189:
		{ /* '189' */
			return ""
		}
	case 19:
		{ /* '19' */
			return "Manufacturer Data"
		}
	case 190:
		{ /* '190' */
			return ""
		}
	case 191:
		{ /* '191' */
			return ""
		}
	case 192:
		{ /* '192' */
			return ""
		}
	case 193:
		{ /* '193' */
			return ""
		}
	case 194:
		{ /* '194' */
			return ""
		}
	case 195:
		{ /* '195' */
			return ""
		}
	case 196:
		{ /* '196' */
			return ""
		}
	case 197:
		{ /* '197' */
			return ""
		}
	case 198:
		{ /* '198' */
			return ""
		}
	case 199:
		{ /* '199' */
			return ""
		}
	case 2:
		{ /* '2' */
			return "Interface Object Name"
		}
	case 20:
		{ /* '20' */
			return ""
		}
	case 200:
		{ /* '200' */
			return ""
		}
	case 201:
		{ /* '201' */
			return ""
		}
	case 202:
		{ /* '202' */
			return ""
		}
	case 203:
		{ /* '203' */
			return ""
		}
	case 204:
		{ /* '204' */
			return ""
		}
	case 205:
		{ /* '205' */
			return ""
		}
	case 206:
		{ /* '206' */
			return ""
		}
	case 207:
		{ /* '207' */
			return ""
		}
	case 208:
		{ /* '208' */
			return ""
		}
	case 209:
		{ /* '209' */
			return ""
		}
	case 21:
		{ /* '21' */
			return "Description"
		}
	case 210:
		{ /* '210' */
			return ""
		}
	case 211:
		{ /* '211' */
			return ""
		}
	case 212:
		{ /* '212' */
			return ""
		}
	case 213:
		{ /* '213' */
			return ""
		}
	case 214:
		{ /* '214' */
			return ""
		}
	case 215:
		{ /* '215' */
			return ""
		}
	case 22:
		{ /* '22' */
			return ""
		}
	case 23:
		{ /* '23' */
			return "Table"
		}
	case 24:
		{ /* '24' */
			return "Interface Object Link"
		}
	case 25:
		{ /* '25' */
			return "Version"
		}
	case 26:
		{ /* '26' */
			return "Group Address Assignment"
		}
	case 27:
		{ /* '27' */
			return "Memory Control Table"
		}
	case 28:
		{ /* '28' */
			return "Error Code"
		}
	case 29:
		{ /* '29' */
			return "Object Index"
		}
	case 3:
		{ /* '3' */
			return "Semaphor"
		}
	case 30:
		{ /* '30' */
			return "Download Counter"
		}
	case 31:
		{ /* '31' */
			return "Routing Count"
		}
	case 32:
		{ /* '32' */
			return "Maximum Retry Count"
		}
	case 33:
		{ /* '33' */
			return "Error Flags"
		}
	case 34:
		{ /* '34' */
			return "Programming Mode"
		}
	case 35:
		{ /* '35' */
			return "Product Identification"
		}
	case 36:
		{ /* '36' */
			return "Max. APDU-Length"
		}
	case 37:
		{ /* '37' */
			return "Subnetwork Address"
		}
	case 38:
		{ /* '38' */
			return "Device Address"
		}
	case 39:
		{ /* '39' */
			return "Config Link"
		}
	case 4:
		{ /* '4' */
			return "Group Object Reference"
		}
	case 40:
		{ /* '40' */
			return ""
		}
	case 41:
		{ /* '41' */
			return ""
		}
	case 42:
		{ /* '42' */
			return ""
		}
	case 43:
		{ /* '43' */
			return ""
		}
	case 44:
		{ /* '44' */
			return ""
		}
	case 45:
		{ /* '45' */
			return ""
		}
	case 46:
		{ /* '46' */
			return ""
		}
	case 47:
		{ /* '47' */
			return ""
		}
	case 48:
		{ /* '48' */
			return ""
		}
	case 49:
		{ /* '49' */
			return ""
		}
	case 5:
		{ /* '5' */
			return "Load Control"
		}
	case 50:
		{ /* '50' */
			return "Domain Address"
		}
	case 51:
		{ /* '51' */
			return ""
		}
	case 52:
		{ /* '52' */
			return "Management Descriptor 1"
		}
	case 53:
		{ /* '53' */
			return "PL110 Parameters"
		}
	case 54:
		{ /* '54' */
			return ""
		}
	case 55:
		{ /* '55' */
			return ""
		}
	case 56:
		{ /* '56' */
			return ""
		}
	case 57:
		{ /* '57' */
			return ""
		}
	case 58:
		{ /* '58' */
			return "Hardware Type"
		}
	case 59:
		{ /* '59' */
			return ""
		}
	case 6:
		{ /* '6' */
			return "Run Control"
		}
	case 60:
		{ /* '60' */
			return ""
		}
	case 61:
		{ /* '61' */
			return ""
		}
	case 62:
		{ /* '62' */
			return "RF Domain Address"
		}
	case 63:
		{ /* '63' */
			return ""
		}
	case 64:
		{ /* '64' */
			return ""
		}
	case 65:
		{ /* '65' */
			return ""
		}
	case 66:
		{ /* '66' */
			return ""
		}
	case 67:
		{ /* '67' */
			return ""
		}
	case 68:
		{ /* '68' */
			return ""
		}
	case 69:
		{ /* '69' */
			return ""
		}
	case 7:
		{ /* '7' */
			return "Table Reference"
		}
	case 70:
		{ /* '70' */
			return ""
		}
	case 71:
		{ /* '71' */
			return ""
		}
	case 72:
		{ /* '72' */
			return ""
		}
	case 73:
		{ /* '73' */
			return ""
		}
	case 74:
		{ /* '74' */
			return ""
		}
	case 75:
		{ /* '75' */
			return ""
		}
	case 76:
		{ /* '76' */
			return ""
		}
	case 77:
		{ /* '77' */
			return ""
		}
	case 78:
		{ /* '78' */
			return ""
		}
	case 79:
		{ /* '79' */
			return ""
		}
	case 8:
		{ /* '8' */
			return "Service Control"
		}
	case 80:
		{ /* '80' */
			return "Project Installation Identification"
		}
	case 81:
		{ /* '81' */
			return "KNX Individual Address"
		}
	case 82:
		{ /* '82' */
			return "Additional Individual Addresses"
		}
	case 83:
		{ /* '83' */
			return ""
		}
	case 84:
		{ /* '84' */
			return ""
		}
	case 85:
		{ /* '85' */
			return ""
		}
	case 86:
		{ /* '86' */
			return ""
		}
	case 87:
		{ /* '87' */
			return ""
		}
	case 88:
		{ /* '88' */
			return ""
		}
	case 89:
		{ /* '89' */
			return ""
		}
	case 9:
		{ /* '9' */
			return "Firmware Revision"
		}
	case 90:
		{ /* '90' */
			return ""
		}
	case 91:
		{ /* '91' */
			return ""
		}
	case 92:
		{ /* '92' */
			return ""
		}
	case 93:
		{ /* '93' */
			return ""
		}
	case 94:
		{ /* '94' */
			return ""
		}
	case 95:
		{ /* '95' */
			return ""
		}
	case 96:
		{ /* '96' */
			return ""
		}
	case 97:
		{ /* '97' */
			return ""
		}
	case 98:
		{ /* '98' */
			return ""
		}
	case 99:
		{ /* '99' */
			return ""
		}
	default:
		{
			return ""
		}
	}
}

func (e KnxInterfaceObjectProperty) PropertyId() uint8 {
	switch e {
	case 0:
		{ /* '0' */
			return 0
		}
	case 1:
		{ /* '1' */
			return 1
		}
	case 10:
		{ /* '10' */
			return 10
		}
	case 100:
		{ /* '100' */
			return 71
		}
	case 101:
		{ /* '101' */
			return 72
		}
	case 102:
		{ /* '102' */
			return 73
		}
	case 103:
		{ /* '103' */
			return 74
		}
	case 104:
		{ /* '104' */
			return 75
		}
	case 105:
		{ /* '105' */
			return 76
		}
	case 106:
		{ /* '106' */
			return 91
		}
	case 107:
		{ /* '107' */
			return 92
		}
	case 108:
		{ /* '108' */
			return 93
		}
	case 109:
		{ /* '109' */
			return 94
		}
	case 11:
		{ /* '11' */
			return 11
		}
	case 110:
		{ /* '110' */
			return 95
		}
	case 111:
		{ /* '111' */
			return 96
		}
	case 112:
		{ /* '112' */
			return 97
		}
	case 113:
		{ /* '113' */
			return 51
		}
	case 114:
		{ /* '114' */
			return 52
		}
	case 115:
		{ /* '115' */
			return 53
		}
	case 116:
		{ /* '116' */
			return 54
		}
	case 117:
		{ /* '117' */
			return 55
		}
	case 118:
		{ /* '118' */
			return 56
		}
	case 119:
		{ /* '119' */
			return 57
		}
	case 12:
		{ /* '12' */
			return 12
		}
	case 120:
		{ /* '120' */
			return 58
		}
	case 121:
		{ /* '121' */
			return 59
		}
	case 122:
		{ /* '122' */
			return 60
		}
	case 123:
		{ /* '123' */
			return 61
		}
	case 124:
		{ /* '124' */
			return 51
		}
	case 125:
		{ /* '125' */
			return 56
		}
	case 126:
		{ /* '126' */
			return 57
		}
	case 127:
		{ /* '127' */
			return 58
		}
	case 128:
		{ /* '128' */
			return 59
		}
	case 129:
		{ /* '129' */
			return 60
		}
	case 13:
		{ /* '13' */
			return 13
		}
	case 130:
		{ /* '130' */
			return 61
		}
	case 131:
		{ /* '131' */
			return 62
		}
	case 132:
		{ /* '132' */
			return 63
		}
	case 133:
		{ /* '133' */
			return 110
		}
	case 134:
		{ /* '134' */
			return 111
		}
	case 135:
		{ /* '135' */
			return 110
		}
	case 136:
		{ /* '136' */
			return 111
		}
	case 137:
		{ /* '137' */
			return 101
		}
	case 138:
		{ /* '138' */
			return 102
		}
	case 139:
		{ /* '139' */
			return 103
		}
	case 14:
		{ /* '14' */
			return 14
		}
	case 140:
		{ /* '140' */
			return 104
		}
	case 141:
		{ /* '141' */
			return 105
		}
	case 142:
		{ /* '142' */
			return 106
		}
	case 143:
		{ /* '143' */
			return 107
		}
	case 144:
		{ /* '144' */
			return 108
		}
	case 145:
		{ /* '145' */
			return 109
		}
	case 146:
		{ /* '146' */
			return 110
		}
	case 147:
		{ /* '147' */
			return 111
		}
	case 148:
		{ /* '148' */
			return 112
		}
	case 149:
		{ /* '149' */
			return 113
		}
	case 15:
		{ /* '15' */
			return 15
		}
	case 150:
		{ /* '150' */
			return 114
		}
	case 151:
		{ /* '151' */
			return 115
		}
	case 152:
		{ /* '152' */
			return 116
		}
	case 153:
		{ /* '153' */
			return 117
		}
	case 154:
		{ /* '154' */
			return 118
		}
	case 155:
		{ /* '155' */
			return 119
		}
	case 156:
		{ /* '156' */
			return 120
		}
	case 157:
		{ /* '157' */
			return 101
		}
	case 158:
		{ /* '158' */
			return 102
		}
	case 159:
		{ /* '159' */
			return 103
		}
	case 16:
		{ /* '16' */
			return 16
		}
	case 160:
		{ /* '160' */
			return 104
		}
	case 161:
		{ /* '161' */
			return 105
		}
	case 162:
		{ /* '162' */
			return 106
		}
	case 163:
		{ /* '163' */
			return 107
		}
	case 164:
		{ /* '164' */
			return 108
		}
	case 165:
		{ /* '165' */
			return 109
		}
	case 166:
		{ /* '166' */
			return 110
		}
	case 167:
		{ /* '167' */
			return 111
		}
	case 168:
		{ /* '168' */
			return 112
		}
	case 169:
		{ /* '169' */
			return 113
		}
	case 17:
		{ /* '17' */
			return 17
		}
	case 170:
		{ /* '170' */
			return 114
		}
	case 171:
		{ /* '171' */
			return 115
		}
	case 172:
		{ /* '172' */
			return 116
		}
	case 173:
		{ /* '173' */
			return 117
		}
	case 174:
		{ /* '174' */
			return 118
		}
	case 175:
		{ /* '175' */
			return 119
		}
	case 176:
		{ /* '176' */
			return 120
		}
	case 177:
		{ /* '177' */
			return 121
		}
	case 178:
		{ /* '178' */
			return 122
		}
	case 179:
		{ /* '179' */
			return 123
		}
	case 18:
		{ /* '18' */
			return 18
		}
	case 180:
		{ /* '180' */
			return 124
		}
	case 181:
		{ /* '181' */
			return 125
		}
	case 182:
		{ /* '182' */
			return 126
		}
	case 183:
		{ /* '183' */
			return 127
		}
	case 184:
		{ /* '184' */
			return 128
		}
	case 185:
		{ /* '185' */
			return 129
		}
	case 186:
		{ /* '186' */
			return 130
		}
	case 187:
		{ /* '187' */
			return 131
		}
	case 188:
		{ /* '188' */
			return 132
		}
	case 189:
		{ /* '189' */
			return 133
		}
	case 19:
		{ /* '19' */
			return 19
		}
	case 190:
		{ /* '190' */
			return 134
		}
	case 191:
		{ /* '191' */
			return 51
		}
	case 192:
		{ /* '192' */
			return 52
		}
	case 193:
		{ /* '193' */
			return 53
		}
	case 194:
		{ /* '194' */
			return 51
		}
	case 195:
		{ /* '195' */
			return 52
		}
	case 196:
		{ /* '196' */
			return 51
		}
	case 197:
		{ /* '197' */
			return 52
		}
	case 198:
		{ /* '198' */
			return 53
		}
	case 199:
		{ /* '199' */
			return 54
		}
	case 2:
		{ /* '2' */
			return 2
		}
	case 20:
		{ /* '20' */
			return 20
		}
	case 200:
		{ /* '200' */
			return 55
		}
	case 201:
		{ /* '201' */
			return 57
		}
	case 202:
		{ /* '202' */
			return 58
		}
	case 203:
		{ /* '203' */
			return 60
		}
	case 204:
		{ /* '204' */
			return 61
		}
	case 205:
		{ /* '205' */
			return 62
		}
	case 206:
		{ /* '206' */
			return 63
		}
	case 207:
		{ /* '207' */
			return 64
		}
	case 208:
		{ /* '208' */
			return 65
		}
	case 209:
		{ /* '209' */
			return 66
		}
	case 21:
		{ /* '21' */
			return 21
		}
	case 210:
		{ /* '210' */
			return 67
		}
	case 211:
		{ /* '211' */
			return 68
		}
	case 212:
		{ /* '212' */
			return 69
		}
	case 213:
		{ /* '213' */
			return 51
		}
	case 214:
		{ /* '214' */
			return 52
		}
	case 215:
		{ /* '215' */
			return 53
		}
	case 22:
		{ /* '22' */
			return 22
		}
	case 23:
		{ /* '23' */
			return 23
		}
	case 24:
		{ /* '24' */
			return 24
		}
	case 25:
		{ /* '25' */
			return 25
		}
	case 26:
		{ /* '26' */
			return 26
		}
	case 27:
		{ /* '27' */
			return 27
		}
	case 28:
		{ /* '28' */
			return 28
		}
	case 29:
		{ /* '29' */
			return 29
		}
	case 3:
		{ /* '3' */
			return 3
		}
	case 30:
		{ /* '30' */
			return 30
		}
	case 31:
		{ /* '31' */
			return 51
		}
	case 32:
		{ /* '32' */
			return 52
		}
	case 33:
		{ /* '33' */
			return 53
		}
	case 34:
		{ /* '34' */
			return 54
		}
	case 35:
		{ /* '35' */
			return 55
		}
	case 36:
		{ /* '36' */
			return 56
		}
	case 37:
		{ /* '37' */
			return 57
		}
	case 38:
		{ /* '38' */
			return 58
		}
	case 39:
		{ /* '39' */
			return 59
		}
	case 4:
		{ /* '4' */
			return 4
		}
	case 40:
		{ /* '40' */
			return 60
		}
	case 41:
		{ /* '41' */
			return 61
		}
	case 42:
		{ /* '42' */
			return 62
		}
	case 43:
		{ /* '43' */
			return 63
		}
	case 44:
		{ /* '44' */
			return 64
		}
	case 45:
		{ /* '45' */
			return 65
		}
	case 46:
		{ /* '46' */
			return 66
		}
	case 47:
		{ /* '47' */
			return 67
		}
	case 48:
		{ /* '48' */
			return 68
		}
	case 49:
		{ /* '49' */
			return 69
		}
	case 5:
		{ /* '5' */
			return 5
		}
	case 50:
		{ /* '50' */
			return 70
		}
	case 51:
		{ /* '51' */
			return 71
		}
	case 52:
		{ /* '52' */
			return 72
		}
	case 53:
		{ /* '53' */
			return 73
		}
	case 54:
		{ /* '54' */
			return 74
		}
	case 55:
		{ /* '55' */
			return 75
		}
	case 56:
		{ /* '56' */
			return 76
		}
	case 57:
		{ /* '57' */
			return 77
		}
	case 58:
		{ /* '58' */
			return 78
		}
	case 59:
		{ /* '59' */
			return 79
		}
	case 6:
		{ /* '6' */
			return 6
		}
	case 60:
		{ /* '60' */
			return 80
		}
	case 61:
		{ /* '61' */
			return 81
		}
	case 62:
		{ /* '62' */
			return 82
		}
	case 63:
		{ /* '63' */
			return 83
		}
	case 64:
		{ /* '64' */
			return 84
		}
	case 65:
		{ /* '65' */
			return 85
		}
	case 66:
		{ /* '66' */
			return 86
		}
	case 67:
		{ /* '67' */
			return 51
		}
	case 68:
		{ /* '68' */
			return 52
		}
	case 69:
		{ /* '69' */
			return 51
		}
	case 7:
		{ /* '7' */
			return 7
		}
	case 70:
		{ /* '70' */
			return 52
		}
	case 71:
		{ /* '71' */
			return 53
		}
	case 72:
		{ /* '72' */
			return 54
		}
	case 73:
		{ /* '73' */
			return 55
		}
	case 74:
		{ /* '74' */
			return 56
		}
	case 75:
		{ /* '75' */
			return 57
		}
	case 76:
		{ /* '76' */
			return 58
		}
	case 77:
		{ /* '77' */
			return 63
		}
	case 78:
		{ /* '78' */
			return 67
		}
	case 79:
		{ /* '79' */
			return 112
		}
	case 8:
		{ /* '8' */
			return 8
		}
	case 80:
		{ /* '80' */
			return 51
		}
	case 81:
		{ /* '81' */
			return 52
		}
	case 82:
		{ /* '82' */
			return 53
		}
	case 83:
		{ /* '83' */
			return 54
		}
	case 84:
		{ /* '84' */
			return 55
		}
	case 85:
		{ /* '85' */
			return 56
		}
	case 86:
		{ /* '86' */
			return 57
		}
	case 87:
		{ /* '87' */
			return 58
		}
	case 88:
		{ /* '88' */
			return 59
		}
	case 89:
		{ /* '89' */
			return 60
		}
	case 9:
		{ /* '9' */
			return 9
		}
	case 90:
		{ /* '90' */
			return 61
		}
	case 91:
		{ /* '91' */
			return 62
		}
	case 92:
		{ /* '92' */
			return 63
		}
	case 93:
		{ /* '93' */
			return 64
		}
	case 94:
		{ /* '94' */
			return 65
		}
	case 95:
		{ /* '95' */
			return 66
		}
	case 96:
		{ /* '96' */
			return 67
		}
	case 97:
		{ /* '97' */
			return 68
		}
	case 98:
		{ /* '98' */
			return 69
		}
	case 99:
		{ /* '99' */
			return 70
		}
	default:
		{
			return 0
		}
	}
}

func (e KnxInterfaceObjectProperty) ObjectType() KnxInterfaceObjectType {
	switch e {
	case 0:
		{ /* '0' */
			return KnxInterfaceObjectType_OT_UNKNOWN
		}
	case 1:
		{ /* '1' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 10:
		{ /* '10' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 100:
		{ /* '100' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 101:
		{ /* '101' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 102:
		{ /* '102' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 103:
		{ /* '103' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 104:
		{ /* '104' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 105:
		{ /* '105' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 106:
		{ /* '106' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 107:
		{ /* '107' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 108:
		{ /* '108' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 109:
		{ /* '109' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 11:
		{ /* '11' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 110:
		{ /* '110' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 111:
		{ /* '111' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 112:
		{ /* '112' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 113:
		{ /* '113' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 114:
		{ /* '114' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 115:
		{ /* '115' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 116:
		{ /* '116' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 117:
		{ /* '117' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 118:
		{ /* '118' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 119:
		{ /* '119' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 12:
		{ /* '12' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 120:
		{ /* '120' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 121:
		{ /* '121' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 122:
		{ /* '122' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 123:
		{ /* '123' */
			return KnxInterfaceObjectType_OT_SECURITY
		}
	case 124:
		{ /* '124' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 125:
		{ /* '125' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 126:
		{ /* '126' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 127:
		{ /* '127' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 128:
		{ /* '128' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 129:
		{ /* '129' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 13:
		{ /* '13' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 130:
		{ /* '130' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 131:
		{ /* '131' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 132:
		{ /* '132' */
			return KnxInterfaceObjectType_OT_RF_MEDIUM
		}
	case 133:
		{ /* '133' */
			return KnxInterfaceObjectType_OT_INDOOR_BRIGHTNESS_SENSOR
		}
	case 134:
		{ /* '134' */
			return KnxInterfaceObjectType_OT_INDOOR_BRIGHTNESS_SENSOR
		}
	case 135:
		{ /* '135' */
			return KnxInterfaceObjectType_OT_INDOOR_LUMINANCE_SENSOR
		}
	case 136:
		{ /* '136' */
			return KnxInterfaceObjectType_OT_INDOOR_LUMINANCE_SENSOR
		}
	case 137:
		{ /* '137' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 138:
		{ /* '138' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 139:
		{ /* '139' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 14:
		{ /* '14' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 140:
		{ /* '140' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 141:
		{ /* '141' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 142:
		{ /* '142' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 143:
		{ /* '143' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 144:
		{ /* '144' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 145:
		{ /* '145' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 146:
		{ /* '146' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 147:
		{ /* '147' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 148:
		{ /* '148' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 149:
		{ /* '149' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 15:
		{ /* '15' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 150:
		{ /* '150' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 151:
		{ /* '151' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 152:
		{ /* '152' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 153:
		{ /* '153' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 154:
		{ /* '154' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 155:
		{ /* '155' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 156:
		{ /* '156' */
			return KnxInterfaceObjectType_OT_LIGHT_SWITCHING_ACTUATOR_BASIC
		}
	case 157:
		{ /* '157' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 158:
		{ /* '158' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 159:
		{ /* '159' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 16:
		{ /* '16' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 160:
		{ /* '160' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 161:
		{ /* '161' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 162:
		{ /* '162' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 163:
		{ /* '163' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 164:
		{ /* '164' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 165:
		{ /* '165' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 166:
		{ /* '166' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 167:
		{ /* '167' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 168:
		{ /* '168' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 169:
		{ /* '169' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 17:
		{ /* '17' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 170:
		{ /* '170' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 171:
		{ /* '171' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 172:
		{ /* '172' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 173:
		{ /* '173' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 174:
		{ /* '174' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 175:
		{ /* '175' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 176:
		{ /* '176' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 177:
		{ /* '177' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 178:
		{ /* '178' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 179:
		{ /* '179' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 18:
		{ /* '18' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 180:
		{ /* '180' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 181:
		{ /* '181' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 182:
		{ /* '182' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 183:
		{ /* '183' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 184:
		{ /* '184' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 185:
		{ /* '185' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 186:
		{ /* '186' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 187:
		{ /* '187' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 188:
		{ /* '188' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 189:
		{ /* '189' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 19:
		{ /* '19' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 190:
		{ /* '190' */
			return KnxInterfaceObjectType_OT_DIMMING_ACTUATOR_BASIC
		}
	case 191:
		{ /* '191' */
			return KnxInterfaceObjectType_OT_DIMMING_SENSOR_BASIC
		}
	case 192:
		{ /* '192' */
			return KnxInterfaceObjectType_OT_DIMMING_SENSOR_BASIC
		}
	case 193:
		{ /* '193' */
			return KnxInterfaceObjectType_OT_DIMMING_SENSOR_BASIC
		}
	case 194:
		{ /* '194' */
			return KnxInterfaceObjectType_OT_SWITCHING_SENSOR_BASIC
		}
	case 195:
		{ /* '195' */
			return KnxInterfaceObjectType_OT_SWITCHING_SENSOR_BASIC
		}
	case 196:
		{ /* '196' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 197:
		{ /* '197' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 198:
		{ /* '198' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 199:
		{ /* '199' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 2:
		{ /* '2' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 20:
		{ /* '20' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 200:
		{ /* '200' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 201:
		{ /* '201' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 202:
		{ /* '202' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 203:
		{ /* '203' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 204:
		{ /* '204' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 205:
		{ /* '205' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 206:
		{ /* '206' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 207:
		{ /* '207' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 208:
		{ /* '208' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 209:
		{ /* '209' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 21:
		{ /* '21' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 210:
		{ /* '210' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 211:
		{ /* '211' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 212:
		{ /* '212' */
			return KnxInterfaceObjectType_OT_SUNBLIND_ACTUATOR_BASIC
		}
	case 213:
		{ /* '213' */
			return KnxInterfaceObjectType_OT_SUNBLIND_SENSOR_BASIC
		}
	case 214:
		{ /* '214' */
			return KnxInterfaceObjectType_OT_SUNBLIND_SENSOR_BASIC
		}
	case 215:
		{ /* '215' */
			return KnxInterfaceObjectType_OT_SUNBLIND_SENSOR_BASIC
		}
	case 22:
		{ /* '22' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 23:
		{ /* '23' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 24:
		{ /* '24' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 25:
		{ /* '25' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 26:
		{ /* '26' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 27:
		{ /* '27' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 28:
		{ /* '28' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 29:
		{ /* '29' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 3:
		{ /* '3' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 30:
		{ /* '30' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 31:
		{ /* '31' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 32:
		{ /* '32' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 33:
		{ /* '33' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 34:
		{ /* '34' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 35:
		{ /* '35' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 36:
		{ /* '36' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 37:
		{ /* '37' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 38:
		{ /* '38' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 39:
		{ /* '39' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 4:
		{ /* '4' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 40:
		{ /* '40' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 41:
		{ /* '41' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 42:
		{ /* '42' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 43:
		{ /* '43' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 44:
		{ /* '44' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 45:
		{ /* '45' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 46:
		{ /* '46' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 47:
		{ /* '47' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 48:
		{ /* '48' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 49:
		{ /* '49' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 5:
		{ /* '5' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 50:
		{ /* '50' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 51:
		{ /* '51' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 52:
		{ /* '52' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 53:
		{ /* '53' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 54:
		{ /* '54' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 55:
		{ /* '55' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 56:
		{ /* '56' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 57:
		{ /* '57' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 58:
		{ /* '58' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 59:
		{ /* '59' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 6:
		{ /* '6' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 60:
		{ /* '60' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 61:
		{ /* '61' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 62:
		{ /* '62' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 63:
		{ /* '63' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 64:
		{ /* '64' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 65:
		{ /* '65' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 66:
		{ /* '66' */
			return KnxInterfaceObjectType_OT_DEVICE
		}
	case 67:
		{ /* '67' */
			return KnxInterfaceObjectType_OT_GROUP_OBJECT_TABLE
		}
	case 68:
		{ /* '68' */
			return KnxInterfaceObjectType_OT_GROUP_OBJECT_TABLE
		}
	case 69:
		{ /* '69' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 7:
		{ /* '7' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 70:
		{ /* '70' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 71:
		{ /* '71' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 72:
		{ /* '72' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 73:
		{ /* '73' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 74:
		{ /* '74' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 75:
		{ /* '75' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 76:
		{ /* '76' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 77:
		{ /* '77' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 78:
		{ /* '78' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 79:
		{ /* '79' */
			return KnxInterfaceObjectType_OT_ROUTER
		}
	case 8:
		{ /* '8' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 80:
		{ /* '80' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 81:
		{ /* '81' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 82:
		{ /* '82' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 83:
		{ /* '83' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 84:
		{ /* '84' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 85:
		{ /* '85' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 86:
		{ /* '86' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 87:
		{ /* '87' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 88:
		{ /* '88' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 89:
		{ /* '89' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 9:
		{ /* '9' */
			return KnxInterfaceObjectType_OT_GENERAL
		}
	case 90:
		{ /* '90' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 91:
		{ /* '91' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 92:
		{ /* '92' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 93:
		{ /* '93' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 94:
		{ /* '94' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 95:
		{ /* '95' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 96:
		{ /* '96' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 97:
		{ /* '97' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 98:
		{ /* '98' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	case 99:
		{ /* '99' */
			return KnxInterfaceObjectType_OT_KNXIP_PARAMETER
		}
	default:
		{
			return 0
		}
	}
}
func KnxInterfaceObjectPropertyByValue(value uint32) KnxInterfaceObjectProperty {
	switch value {
	case 0:
		return KnxInterfaceObjectProperty_PID_UNKNOWN
	case 1:
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_TYPE
	case 10:
		return KnxInterfaceObjectProperty_PID_GENERAL_SERVICES_SUPPORTED
	case 100:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED
	case 101:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP
	case 102:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX
	case 103:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP
	case 104:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX
	case 105:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_FRIENDLY_NAME
	case 106:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_BACKBONE_KEY
	case 107:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE
	case 108:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PASSWORD_HASHES
	case 109:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES
	case 11:
		return KnxInterfaceObjectProperty_PID_GENERAL_SERIAL_NUMBER
	case 110:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE
	case 111:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION
	case 112:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TUNNELLING_USERS
	case 113:
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_MODE
	case 114:
		return KnxInterfaceObjectProperty_PID_SECURITY_P2P_KEY_TABLE
	case 115:
		return KnxInterfaceObjectProperty_PID_SECURITY_GRP_KEY_TABLE
	case 116:
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE
	case 117:
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_FAILURES_LOG
	case 118:
		return KnxInterfaceObjectProperty_PID_SECURITY_SKI_TOOL
	case 119:
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT
	case 12:
		return KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_ID
	case 120:
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT_CONTROL
	case 121:
		return KnxInterfaceObjectProperty_PID_SECURITY_SEQUENCE_NUMBER_SENDING
	case 122:
		return KnxInterfaceObjectProperty_PID_SECURITY_ZONE_KEYS_TABLE
	case 123:
		return KnxInterfaceObjectProperty_PID_SECURITY_GO_SECURITY_FLAGS
	case 124:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_MULTI_TYPE
	case 125:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DOMAIN_ADDRESS
	case 126:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_RETRANSMITTER
	case 127:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_SECURITY_REPORT_CONTROL
	case 128:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT
	case 129:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_BIDIR_TIMEOUT
	case 13:
		return KnxInterfaceObjectProperty_PID_GENERAL_PROGRAM_VERSION
	case 130:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE
	case 131:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE
	case 132:
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_PROBE
	case 133:
		return KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE
	case 134:
		return KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME
	case 135:
		return KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE
	case 136:
		return KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME
	case 137:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY
	case 138:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY
	case 139:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION
	case 14:
		return KnxInterfaceObjectProperty_PID_GENERAL_DEVICE_CONTROL
	case 140:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION
	case 141:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME
	case 142:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY
	case 143:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING
	case 144:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING
	case 145:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP
	case 146:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN
	case 147:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE
	case 148:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION
	case 149:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE
	case 15:
		return KnxInterfaceObjectProperty_PID_GENERAL_ORDER_INFO
	case 150:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE
	case 151:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE
	case 152:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE
	case 153:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER
	case 154:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE
	case 155:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE
	case 156:
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2
	case 157:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_ON_DELAY
	case 158:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY
	case 159:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME
	case 16:
		return KnxInterfaceObjectProperty_PID_GENERAL_PEI_TYPE
	case 160:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION
	case 161:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION
	case 162:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME
	case 163:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY
	case 164:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED
	case 165:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME
	case 166:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE
	case 167:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF
	case 168:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE
	case 169:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF
	case 17:
		return KnxInterfaceObjectProperty_PID_GENERAL_PORT_CONFIGURATION
	case 170:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS
	case 171:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE
	case 172:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE
	case 173:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE
	case 174:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION
	case 175:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE
	case 176:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION
	case 177:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION
	case 178:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE
	case 179:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE
	case 18:
		return KnxInterfaceObjectProperty_PID_GENERAL_POLL_GROUP_SETTINGS
	case 180:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING
	case 181:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING
	case 182:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE
	case 183:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE
	case 184:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE
	case 185:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE
	case 186:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE
	case 187:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP
	case 188:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE
	case 189:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN
	case 19:
		return KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_DATA
	case 190:
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE
	case 191:
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION
	case 192:
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case 193:
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE
	case 194:
		return KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION
	case 195:
		return KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case 196:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME
	case 197:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME
	case 198:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME
	case 199:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME
	case 2:
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_NAME
	case 20:
		return KnxInterfaceObjectProperty_PID_GENERAL_ENABLE
	case 200:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT
	case 201:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH
	case 202:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT
	case 203:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE
	case 204:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM
	case 205:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM
	case 206:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM
	case 207:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM
	case 208:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM
	case 209:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM
	case 21:
		return KnxInterfaceObjectProperty_PID_GENERAL_DESCRIPTION
	case 210:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME
	case 211:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE
	case 212:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE
	case 213:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE
	case 214:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION
	case 215:
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case 22:
		return KnxInterfaceObjectProperty_PID_GENERAL_FILE
	case 23:
		return KnxInterfaceObjectProperty_PID_GENERAL_TABLE
	case 24:
		return KnxInterfaceObjectProperty_PID_GENERAL_ENROL
	case 25:
		return KnxInterfaceObjectProperty_PID_GENERAL_VERSION
	case 26:
		return KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_LINK
	case 27:
		return KnxInterfaceObjectProperty_PID_GENERAL_MCB_TABLE
	case 28:
		return KnxInterfaceObjectProperty_PID_GENERAL_ERROR_CODE
	case 29:
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_INDEX
	case 3:
		return KnxInterfaceObjectProperty_PID_GENERAL_SEMAPHOR
	case 30:
		return KnxInterfaceObjectProperty_PID_GENERAL_DOWNLOAD_COUNTER
	case 31:
		return KnxInterfaceObjectProperty_PID_DEVICE_ROUTING_COUNT
	case 32:
		return KnxInterfaceObjectProperty_PID_DEVICE_MAX_RETRY_COUNT
	case 33:
		return KnxInterfaceObjectProperty_PID_DEVICE_ERROR_FLAGS
	case 34:
		return KnxInterfaceObjectProperty_PID_DEVICE_PROGMODE
	case 35:
		return KnxInterfaceObjectProperty_PID_DEVICE_PRODUCT_ID
	case 36:
		return KnxInterfaceObjectProperty_PID_DEVICE_MAX_APDULENGTH
	case 37:
		return KnxInterfaceObjectProperty_PID_DEVICE_SUBNET_ADDR
	case 38:
		return KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_ADDR
	case 39:
		return KnxInterfaceObjectProperty_PID_DEVICE_PB_CONFIG
	case 4:
		return KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_REFERENCE
	case 40:
		return KnxInterfaceObjectProperty_PID_DEVICE_ADDR_REPORT
	case 41:
		return KnxInterfaceObjectProperty_PID_DEVICE_ADDR_CHECK
	case 42:
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECT_VALUE
	case 43:
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECTLINK
	case 44:
		return KnxInterfaceObjectProperty_PID_DEVICE_APPLICATION
	case 45:
		return KnxInterfaceObjectProperty_PID_DEVICE_PARAMETER
	case 46:
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECTADDRESS
	case 47:
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_TYPE
	case 48:
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_STATUS
	case 49:
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_ENABLE
	case 5:
		return KnxInterfaceObjectProperty_PID_GENERAL_LOAD_STATE_CONTROL
	case 50:
		return KnxInterfaceObjectProperty_PID_DEVICE_DOMAIN_ADDRESS
	case 51:
		return KnxInterfaceObjectProperty_PID_DEVICE_IO_LIST
	case 52:
		return KnxInterfaceObjectProperty_PID_DEVICE_MGT_DESCRIPTOR_01
	case 53:
		return KnxInterfaceObjectProperty_PID_DEVICE_PL110_PARAM
	case 54:
		return KnxInterfaceObjectProperty_PID_DEVICE_RF_REPEAT_COUNTER
	case 55:
		return KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_TABLE
	case 56:
		return KnxInterfaceObjectProperty_PID_DEVICE_RANDOM_PAUSE_TABLE
	case 57:
		return KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_NR
	case 58:
		return KnxInterfaceObjectProperty_PID_DEVICE_HARDWARE_TYPE
	case 59:
		return KnxInterfaceObjectProperty_PID_DEVICE_RETRANSMITTER_NUMBER
	case 6:
		return KnxInterfaceObjectProperty_PID_GENERAL_RUN_STATE_CONTROL
	case 60:
		return KnxInterfaceObjectProperty_PID_DEVICE_SERIAL_NR_TABLE
	case 61:
		return KnxInterfaceObjectProperty_PID_DEVICE_BIBATMASTER_ADDRESS
	case 62:
		return KnxInterfaceObjectProperty_PID_DEVICE_RF_DOMAIN_ADDRESS
	case 63:
		return KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_DESCRIPTOR
	case 64:
		return KnxInterfaceObjectProperty_PID_DEVICE_METERING_FILTER_TABLE
	case 65:
		return KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE
	case 66:
		return KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR
	case 67:
		return KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_GRPOBJTABLE
	case 68:
		return KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE
	case 69:
		return KnxInterfaceObjectProperty_PID_ROUTER_LINE_STATUS
	case 7:
		return KnxInterfaceObjectProperty_PID_GENERAL_TABLE_REFERENCE
	case 70:
		return KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCCONFIG
	case 71:
		return KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCCONFIG
	case 72:
		return KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCGRPCONFIG
	case 73:
		return KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCGRPCONFIG
	case 74:
		return KnxInterfaceObjectProperty_PID_ROUTER_ROUTETABLE_CONTROL
	case 75:
		return KnxInterfaceObjectProperty_PID_ROUTER_COUPL_SERV_CONTROL
	case 76:
		return KnxInterfaceObjectProperty_PID_ROUTER_MAX_ROUTER_APDU_LENGTH
	case 77:
		return KnxInterfaceObjectProperty_PID_ROUTER_MEDIUM
	case 78:
		return KnxInterfaceObjectProperty_PID_ROUTER_FILTER_TABLE_USE
	case 79:
		return KnxInterfaceObjectProperty_PID_ROUTER_RF_ENABLE_SBC
	case 8:
		return KnxInterfaceObjectProperty_PID_GENERAL_SERVICE_CONTROL
	case 80:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID
	case 81:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS
	case 82:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES
	case 83:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD
	case 84:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD
	case 85:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_CAPABILITIES
	case 86:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS
	case 87:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK
	case 88:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY
	case 89:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ADDRESS
	case 9:
		return KnxInterfaceObjectProperty_PID_GENERAL_FIRMWARE_REVISION
	case 90:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SUBNET_MASK
	case 91:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEFAULT_GATEWAY
	case 92:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER
	case 93:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MAC_ADDRESS
	case 94:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS
	case 95:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS
	case 96:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TTL
	case 97:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES
	case 98:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE
	case 99:
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES
	}
	return 0
}

func KnxInterfaceObjectPropertyByName(value string) KnxInterfaceObjectProperty {
	switch value {
	case "PID_UNKNOWN":
		return KnxInterfaceObjectProperty_PID_UNKNOWN
	case "PID_GENERAL_OBJECT_TYPE":
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_TYPE
	case "PID_GENERAL_SERVICES_SUPPORTED":
		return KnxInterfaceObjectProperty_PID_GENERAL_SERVICES_SUPPORTED
	case "PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED
	case "PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP
	case "PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX
	case "PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP
	case "PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX
	case "PID_KNXIP_PARAMETER_FRIENDLY_NAME":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_FRIENDLY_NAME
	case "PID_KNXIP_PARAMETER_BACKBONE_KEY":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_BACKBONE_KEY
	case "PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE
	case "PID_KNXIP_PARAMETER_PASSWORD_HASHES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PASSWORD_HASHES
	case "PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES
	case "PID_GENERAL_SERIAL_NUMBER":
		return KnxInterfaceObjectProperty_PID_GENERAL_SERIAL_NUMBER
	case "PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE
	case "PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION
	case "PID_KNXIP_PARAMETER_TUNNELLING_USERS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TUNNELLING_USERS
	case "PID_SECURITY_SECURITY_MODE":
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_MODE
	case "PID_SECURITY_P2P_KEY_TABLE":
		return KnxInterfaceObjectProperty_PID_SECURITY_P2P_KEY_TABLE
	case "PID_SECURITY_GRP_KEY_TABLE":
		return KnxInterfaceObjectProperty_PID_SECURITY_GRP_KEY_TABLE
	case "PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE":
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE
	case "PID_SECURITY_SECURITY_FAILURES_LOG":
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_FAILURES_LOG
	case "PID_SECURITY_SKI_TOOL":
		return KnxInterfaceObjectProperty_PID_SECURITY_SKI_TOOL
	case "PID_SECURITY_SECURITY_REPORT":
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT
	case "PID_GENERAL_MANUFACTURER_ID":
		return KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_ID
	case "PID_SECURITY_SECURITY_REPORT_CONTROL":
		return KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT_CONTROL
	case "PID_SECURITY_SEQUENCE_NUMBER_SENDING":
		return KnxInterfaceObjectProperty_PID_SECURITY_SEQUENCE_NUMBER_SENDING
	case "PID_SECURITY_ZONE_KEYS_TABLE":
		return KnxInterfaceObjectProperty_PID_SECURITY_ZONE_KEYS_TABLE
	case "PID_SECURITY_GO_SECURITY_FLAGS":
		return KnxInterfaceObjectProperty_PID_SECURITY_GO_SECURITY_FLAGS
	case "PID_RF_MEDIUM_RF_MULTI_TYPE":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_MULTI_TYPE
	case "PID_RF_MEDIUM_RF_DOMAIN_ADDRESS":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DOMAIN_ADDRESS
	case "PID_RF_MEDIUM_RF_RETRANSMITTER":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_RETRANSMITTER
	case "PID_RF_MEDIUM_SECURITY_REPORT_CONTROL":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_SECURITY_REPORT_CONTROL
	case "PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT
	case "PID_RF_MEDIUM_RF_BIDIR_TIMEOUT":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_BIDIR_TIMEOUT
	case "PID_GENERAL_PROGRAM_VERSION":
		return KnxInterfaceObjectProperty_PID_GENERAL_PROGRAM_VERSION
	case "PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE
	case "PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE
	case "PID_RF_MEDIUM_RF_DIAG_PROBE":
		return KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_PROBE
	case "PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE":
		return KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE
	case "PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME":
		return KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME
	case "PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE":
		return KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE
	case "PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME":
		return KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION
	case "PID_GENERAL_DEVICE_CONTROL":
		return KnxInterfaceObjectProperty_PID_GENERAL_DEVICE_CONTROL
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE
	case "PID_GENERAL_ORDER_INFO":
		return KnxInterfaceObjectProperty_PID_GENERAL_ORDER_INFO
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE
	case "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2":
		return KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2
	case "PID_DIMMING_ACTUATOR_BASIC_ON_DELAY":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_ON_DELAY
	case "PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY
	case "PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME
	case "PID_GENERAL_PEI_TYPE":
		return KnxInterfaceObjectProperty_PID_GENERAL_PEI_TYPE
	case "PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION
	case "PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION
	case "PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME
	case "PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF
	case "PID_GENERAL_PORT_CONFIGURATION":
		return KnxInterfaceObjectProperty_PID_GENERAL_PORT_CONFIGURATION
	case "PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS
	case "PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION
	case "PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE
	case "PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION
	case "PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION
	case "PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE
	case "PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE
	case "PID_GENERAL_POLL_GROUP_SETTINGS":
		return KnxInterfaceObjectProperty_PID_GENERAL_POLL_GROUP_SETTINGS
	case "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING
	case "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING
	case "PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE
	case "PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE
	case "PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE
	case "PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE
	case "PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP
	case "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE
	case "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN
	case "PID_GENERAL_MANUFACTURER_DATA":
		return KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_DATA
	case "PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE
	case "PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION":
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION
	case "PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE":
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case "PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE":
		return KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE
	case "PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION":
		return KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION
	case "PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE":
		return KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case "PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME
	case "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME
	case "PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME
	case "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME
	case "PID_GENERAL_OBJECT_NAME":
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_NAME
	case "PID_GENERAL_ENABLE":
		return KnxInterfaceObjectProperty_PID_GENERAL_ENABLE
	case "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT
	case "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH
	case "PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT
	case "PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE
	case "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM
	case "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM
	case "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM
	case "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM
	case "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM
	case "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM
	case "PID_GENERAL_DESCRIPTION":
		return KnxInterfaceObjectProperty_PID_GENERAL_DESCRIPTION
	case "PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME
	case "PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE
	case "PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE
	case "PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE
	case "PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION
	case "PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE":
		return KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE
	case "PID_GENERAL_FILE":
		return KnxInterfaceObjectProperty_PID_GENERAL_FILE
	case "PID_GENERAL_TABLE":
		return KnxInterfaceObjectProperty_PID_GENERAL_TABLE
	case "PID_GENERAL_ENROL":
		return KnxInterfaceObjectProperty_PID_GENERAL_ENROL
	case "PID_GENERAL_VERSION":
		return KnxInterfaceObjectProperty_PID_GENERAL_VERSION
	case "PID_GENERAL_GROUP_OBJECT_LINK":
		return KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_LINK
	case "PID_GENERAL_MCB_TABLE":
		return KnxInterfaceObjectProperty_PID_GENERAL_MCB_TABLE
	case "PID_GENERAL_ERROR_CODE":
		return KnxInterfaceObjectProperty_PID_GENERAL_ERROR_CODE
	case "PID_GENERAL_OBJECT_INDEX":
		return KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_INDEX
	case "PID_GENERAL_SEMAPHOR":
		return KnxInterfaceObjectProperty_PID_GENERAL_SEMAPHOR
	case "PID_GENERAL_DOWNLOAD_COUNTER":
		return KnxInterfaceObjectProperty_PID_GENERAL_DOWNLOAD_COUNTER
	case "PID_DEVICE_ROUTING_COUNT":
		return KnxInterfaceObjectProperty_PID_DEVICE_ROUTING_COUNT
	case "PID_DEVICE_MAX_RETRY_COUNT":
		return KnxInterfaceObjectProperty_PID_DEVICE_MAX_RETRY_COUNT
	case "PID_DEVICE_ERROR_FLAGS":
		return KnxInterfaceObjectProperty_PID_DEVICE_ERROR_FLAGS
	case "PID_DEVICE_PROGMODE":
		return KnxInterfaceObjectProperty_PID_DEVICE_PROGMODE
	case "PID_DEVICE_PRODUCT_ID":
		return KnxInterfaceObjectProperty_PID_DEVICE_PRODUCT_ID
	case "PID_DEVICE_MAX_APDULENGTH":
		return KnxInterfaceObjectProperty_PID_DEVICE_MAX_APDULENGTH
	case "PID_DEVICE_SUBNET_ADDR":
		return KnxInterfaceObjectProperty_PID_DEVICE_SUBNET_ADDR
	case "PID_DEVICE_DEVICE_ADDR":
		return KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_ADDR
	case "PID_DEVICE_PB_CONFIG":
		return KnxInterfaceObjectProperty_PID_DEVICE_PB_CONFIG
	case "PID_GENERAL_GROUP_OBJECT_REFERENCE":
		return KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_REFERENCE
	case "PID_DEVICE_ADDR_REPORT":
		return KnxInterfaceObjectProperty_PID_DEVICE_ADDR_REPORT
	case "PID_DEVICE_ADDR_CHECK":
		return KnxInterfaceObjectProperty_PID_DEVICE_ADDR_CHECK
	case "PID_DEVICE_OBJECT_VALUE":
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECT_VALUE
	case "PID_DEVICE_OBJECTLINK":
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECTLINK
	case "PID_DEVICE_APPLICATION":
		return KnxInterfaceObjectProperty_PID_DEVICE_APPLICATION
	case "PID_DEVICE_PARAMETER":
		return KnxInterfaceObjectProperty_PID_DEVICE_PARAMETER
	case "PID_DEVICE_OBJECTADDRESS":
		return KnxInterfaceObjectProperty_PID_DEVICE_OBJECTADDRESS
	case "PID_DEVICE_PSU_TYPE":
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_TYPE
	case "PID_DEVICE_PSU_STATUS":
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_STATUS
	case "PID_DEVICE_PSU_ENABLE":
		return KnxInterfaceObjectProperty_PID_DEVICE_PSU_ENABLE
	case "PID_GENERAL_LOAD_STATE_CONTROL":
		return KnxInterfaceObjectProperty_PID_GENERAL_LOAD_STATE_CONTROL
	case "PID_DEVICE_DOMAIN_ADDRESS":
		return KnxInterfaceObjectProperty_PID_DEVICE_DOMAIN_ADDRESS
	case "PID_DEVICE_IO_LIST":
		return KnxInterfaceObjectProperty_PID_DEVICE_IO_LIST
	case "PID_DEVICE_MGT_DESCRIPTOR_01":
		return KnxInterfaceObjectProperty_PID_DEVICE_MGT_DESCRIPTOR_01
	case "PID_DEVICE_PL110_PARAM":
		return KnxInterfaceObjectProperty_PID_DEVICE_PL110_PARAM
	case "PID_DEVICE_RF_REPEAT_COUNTER":
		return KnxInterfaceObjectProperty_PID_DEVICE_RF_REPEAT_COUNTER
	case "PID_DEVICE_RECEIVE_BLOCK_TABLE":
		return KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_TABLE
	case "PID_DEVICE_RANDOM_PAUSE_TABLE":
		return KnxInterfaceObjectProperty_PID_DEVICE_RANDOM_PAUSE_TABLE
	case "PID_DEVICE_RECEIVE_BLOCK_NR":
		return KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_NR
	case "PID_DEVICE_HARDWARE_TYPE":
		return KnxInterfaceObjectProperty_PID_DEVICE_HARDWARE_TYPE
	case "PID_DEVICE_RETRANSMITTER_NUMBER":
		return KnxInterfaceObjectProperty_PID_DEVICE_RETRANSMITTER_NUMBER
	case "PID_GENERAL_RUN_STATE_CONTROL":
		return KnxInterfaceObjectProperty_PID_GENERAL_RUN_STATE_CONTROL
	case "PID_DEVICE_SERIAL_NR_TABLE":
		return KnxInterfaceObjectProperty_PID_DEVICE_SERIAL_NR_TABLE
	case "PID_DEVICE_BIBATMASTER_ADDRESS":
		return KnxInterfaceObjectProperty_PID_DEVICE_BIBATMASTER_ADDRESS
	case "PID_DEVICE_RF_DOMAIN_ADDRESS":
		return KnxInterfaceObjectProperty_PID_DEVICE_RF_DOMAIN_ADDRESS
	case "PID_DEVICE_DEVICE_DESCRIPTOR":
		return KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_DESCRIPTOR
	case "PID_DEVICE_METERING_FILTER_TABLE":
		return KnxInterfaceObjectProperty_PID_DEVICE_METERING_FILTER_TABLE
	case "PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE":
		return KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE
	case "PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR":
		return KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR
	case "PID_GROUP_OBJECT_TABLE_GRPOBJTABLE":
		return KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_GRPOBJTABLE
	case "PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE":
		return KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE
	case "PID_ROUTER_LINE_STATUS":
		return KnxInterfaceObjectProperty_PID_ROUTER_LINE_STATUS
	case "PID_GENERAL_TABLE_REFERENCE":
		return KnxInterfaceObjectProperty_PID_GENERAL_TABLE_REFERENCE
	case "PID_ROUTER_MAIN_LCCONFIG":
		return KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCCONFIG
	case "PID_ROUTER_SUB_LCCONFIG":
		return KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCCONFIG
	case "PID_ROUTER_MAIN_LCGRPCONFIG":
		return KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCGRPCONFIG
	case "PID_ROUTER_SUB_LCGRPCONFIG":
		return KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCGRPCONFIG
	case "PID_ROUTER_ROUTETABLE_CONTROL":
		return KnxInterfaceObjectProperty_PID_ROUTER_ROUTETABLE_CONTROL
	case "PID_ROUTER_COUPL_SERV_CONTROL":
		return KnxInterfaceObjectProperty_PID_ROUTER_COUPL_SERV_CONTROL
	case "PID_ROUTER_MAX_ROUTER_APDU_LENGTH":
		return KnxInterfaceObjectProperty_PID_ROUTER_MAX_ROUTER_APDU_LENGTH
	case "PID_ROUTER_MEDIUM":
		return KnxInterfaceObjectProperty_PID_ROUTER_MEDIUM
	case "PID_ROUTER_FILTER_TABLE_USE":
		return KnxInterfaceObjectProperty_PID_ROUTER_FILTER_TABLE_USE
	case "PID_ROUTER_RF_ENABLE_SBC":
		return KnxInterfaceObjectProperty_PID_ROUTER_RF_ENABLE_SBC
	case "PID_GENERAL_SERVICE_CONTROL":
		return KnxInterfaceObjectProperty_PID_GENERAL_SERVICE_CONTROL
	case "PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID
	case "PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS
	case "PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES
	case "PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD
	case "PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD
	case "PID_KNXIP_PARAMETER_IP_CAPABILITIES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_CAPABILITIES
	case "PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS
	case "PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK
	case "PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY
	case "PID_KNXIP_PARAMETER_IP_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ADDRESS
	case "PID_GENERAL_FIRMWARE_REVISION":
		return KnxInterfaceObjectProperty_PID_GENERAL_FIRMWARE_REVISION
	case "PID_KNXIP_PARAMETER_SUBNET_MASK":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SUBNET_MASK
	case "PID_KNXIP_PARAMETER_DEFAULT_GATEWAY":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEFAULT_GATEWAY
	case "PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER
	case "PID_KNXIP_PARAMETER_MAC_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MAC_ADDRESS
	case "PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS
	case "PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS
	case "PID_KNXIP_PARAMETER_TTL":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TTL
	case "PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES
	case "PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE
	case "PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES":
		return KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES
	}
	return 0
}

func CastKnxInterfaceObjectProperty(structType interface{}) KnxInterfaceObjectProperty {
	castFunc := func(typ interface{}) KnxInterfaceObjectProperty {
		if sKnxInterfaceObjectProperty, ok := typ.(KnxInterfaceObjectProperty); ok {
			return sKnxInterfaceObjectProperty
		}
		return 0
	}
	return castFunc(structType)
}

func (m KnxInterfaceObjectProperty) LengthInBits() uint16 {
	return 32
}

func (m KnxInterfaceObjectProperty) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxInterfaceObjectPropertyParse(io *utils.ReadBuffer) (KnxInterfaceObjectProperty, error) {
	val, err := io.ReadUint32(32)
	if err != nil {
		return 0, nil
	}
	return KnxInterfaceObjectPropertyByValue(val), nil
}

func (e KnxInterfaceObjectProperty) Serialize(io utils.WriteBuffer) error {
	err := io.WriteUint32(32, uint32(e))
	return err
}

func (e KnxInterfaceObjectProperty) String() string {
	switch e {
	case KnxInterfaceObjectProperty_PID_UNKNOWN:
		return "PID_UNKNOWN"
	case KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_TYPE:
		return "PID_GENERAL_OBJECT_TYPE"
	case KnxInterfaceObjectProperty_PID_GENERAL_SERVICES_SUPPORTED:
		return "PID_GENERAL_SERVICES_SUPPORTED"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED:
		return "PID_KNXIP_PARAMETER_PRIORITY_FIFO_ENABLED"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP:
		return "PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_IP"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX:
		return "PID_KNXIP_PARAMETER_QUEUE_OVERFLOW_TO_KNX"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP:
		return "PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_IP"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX:
		return "PID_KNXIP_PARAMETER_MSG_TRANSMIT_TO_KNX"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_FRIENDLY_NAME:
		return "PID_KNXIP_PARAMETER_FRIENDLY_NAME"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_BACKBONE_KEY:
		return "PID_KNXIP_PARAMETER_BACKBONE_KEY"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE:
		return "PID_KNXIP_PARAMETER_DEVICE_AUTHENTICATION_CODE"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PASSWORD_HASHES:
		return "PID_KNXIP_PARAMETER_PASSWORD_HASHES"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES:
		return "PID_KNXIP_PARAMETER_SECURED_SERVICE_FAMILIES"
	case KnxInterfaceObjectProperty_PID_GENERAL_SERIAL_NUMBER:
		return "PID_GENERAL_SERIAL_NUMBER"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE:
		return "PID_KNXIP_PARAMETER_MULTICAST_LATENCY_TOLERANCE"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION:
		return "PID_KNXIP_PARAMETER_SYNC_LATENCY_FRACTION"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TUNNELLING_USERS:
		return "PID_KNXIP_PARAMETER_TUNNELLING_USERS"
	case KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_MODE:
		return "PID_SECURITY_SECURITY_MODE"
	case KnxInterfaceObjectProperty_PID_SECURITY_P2P_KEY_TABLE:
		return "PID_SECURITY_P2P_KEY_TABLE"
	case KnxInterfaceObjectProperty_PID_SECURITY_GRP_KEY_TABLE:
		return "PID_SECURITY_GRP_KEY_TABLE"
	case KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE:
		return "PID_SECURITY_SECURITY_INDIVIDUAL_ADDRESS_TABLE"
	case KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_FAILURES_LOG:
		return "PID_SECURITY_SECURITY_FAILURES_LOG"
	case KnxInterfaceObjectProperty_PID_SECURITY_SKI_TOOL:
		return "PID_SECURITY_SKI_TOOL"
	case KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT:
		return "PID_SECURITY_SECURITY_REPORT"
	case KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_ID:
		return "PID_GENERAL_MANUFACTURER_ID"
	case KnxInterfaceObjectProperty_PID_SECURITY_SECURITY_REPORT_CONTROL:
		return "PID_SECURITY_SECURITY_REPORT_CONTROL"
	case KnxInterfaceObjectProperty_PID_SECURITY_SEQUENCE_NUMBER_SENDING:
		return "PID_SECURITY_SEQUENCE_NUMBER_SENDING"
	case KnxInterfaceObjectProperty_PID_SECURITY_ZONE_KEYS_TABLE:
		return "PID_SECURITY_ZONE_KEYS_TABLE"
	case KnxInterfaceObjectProperty_PID_SECURITY_GO_SECURITY_FLAGS:
		return "PID_SECURITY_GO_SECURITY_FLAGS"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_MULTI_TYPE:
		return "PID_RF_MEDIUM_RF_MULTI_TYPE"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DOMAIN_ADDRESS:
		return "PID_RF_MEDIUM_RF_DOMAIN_ADDRESS"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_RETRANSMITTER:
		return "PID_RF_MEDIUM_RF_RETRANSMITTER"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_SECURITY_REPORT_CONTROL:
		return "PID_RF_MEDIUM_SECURITY_REPORT_CONTROL"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT:
		return "PID_RF_MEDIUM_RF_FILTERING_MODE_SELECT"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_BIDIR_TIMEOUT:
		return "PID_RF_MEDIUM_RF_BIDIR_TIMEOUT"
	case KnxInterfaceObjectProperty_PID_GENERAL_PROGRAM_VERSION:
		return "PID_GENERAL_PROGRAM_VERSION"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE:
		return "PID_RF_MEDIUM_RF_DIAG_SA_FILTER_TABLE"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE:
		return "PID_RF_MEDIUM_RF_DIAG_QUALITY_TABLE"
	case KnxInterfaceObjectProperty_PID_RF_MEDIUM_RF_DIAG_PROBE:
		return "PID_RF_MEDIUM_RF_DIAG_PROBE"
	case KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE:
		return "PID_INDOOR_BRIGHTNESS_SENSOR_CHANGE_OF_VALUE"
	case KnxInterfaceObjectProperty_PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME:
		return "PID_INDOOR_BRIGHTNESS_SENSOR_REPETITION_TIME"
	case KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE:
		return "PID_INDOOR_LUMINANCE_SENSOR_CHANGE_OF_VALUE"
	case KnxInterfaceObjectProperty_PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME:
		return "PID_INDOOR_LUMINANCE_SENSOR_REPETITION_TIME"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_ON_DELAY"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_OFF_DELAY"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_DURATION"
	case KnxInterfaceObjectProperty_PID_GENERAL_DEVICE_CONTROL:
		return "PID_GENERAL_DEVICE_CONTROL"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_PREWARNING_DURATION"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_OUTPUT_STATE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE"
	case KnxInterfaceObjectProperty_PID_GENERAL_ORDER_INFO:
		return "PID_GENERAL_ORDER_INFO"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_LOCK_STATE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_UNLOCK_STATE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STATE_FOR_SCENE_NUMBER"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BUS_POWER_UP_STATE"
	case KnxInterfaceObjectProperty_PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2:
		return "PID_LIGHT_SWITCHING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_2"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_ON_DELAY:
		return "PID_DIMMING_ACTUATOR_BASIC_ON_DELAY"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY:
		return "PID_DIMMING_ACTUATOR_BASIC_OFF_DELAY"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME:
		return "PID_DIMMING_ACTUATOR_BASIC_SWITCH_OFF_BRIGHTNESS_DELAY_TIME"
	case KnxInterfaceObjectProperty_PID_GENERAL_PEI_TYPE:
		return "PID_GENERAL_PEI_TYPE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION:
		return "PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_DURATION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION:
		return "PID_DIMMING_ACTUATOR_BASIC_PREWARNING_DURATION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME:
		return "PID_DIMMING_ACTUATOR_BASIC_TRANSMISSION_CYCLE_TIME"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY:
		return "PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_UP_MESSAGE_DELAY"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_ON_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_SPEED_FOR_SWITCH_OFF"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_ON_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMMING_STEP_TIME_FOR_SWITCH_OFF"
	case KnxInterfaceObjectProperty_PID_GENERAL_PORT_CONFIGURATION:
		return "PID_GENERAL_PORT_CONFIGURATION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS:
		return "PID_DIMMING_ACTUATOR_BASIC_SWITCFH_OFF_BRIGHTNESS"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_MINIMUM_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_MAXIMUM_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_SWITCH_ON_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION:
		return "PID_DIMMING_ACTUATOR_BASIC_DIMM_MODE_SELECTION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE:
		return "PID_DIMMING_ACTUATOR_BASIC_RELATIV_OFF_ENABLE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION:
		return "PID_DIMMING_ACTUATOR_BASIC_MEMORY_FUNCTION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION:
		return "PID_DIMMING_ACTUATOR_BASIC_TIMED_ON_RETRIGGER_FUNCTION"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE:
		return "PID_DIMMING_ACTUATOR_BASIC_MANUAL_OFF_ENABLE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE:
		return "PID_DIMMING_ACTUATOR_BASIC_INVERT_LOCK_DEVICE"
	case KnxInterfaceObjectProperty_PID_GENERAL_POLL_GROUP_SETTINGS:
		return "PID_GENERAL_POLL_GROUP_SETTINGS"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING:
		return "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_LOCKING"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING:
		return "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_AT_UNLOCKING"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_LOCK_SETVALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_UNLOCK_SETVALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE:
		return "PID_DIMMING_ACTUATOR_BASIC_BIGHTNESS_FOR_SCENE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE:
		return "PID_DIMMING_ACTUATOR_BASIC_STORAGE_FUNCTION_FOR_SCENE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_DELTA_DIMMING_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP:
		return "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_UP_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN:
		return "PID_DIMMING_ACTUATOR_BASIC_BEHAVIOUR_BUS_POWER_DOWN"
	case KnxInterfaceObjectProperty_PID_GENERAL_MANUFACTURER_DATA:
		return "PID_GENERAL_MANUFACTURER_DATA"
	case KnxInterfaceObjectProperty_PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE:
		return "PID_DIMMING_ACTUATOR_BASIC_BUS_POWER_DOWN_SET_VALUE"
	case KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION:
		return "PID_DIMMING_SENSOR_BASIC_ON_OFF_ACTION"
	case KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE:
		return "PID_DIMMING_SENSOR_BASIC_ENABLE_TOGGLE_MODE"
	case KnxInterfaceObjectProperty_PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE:
		return "PID_DIMMING_SENSOR_BASIC_ABSOLUTE_SETVALUE"
	case KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION:
		return "PID_SWITCHING_SENSOR_BASIC_ON_OFF_ACTION"
	case KnxInterfaceObjectProperty_PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE:
		return "PID_SWITCHING_SENSOR_BASIC_ENABLE_TOGGLE_MODE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME:
		return "PID_SUNBLIND_ACTUATOR_BASIC_REVERSION_PAUSE_TIME"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME:
		return "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_UP_DOWN_TIME"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME:
		return "PID_SUNBLIND_ACTUATOR_BASIC_SLAT_STEP_TIME"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME:
		return "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_PRESET_POSITION_TIME"
	case KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_NAME:
		return "PID_GENERAL_OBJECT_NAME"
	case KnxInterfaceObjectProperty_PID_GENERAL_ENABLE:
		return "PID_GENERAL_ENABLE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT:
		return "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_IN_PERCENT"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH:
		return "PID_SUNBLIND_ACTUATOR_BASIC_MOVE_TO_PRESET_POSITION_LENGTH"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT:
		return "PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_PERCENT"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE:
		return "PID_SUNBLIND_ACTUATOR_BASIC_PRESET_SLAT_POSITION_ANGLE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_WIND_ALARM"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_WIND_ALARM"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_ON_RAIN_ALARM"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_RAIN_ALARM"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_REACTION_FROST_ALARM"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM:
		return "PID_SUNBLIND_ACTUATOR_BASIC_HEARTBEAT_FROST_ALARM"
	case KnxInterfaceObjectProperty_PID_GENERAL_DESCRIPTION:
		return "PID_GENERAL_DESCRIPTION"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME:
		return "PID_SUNBLIND_ACTUATOR_BASIC_MAX_SLAT_MOVE_TIME"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE:
		return "PID_SUNBLIND_ACTUATOR_BASIC_ENABLE_BLINDS_MODE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE:
		return "PID_SUNBLIND_ACTUATOR_BASIC_STORAGE_FUNCTIONS_FOR_SCENE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE:
		return "PID_SUNBLIND_SENSOR_BASIC_ENABLE_BLINDS_MODE"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION:
		return "PID_SUNBLIND_SENSOR_BASIC_UP_DOWN_ACTION"
	case KnxInterfaceObjectProperty_PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE:
		return "PID_SUNBLIND_SENSOR_BASIC_ENABLE_TOGGLE_MODE"
	case KnxInterfaceObjectProperty_PID_GENERAL_FILE:
		return "PID_GENERAL_FILE"
	case KnxInterfaceObjectProperty_PID_GENERAL_TABLE:
		return "PID_GENERAL_TABLE"
	case KnxInterfaceObjectProperty_PID_GENERAL_ENROL:
		return "PID_GENERAL_ENROL"
	case KnxInterfaceObjectProperty_PID_GENERAL_VERSION:
		return "PID_GENERAL_VERSION"
	case KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_LINK:
		return "PID_GENERAL_GROUP_OBJECT_LINK"
	case KnxInterfaceObjectProperty_PID_GENERAL_MCB_TABLE:
		return "PID_GENERAL_MCB_TABLE"
	case KnxInterfaceObjectProperty_PID_GENERAL_ERROR_CODE:
		return "PID_GENERAL_ERROR_CODE"
	case KnxInterfaceObjectProperty_PID_GENERAL_OBJECT_INDEX:
		return "PID_GENERAL_OBJECT_INDEX"
	case KnxInterfaceObjectProperty_PID_GENERAL_SEMAPHOR:
		return "PID_GENERAL_SEMAPHOR"
	case KnxInterfaceObjectProperty_PID_GENERAL_DOWNLOAD_COUNTER:
		return "PID_GENERAL_DOWNLOAD_COUNTER"
	case KnxInterfaceObjectProperty_PID_DEVICE_ROUTING_COUNT:
		return "PID_DEVICE_ROUTING_COUNT"
	case KnxInterfaceObjectProperty_PID_DEVICE_MAX_RETRY_COUNT:
		return "PID_DEVICE_MAX_RETRY_COUNT"
	case KnxInterfaceObjectProperty_PID_DEVICE_ERROR_FLAGS:
		return "PID_DEVICE_ERROR_FLAGS"
	case KnxInterfaceObjectProperty_PID_DEVICE_PROGMODE:
		return "PID_DEVICE_PROGMODE"
	case KnxInterfaceObjectProperty_PID_DEVICE_PRODUCT_ID:
		return "PID_DEVICE_PRODUCT_ID"
	case KnxInterfaceObjectProperty_PID_DEVICE_MAX_APDULENGTH:
		return "PID_DEVICE_MAX_APDULENGTH"
	case KnxInterfaceObjectProperty_PID_DEVICE_SUBNET_ADDR:
		return "PID_DEVICE_SUBNET_ADDR"
	case KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_ADDR:
		return "PID_DEVICE_DEVICE_ADDR"
	case KnxInterfaceObjectProperty_PID_DEVICE_PB_CONFIG:
		return "PID_DEVICE_PB_CONFIG"
	case KnxInterfaceObjectProperty_PID_GENERAL_GROUP_OBJECT_REFERENCE:
		return "PID_GENERAL_GROUP_OBJECT_REFERENCE"
	case KnxInterfaceObjectProperty_PID_DEVICE_ADDR_REPORT:
		return "PID_DEVICE_ADDR_REPORT"
	case KnxInterfaceObjectProperty_PID_DEVICE_ADDR_CHECK:
		return "PID_DEVICE_ADDR_CHECK"
	case KnxInterfaceObjectProperty_PID_DEVICE_OBJECT_VALUE:
		return "PID_DEVICE_OBJECT_VALUE"
	case KnxInterfaceObjectProperty_PID_DEVICE_OBJECTLINK:
		return "PID_DEVICE_OBJECTLINK"
	case KnxInterfaceObjectProperty_PID_DEVICE_APPLICATION:
		return "PID_DEVICE_APPLICATION"
	case KnxInterfaceObjectProperty_PID_DEVICE_PARAMETER:
		return "PID_DEVICE_PARAMETER"
	case KnxInterfaceObjectProperty_PID_DEVICE_OBJECTADDRESS:
		return "PID_DEVICE_OBJECTADDRESS"
	case KnxInterfaceObjectProperty_PID_DEVICE_PSU_TYPE:
		return "PID_DEVICE_PSU_TYPE"
	case KnxInterfaceObjectProperty_PID_DEVICE_PSU_STATUS:
		return "PID_DEVICE_PSU_STATUS"
	case KnxInterfaceObjectProperty_PID_DEVICE_PSU_ENABLE:
		return "PID_DEVICE_PSU_ENABLE"
	case KnxInterfaceObjectProperty_PID_GENERAL_LOAD_STATE_CONTROL:
		return "PID_GENERAL_LOAD_STATE_CONTROL"
	case KnxInterfaceObjectProperty_PID_DEVICE_DOMAIN_ADDRESS:
		return "PID_DEVICE_DOMAIN_ADDRESS"
	case KnxInterfaceObjectProperty_PID_DEVICE_IO_LIST:
		return "PID_DEVICE_IO_LIST"
	case KnxInterfaceObjectProperty_PID_DEVICE_MGT_DESCRIPTOR_01:
		return "PID_DEVICE_MGT_DESCRIPTOR_01"
	case KnxInterfaceObjectProperty_PID_DEVICE_PL110_PARAM:
		return "PID_DEVICE_PL110_PARAM"
	case KnxInterfaceObjectProperty_PID_DEVICE_RF_REPEAT_COUNTER:
		return "PID_DEVICE_RF_REPEAT_COUNTER"
	case KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_TABLE:
		return "PID_DEVICE_RECEIVE_BLOCK_TABLE"
	case KnxInterfaceObjectProperty_PID_DEVICE_RANDOM_PAUSE_TABLE:
		return "PID_DEVICE_RANDOM_PAUSE_TABLE"
	case KnxInterfaceObjectProperty_PID_DEVICE_RECEIVE_BLOCK_NR:
		return "PID_DEVICE_RECEIVE_BLOCK_NR"
	case KnxInterfaceObjectProperty_PID_DEVICE_HARDWARE_TYPE:
		return "PID_DEVICE_HARDWARE_TYPE"
	case KnxInterfaceObjectProperty_PID_DEVICE_RETRANSMITTER_NUMBER:
		return "PID_DEVICE_RETRANSMITTER_NUMBER"
	case KnxInterfaceObjectProperty_PID_GENERAL_RUN_STATE_CONTROL:
		return "PID_GENERAL_RUN_STATE_CONTROL"
	case KnxInterfaceObjectProperty_PID_DEVICE_SERIAL_NR_TABLE:
		return "PID_DEVICE_SERIAL_NR_TABLE"
	case KnxInterfaceObjectProperty_PID_DEVICE_BIBATMASTER_ADDRESS:
		return "PID_DEVICE_BIBATMASTER_ADDRESS"
	case KnxInterfaceObjectProperty_PID_DEVICE_RF_DOMAIN_ADDRESS:
		return "PID_DEVICE_RF_DOMAIN_ADDRESS"
	case KnxInterfaceObjectProperty_PID_DEVICE_DEVICE_DESCRIPTOR:
		return "PID_DEVICE_DEVICE_DESCRIPTOR"
	case KnxInterfaceObjectProperty_PID_DEVICE_METERING_FILTER_TABLE:
		return "PID_DEVICE_METERING_FILTER_TABLE"
	case KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE:
		return "PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_TIME_BASE"
	case KnxInterfaceObjectProperty_PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR:
		return "PID_DEVICE_GROUP_TELEGR_RATE_LIMIT_NO_OF_TELEGR"
	case KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_GRPOBJTABLE:
		return "PID_GROUP_OBJECT_TABLE_GRPOBJTABLE"
	case KnxInterfaceObjectProperty_PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE:
		return "PID_GROUP_OBJECT_TABLE_EXT_GRPOBJREFERENCE"
	case KnxInterfaceObjectProperty_PID_ROUTER_LINE_STATUS:
		return "PID_ROUTER_LINE_STATUS"
	case KnxInterfaceObjectProperty_PID_GENERAL_TABLE_REFERENCE:
		return "PID_GENERAL_TABLE_REFERENCE"
	case KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCCONFIG:
		return "PID_ROUTER_MAIN_LCCONFIG"
	case KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCCONFIG:
		return "PID_ROUTER_SUB_LCCONFIG"
	case KnxInterfaceObjectProperty_PID_ROUTER_MAIN_LCGRPCONFIG:
		return "PID_ROUTER_MAIN_LCGRPCONFIG"
	case KnxInterfaceObjectProperty_PID_ROUTER_SUB_LCGRPCONFIG:
		return "PID_ROUTER_SUB_LCGRPCONFIG"
	case KnxInterfaceObjectProperty_PID_ROUTER_ROUTETABLE_CONTROL:
		return "PID_ROUTER_ROUTETABLE_CONTROL"
	case KnxInterfaceObjectProperty_PID_ROUTER_COUPL_SERV_CONTROL:
		return "PID_ROUTER_COUPL_SERV_CONTROL"
	case KnxInterfaceObjectProperty_PID_ROUTER_MAX_ROUTER_APDU_LENGTH:
		return "PID_ROUTER_MAX_ROUTER_APDU_LENGTH"
	case KnxInterfaceObjectProperty_PID_ROUTER_MEDIUM:
		return "PID_ROUTER_MEDIUM"
	case KnxInterfaceObjectProperty_PID_ROUTER_FILTER_TABLE_USE:
		return "PID_ROUTER_FILTER_TABLE_USE"
	case KnxInterfaceObjectProperty_PID_ROUTER_RF_ENABLE_SBC:
		return "PID_ROUTER_RF_ENABLE_SBC"
	case KnxInterfaceObjectProperty_PID_GENERAL_SERVICE_CONTROL:
		return "PID_GENERAL_SERVICE_CONTROL"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID:
		return "PID_KNXIP_PARAMETER_PROJECT_INSTALLATION_ID"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS:
		return "PID_KNXIP_PARAMETER_KNX_INDIVIDUAL_ADDRESS"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES:
		return "PID_KNXIP_PARAMETER_ADDITIONAL_INDIVIDUAL_ADDRESSES"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD:
		return "PID_KNXIP_PARAMETER_CURRENT_IP_ASSIGNMENT_METHOD"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD:
		return "PID_KNXIP_PARAMETER_IP_ASSIGNMENT_METHOD"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_CAPABILITIES:
		return "PID_KNXIP_PARAMETER_IP_CAPABILITIES"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS:
		return "PID_KNXIP_PARAMETER_CURRENT_IP_ADDRESS"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK:
		return "PID_KNXIP_PARAMETER_CURRENT_SUBNET_MASK"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY:
		return "PID_KNXIP_PARAMETER_CURRENT_DEFAULT_GATEWAY"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_IP_ADDRESS:
		return "PID_KNXIP_PARAMETER_IP_ADDRESS"
	case KnxInterfaceObjectProperty_PID_GENERAL_FIRMWARE_REVISION:
		return "PID_GENERAL_FIRMWARE_REVISION"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SUBNET_MASK:
		return "PID_KNXIP_PARAMETER_SUBNET_MASK"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DEFAULT_GATEWAY:
		return "PID_KNXIP_PARAMETER_DEFAULT_GATEWAY"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER:
		return "PID_KNXIP_PARAMETER_DHCP_BOOTP_SERVER"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_MAC_ADDRESS:
		return "PID_KNXIP_PARAMETER_MAC_ADDRESS"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS:
		return "PID_KNXIP_PARAMETER_SYSTEM_SETUP_MULTICAST_ADDRESS"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS:
		return "PID_KNXIP_PARAMETER_ROUTING_MULTICAST_ADDRESS"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_TTL:
		return "PID_KNXIP_PARAMETER_TTL"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES:
		return "PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_CAPABILITIES"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE:
		return "PID_KNXIP_PARAMETER_KNXNETIP_DEVICE_STATE"
	case KnxInterfaceObjectProperty_PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES:
		return "PID_KNXIP_PARAMETER_KNXNETIP_ROUTING_CAPABILITIES"
	}
	return ""
}
