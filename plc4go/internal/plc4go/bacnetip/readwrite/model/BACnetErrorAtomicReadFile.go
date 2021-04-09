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
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetErrorAtomicReadFile struct {
    BACnetError
}

// The corresponding interface
type IBACnetErrorAtomicReadFile interface {
    IBACnetError
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m BACnetErrorAtomicReadFile) ServiceChoice() uint8 {
    return 0x06
}

func (m BACnetErrorAtomicReadFile) initialize() spi.Message {
    return m
}

func NewBACnetErrorAtomicReadFile() BACnetErrorInitializer {
    return &BACnetErrorAtomicReadFile{}
}

func CastIBACnetErrorAtomicReadFile(structType interface{}) IBACnetErrorAtomicReadFile {
    castFunc := func(typ interface{}) IBACnetErrorAtomicReadFile {
        if iBACnetErrorAtomicReadFile, ok := typ.(IBACnetErrorAtomicReadFile); ok {
            return iBACnetErrorAtomicReadFile
        }
        return nil
    }
    return castFunc(structType)
}

func CastBACnetErrorAtomicReadFile(structType interface{}) BACnetErrorAtomicReadFile {
    castFunc := func(typ interface{}) BACnetErrorAtomicReadFile {
        if sBACnetErrorAtomicReadFile, ok := typ.(BACnetErrorAtomicReadFile); ok {
            return sBACnetErrorAtomicReadFile
        }
        if sBACnetErrorAtomicReadFile, ok := typ.(*BACnetErrorAtomicReadFile); ok {
            return *sBACnetErrorAtomicReadFile
        }
        return BACnetErrorAtomicReadFile{}
    }
    return castFunc(structType)
}

func (m BACnetErrorAtomicReadFile) LengthInBits() uint16 {
    var lengthInBits uint16 = m.BACnetError.LengthInBits()

    return lengthInBits
}

func (m BACnetErrorAtomicReadFile) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetErrorAtomicReadFileParse(io *utils.ReadBuffer) (BACnetErrorInitializer, error) {

    // Create the instance
    return NewBACnetErrorAtomicReadFile(), nil
}

func (m BACnetErrorAtomicReadFile) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return BACnetErrorSerialize(io, m.BACnetError, CastIBACnetError(m), ser)
}

func (m *BACnetErrorAtomicReadFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m BACnetErrorAtomicReadFile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicReadFile"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

