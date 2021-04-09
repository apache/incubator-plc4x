//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

[discriminatedType 'KNXNetIPMessage'
    [implicit      uint 8  'headerLength'    '6']
    [const         uint 8  'protocolVersion' '0x10']
    [discriminator uint 16 'msgType']
    [implicit      uint 16 'totalLength'     'lengthInBytes']
    [typeSwitch 'msgType'
        ['0x0201' SearchRequest
            [simple HPAIDiscoveryEndpoint 'hpaiIDiscoveryEndpoint']
        ]
        ['0x0202' SearchResponse
            [simple HPAIControlEndpoint 'hpaiControlEndpoint']
            [simple DIBDeviceInfo       'dibDeviceInfo']
            [simple DIBSuppSvcFamilies  'dibSuppSvcFamilies']
        ]
        ['0x0203' DescriptionRequest
            [simple HPAIControlEndpoint 'hpaiControlEndpoint']
        ]
        ['0x0204' DescriptionResponse
            [simple DIBDeviceInfo       'dibDeviceInfo']
            [simple DIBSuppSvcFamilies  'dibSuppSvcFamilies']
        ]
        ['0x0205' ConnectionRequest
            [simple HPAIDiscoveryEndpoint        'hpaiDiscoveryEndpoint']
            [simple HPAIDataEndpoint             'hpaiDataEndpoint']
            [simple ConnectionRequestInformation 'connectionRequestInformation']
        ]
        ['0x0206' ConnectionResponse
            [simple   uint 8 'communicationChannelId']
            [enum     Status 'status']
            [optional HPAIDataEndpoint            'hpaiDataEndpoint'            'status == Status.NO_ERROR']
            [optional ConnectionResponseDataBlock 'connectionResponseDataBlock' 'status == Status.NO_ERROR']
        ]
        ['0x0207' ConnectionStateRequest
            [simple   uint 8 'communicationChannelId']
            [reserved uint 8 '0x00']
            [simple   HPAIControlEndpoint 'hpaiControlEndpoint']
        ]
        ['0x0208' ConnectionStateResponse
            [simple uint 8 'communicationChannelId']
            [enum   Status 'status']
        ]
        ['0x0209' DisconnectRequest
            [simple   uint 8 'communicationChannelId']
            [reserved uint 8 '0x00']
            [simple   HPAIControlEndpoint 'hpaiControlEndpoint']
        ]
        ['0x020A' DisconnectResponse
            [simple uint 8 'communicationChannelId']
            [enum   Status 'status']
        ]
        ['0x020B' UnknownMessage [uint 16 'totalLength']
            [array int 8 'unknownData' count 'totalLength - 6']
        ]
        ['0x0310' DeviceConfigurationRequest [uint 16 'totalLength']
            [simple DeviceConfigurationRequestDataBlock 'deviceConfigurationRequestDataBlock']
            [simple CEMI                                'cemi' ['totalLength - (6 + deviceConfigurationRequestDataBlock.lengthInBytes)']]
        ]
        ['0x0311' DeviceConfigurationAck
            [simple DeviceConfigurationAckDataBlock 'deviceConfigurationAckDataBlock']
        ]
        ['0x0420' TunnelingRequest [uint 16 'totalLength']
            [simple TunnelingRequestDataBlock 'tunnelingRequestDataBlock']
            [simple CEMI                      'cemi' ['totalLength - (6 + tunnelingRequestDataBlock.lengthInBytes)']]
        ]
        ['0x0421' TunnelingResponse
            [simple TunnelingResponseDataBlock 'tunnelingResponseDataBlock']
        ]
        ['0x0530' RoutingIndication
        ]
    ]
]

[type 'HPAIDiscoveryEndpoint'
    [implicit uint 8           'structureLength' 'lengthInBytes']
    [enum     HostProtocolCode 'hostProtocolCode']
    [simple   IPAddress        'ipAddress']
    [simple   uint 16          'ipPort']
]

[type 'HPAIControlEndpoint'
    [implicit uint 8           'structureLength' 'lengthInBytes']
    [enum     HostProtocolCode 'hostProtocolCode']
    [simple   IPAddress        'ipAddress']
    [simple   uint 16          'ipPort']
]

[type 'DIBDeviceInfo'
    [implicit uint 8       'structureLength' 'lengthInBytes']
    [simple   uint 8       'descriptionType']
    [simple   uint 8       'knxMedium']
    [simple   DeviceStatus 'deviceStatus']
    [simple   KNXAddress   'knxAddress']
    [simple   ProjectInstallationIdentifier 'projectInstallationIdentifier']
    [array    int 8        'knxNetIpDeviceSerialNumber' count '6']
    [simple   IPAddress    'knxNetIpDeviceMulticastAddress']
    [simple   MACAddress   'knxNetIpDeviceMacAddress']
    [array    int 8        'deviceFriendlyName'         count '30']
]

