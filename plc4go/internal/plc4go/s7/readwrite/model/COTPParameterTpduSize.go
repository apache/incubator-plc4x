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
    "errors"
    "io"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

// The data-structure of this message
type COTPParameterTpduSize struct {
    TpduSize COTPTpduSize
    Parent *COTPParameter
    ICOTPParameterTpduSize
}

// The corresponding interface
type ICOTPParameterTpduSize interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPParameterTpduSize) ParameterType() uint8 {
    return 0xC0
}


func (m *COTPParameterTpduSize) InitializeParent(parent *COTPParameter) {
}

func NewCOTPParameterTpduSize(tpduSize COTPTpduSize, ) *COTPParameter {
    child := &COTPParameterTpduSize{
        TpduSize: tpduSize,
        Parent: NewCOTPParameter(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastCOTPParameterTpduSize(structType interface{}) COTPParameterTpduSize {
    castFunc := func(typ interface{}) COTPParameterTpduSize {
        if casted, ok := typ.(COTPParameterTpduSize); ok {
            return casted
        }
        if casted, ok := typ.(*COTPParameterTpduSize); ok {
            return *casted
        }
        if casted, ok := typ.(COTPParameter); ok {
            return CastCOTPParameterTpduSize(casted.Child)
        }
        if casted, ok := typ.(*COTPParameter); ok {
            return CastCOTPParameterTpduSize(casted.Child)
        }
        return COTPParameterTpduSize{}
    }
    return castFunc(structType)
}

func (m *COTPParameterTpduSize) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Enum Field (tpduSize)
    lengthInBits += 8

    return lengthInBits
}

func (m *COTPParameterTpduSize) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPParameterTpduSizeParse(io *utils.ReadBuffer) (*COTPParameter, error) {

    // Enum field (tpduSize)
    tpduSize, _tpduSizeErr := COTPTpduSizeParse(io)
    if _tpduSizeErr != nil {
        return nil, errors.New("Error parsing 'tpduSize' field " + _tpduSizeErr.Error())
    }

    // Create a partially initialized instance
    _child := &COTPParameterTpduSize{
        TpduSize: tpduSize,
        Parent: &COTPParameter{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *COTPParameterTpduSize) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Enum field (tpduSize)
    tpduSize := CastCOTPTpduSize(m.TpduSize)
    _tpduSizeErr := tpduSize.Serialize(io)
    if _tpduSizeErr != nil {
        return errors.New("Error serializing 'tpduSize' field " + _tpduSizeErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPParameterTpduSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "tpduSize":
                var data COTPTpduSize
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.TpduSize = data
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

func (m *COTPParameterTpduSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.TpduSize, xml.StartElement{Name: xml.Name{Local: "tpduSize"}}); err != nil {
        return err
    }
    return nil
}

