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
type S7MessageResponse struct {
    ErrorClass uint8
    ErrorCode uint8
    S7Message
}

// The corresponding interface
type IS7MessageResponse interface {
    IS7Message
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m S7MessageResponse) MessageType() uint8 {
    return 0x02
}

func (m S7MessageResponse) initialize(tpduReference uint16, parameter *IS7Parameter, payload *IS7Payload) spi.Message {
    m.TpduReference = tpduReference
    m.Parameter = parameter
    m.Payload = payload
    return m
}

func NewS7MessageResponse(errorClass uint8, errorCode uint8) S7MessageInitializer {
    return &S7MessageResponse{ErrorClass: errorClass, ErrorCode: errorCode}
}

func CastIS7MessageResponse(structType interface{}) IS7MessageResponse {
    castFunc := func(typ interface{}) IS7MessageResponse {
        if iS7MessageResponse, ok := typ.(IS7MessageResponse); ok {
            return iS7MessageResponse
        }
        return nil
    }
    return castFunc(structType)
}

func CastS7MessageResponse(structType interface{}) S7MessageResponse {
    castFunc := func(typ interface{}) S7MessageResponse {
        if sS7MessageResponse, ok := typ.(S7MessageResponse); ok {
            return sS7MessageResponse
        }
        if sS7MessageResponse, ok := typ.(*S7MessageResponse); ok {
            return *sS7MessageResponse
        }
        return S7MessageResponse{}
    }
    return castFunc(structType)
}

func (m S7MessageResponse) LengthInBits() uint16 {
    var lengthInBits uint16 = m.S7Message.LengthInBits()

    // Simple field (errorClass)
    lengthInBits += 8

    // Simple field (errorCode)
    lengthInBits += 8

    return lengthInBits
}

func (m S7MessageResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func S7MessageResponseParse(io *utils.ReadBuffer) (S7MessageInitializer, error) {

    // Simple Field (errorClass)
    errorClass, _errorClassErr := io.ReadUint8(8)
    if _errorClassErr != nil {
        return nil, errors.New("Error parsing 'errorClass' field " + _errorClassErr.Error())
    }

    // Simple Field (errorCode)
    errorCode, _errorCodeErr := io.ReadUint8(8)
    if _errorCodeErr != nil {
        return nil, errors.New("Error parsing 'errorCode' field " + _errorCodeErr.Error())
    }

    // Create the instance
    return NewS7MessageResponse(errorClass, errorCode), nil
}

func (m S7MessageResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (errorClass)
    errorClass := uint8(m.ErrorClass)
    _errorClassErr := io.WriteUint8(8, (errorClass))
    if _errorClassErr != nil {
        return errors.New("Error serializing 'errorClass' field " + _errorClassErr.Error())
    }

    // Simple Field (errorCode)
    errorCode := uint8(m.ErrorCode)
    _errorCodeErr := io.WriteUint8(8, (errorCode))
    if _errorCodeErr != nil {
        return errors.New("Error serializing 'errorCode' field " + _errorCodeErr.Error())
    }

        return nil
    }
    return S7MessageSerialize(io, m.S7Message, CastIS7Message(m), ser)
}

func (m *S7MessageResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "errorClass":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ErrorClass = data
            case "errorCode":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ErrorCode = data
            }
        }
    }
}

func (m S7MessageResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.s7.readwrite.S7MessageResponse"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ErrorClass, xml.StartElement{Name: xml.Name{Local: "errorClass"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ErrorCode, xml.StartElement{Name: xml.Name{Local: "errorCode"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

