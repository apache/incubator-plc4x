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
type COTPParameterChecksum struct {
    Crc uint8
    COTPParameter
}

// The corresponding interface
type ICOTPParameterChecksum interface {
    ICOTPParameter
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m COTPParameterChecksum) ParameterType() uint8 {
    return 0xC3
}

func (m COTPParameterChecksum) initialize() spi.Message {
    return m
}

func NewCOTPParameterChecksum(crc uint8) COTPParameterInitializer {
    return &COTPParameterChecksum{Crc: crc}
}

func CastICOTPParameterChecksum(structType interface{}) ICOTPParameterChecksum {
    castFunc := func(typ interface{}) ICOTPParameterChecksum {
        if iCOTPParameterChecksum, ok := typ.(ICOTPParameterChecksum); ok {
            return iCOTPParameterChecksum
        }
        return nil
    }
    return castFunc(structType)
}

func CastCOTPParameterChecksum(structType interface{}) COTPParameterChecksum {
    castFunc := func(typ interface{}) COTPParameterChecksum {
        if sCOTPParameterChecksum, ok := typ.(COTPParameterChecksum); ok {
            return sCOTPParameterChecksum
        }
        if sCOTPParameterChecksum, ok := typ.(*COTPParameterChecksum); ok {
            return *sCOTPParameterChecksum
        }
        return COTPParameterChecksum{}
    }
    return castFunc(structType)
}

func (m COTPParameterChecksum) LengthInBits() uint16 {
    var lengthInBits uint16 = m.COTPParameter.LengthInBits()

    // Simple field (crc)
    lengthInBits += 8

    return lengthInBits
}

func (m COTPParameterChecksum) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPParameterChecksumParse(io *utils.ReadBuffer) (COTPParameterInitializer, error) {

    // Simple Field (crc)
    crc, _crcErr := io.ReadUint8(8)
    if _crcErr != nil {
        return nil, errors.New("Error parsing 'crc' field " + _crcErr.Error())
    }

    // Create the instance
    return NewCOTPParameterChecksum(crc), nil
}

func (m COTPParameterChecksum) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (crc)
    crc := uint8(m.Crc)
    _crcErr := io.WriteUint8(8, (crc))
    if _crcErr != nil {
        return errors.New("Error serializing 'crc' field " + _crcErr.Error())
    }

        return nil
    }
    return COTPParameterSerialize(io, m.COTPParameter, CastICOTPParameter(m), ser)
}

func (m *COTPParameterChecksum) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "crc":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Crc = data
            }
        }
    }
}

func (m COTPParameterChecksum) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.s7.readwrite.COTPParameterChecksum"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

