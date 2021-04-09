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
type UnknownMessage struct {
    UnknownData []int8
    Parent *KNXNetIPMessage
    IUnknownMessage
}

// The corresponding interface
type IUnknownMessage interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *UnknownMessage) MsgType() uint16 {
    return 0x020B
}


func (m *UnknownMessage) InitializeParent(parent *KNXNetIPMessage) {
}

func NewUnknownMessage(unknownData []int8, ) *KNXNetIPMessage {
    child := &UnknownMessage{
        UnknownData: unknownData,
        Parent: NewKNXNetIPMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastUnknownMessage(structType interface{}) UnknownMessage {
    castFunc := func(typ interface{}) UnknownMessage {
        if casted, ok := typ.(UnknownMessage); ok {
            return casted
        }
        if casted, ok := typ.(*UnknownMessage); ok {
            return *casted
        }
        if casted, ok := typ.(KNXNetIPMessage); ok {
            return CastUnknownMessage(casted.Child)
        }
        if casted, ok := typ.(*KNXNetIPMessage); ok {
            return CastUnknownMessage(casted.Child)
        }
        return UnknownMessage{}
    }
    return castFunc(structType)
}

func (m *UnknownMessage) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Array field
    if len(m.UnknownData) > 0 {
        lengthInBits += 8 * uint16(len(m.UnknownData))
    }

    return lengthInBits
}

func (m *UnknownMessage) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func UnknownMessageParse(io *utils.ReadBuffer, totalLength uint16) (*KNXNetIPMessage, error) {

    // Array field (unknownData)
    // Count array
    unknownData := make([]int8, uint16(totalLength) - uint16(uint16(6)))
    for curItem := uint16(0); curItem < uint16(uint16(totalLength) - uint16(uint16(6))); curItem++ {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'unknownData' field " + _err.Error())
        }
        unknownData[curItem] = _item
    }

    // Create a partially initialized instance
    _child := &UnknownMessage{
        UnknownData: unknownData,
        Parent: &KNXNetIPMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *UnknownMessage) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Array Field (unknownData)
    if m.UnknownData != nil {
        for _, _element := range m.UnknownData {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'unknownData' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *UnknownMessage) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "unknownData":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.UnknownData = utils.ByteToInt8(_decoded[0:_len])
            }
        }
    }
}

func (m *UnknownMessage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.UnknownMessage"},
        }}); err != nil {
        return err
    }
    _encodedUnknownData := make([]byte, base64.StdEncoding.EncodedLen(len(m.UnknownData)))
    base64.StdEncoding.Encode(_encodedUnknownData, utils.Int8ToByte(m.UnknownData))
    if err := e.EncodeElement(_encodedUnknownData, xml.StartElement{Name: xml.Name{Local: "unknownData"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

