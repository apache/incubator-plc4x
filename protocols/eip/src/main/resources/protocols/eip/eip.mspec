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

 //////////////////////////////////////////////////////////////////
 ///EthernetIP Header of size 24
 /////////////////////////////////////////////////////////////////

[discriminatedType 'EipPacket'
    [discriminator uint 16 'command']
    [implicit      uint 16 'len' 'lengthInBytes - 24']
    [simple        uint 32 'sessionHandle']
    [simple        uint 32 'status']
    [array         uint 8  'senderContext' count '8']
    [simple        uint 32 'options']
    [typeSwitch 'command'
            ['0x0065' EipConnectionRequest
                [const  uint    16   'protocolVersion'   '0x01']
                [const  uint    16   'flags'             '0x00']
            ]
            ['0x0066' EipDisconnectRequest
            ]
            ['0x006F' CipRRData
                [reserved  uint    32    '0x00000000']
                [reserved  uint    16    '0x0000']
                [simple CipExchange 'exchange']
            ]
        ]
]
[discriminatedType  'CipExchange'
    [const          uint        16      'itemCount'           '0x0002']     //2 items
    [const          uint        32      'nullPtr'             '0x0']          //NullPointerAddress
    [const          uint        16      'UnconnectedData'     '0x00B2']   //Connection Manager
    [implicit       uint        16      'size'                'lengthInBytes - 8 - 2'] //remove fields above and routing
    [discriminator  uint        8       'CipServiceCode']
    [typeSwitch 'CipServiceCode'
       ['0xCC' CipReadResponse
            [reserved   uint            16   '0x0000']
            [simple     uint            8   'status']
            [enum     CIPDataTypeCode   'dataType']
            [array      int             8   'data'  length  'dataType.size']
        ]
        ['0x0052'   CipUnconnectedRequest
            [reserved   uint    8   '0x02']
            [reserved   uint    8   '0x20']   // setRequestPathLogicalClassSegment
            [reserved   uint    8   '0x06']   // set request class path
            [reserved   uint    8   '0x24']   // setRequestPathLogicalInstanceSegment
            [reserved   uint    8   '0x01']   // setRequestPathInstance
            [reserved   uint    16  '0x9D05']   //Timeout 5s
            [simple     CipService  'service']
        ]
    ]
]

[discriminatedType  'CipService'
    [implicit       uint    16   'messageSize'   'lengthInBytes - 4 - 2']   //substract messageSize and routing
    [discriminator  uint    8   'service']
    [typeSwitch 'service'
            ['0x4C' CipReadRequest
                [simple     int     8   'RequestPathSize']
                [array      int     8   'tag'   length  '(RequestPathSize*2)']
                [simple     uint    16  'elementNb']
                [const      uint    16  'route' '0x0001']
                [simple     int     8   'backPlane']
                [simple     int     8   'slot']
            ]
        ]
]


[type 'CipItem'
    [simple uint    16  'typeID']
    [simple uint    16  'size']
]

[enum uint   16   'CIPDataTypeCode' [uint 8  'size']
    ['0X00C1'   BOOL            ['1']]
    ['0X00CA'   REAL            ['4']]
    ['0X00C4'   DINT            ['4']]
    ['0X00C3'   INT             ['2']]
    ['0X00C2'   SINT            ['1']]
    ['0X02A0'   STRUCTURED      ['88']]
    ['0X02A0'   STRING          ['88']]
    ['0X02A0'   STRING36        ['40']]
    ['-1'       UNKNOWN         ['-1']]
]
