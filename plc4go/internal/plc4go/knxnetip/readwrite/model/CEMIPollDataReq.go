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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type CEMIPollDataReq struct {
    Parent *CEMI
    ICEMIPollDataReq
}

// The corresponding interface
type ICEMIPollDataReq interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CEMIPollDataReq) MessageCode() uint8 {
    return 0x13
}


func (m *CEMIPollDataReq) InitializeParent(parent *CEMI) {
}

func NewCEMIPollDataReq() *CEMI {
    child := &CEMIPollDataReq{
        Parent: NewCEMI(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastCEMIPollDataReq(structType interface{}) CEMIPollDataReq {
    castFunc := func(typ interface{}) CEMIPollDataReq {
        if casted, ok := typ.(CEMIPollDataReq); ok {
            return casted
        }
        if casted, ok := typ.(*CEMIPollDataReq); ok {
            return *casted
        }
        if casted, ok := typ.(CEMI); ok {
            return CastCEMIPollDataReq(casted.Child)
        }
        if casted, ok := typ.(*CEMI); ok {
            return CastCEMIPollDataReq(casted.Child)
        }
        return CEMIPollDataReq{}
    }
    return castFunc(structType)
}

func (m *CEMIPollDataReq) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *CEMIPollDataReq) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIPollDataReqParse(io *utils.ReadBuffer) (*CEMI, error) {

    // Create a partially initialized instance
    _child := &CEMIPollDataReq{
        Parent: &CEMI{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *CEMIPollDataReq) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *CEMIPollDataReq) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
        if err != nil {
            if err == io.EOF {
                return nil
            }
            return err
        }
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            }
        }
    }
}

func (m *CEMIPollDataReq) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.CEMIPollDataReq"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

