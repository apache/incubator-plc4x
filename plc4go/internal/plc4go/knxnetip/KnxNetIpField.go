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
package knxnetip

import (
    "errors"
    driverModel "github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
    apiModel "github.com/apache/plc4x/plc4go/pkg/plc4go/model"
    "strconv"
    "strings"
)

type KnxNetIpField interface {
    IsPatternField() bool
    GetFieldType() *driverModel.KnxDatapointType
    matches(knxGroupAddress *driverModel.KnxGroupAddress) bool
    toGroupAddress() *driverModel.KnxGroupAddress
    apiModel.PlcField
}

type KnxNetIpGroupAddress3LevelPlcField struct {
    FieldType *driverModel.KnxDatapointType
    // 5 Bits: Values 0-31
    MainGroup string
    // 3 Bits: values 0-7
    MiddleGroup string
    // 8 Bits
    SubGroup string
    KnxNetIpField
}

func NewKnxNetIpGroupAddress3LevelPlcField(fieldType *driverModel.KnxDatapointType, mainGroup string, middleGroup string, subGroup string) KnxNetIpGroupAddress3LevelPlcField {
    return KnxNetIpGroupAddress3LevelPlcField{
        FieldType:   fieldType,
        MainGroup:   mainGroup,
        MiddleGroup: middleGroup,
        SubGroup:    subGroup,
    }
}

func (k KnxNetIpGroupAddress3LevelPlcField) GetTypeName() string {
    return k.FieldType.Name()
}

func (k KnxNetIpGroupAddress3LevelPlcField) GetFieldType() *driverModel.KnxDatapointType {
    return k.FieldType
}

func (k KnxNetIpGroupAddress3LevelPlcField) GetQuantity() uint16 {
    return 1
}

func (k KnxNetIpGroupAddress3LevelPlcField) IsPatternField() bool {
    _, err := strconv.Atoi(k.MainGroup)
    if err == nil {
        _, err = strconv.Atoi(k.MiddleGroup)
        if err == nil {
            _, err = strconv.Atoi(k.SubGroup)
            if err == nil {
                return false
            }
        }
    }
    return true
}

func (k KnxNetIpGroupAddress3LevelPlcField) matches(knxGroupAddress *driverModel.KnxGroupAddress) bool {
    level3KnxGroupAddress := driverModel.CastKnxGroupAddress3Level(knxGroupAddress)
    if level3KnxGroupAddress == nil {
        return false
    }
    return matches(k.MainGroup, strconv.Itoa(int(level3KnxGroupAddress.MainGroup))) &&
        matches(k.MiddleGroup, strconv.Itoa(int(level3KnxGroupAddress.MiddleGroup))) &&
        matches(k.SubGroup, strconv.Itoa(int(level3KnxGroupAddress.SubGroup)))
}

func (k KnxNetIpGroupAddress3LevelPlcField) toGroupAddress() *driverModel.KnxGroupAddress {
    mainGroup, err := strconv.Atoi(k.MainGroup)
    if err != nil {
        return nil
    }
    middleGroup, err := strconv.Atoi(k.MiddleGroup)
    if err != nil {
        return nil
    }
    subGroup, err := strconv.Atoi(k.SubGroup)
    if err != nil {
        return nil
    }
    ga := &driverModel.KnxGroupAddress{}
    l3 := &driverModel.KnxGroupAddress3Level{
        MainGroup:   uint8(mainGroup),
        MiddleGroup: uint8(middleGroup),
        SubGroup:    uint8(subGroup),
        Parent:      ga,
    }
    ga.Child = l3
    return ga
}

type KnxNetIpGroupAddress2LevelPlcField struct {
    FieldType *driverModel.KnxDatapointType
    // 5 Bits: Values 0-31
    MainGroup string
    // 11 Bits
    SubGroup string
    KnxNetIpField
}

func NewKnxNetIpGroupAddress2LevelPlcField(fieldType *driverModel.KnxDatapointType, mainGroup string, subGroup string) KnxNetIpGroupAddress2LevelPlcField {
    return KnxNetIpGroupAddress2LevelPlcField{
        FieldType: fieldType,
        MainGroup: mainGroup,
        SubGroup:  subGroup,
    }
}

func (k KnxNetIpGroupAddress2LevelPlcField) GetTypeName() string {
    return k.FieldType.Name()
}

func (k KnxNetIpGroupAddress2LevelPlcField) GetFieldType() *driverModel.KnxDatapointType {
    return k.FieldType
}

func (k KnxNetIpGroupAddress2LevelPlcField) GetQuantity() uint16 {
    return 1
}

func (k KnxNetIpGroupAddress2LevelPlcField) IsPatternField() bool {
    _, err := strconv.Atoi(k.MainGroup)
    if err == nil {
        _, err = strconv.Atoi(k.SubGroup)
        if err == nil {
            return false
        }
    }
    return true
}

