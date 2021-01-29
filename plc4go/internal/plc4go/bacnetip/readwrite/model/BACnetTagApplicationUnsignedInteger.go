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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
)

// The data-structure of this message
type BACnetTagApplicationUnsignedInteger struct {
    Data []int8
    Parent *BACnetTag
    IBACnetTagApplicationUnsignedInteger
}

// The corresponding interface
type IBACnetTagApplicationUnsignedInteger interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetTagApplicationUnsignedInteger) ContextSpecificTag() uint8 {
    return 0
}


func (m *BACnetTagApplicationUnsignedInteger) InitializeParent(parent *BACnetTag, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) {
    m.Parent.TypeOrTagNumber = typeOrTagNumber
    m.Parent.LengthValueType = lengthValueType
    m.Parent.ExtTagNumber = extTagNumber
    m.Parent.ExtLength = extLength
}

func NewBACnetTagApplicationUnsignedInteger(data []int8, typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) *BACnetTag {
    child := &BACnetTagApplicationUnsignedInteger{
        Data: data,
        Parent: NewBACnetTag(typeOrTagNumber, lengthValueType, extTagNumber, extLength),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetTagApplicationUnsignedInteger(structType interface{}) *BACnetTagApplicationUnsignedInteger {
    castFunc := func(typ interface{}) *BACnetTagApplicationUnsignedInteger {
        if casted, ok := typ.(BACnetTagApplicationUnsignedInteger); ok {
            return &casted
        }
        if casted, ok := typ.(*BACnetTagApplicationUnsignedInteger); ok {
            return casted
        }
        if casted, ok := typ.(BACnetTag); ok {
            return CastBACnetTagApplicationUnsignedInteger(casted.Child)
        }
        if casted, ok := typ.(*BACnetTag); ok {
            return CastBACnetTagApplicationUnsignedInteger(casted.Child)
        }
        return nil
    }
    return castFunc(structType)
}

func (m *BACnetTagApplicationUnsignedInteger) GetTypeName() string {
    return "BACnetTagApplicationUnsignedInteger"
}

func (m *BACnetTagApplicationUnsignedInteger) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m *BACnetTagApplicationUnsignedInteger) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetTagApplicationUnsignedIntegerParse(io *utils.ReadBuffer, lengthValueType uint8, extLength uint8) (*BACnetTag, error) {

    // Array field (data)
    // Length array
    data := make([]int8, 0)
    _dataLength := utils.InlineIf(bool(bool((lengthValueType) == ((5)))), uint16(extLength), uint16(lengthValueType))
    _dataEndPos := io.GetPos() + uint16(_dataLength)
    for ;io.GetPos() < _dataEndPos; {
        _item, _err := io.ReadInt8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'data' field " + _err.Error())
        }
        data = append(data, _item)
    }

    // Create a partially initialized instance
    _child := &BACnetTagApplicationUnsignedInteger{
        Data: data,
        Parent: &BACnetTag{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetTagApplicationUnsignedInteger) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Array Field (data)
    if m.Data != nil {
        for _, _element := range m.Data {
            _elementErr := io.WriteInt8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'data' field " + _elementErr.Error())
            }
        }
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetTagApplicationUnsignedInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
            case "data":
                var _encoded string
                if err := d.DecodeElement(&_encoded, &tok); err != nil {
                    return err
                }
                _decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
                _len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
                if err != nil {
                    return err
                }
                m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *BACnetTagApplicationUnsignedInteger) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    _encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(m.Data)))
    base64.StdEncoding.Encode(_encodedData, utils.Int8ArrayToByteArray(m.Data))
    if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    return nil
}

