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
    log "github.com/sirupsen/logrus"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

// The data-structure of this message
type APDUSegmentAck struct {
    NegativeAck bool
    Server bool
    OriginalInvokeId uint8
    SequenceNumber uint8
    ProposedWindowSize uint8
    Parent *APDU
    IAPDUSegmentAck
}

// The corresponding interface
type IAPDUSegmentAck interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUSegmentAck) ApduType() uint8 {
    return 0x4
}


func (m *APDUSegmentAck) InitializeParent(parent *APDU) {
}

func NewAPDUSegmentAck(negativeAck bool, server bool, originalInvokeId uint8, sequenceNumber uint8, proposedWindowSize uint8, ) *APDU {
    child := &APDUSegmentAck{
        NegativeAck: negativeAck,
        Server: server,
        OriginalInvokeId: originalInvokeId,
        SequenceNumber: sequenceNumber,
        ProposedWindowSize: proposedWindowSize,
        Parent: NewAPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAPDUSegmentAck(structType interface{}) *APDUSegmentAck {
    castFunc := func(typ interface{}) *APDUSegmentAck {
        if casted, ok := typ.(APDUSegmentAck); ok {
            return &casted
        }
        if casted, ok := typ.(*APDUSegmentAck); ok {
            return casted
        }
        if casted, ok := typ.(APDU); ok {
            return CastAPDUSegmentAck(casted.Child)
        }
        if casted, ok := typ.(*APDU); ok {
            return CastAPDUSegmentAck(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *APDUSegmentAck) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Reserved Field (reserved)
    lengthInBits += 2

    // Simple field (negativeAck)
    lengthInBits += 1

    // Simple field (server)
    lengthInBits += 1

    // Simple field (originalInvokeId)
    lengthInBits += 8

    // Simple field (sequenceNumber)
    lengthInBits += 8

    // Simple field (proposedWindowSize)
    lengthInBits += 8

    return lengthInBits
}

func (m *APDUSegmentAck) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func APDUSegmentAckParse(io *utils.ReadBuffer) (*APDU, error) {

    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
        reserved, _err := io.ReadUint8(2)
        if _err != nil {
            return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
        }
        if reserved != uint8(0x00) {
            log.WithFields(log.Fields{
                "expected value": uint8(0x00),
                "got value": reserved,
            }).Info("Got unexpected response.")
        }
    }

    // Simple Field (negativeAck)
    negativeAck, _negativeAckErr := io.ReadBit()
    if _negativeAckErr != nil {
        return nil, errors.New("Error parsing 'negativeAck' field " + _negativeAckErr.Error())
    }

    // Simple Field (server)
    server, _serverErr := io.ReadBit()
    if _serverErr != nil {
        return nil, errors.New("Error parsing 'server' field " + _serverErr.Error())
    }

    // Simple Field (originalInvokeId)
    originalInvokeId, _originalInvokeIdErr := io.ReadUint8(8)
    if _originalInvokeIdErr != nil {
        return nil, errors.New("Error parsing 'originalInvokeId' field " + _originalInvokeIdErr.Error())
    }

    // Simple Field (sequenceNumber)
    sequenceNumber, _sequenceNumberErr := io.ReadUint8(8)
    if _sequenceNumberErr != nil {
        return nil, errors.New("Error parsing 'sequenceNumber' field " + _sequenceNumberErr.Error())
    }

    // Simple Field (proposedWindowSize)
    proposedWindowSize, _proposedWindowSizeErr := io.ReadUint8(8)
    if _proposedWindowSizeErr != nil {
        return nil, errors.New("Error parsing 'proposedWindowSize' field " + _proposedWindowSizeErr.Error())
    }

    // Create a partially initialized instance
    _child := &APDUSegmentAck{
        NegativeAck: negativeAck,
        Server: server,
        OriginalInvokeId: originalInvokeId,
        SequenceNumber: sequenceNumber,
        ProposedWindowSize: proposedWindowSize,
        Parent: &APDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *APDUSegmentAck) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Reserved Field (reserved)
    {
        _err := io.WriteUint8(2, uint8(0x00))
        if _err != nil {
            return errors.New("Error serializing 'reserved' field " + _err.Error())
        }
    }

    // Simple Field (negativeAck)
    negativeAck := bool(m.NegativeAck)
    _negativeAckErr := io.WriteBit((negativeAck))
    if _negativeAckErr != nil {
        return errors.New("Error serializing 'negativeAck' field " + _negativeAckErr.Error())
    }

    // Simple Field (server)
    server := bool(m.Server)
    _serverErr := io.WriteBit((server))
    if _serverErr != nil {
        return errors.New("Error serializing 'server' field " + _serverErr.Error())
    }

    // Simple Field (originalInvokeId)
    originalInvokeId := uint8(m.OriginalInvokeId)
    _originalInvokeIdErr := io.WriteUint8(8, (originalInvokeId))
    if _originalInvokeIdErr != nil {
        return errors.New("Error serializing 'originalInvokeId' field " + _originalInvokeIdErr.Error())
    }

    // Simple Field (sequenceNumber)
    sequenceNumber := uint8(m.SequenceNumber)
    _sequenceNumberErr := io.WriteUint8(8, (sequenceNumber))
    if _sequenceNumberErr != nil {
        return errors.New("Error serializing 'sequenceNumber' field " + _sequenceNumberErr.Error())
    }

    // Simple Field (proposedWindowSize)
    proposedWindowSize := uint8(m.ProposedWindowSize)
    _proposedWindowSizeErr := io.WriteUint8(8, (proposedWindowSize))
    if _proposedWindowSizeErr != nil {
        return errors.New("Error serializing 'proposedWindowSize' field " + _proposedWindowSizeErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *APDUSegmentAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "negativeAck":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.NegativeAck = data
            case "server":
                var data bool
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Server = data
            case "originalInvokeId":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.OriginalInvokeId = data
            case "sequenceNumber":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.SequenceNumber = data
            case "proposedWindowSize":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ProposedWindowSize = data
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

func (m *APDUSegmentAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeElement(m.NegativeAck, xml.StartElement{Name: xml.Name{Local: "negativeAck"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Server, xml.StartElement{Name: xml.Name{Local: "server"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.OriginalInvokeId, xml.StartElement{Name: xml.Name{Local: "originalInvokeId"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.SequenceNumber, xml.StartElement{Name: xml.Name{Local: "sequenceNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ProposedWindowSize, xml.StartElement{Name: xml.Name{Local: "proposedWindowSize"}}); err != nil {
        return err
    }
    return nil
}

