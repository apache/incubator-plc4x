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
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
)

type SzlSublist uint8

type ISzlSublist interface {
    Serialize(io utils.WriteBuffer) error
}

const(
    SzlSublist_MODULE_IDENTIFICATION SzlSublist = 0x11
    SzlSublist_CPU_FEATURES SzlSublist = 0x12
    SzlSublist_USER_MEMORY_AREA SzlSublist = 0x13
    SzlSublist_SYSTEM_AREAS SzlSublist = 0x14
    SzlSublist_BLOCK_TYPES SzlSublist = 0x15
    SzlSublist_STATUS_MODULE_LEDS SzlSublist = 0x19
    SzlSublist_COMPONENT_IDENTIFICATION SzlSublist = 0x1C
    SzlSublist_INTERRUPT_STATUS SzlSublist = 0x22
    SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS SzlSublist = 0x25
    SzlSublist_COMMUNICATION_STATUS_DATA SzlSublist = 0x32
    SzlSublist_STATUS_SINGLE_MODULE_LED SzlSublist = 0x74
    SzlSublist_DP_MASTER_SYSTEM_INFORMATION SzlSublist = 0x90
    SzlSublist_MODULE_STATUS_INFORMATION SzlSublist = 0x91
    SzlSublist_RACK_OR_STATION_STATUS_INFORMATION SzlSublist = 0x92
    SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2 SzlSublist = 0x94
    SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION SzlSublist = 0x95
    SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP SzlSublist = 0x96
    SzlSublist_DIAGNOSTIC_BUFFER SzlSublist = 0xA0
    SzlSublist_MODULE_DIAGNOSTIC_DATA SzlSublist = 0xB1
)

func SzlSublistValueOf(value uint8) SzlSublist {
    switch value {
        case 0x11:
            return SzlSublist_MODULE_IDENTIFICATION
        case 0x12:
            return SzlSublist_CPU_FEATURES
        case 0x13:
            return SzlSublist_USER_MEMORY_AREA
        case 0x14:
            return SzlSublist_SYSTEM_AREAS
        case 0x15:
            return SzlSublist_BLOCK_TYPES
        case 0x19:
            return SzlSublist_STATUS_MODULE_LEDS
        case 0x1C:
            return SzlSublist_COMPONENT_IDENTIFICATION
        case 0x22:
            return SzlSublist_INTERRUPT_STATUS
        case 0x25:
            return SzlSublist_ASSIGNMENT_BETWEEN_PROCESS_IMAGE_PARTITIONS_AND_OBS
        case 0x32:
            return SzlSublist_COMMUNICATION_STATUS_DATA
        case 0x74:
            return SzlSublist_STATUS_SINGLE_MODULE_LED
        case 0x90:
            return SzlSublist_DP_MASTER_SYSTEM_INFORMATION
        case 0x91:
            return SzlSublist_MODULE_STATUS_INFORMATION
        case 0x92:
            return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION
        case 0x94:
            return SzlSublist_RACK_OR_STATION_STATUS_INFORMATION_2
        case 0x95:
            return SzlSublist_ADDITIONAL_DP_MASTER_SYSTEM_OR_PROFINET_IO_SYSTEM_INFORMATION
        case 0x96:
            return SzlSublist_MODULE_STATUS_INFORMATION_PROFINET_IO_AND_PROFIBUS_DP
        case 0xA0:
            return SzlSublist_DIAGNOSTIC_BUFFER
        case 0xB1:
            return SzlSublist_MODULE_DIAGNOSTIC_DATA
    }
    return 0
}

func CastSzlSublist(structType interface{}) SzlSublist {
    castFunc := func(typ interface{}) SzlSublist {
        if sSzlSublist, ok := typ.(SzlSublist); ok {
            return sSzlSublist
        }
        return 0
    }
    return castFunc(structType)
}

func (m SzlSublist) LengthInBits() uint16 {
    return 8
}

func (m SzlSublist) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func SzlSublistParse(io *utils.ReadBuffer) (SzlSublist, error) {
    val, err := io.ReadUint8(8)
    if err != nil {
        return 0, nil
    }
    return SzlSublistValueOf(val), nil
}

func (e SzlSublist) Serialize(io utils.WriteBuffer) error {
    err := io.WriteUint8(8, uint8(e))
    return err
}
