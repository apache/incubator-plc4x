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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BACnetServiceAckAtomicWriteFile struct {
    Parent *BACnetServiceAck
    IBACnetServiceAckAtomicWriteFile
}

// The corresponding interface
type IBACnetServiceAckAtomicWriteFile interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicWriteFile) ServiceChoice() uint8 {
    return 0x07
}


func (m *BACnetServiceAckAtomicWriteFile) InitializeParent(parent *BACnetServiceAck) {
}

func NewBACnetServiceAckAtomicWriteFile() *BACnetServiceAck {
    child := &BACnetServiceAckAtomicWriteFile{
        Parent: NewBACnetServiceAck(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetServiceAckAtomicWriteFile(structType interface{}) BACnetServiceAckAtomicWriteFile {
    castFunc := func(typ interface{}) BACnetServiceAckAtomicWriteFile {
        if casted, ok := typ.(BACnetServiceAckAtomicWriteFile); ok {
            return casted
        }
        if casted, ok := typ.(*BACnetServiceAckAtomicWriteFile); ok {
            return *casted
        }
        if casted, ok := typ.(BACnetServiceAck); ok {
            return CastBACnetServiceAckAtomicWriteFile(casted.Child)
        }
        if casted, ok := typ.(*BACnetServiceAck); ok {
            return CastBACnetServiceAckAtomicWriteFile(casted.Child)
        }
        return BACnetServiceAckAtomicWriteFile{}
    }
    return castFunc(structType)
}

func (m *BACnetServiceAckAtomicWriteFile) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BACnetServiceAckAtomicWriteFile) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetServiceAckAtomicWriteFileParse(io *utils.ReadBuffer) (*BACnetServiceAck, error) {

    // Create a partially initialized instance
    _child := &BACnetServiceAckAtomicWriteFile{
        Parent: &BACnetServiceAck{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetServiceAckAtomicWriteFile) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetServiceAckAtomicWriteFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetServiceAckAtomicWriteFile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BACnetServiceAckAtomicWriteFile"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

