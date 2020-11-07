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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type COTPParameterDisconnectAdditionalInformation struct {
    Data []uint8
    Parent *COTPParameter
    ICOTPParameterDisconnectAdditionalInformation
}

// The corresponding interface
type ICOTPParameterDisconnectAdditionalInformation interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPParameterDisconnectAdditionalInformation) ParameterType() uint8 {
    return 0xE0
}


func (m *COTPParameterDisconnectAdditionalInformation) InitializeParent(parent *COTPParameter) {
}

func NewCOTPParameterDisconnectAdditionalInformation(data []uint8, ) *COTPParameter {
    child := &COTPParameterDisconnectAdditionalInformation{
        Data: data,
        Parent: NewCOTPParameter(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastCOTPParameterDisconnectAdditionalInformation(structType interface{}) COTPParameterDisconnectAdditionalInformation {
    castFunc := func(typ interface{}) COTPParameterDisconnectAdditionalInformation {
        if casted, ok := typ.(COTPParameterDisconnectAdditionalInformation); ok {
            return casted
        }
        if casted, ok := typ.(*COTPParameterDisconnectAdditionalInformation); ok {
            return *casted
        }
        if casted, ok := typ.(COTPParameter); ok {
            return CastCOTPParameterDisconnectAdditionalInformation(casted.Child)
        }
        if casted, ok := typ.(*COTPParameter); ok {
            return CastCOTPParameterDisconnectAdditionalInformation(casted.Child)
        }
        return COTPParameterDisconnectAdditionalInformation{}
    }
    return castFunc(structType)
}

func (m *COTPParameterDisconnectAdditionalInformation) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m *COTPParameterDisconnectAdditionalInformation) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPParameterDisconnectAdditionalInformationParse(io *utils.ReadBuffer, rest uint8) (*COTPParameter, error) {

    // Array field (data)
    // Count array
    data := make([]uint8, rest)
    for curItem := uint16(0); curItem < uint16(rest); curItem++ {
        _item, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &COTPParameterDisconnectAdditionalInformation{
        Data: data,
        Parent: &COTPParameter{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *COTPParameterDisconnectAdditionalInformation) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteUint8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPParameterDisconnectAdditionalInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "data":
                var data []uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Data = data
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

func (m *COTPParameterDisconnectAdditionalInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

