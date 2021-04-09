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

// Remark: The different fields are encoded in Big-endian.

[enum string '-1' 'OpcuaDataType'
    ['IEC61131_NULL' NULL ]
    ['IEC61131_BOOL' BOOL ]
    ['IEC61131_BYTE' BYTE ]
    ['IEC61131_WORD' WORD ]
    ['IEC61131_DWORD' DWORD ]
    ['IEC61131_LWORD' LWORD ]
    ['IEC61131_SINT' SINT ]
    ['IEC61131_INT' INT ]
    ['IEC61131_DINT' DINT ]
    ['IEC61131_LINT' LINT ]
    ['IEC61131_USINT' USINT ]
    ['IEC61131_UINT' UINT ]
    ['IEC61131_UDINT' UDINT ]
    ['IEC61131_ULINT' ULINT ]
    ['IEC61131_REAL' REAL ]
    ['IEC61131_LREAL' LREAL ]
    ['IEC61131_TIME' TIME ]
    ['IEC61131_LTIME' LTIME ]
    ['IEC61131_DATE' DATE ]
    ['IEC61131_LDATE' LDATE ]
    ['IEC61131_TIME_OF_DAY' TIME_OF_DAY ]
    ['IEC61131_LTIME_OF_DAY' LTIME_OF_DAY ]
    ['IEC61131_DATE_AND_TIME' DATE_AND_TIME ]
    ['IEC61131_LDATE_AND_TIME' LDATE_AND_TIME ]
    ['IEC61131_CHAR' CHAR ]
    ['IEC61131_WCHAR' WCHAR ]
    ['IEC61131_STRING' STRING ]
    ['IEC61131_WSTRING' WSTRING ]
]


[enum string '-1' 'OpcuaIdentifierType'
    ['s' STRING_IDENTIFIER]
    ['i' NUMBER_IDENTIFIER]
    ['g' GUID_IDENTIFIER]
    ['b' BINARY_IDENTIFIER]
]