func (k KnxNetIpGroupAddress2LevelPlcField) matches(knxGroupAddress *driverModel.KnxGroupAddress) bool {
    level2KnxGroupAddress := driverModel.CastKnxGroupAddress2Level(knxGroupAddress)
    if level2KnxGroupAddress == nil {
        return false
    }
    return matches(k.MainGroup, strconv.Itoa(int(level2KnxGroupAddress.MainGroup))) &&
        matches(k.SubGroup, strconv.Itoa(int(level2KnxGroupAddress.SubGroup)))
}

func (k KnxNetIpGroupAddress2LevelPlcField) toGroupAddress() *driverModel.KnxGroupAddress {
    mainGroup, err := strconv.Atoi(k.MainGroup)
    if err != nil {
        return nil
    }
    subGroup, err := strconv.Atoi(k.SubGroup)
    if err != nil {
        return nil
    }
    ga := &driverModel.KnxGroupAddress{}
    l3 := &driverModel.KnxGroupAddress2Level{
        MainGroup: uint8(mainGroup),
        SubGroup:  uint16(subGroup),
        Parent:    ga,
    }
    ga.Child = l3
    return ga
}

type KnxNetIpGroupAddress1LevelPlcField struct {
    FieldType *driverModel.KnxDatapointType
    // 16 Bits
    MainGroup string
    KnxNetIpField
}

func NewKnxNetIpGroupAddress1LevelPlcField(fieldType *driverModel.KnxDatapointType, mainGroup string) KnxNetIpGroupAddress1LevelPlcField {
    return KnxNetIpGroupAddress1LevelPlcField{
        FieldType: fieldType,
        MainGroup: mainGroup,
    }
}

func (k KnxNetIpGroupAddress1LevelPlcField) GetTypeName() string {
    return k.FieldType.Name()
}

func (k KnxNetIpGroupAddress1LevelPlcField) GetFieldType() *driverModel.KnxDatapointType {
    return k.FieldType
}

func (k KnxNetIpGroupAddress1LevelPlcField) GetQuantity() uint16 {
    return 1
}

func (k KnxNetIpGroupAddress1LevelPlcField) IsPatternField() bool {
    _, err := strconv.Atoi(k.MainGroup)
    if err == nil {
        return false
    }
    return true
}

func (k KnxNetIpGroupAddress1LevelPlcField) matches(knxGroupAddress *driverModel.KnxGroupAddress) bool {
    level1KnxGroupAddress := driverModel.CastKnxGroupAddressFreeLevel(knxGroupAddress)
    if level1KnxGroupAddress == nil {
        return false
    }
    return matches(k.MainGroup, strconv.Itoa(int(level1KnxGroupAddress.SubGroup)))
}

func (k KnxNetIpGroupAddress1LevelPlcField) toGroupAddress() *driverModel.KnxGroupAddress {
    mainGroup, err := strconv.Atoi(k.MainGroup)
    if err != nil {
        return nil
    }
    ga := &driverModel.KnxGroupAddress{}
    l3 := &driverModel.KnxGroupAddressFreeLevel{
        SubGroup: uint16(mainGroup),
        Parent:   ga,
    }
    ga.Child = l3
    return ga
}

type KnxNetIpDevicePropertyAddressPlcField struct {
    // 5 Bits: Values 0-31
    MainGroup string
    // 3 Bits: values 0-7
    MiddleGroup string
    // 8 Bits
    SubGroup   string
    ObjectId   string
    PropertyId string
    KnxNetIpField
}

func NewKnxNetIpDevicePropertyAddressPlcField(mainGroup string, middleGroup string, subGroup string, objectId string, propertyId string) KnxNetIpDevicePropertyAddressPlcField {
    return KnxNetIpDevicePropertyAddressPlcField{
        MainGroup:   mainGroup,
        MiddleGroup: middleGroup,
        SubGroup:    subGroup,
        ObjectId:    objectId,
        PropertyId:  propertyId,
    }
}

func (k KnxNetIpDevicePropertyAddressPlcField) GetTypeName() string {
	return ""
}

func (k KnxNetIpDevicePropertyAddressPlcField) GetFieldType() *driverModel.KnxDatapointType {
    return nil
}

func (k KnxNetIpDevicePropertyAddressPlcField) GetQuantity() uint16 {
	return 1
}

func (k KnxNetIpDevicePropertyAddressPlcField) IsPatternField() bool {
    _, err := strconv.Atoi(k.MainGroup)
    if err == nil {
        _, err = strconv.Atoi(k.MiddleGroup)
        if err == nil {
            _, err = strconv.Atoi(k.SubGroup)
            if err == nil {
                return false
            }
        }
    }
    return true
}

