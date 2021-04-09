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
type BACnetAddress struct {
    Address []uint8
    Port uint16

}

// The corresponding interface
type IBACnetAddress interface {
    spi.Message
    Serialize(io utils.WriteBuffer) error
}


func NewBACnetAddress(address []uint8, port uint16) spi.Message {
    return &BACnetAddress{Address: address, Port: port}
}

func CastIBACnetAddress(structType interface{}) IBACnetAddress {
    castFunc := func(typ interface{}) IBACnetAddress {
        if iBACnetAddress, ok := typ.(IBACnetAddress); ok {
            return iBACnetAddress
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetAddress(structType interface{}) BACnetAddress {
    castFunc := func(typ interface{}) BACnetAddress {
        if sBACnetAddress, ok := typ.(BACnetAddress); ok {
            return sBACnetAddress
        }
        if sBACnetAddress, ok := typ.(*BACnetAddress); ok {
            return *sBACnetAddress
        }
        return BACnetAddress{}
    }
    return castFunc(structType)
}

func (m BACnetAddress) LengthInBits() uint16 {
    var lengthInBits uint16 = 0

    // Array field
    if len(m.Address) > 0 {
        lengthInBits += 8 * uint16(len(m.Address))
    }

    // Simple field (port)
    lengthInBits += 16

    return lengthInBits
}

func (m BACnetAddress) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetAddressParse(io *utils.ReadBuffer) (spi.Message, error) {

    // Array field (address)
    // Count array
    address := make([]uint8, uint16(4))
    for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {

        _item, _err := io.ReadUint8(8)
        if _err != nil {
            return nil, errors.New("Error parsing 'address' field " + _err.Error())
        }
        address[curItem] = _item
    }

    // Simple Field (port)
    port, _portErr := io.ReadUint16(16)
    if _portErr != nil {
        return nil, errors.New("Error parsing 'port' field " + _portErr.Error())
    }

    // Create the instance
    return NewBACnetAddress(address, port), nil
}

func (m BACnetAddress) Serialize(io utils.WriteBuffer) error {

    // Array Field (address)
    if m.Address != nil {
        for _, _element := range m.Address {
            _elementErr := io.WriteUint8(8, _element)
            if _elementErr != nil {
                return errors.New("Error serializing 'address' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (port)
    port := uint16(m.Port)
    _portErr := io.WriteUint16(16, (port))
    if _portErr != nil {
        return errors.New("Error serializing 'port' field " + _portErr.Error())
    }

    return nil
}

func (m *BACnetAddress) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "address":
                var data []uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Address = data
            case "port":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.Port = data
            }
        }
    }
}

func (m BACnetAddress) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetAddress"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Address, xml.StartElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "address"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.Port, xml.StartElement{Name: xml.Name{Local: "port"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

