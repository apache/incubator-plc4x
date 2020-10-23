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
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

type CEMIPriority uint8

type ICEMIPriority interface {
    spi.Message
    Serialize(io utils.WriteBuffer) error
}

const(
    CEMIPriority_SYSTEM CEMIPriority = 0x0
    CEMIPriority_NORMAL CEMIPriority = 0x1
    CEMIPriority_URGENT CEMIPriority = 0x2
    CEMIPriority_LOW CEMIPriority = 0x3
)

func CEMIPriorityValueOf(value uint8) CEMIPriority {
    switch value {
        case 0x0:
            return CEMIPriority_SYSTEM
        case 0x1:
            return CEMIPriority_NORMAL
        case 0x2:
            return CEMIPriority_URGENT
        case 0x3:
            return CEMIPriority_LOW
    }
    return 0
}

func CastCEMIPriority(structType interface{}) CEMIPriority {
    castFunc := func(typ interface{}) CEMIPriority {
        if sCEMIPriority, ok := typ.(CEMIPriority); ok {
            return sCEMIPriority
        }
        return 0
    }
    return castFunc(structType)
}

func (m CEMIPriority) LengthInBits() uint16 {
    return 2
}

func (m CEMIPriority) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIPriorityParse(io *utils.ReadBuffer) (CEMIPriority, error) {
    // TODO: Implement ...
    return 0, nil
}

func (e CEMIPriority) Serialize(io utils.WriteBuffer) error {
    // TODO: Implement ...
    return nil
}
