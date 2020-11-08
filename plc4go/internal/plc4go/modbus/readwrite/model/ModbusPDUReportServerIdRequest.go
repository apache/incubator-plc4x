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
    "encoding/xml"
    "io"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUReportServerIdRequest struct {
    Parent *ModbusPDU
    IModbusPDUReportServerIdRequest
}

// The corresponding interface
type IModbusPDUReportServerIdRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReportServerIdRequest) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUReportServerIdRequest) FunctionFlag() uint8 {
    return 0x11
}

func (m *ModbusPDUReportServerIdRequest) Response() bool {
    return false
}


func (m *ModbusPDUReportServerIdRequest) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUReportServerIdRequest() *ModbusPDU {
    child := &ModbusPDUReportServerIdRequest{
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUReportServerIdRequest(structType interface{}) ModbusPDUReportServerIdRequest {
    castFunc := func(typ interface{}) ModbusPDUReportServerIdRequest {
        if casted, ok := typ.(ModbusPDUReportServerIdRequest); ok {
            return casted
        }
        if casted, ok := typ.(*ModbusPDUReportServerIdRequest); ok {
            return *casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUReportServerIdRequest(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUReportServerIdRequest(casted.Child)
        }
        return ModbusPDUReportServerIdRequest{}
    }
    return castFunc(structType)
}

func (m *ModbusPDUReportServerIdRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *ModbusPDUReportServerIdRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUReportServerIdRequestParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Create a partially initialized instance
    _child := &ModbusPDUReportServerIdRequest{
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUReportServerIdRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUReportServerIdRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
        token, err = d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
    }
}

func (m *ModbusPDUReportServerIdRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