[type 'DIBSuppSvcFamilies'
    [implicit uint 8       'structureLength' 'lengthInBytes']
    [simple   uint 8       'descriptionType']
    [array    ServiceId    'serviceIds' count '3']
]

[type 'HPAIDataEndpoint'
    [implicit uint 8           'structureLength' 'lengthInBytes']
    [enum     HostProtocolCode 'hostProtocolCode']
    [simple   IPAddress        'ipAddress']
    [simple   uint 16          'ipPort']
]

[discriminatedType 'ConnectionRequestInformation'
    [implicit      uint 8    'structureLength' 'lengthInBytes']
    [discriminator uint 8    'connectionType']
    [typeSwitch 'connectionType'
        ['0x03' ConnectionRequestInformationDeviceManagement
        ]
        ['0x04' ConnectionRequestInformationTunnelConnection
            [enum     KnxLayer  'knxLayer']
            [reserved uint 8    '0x00']
        ]
    ]
]

[discriminatedType 'ConnectionResponseDataBlock'
    [implicit      uint 8     'structureLength' 'lengthInBytes']
    [discriminator uint 8     'connectionType']
    [typeSwitch 'connectionType'
        ['0x03' ConnectionResponseDataBlockDeviceManagement
        ]
        ['0x04' ConnectionResponseDataBlockTunnelConnection
            [simple KNXAddress 'knxAddress']
        ]
    ]
]

[type 'DeviceConfigurationRequestDataBlock'
    [implicit uint 8 'structureLength' 'lengthInBytes']
    [simple   uint 8 'communicationChannelId']
    [simple   uint 8 'sequenceCounter']
    [reserved uint 8 '0x00']
]

[type 'DeviceConfigurationAckDataBlock'
    [implicit uint 8 'structureLength' 'lengthInBytes']
    [simple   uint 8 'communicationChannelId']
    [simple   uint 8 'sequenceCounter']
    [enum     Status 'status']
]

[type 'TunnelingRequestDataBlock'
    [implicit uint 8 'structureLength' 'lengthInBytes']
    [simple   uint 8 'communicationChannelId']
    [simple   uint 8 'sequenceCounter']
    [reserved uint 8 '0x00']
]

[type 'TunnelingResponseDataBlock'
    [implicit uint 8 'structureLength' 'lengthInBytes']
    [simple   uint 8 'communicationChannelId']
    [simple   uint 8 'sequenceCounter']
    [enum     Status 'status']
]

[type 'IPAddress'
    [array int 8 'addr' count '4']
]

[type 'MACAddress'
    [array int 8 'addr' count '6']
]

[type 'KNXAddress'
    [simple uint 4 'mainGroup']
    [simple uint 4 'middleGroup']
    [simple uint 8 'subGroup']
]

[type 'DeviceStatus'
    [reserved uint 7 '0x00']
    [simple   bit    'programMode']
]

[type 'ProjectInstallationIdentifier'
    [simple uint 8 'projectNumber']
    [simple uint 8 'installationNumber']
]

[discriminatedType 'ServiceId'
    [discriminator uint 8 'serviceType']
    [typeSwitch 'serviceType'
        ['0x02' KnxNetIpCore
            [simple uint 8 'version']
        ]
        ['0x03' KnxNetIpDeviceManagement
            [simple uint 8 'version']
        ]
        ['0x04' KnxNetIpTunneling
            [simple uint 8 'version']
        ]
        ['0x06' KnxNetRemoteLogging
            [simple uint 8 'version']
        ]
        ['0x07' KnxNetRemoteConfigurationAndDiagnosis
            [simple uint 8 'version']
        ]
        ['0x08' KnxNetObjectServer
            [simple uint 8 'version']
        ]
    ]
]