func (k KnxNetIpDevicePropertyAddressPlcField) matches(knxGroupAddress *driverModel.KnxGroupAddress) bool {
	level3KnxGroupAddress := driverModel.CastKnxGroupAddress3Level(knxGroupAddress)
	if level3KnxGroupAddress == nil {
		return false
	}
	return matches(k.MainGroup, strconv.Itoa(int(level3KnxGroupAddress.MainGroup))) &&
		matches(k.MiddleGroup, strconv.Itoa(int(level3KnxGroupAddress.MiddleGroup))) &&
		matches(k.SubGroup, strconv.Itoa(int(level3KnxGroupAddress.SubGroup)))
}

func (k KnxNetIpDevicePropertyAddressPlcField) toKnxAddress() *driverModel.KnxAddress {
    mainGroup, err := strconv.Atoi(k.MainGroup)
    if err != nil {
        return nil
    }
    middleGroup, err := strconv.Atoi(k.MiddleGroup)
    if err != nil {
        return nil
    }
    subGroup, err := strconv.Atoi(k.SubGroup)
    if err != nil {
        return nil
    }
    ga := &driverModel.KnxAddress{
        MainGroup:   uint8(mainGroup),
        MiddleGroup: uint8(middleGroup),
        SubGroup:    uint8(subGroup),
    }
    return ga
}

type KnxNetIpDeviceMemoryAddressPlcField struct {
    FieldType *driverModel.KnxDatapointType
    // 5 Bits: Values 0-31
    MainGroup uint8
    // 3 Bits: values 0-7
    MiddleGroup uint8
    // 8 Bits
    SubGroup         uint8
    Address          uint16
    NumberOfElements uint8
    KnxNetIpField
}

func NewKnxNetIpDeviceMemoryAddressPlcField(fieldType *driverModel.KnxDatapointType, mainGroup uint8, middleGroup uint8, subGroup uint8, address uint16, numberOfElements uint8) KnxNetIpDeviceMemoryAddressPlcField {
    return KnxNetIpDeviceMemoryAddressPlcField{
        FieldType:        fieldType,
        MainGroup:        mainGroup,
        MiddleGroup:      middleGroup,
        SubGroup:         subGroup,
        Address:          address,
        NumberOfElements: numberOfElements,
    }
}

func (k KnxNetIpDeviceMemoryAddressPlcField) GetTypeName() string {
    return k.FieldType.Name()
}

func (k KnxNetIpDeviceMemoryAddressPlcField) GetFieldType() *driverModel.KnxDatapointType {
    return k.FieldType
}

func (k KnxNetIpDeviceMemoryAddressPlcField) GetQuantity() uint16 {
    return uint16(k.NumberOfElements)
}

func (k KnxNetIpDeviceMemoryAddressPlcField) IsPatternField() bool {
    return false
}

func (k KnxNetIpDeviceMemoryAddressPlcField) matches(knxGroupAddress *driverModel.KnxGroupAddress) bool {
    return false
}

func (k KnxNetIpDeviceMemoryAddressPlcField) toKnxAddress() *driverModel.KnxAddress {
    ga := &driverModel.KnxAddress{
        MainGroup:   k.MainGroup,
        MiddleGroup: k.MiddleGroup,
        SubGroup:    k.SubGroup,
    }
    return ga
}

func CastToKnxNetIpFieldFromPlcField(plcField apiModel.PlcField) (KnxNetIpField, error) {
    if knxNetIpField, ok := plcField.(KnxNetIpField); ok {
        return knxNetIpField, nil
    }
    return nil, errors.New("couldn't cast to KnxNetIpField")
}

func matches(pattern string, groupAddressPart string) bool {
    // A "*" simply matches everything
    if pattern == "*" {
        return true
    }
    // If the pattern starts and ends with square brackets, it's a list of values or range queries
    if strings.HasPrefix(pattern, "[") && strings.HasSuffix(pattern, "]") {
        matches := false
        for _, segment := range strings.Split(pattern, ",") {
            if strings.Contains(segment, "-") {
                // If the segment contains a "-", then it's a range query
                split := strings.Split(segment, "-")
                if len(split) == 2 {
                    if val, err := strconv.Atoi(groupAddressPart); err != nil {
                        var err error
                        var from int
                        if from, err = strconv.Atoi(split[0]); err != nil {
                            continue
                        }
                        if val < from {
                            continue
                        }
                        var to int
                        if to, err = strconv.Atoi(split[1]); err == nil {
                            continue
                        }
                        if val > to {
                            continue
                        }
                        matches = true
                    }
                }
            } else if segment == groupAddressPart {
                // In all other cases it's an explicit value
                matches = true
            }
        }
        return matches
    } else {
        return pattern == groupAddressPart
    }
}
