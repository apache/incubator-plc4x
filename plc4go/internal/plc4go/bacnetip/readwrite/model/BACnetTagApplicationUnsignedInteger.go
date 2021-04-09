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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetTagApplicationUnsignedInteger struct {
    Data []int8
    BACnetTag
}

// The corresponding interface
type IBACnetTagApplicationUnsignedInteger interface {
    IBACnetTag
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetTagApplicationUnsignedInteger) ContextSpecificTag() uint8 {
    return 0
}

func (m BACnetTagApplicationUnsignedInteger) initialize(typeOrTagNumber uint8, lengthValueType uint8, extTagNumber *uint8, extLength *uint8) spi.Message {
    m.TypeOrTagNumber = typeOrTagNumber
    m.LengthValueType = lengthValueType
    m.ExtTagNumber = extTagNumber
    m.ExtLength = extLength
    return m
}

func NewBACnetTagApplicationUnsignedInteger(data []int8) BACnetTagInitializer {
    return &BACnetTagApplicationUnsignedInteger{Data: data}
}

func CastIBACnetTagApplicationUnsignedInteger(structType interface{}) IBACnetTagApplicationUnsignedInteger {
    castFunc := func(typ interface{}) IBACnetTagApplicationUnsignedInteger {
        if iBACnetTagApplicationUnsignedInteger, ok := typ.(IBACnetTagApplicationUnsignedInteger); ok {
            return iBACnetTagApplicationUnsignedInteger
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetTagApplicationUnsignedInteger(structType interface{}) BACnetTagApplicationUnsignedInteger {
    castFunc := func(typ interface{}) BACnetTagApplicationUnsignedInteger {
        if sBACnetTagApplicationUnsignedInteger, ok := typ.(BACnetTagApplicationUnsignedInteger); ok {
            return sBACnetTagApplicationUnsignedInteger
        }
        if sBACnetTagApplicationUnsignedInteger, ok := typ.(*BACnetTagApplicationUnsignedInteger); ok {
            return *sBACnetTagApplicationUnsignedInteger
        }
        return BACnetTagApplicationUnsignedInteger{}
    }
    return castFunc(structType)
}

func (m BACnetTagApplicationUnsignedInteger) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetTag.LengthInBits()

    // Array field
    if len(m.Data) > 0 {
        lengthInBits += 8 * uint16(len(m.Data))
    }

    return lengthInBits
}

func (m BACnetTagApplicationUnsignedInteger) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetTagApplicationUnsignedIntegerParse(io *utils.ReadBuffer, lengthValueType uint8, extLength uint8) (BACnetTagInitializer, error) {

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

    // Create the instance
    return NewBACnetTagApplicationUnsignedInteger(data), nil
}

func (m BACnetTagApplicationUnsignedInteger) Serialize(io utils.WriteBuffer) error {
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
    return BACnetTagSerialize(io, m.BACnetTag, CastIBACnetTag(m), ser)
}

func (m *BACnetTagApplicationUnsignedInteger) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
                m.Data = utils.ByteToInt8(_decoded[0:_len])
            }
        }
    }
}

func (m BACnetTagApplicationUnsignedInteger) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetTagApplicationUnsignedInteger"},
        }}); err != nil {
        return err
    }
    _encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(m.Data)))
    base64.StdEncoding.Encode(_encodedData, utils.Int8ToByte(m.Data))
    if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