/* The CEMI part is described in the document "03_06_03 EMI_IMI v01.03.03 AS" */
[discriminatedType 'CEMI' [uint 8 'size']
    [discriminator uint 8 'messageCode']
    [typeSwitch 'messageCode'
        ['0x11' CEMIDataReq
            [simple uint 8                    'additionalInformationLength']
            [array  CEMIAdditionalInformation 'additionalInformation' length 'additionalInformationLength']
            [simple CEMIDataFrame             'cemiDataFrame']
        ]
        ['0x2E' CEMIDataCon
            [simple uint 8                    'additionalInformationLength']
            [array  CEMIAdditionalInformation 'additionalInformation' length 'additionalInformationLength']
            [simple CEMIDataFrame             'cemiDataFrame']
        ]
        ['0x29' CEMIDataInd
            [simple uint 8                    'additionalInformationLength']
            [array  CEMIAdditionalInformation 'additionalInformation' length 'additionalInformationLength']
            [simple CEMIDataFrame             'cemiDataFrame']
        ]

        ['0x10' CEMIRawReq
        ]
        ['0x2F' CEMIRawCon
        ]
        ['0x2D' CEMIRawInd
        ]

        ['0x13' CEMIPollDataReq
        ]
        ['0x25' CEMIPollDataCon
        ]

        ['0x2B' CEMIBusmonInd
            [simple uint 8                    'additionalInformationLength']
            [array  CEMIAdditionalInformation 'additionalInformation' length 'additionalInformationLength']
            [simple CEMIFrame                 'cemiFrame']
        ]

        ['0xFC' CEMIMPropReadReq
            [simple uint 16 'interfaceObjectType']
            [simple uint  8 'objectInstance']
            [simple uint  8 'propertyId']
            [simple uint  4 'numberOfElements']
            [simple uint 12 'startIndex']
        ]
        ['0xFB' CEMIMPropReadCon
            [simple uint 16 'interfaceObjectType']
            [simple uint  8 'objectInstance']
            [simple uint  8 'propertyId']
            [simple uint  4 'numberOfElements']
            [simple uint 12 'startIndex']
            [simple uint 16 'unknown']
        ]
    ]
]

[discriminatedType 'CEMIAdditionalInformation'
    [discriminator uint 8 'additionalInformationType']
    [typeSwitch 'additionalInformationType'
        ['0x03' CEMIAdditionalInformationBusmonitorInfo
            [const     uint 8 'len' '1']
            [simple    bit    'frameErrorFlag']
            [simple    bit    'bitErrorFlag']
            [simple    bit    'parityErrorFlag']
            [simple    bit    'unknownFlag']
            [simple    bit    'lostFlag']
            [simple    uint 3 'sequenceNumber']
        ]
        ['0x04' CEMIAdditionalInformationRelativeTimestamp
            [const    uint 8            'len' '2']
            [simple   RelativeTimestamp 'relativeTimestamp']
        ]
    ]
]

[type 'CEMIDataFrame'
    [simple        bit          'standardFrame']
    [simple        bit          'polling']
    [simple        bit          'notRepeated']
    [simple        bit          'notAckFrame']
    [enum          CEMIPriority 'priority']
    [simple        bit          'acknowledgeRequested']
    [simple        bit          'errorFlag']
    [simple        bit          'groupDestinationAddress']
    [simple        uint 3       'hopCount']
    [simple        uint 4       'extendedFrameFormat']
    [simple        KNXAddress   'sourceAddress']
    [array         int 8        'destinationAddress' count '2']
    [simple        uint 8       'dataLength']
    [enum          TPCI         'tcpi']
    [simple        uint 4       'counter']
    [enum          APCI         'apci']
    [simple        int 6        'dataFirstByte']
    [array         int 8        'data' count 'dataLength - 1']
]

/* The CEMI part is described in the document "03_06_03 EMI_IMI v01.03.03 AS" Page 73
"03_02_02 Communication Medium TP1 v01.02.02 AS" Page 27 */
[discriminatedType 'CEMIFrame'
    [discriminator bit          'standardFrame']
    [discriminator bit          'polling']
    [simple        bit          'repeated']
    [discriminator bit          'notAckFrame']
    [enum          CEMIPriority 'priority']
    [simple        bit          'acknowledgeRequested']
    [simple        bit          'errorFlag']
    [typeSwitch 'notAckFrame','standardFrame','polling'
        ['false' CEMIFrameAck
        ]
        ['true','true','false' CEMIFrameData
            [simple   KNXAddress      'sourceAddress']
            [array    int 8           'destinationAddress' count '2']
            [simple   bit             'groupAddress']
            [simple   uint 3          'hopCount']
            [simple   uint 4          'dataLength']
            [enum     TPCI            'tcpi']
            [simple   uint 4          'counter']
            [enum     APCI            'apci']
            [simple   int 6           'dataFirstByte']
            [array    int 8           'data' count 'dataLength - 1']
            [simple   uint 8          'crc']
        ]
        ['true','true','true' CEMIFramePollingData
        ]
        ['true','false','false' CEMIFrameDataExt
            [simple   bit             'groupAddress']
            [simple   uint 3          'hopCount']
            [simple   uint 4          'extendedFrameFormat']
            [simple   KNXAddress      'sourceAddress']
            [array    int 8           'destinationAddress' count '2']
            [simple   uint 8          'dataLength']
            [enum     TPCI            'tcpi']
            [simple   uint 4          'counter']
            [enum     APCI            'apci']
            [simple   int 6           'dataFirstByte']
            [array    int 8           'data' count 'dataLength - 1']
            [simple   uint 8          'crc']
        ]
        ['true','false','true' CEMIFramePollingDataExt
        ]
    ]
]

