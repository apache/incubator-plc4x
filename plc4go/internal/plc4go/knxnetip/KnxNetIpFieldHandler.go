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
    model2 "github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi"
    "github.com/apache/plc4x/plc4go/internal/plc4go/utils"
    "github.com/apache/plc4x/plc4go/pkg/plc4go/model"
    "regexp"
)

type KnxNetIpFieldType uint8

const (
    KNX_NET_IP_FIELD_COIL KnxNetIpFieldType = 0x00
)

func (m KnxNetIpFieldType) GetName() string {
    switch m {
    case KNX_NET_IP_FIELD_COIL:
        return "ModbusFieldHoldingRegister"
    }
    return ""
}

type FieldHandler struct {
    knxNetIpGroupAddress3Level *regexp.Regexp
    knxNetIpGroupAddress2Level *regexp.Regexp
    knxNetIpGroupAddress1Level *regexp.Regexp
    spi.PlcFieldHandler
}

func NewFieldHandler() FieldHandler {
    return FieldHandler{
        knxNetIpGroupAddress3Level: regexp.MustCompile(`^(?P<mainGroup>(\d{1,2}|\*|\[(\d{1,2}|\d{1,2}\-\d{1,2})(,(\d{1,2}|\d{1,2}\-\d{1,2}))*]))\/(?P<middleGroup>(\d|\*|\[(\d|\d\-\d)(,(\d|\d\-\d))*]))\/(?P<subGroup>(\d{1,3}|\*|\[(\d{1,3}|\d{1,3}\-\d{1,3})(,(\d{1,3}|\d{1,3}\-\d{1,3}))*]))(:(?P<datatype>[a-zA-Z_]+))?`),
        knxNetIpGroupAddress2Level: regexp.MustCompile(`^(?P<mainGroup>(\d{1,2}|\*|\[(\d{1,2}|\d{1,2}\-\d{1,2})(,(\d{1,2}|\d{1,2}\-\d{1,2}))*]))/(?P<subGroup>(\d{1,4}|\*|\[(\d{1,4}|\d{1,4}\-\d{1,4})(,(\d{1,4}|\d{1,4}\-\d{1,4}))*]))(:(?P<datatype>[a-zA-Z_]+))?`),
        knxNetIpGroupAddress1Level: regexp.MustCompile(`^(?P<mainGroup>(\d{1,5}|\*|\[(\d{1,5}|\d{1,5}\-\d{1,5})(,(\d{1,5}|\d{1,5}\-\d{1,5}))*]))(:(?P<datatype>[a-zA-Z_]+))?`),
    }
}

func (m FieldHandler) ParseQuery(query string) (model.PlcField, error) {
    if match := utils.GetSubgropMatches(m.knxNetIpGroupAddress3Level, query); match != nil {
        fieldTypeName, ok := match["datatype"]
        var fieldType model2.KnxDatapointType
        if ok {
            fieldType = model2.KnxDatapointTypeValueOf(fieldTypeName)
        }
        return NewKnxNetIpGroupAddress3LevelPlcField(&fieldType, match["mainGroup"], match["middleGroup"], match["subGroup"]), nil
    } else if match := utils.GetSubgropMatches(m.knxNetIpGroupAddress2Level, query); match != nil {
        fieldTypeName, ok := match["datatype"]
        var fieldType model2.KnxDatapointType
        if ok {
            fieldType = model2.KnxDatapointTypeValueOf(fieldTypeName)
        }
        return NewKnxNetIpGroupAddress2LevelPlcField(&fieldType, match["mainGroup"], match["subGroup"]), nil
    } else if match := utils.GetSubgropMatches(m.knxNetIpGroupAddress1Level, query); match != nil {
        fieldTypeName, ok := match["datatype"]
        var fieldType model2.KnxDatapointType
        if ok {
            fieldType = model2.KnxDatapointTypeValueOf(fieldTypeName)
        }
        return NewKnxNetIpGroupAddress1LevelPlcField(&fieldType, match["mainGroup"]), nil
    }
    return nil, errors.New("Invalid address format for address '" + query + "'")
}
