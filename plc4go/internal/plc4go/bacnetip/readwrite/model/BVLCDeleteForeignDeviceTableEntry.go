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
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
)

// The data-structure of this message
type BVLCDeleteForeignDeviceTableEntry struct {
    Parent *BVLC
    IBVLCDeleteForeignDeviceTableEntry
}

// The corresponding interface
type IBVLCDeleteForeignDeviceTableEntry interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCDeleteForeignDeviceTableEntry) BvlcFunction() uint8 {
    return 0x08
}


func (m *BVLCDeleteForeignDeviceTableEntry) InitializeParent(parent *BVLC) {
}

func NewBVLCDeleteForeignDeviceTableEntry() *BVLC {
    child := &BVLCDeleteForeignDeviceTableEntry{
        Parent: NewBVLC(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBVLCDeleteForeignDeviceTableEntry(structType interface{}) BVLCDeleteForeignDeviceTableEntry {
    castFunc := func(typ interface{}) BVLCDeleteForeignDeviceTableEntry {
        if casted, ok := typ.(BVLCDeleteForeignDeviceTableEntry); ok {
            return casted
        }
        if casted, ok := typ.(*BVLCDeleteForeignDeviceTableEntry); ok {
            return *casted
        }
        if casted, ok := typ.(BVLC); ok {
            return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
        }
        if casted, ok := typ.(*BVLC); ok {
            return CastBVLCDeleteForeignDeviceTableEntry(casted.Child)
        }
        return BVLCDeleteForeignDeviceTableEntry{}
    }
    return castFunc(structType)
}

func (m *BVLCDeleteForeignDeviceTableEntry) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BVLCDeleteForeignDeviceTableEntry) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BVLCDeleteForeignDeviceTableEntryParse(io *utils.ReadBuffer) (*BVLC, error) {

    // Create a partially initialized instance
    _child := &BVLCDeleteForeignDeviceTableEntry{
        Parent: &BVLC{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BVLCDeleteForeignDeviceTableEntry) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCDeleteForeignDeviceTableEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    token = start
    for {
        switch token.(type) {
        case xml.StartElement:
            tok := token.(xml.StartElement)
            switch tok.Name.Local {
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

func (m *BVLCDeleteForeignDeviceTableEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