[type 'RelativeTimestamp'
    [simple   uint 16 'timestamp']
]

[discriminatedType 'KNXGroupAddress' [uint 2 'numLevels']
    [typeSwitch 'numLevels'
        ['1' KNXGroupAddressFreeLevel
            [simple uint 16 'subGroup']
        ]
        ['2' KNXGroupAddress2Level
            [simple uint 5  'mainGroup']
            [simple uint 11 'subGroup']
        ]
        ['3' KNXGroupAddress3Level
            [simple uint 5 'mainGroup']
            [simple uint 3 'middleGroup']
            [simple uint 8 'subGroup']
        ]
    ]
]

[dataIo 'KnxDatapoint' [uint 10 'mainNumber', uint 10 'subNumber']
    [typeSwitch 'mainNumber','subNumber'
        ['1' BOOL
            [reserved uint 7 '0x0']
            [simple   bit    'value']
        ]
        ['2' BOOL
            [reserved uint 6 '0x0']
            [simple   bit    'control']
            [simple   bit    'value']
        ]
        ['21' Struct
            [simple   bit    'b7']
            [simple   bit    'b6']
            [simple   bit    'b5']
            [simple   bit    'b4']
            [simple   bit    'b3']
            [simple   bit    'b2']
            [simple   bit    'b1']
            [simple   bit    'b0']
        ]
        ['3' USINT
            [reserved uint 4 '0x0']
            [simple   bit    'control']
            [simple   uint 3 'value']
        ]
        ['18' USINT
            [simple   bit    'control']
            [reserved uint 1 '0x0']
            [simple   uint 6 'value']
        ]
        ['17' USINT
            [reserved uint 2 '0x0']
            [simple   uint 6 'value']
        ]
        ['5' USINT
            [reserved uint 8 '0x0']
            [simple   uint 8 'value']
        ]
        ['7' UINT
            [reserved uint 8 '0x0']
            [simple uint 16 'value']
        ]
        ['12' UDINT
            [reserved uint 8 '0x0']
            [simple uint 32 'value']
        ]
        ['6','20' SINT
            [simple   bit   'a']
            [simple   bit   'b']
            [simple   bit   'c']
            [simple   bit   'd']
            [simple   bit   'e']
            [simple   int 8 'value']
        ]
        ['6' SINT
            [reserved uint 8 '0x0']
            [simple   int  8 'value']
        ]
        ['8' INT
            [reserved uint 8  '0x0']
            [simple   int  16 'value']
        ]
        ['13' DINT
            [reserved uint 8  '0x0']
            [simple   int  32 'value']
        ]
        ['9' REAL
            [reserved uint  8    '0x0']
            [manual   float 4.11 'value' 'STATIC_CALL("org.apache.plc4x.java.knxnetip.utils.KnxHelper.bytesToF16", io)' 'STATIC_CALL("org.apache.plc4x.java.knxnetip.utils.KnxHelper.f16toBytes", io, object)' '16']
        ]
        ['14' REAL
            [reserved uint  8    '0x0']
            [simple   float 8.23 'value']
        ]
        ['4' String
            [reserved uint   8 '0x0']
            [simple   string 8 'UTF-8' 'value']
        ]
        ['16' String
            [reserved uint   8   '0x0']
            [simple   string 112 'UTF-8' 'value']
        ]
        ['10' Time
            [simple   uint 3 'day']
            [simple   uint 5 'hours']
            [reserved uint 2 '0x0']
            [simple   uint 6 'minutes']
            [reserved uint 2 '0x0']
            [simple   uint 6 'seconds']
        ]
        ['11' Date
            [reserved uint 3 '0x0']
            [simple   uint 5 'day']
            [reserved uint 4 '0x0']
            [simple   uint 4 'month']
            [reserved uint 1 '0x0']
            [simple   uint 6 'year']
        ]
        ['19' DateTime
            [simple   uint 8 'year']
            [reserved uint 4 '0x0']
            [simple   uint 4 'month']
            [reserved uint 3 '0x0']
            [simple   uint 5 'day']
            [simple   uint 3 'dayOfWeek']
            [simple   uint 5 'hours']
            [reserved uint 2 '0x0']
            [simple   uint 6 'minutes']
            [reserved uint 2 '0x0']
            [simple   uint 6 'seconds']
            [simple   bit    'fault']
            [simple   bit    'workingDay']
            [simple   bit    'workingDayValid']
            [simple   bit    'yearValid']
            [simple   bit    'dayAndMonthValid']
            [simple   bit    'dayOfWeekValid']
            [simple   bit    'timeValid']
            [simple   bit    'standardSummerTime']
            [simple   bit    'clockQuality']
        ]
        ['15' Struct
            [simple   uint 4 'D6']
            [simple   uint 4 'D5']
            [simple   uint 4 'D4']
            [simple   uint 4 'D3']
            [simple   uint 4 'D2']
            [simple   uint 4 'D1']
            [simple   bit    'BE']
            [simple   bit    'BP']
            [simple   bit    'BD']
            [simple   bit    'BC']
            [simple   uint 4 'index']
        ]
    ]
]

