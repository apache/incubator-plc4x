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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type COTPParameterCalledTsap struct {
    TsapId uint16
    COTPParameter
}

// The corresponding interface
type ICOTPParameterCalledTsap interface {
    ICOTPParameter
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m COTPParameterCalledTsap) ParameterType() uint8 {
    return 0xC2
}

func (m COTPParameterCalledTsap) initialize() spi.Message {
    return m
}

func NewCOTPParameterCalledTsap(tsapId uint16) COTPParameterInitializer {
    return &COTPParameterCalledTsap{TsapId: tsapId}
}

func CastICOTPParameterCalledTsap(structType interface{}) ICOTPParameterCalledTsap {
    castFunc := func(typ interface{}) ICOTPParameterCalledTsap {
        if iCOTPParameterCalledTsap, ok := typ.(ICOTPParameterCalledTsap); ok {
            return iCOTPParameterCalledTsap
        }
        return nil
    }
    return castFunc(structType)
}

func CastCOTPParameterCalledTsap(structType interface{}) COTPParameterCalledTsap {
    castFunc := func(typ interface{}) COTPParameterCalledTsap {
        if sCOTPParameterCalledTsap, ok := typ.(COTPParameterCalledTsap); ok {
            return sCOTPParameterCalledTsap
        }
        if sCOTPParameterCalledTsap, ok := typ.(*COTPParameterCalledTsap); ok {
            return *sCOTPParameterCalledTsap
        }
        return COTPParameterCalledTsap{}
    }
    return castFunc(structType)
}

func (m COTPParameterCalledTsap) LengthInBits() uint16 {
    var lengthInBits uint16 = m.COTPParameter.LengthInBits()

    // Simple field (tsapId)
    lengthInBits += 16

    return lengthInBits
}

func (m COTPParameterCalledTsap) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPParameterCalledTsapParse(io *utils.ReadBuffer) (COTPParameterInitializer, error) {

    // Simple Field (tsapId)
    tsapId, _tsapIdErr := io.ReadUint16(16)
    if _tsapIdErr != nil {
        return nil, errors.New("Error parsing 'tsapId' field " + _tsapIdErr.Error())
    }

    // Create the instance
    return NewCOTPParameterCalledTsap(tsapId), nil
}

func (m COTPParameterCalledTsap) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (tsapId)
    tsapId := uint16(m.TsapId)
    _tsapIdErr := io.WriteUint16(16, (tsapId))
    if _tsapIdErr != nil {
        return errors.New("Error serializing 'tsapId' field " + _tsapIdErr.Error())
    }

        return nil
    }
    return COTPParameterSerialize(io, m.COTPParameter, CastICOTPParameter(m), ser)
}

func (m *COTPParameterCalledTsap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "tsapId":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.TsapId = data
            }
        }
    }
}

func (m COTPParameterCalledTsap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.s7.readwrite.COTPParameterCalledTsap"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.TsapId, xml.StartElement{Name: xml.Name{Local: "tsapId"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

