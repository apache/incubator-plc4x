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
type S7VarRequestParameterItem struct {

}

// The corresponding interface
type IS7VarRequestParameterItem interface {
    spi.Message
    ItemType() uint8
    Serialize(io utils.WriteBuffer) error
}

type S7VarRequestParameterItemInitializer interface {
    initialize() spi.Message
}

func S7VarRequestParameterItemItemType(m IS7VarRequestParameterItem) uint8 {
    return m.ItemType()
}


func CastIS7VarRequestParameterItem(structType interface{}) IS7VarRequestParameterItem {
    castFunc := func(typ interface{}) IS7VarRequestParameterItem {
        if iS7VarRequestParameterItem, ok := typ.(IS7VarRequestParameterItem); ok {
            return iS7VarRequestParameterItem
        }
        return nil
    }
    return castFunc(structType)
}

func CastS7VarRequestParameterItem(structType interface{}) S7VarRequestParameterItem {
    castFunc := func(typ interface{}) S7VarRequestParameterItem {
        if sS7VarRequestParameterItem, ok := typ.(S7VarRequestParameterItem); ok {
            return sS7VarRequestParameterItem
        }
        if sS7VarRequestParameterItem, ok := typ.(*S7VarRequestParameterItem); ok {
            return *sS7VarRequestParameterItem
        }
        return S7VarRequestParameterItem{}
    }
    return castFunc(structType)
}

func (m S7VarRequestParameterItem) LengthInBits() uint16 {
    var lengthInBits uint16 = 0

    // Discriminator Field (itemType)
    lengthInBits += 8

    // Length of sub-type elements will be added by sub-type...

    return lengthInBits
}

func (m S7VarRequestParameterItem) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func S7VarRequestParameterItemParse(io *utils.ReadBuffer) (spi.Message, error) {

    // Discriminator Field (itemType) (Used as input to a switch field)
    itemType, _itemTypeErr := io.ReadUint8(8)
    if _itemTypeErr != nil {
        return nil, errors.New("Error parsing 'itemType' field " + _itemTypeErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var initializer S7VarRequestParameterItemInitializer
    var typeSwitchError error
    switch {
    case itemType == 0x12:
        initializer, typeSwitchError = S7VarRequestParameterItemAddressParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Create the instance
    return initializer.initialize(), nil
}

func S7VarRequestParameterItemSerialize(io utils.WriteBuffer, m S7VarRequestParameterItem, i IS7VarRequestParameterItem, childSerialize func() error) error {

    // Discriminator Field (itemType) (Used as input to a switch field)
    itemType := uint8(i.ItemType())
    _itemTypeErr := io.WriteUint8(8, (itemType))
    if _itemTypeErr != nil {
        return errors.New("Error serializing 'itemType' field " + _itemTypeErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := childSerialize()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *S7VarRequestParameterItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            }
        }
    }
}

func (m S7VarRequestParameterItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.s7.readwrite.S7VarRequestParameterItem"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