[enum uint 2 'CEMIPriority'
    ['0x0' SYSTEM]
    ['0x1' NORMAL]
    ['0x2' URGENT]
    ['0x3' LOW]
]

[enum uint 8 'Status'
    ['0x00' NO_ERROR]
    ['0x01' PROTOCOL_TYPE_NOT_SUPPORTED]
    ['0x02' UNSUPPORTED_PROTOCOL_VERSION]
    ['0x04' OUT_OF_ORDER_SEQUENCE_NUMBER]
    ['0x21' INVALID_CONNECTION_ID]
    ['0x22' CONNECTION_TYPE_NOT_SUPPORTED]
    ['0x23' CONNECTION_OPTION_NOT_SUPPORTED]
    ['0x24' NO_MORE_CONNECTIONS]
    ['0x25' NO_MORE_UNIQUE_CONNECTIONS]
    ['0x26' DATA_CONNECTION]
    ['0x27' KNX_CONNECTION]
    ['0x29' TUNNELLING_LAYER_NOT_SUPPORTED]
]

[enum uint 8 'HostProtocolCode'
    ['0x01' IPV4_UDP]
    ['0x02' IPV4_TCP]
]

/*
 The mode in which the connection should be established:
 TUNNEL_LINK_LAYER: The gateway assigns a unique KNX address to the client.
                    The client can then actively participate in communicating
                    with other KNX devices.
 TUNNEL_RAW:        The gateway will just pass along the packets and not
                    automatically generate Ack frames for the packets it
                    receives for a given client.
 TUNNEL_BUSMONITOR: The client becomes a passive participant and all frames
                    on the KNX bus get forwarded to the client. Only one
                    Busmonitor connection is allowed at any given time.
*/
[enum uint 8 'KnxLayer'
    ['0x02' TUNNEL_LINK_LAYER]
    ['0x04' TUNNEL_RAW]
    ['0x80' TUNNEL_BUSMONITOR]
]

[enum uint 2 'TPCI'
    ['0x0' UNNUMBERED_DATA_PACKET]
    ['0x1' UNNUMBERED]
    ['0x2' NUMBERED_DATA_PACKET]
    ['0x3' NUMBERED_CONTROL_DATA]
]

[enum uint 4 'APCI'
    ['0x0' GROUP_VALUE_READ_PDU]
    ['0x1' GROUP_VALUE_RESPONSE_PDU]
    ['0x2' GROUP_VALUE_WRITE_PDU]
    ['0x3' INDIVIDUAL_ADDRESS_WRITE_PDU]
    ['0x4' INDIVIDUAL_ADDRESS_READ_PDU]
    ['0x5' INDIVIDUAL_ADDRESS_RESPONSE_PDU]
    ['0x6' ADC_READ_PDU]
    ['0x7' ADC_RESPONSE_PDU]
    ['0x8' MEMORY_READ_PDU]
    ['0x9' MEMORY_RESPONSE_PDU]
    ['0xA' MEMORY_WRITE_PDU]
    ['0xB' USER_MESSAGE_PDU]
    ['0xC' DEVICE_DESCRIPTOR_READ_PDU]
    ['0xD' DEVICE_DESCRIPTOR_RESPONSE_PDU]
    ['0xE' RESTART_PDU]
    ['0xF' OTHER_PDU]
]


