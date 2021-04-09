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
    "encoding/base64"
    "encoding/xml"
    "errors"
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUGetComEventLogResponse struct {
    Status uint16
    EventCount uint16
    MessageCount uint16
    Events []int8
    Parent *ModbusPDU
    IModbusPDUGetComEventLogResponse
}

// The corresponding interface
type IModbusPDUGetComEventLogResponse interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUGetComEventLogResponse) ErrorFlag() bool {
    return false
}

func (m *ModbusPDUGetComEventLogResponse) FunctionFlag() uint8 {
    return 0x0C
}

func (m *ModbusPDUGetComEventLogResponse) Response() bool {
    return true
}


func (m *ModbusPDUGetComEventLogResponse) InitializeParent(parent *ModbusPDU) {
}

func NewModbusPDUGetComEventLogResponse(status uint16, eventCount uint16, messageCount uint16, events []int8, ) *ModbusPDU {
    child := &ModbusPDUGetComEventLogResponse{
        Status: status,
        EventCount: eventCount,
        MessageCount: messageCount,
        Events: events,
        Parent: NewModbusPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastModbusPDUGetComEventLogResponse(structType interface{}) ModbusPDUGetComEventLogResponse {
    castFunc := func(typ interface{}) ModbusPDUGetComEventLogResponse {
        if casted, ok := typ.(ModbusPDUGetComEventLogResponse); ok {
            return casted
        }
        if casted, ok := typ.(*ModbusPDUGetComEventLogResponse); ok {
            return *casted
        }
        if casted, ok := typ.(ModbusPDU); ok {
            return CastModbusPDUGetComEventLogResponse(casted.Child)
        }
        if casted, ok := typ.(*ModbusPDU); ok {
            return CastModbusPDUGetComEventLogResponse(casted.Child)
        }
        return ModbusPDUGetComEventLogResponse{}
    }
    return castFunc(structType)
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Implicit Field (byteCount)
    lengthInBits += 8

    // Simple field (status)
    lengthInBits += 16

    // Simple field (eventCount)
    lengthInBits += 16

    // Simple field (messageCount)
    lengthInBits += 16

    // Array field
    if len(m.Events) > 0 {
        lengthInBits += 8 * uint16(len(m.Events))
    }

    return lengthInBits
}

func (m *ModbusPDUGetComEventLogResponse) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUGetComEventLogResponseParse(io *utils.ReadBuffer) (*ModbusPDU, error) {

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    byteCount, _byteCountErr := io.ReadUint8(8)
    if _byteCountErr != nil {
        return nil, errors.New("Error parsing 'byteCount' field " + _byteCountErr.Error())
    }

    // Simple Field (status)
    status, _statusErr := io.ReadUint16(16)
    if _statusErr != nil {
        return nil, errors.New("Error parsing 'status' field " + _statusErr.Error())
    }

    // Simple Field (eventCount)
    eventCount, _eventCountErr := io.ReadUint16(16)
    if _eventCountErr != nil {
        return nil, errors.New("Error parsing 'eventCount' field " + _eventCountErr.Error())
    }

    // Simple Field (messageCount)
    messageCount, _messageCountErr := io.ReadUint16(16)
    if _messageCountErr != nil {
        return nil, errors.New("Error parsing 'messageCount' field " + _messageCountErr.Error())
    }

    // Array field (events)
    // Count array
    events := make([]int8, uint16(byteCount) - uint16(uint16(6)))
    for curItem := uint16(0); curItem < uint16(uint16(byteCount) - uint16(uint16(6))); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'events' field " + _err.Error())
        }
        events[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &ModbusPDUGetComEventLogResponse{
        Status: status,
        EventCount: eventCount,
        MessageCount: messageCount,
        Events: events,
        Parent: &ModbusPDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *ModbusPDUGetComEventLogResponse) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    byteCount := uint8(uint8(uint8(len(m.Events))) + uint8(uint8(6)))
    _byteCountErr := io.WriteUint8(8, (byteCount))
    if _byteCountErr != nil {
        return errors.New("Error serializing 'byteCount' field " + _byteCountErr.Error())
    }

    // Simple Field (status)
    status := uint16(m.Status)
    _statusErr := io.WriteUint16(16, (status))
    if _statusErr != nil {
        return errors.New("Error serializing 'status' field " + _statusErr.Error())
    }

    // Simple Field (eventCount)
    eventCount := uint16(m.EventCount)
    _eventCountErr := io.WriteUint16(16, (eventCount))
    if _eventCountErr != nil {
        return errors.New("Error serializing 'eventCount' field " + _eventCountErr.Error())
    }

    // Simple Field (messageCount)
    messageCount := uint16(m.MessageCount)
    _messageCountErr := io.WriteUint16(16, (messageCount))
    if _messageCountErr != nil {
        return errors.New("Error serializing 'messageCount' field " + _messageCountErr.Error())
    }

    // Array Field (events)
    if m.Events != nil {
        for _, _element := range m.Events {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'events' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *ModbusPDUGetComEventLogResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "status":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Status = data
            case "eventCount":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.EventCount = data
            case "messageCount":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.MessageCount = data
            case "events":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.Events = utils.ByteToInt8(_decoded[0:_len])
            }
        }
    }
}

func (m *ModbusPDUGetComEventLogResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.modbus.readwrite.ModbusPDUGetComEventLogResponse"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Status, xml.StartElement{Name: xml.Name{Local: "status"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.EventCount, xml.StartElement{Name: xml.Name{Local: "eventCount"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.MessageCount, xml.StartElement{Name: xml.Name{Local: "messageCount"}}); err != nil {
        return err
    }
    _encodedEvents := make([]byte, base64.StdEncoding.EncodedLen(len(m.Events)))
    base64.StdEncoding.Encode(_encodedEvents, utils.Int8ToByte(m.Events))
    if err := e.EncodeElement(_encodedEvents, xml.StartElement{Name: xml.Name{Local: "events"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

