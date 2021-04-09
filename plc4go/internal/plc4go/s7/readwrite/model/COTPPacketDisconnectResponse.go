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
type COTPPacketDisconnectResponse struct {
    DestinationReference uint16
    SourceReference uint16
    Parent *COTPPacket
    ICOTPPacketDisconnectResponse
}

// The corresponding interface
type ICOTPPacketDisconnectResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketDisconnectResponse) TpduCode() uint8 {
    return 0xC0
}


func (m *COTPPacketDisconnectResponse) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
    m.Parent.Parameters = parameters
    m.Parent.Payload = payload
}

func NewCOTPPacketDisconnectResponse(destinationReference uint16, sourceReference uint16, parameters []*COTPParameter, payload *S7Message) *COTPPacket {
    child := &COTPPacketDisconnectResponse{
        DestinationReference: destinationReference,
        SourceReference: sourceReference,
        Parent: NewCOTPPacket(parameters, payload),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastCOTPPacketDisconnectResponse(structType interface{}) COTPPacketDisconnectResponse {
    castFunc := func(typ interface{}) COTPPacketDisconnectResponse {
        if casted, ok := typ.(COTPPacketDisconnectResponse); ok {
            return casted
        }
        if casted, ok := typ.(*COTPPacketDisconnectResponse); ok {
            return *casted
        }
        if casted, ok := typ.(COTPPacket); ok {
            return CastCOTPPacketDisconnectResponse(casted.Child)
        }
        if casted, ok := typ.(*COTPPacket); ok {
            return CastCOTPPacketDisconnectResponse(casted.Child)
        }
        return COTPPacketDisconnectResponse{}
    }
    return castFunc(structType)
}

func (m *COTPPacketDisconnectResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (destinationReference)
    lengthInBits += 16

    // Simple field (sourceReference)
    lengthInBits += 16

    return lengthInBits
}

func (m *COTPPacketDisconnectResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func COTPPacketDisconnectResponseParse(io *utils.ReadBuffer) (*COTPPacket, error) {

    // Simple Field (destinationReference)
    destinationReference, _destinationReferenceErr := io.ReadUint16(16)
    if _destinationReferenceErr != nil {
        return nil, errors.New("Error parsing 'destinationReference' field " + _destinationReferenceErr.Error())
    }

    // Simple Field (sourceReference)
    sourceReference, _sourceReferenceErr := io.ReadUint16(16)
    if _sourceReferenceErr != nil {
        return nil, errors.New("Error parsing 'sourceReference' field " + _sourceReferenceErr.Error())
    }

    // Create a partially initialized instance
    _child := &COTPPacketDisconnectResponse{
        DestinationReference: destinationReference,
        SourceReference: sourceReference,
        Parent: &COTPPacket{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *COTPPacketDisconnectResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (destinationReference)
    destinationReference := uint16(m.DestinationReference)
    _destinationReferenceErr := io.WriteUint16(16, (destinationReference))
    if _destinationReferenceErr != nil {
        return errors.New("Error serializing 'destinationReference' field " + _destinationReferenceErr.Error())
    }

    // Simple Field (sourceReference)
    sourceReference := uint16(m.SourceReference)
    _sourceReferenceErr := io.WriteUint16(16, (sourceReference))
    if _sourceReferenceErr != nil {
        return errors.New("Error serializing 'sourceReference' field " + _sourceReferenceErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPPacketDisconnectResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "destinationReference":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.DestinationReference = data
            case "sourceReference":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SourceReference = data
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

func (m *COTPPacketDisconnectResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.DestinationReference, xml.StartElement{Name: xml.Name{Local: "destinationReference"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SourceReference, xml.StartElement{Name: xml.Name{Local: "sourceReference"}}); err != nil {
        return err
    }
    return nil
}

