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
type BVLCDistributeBroadcastToNetwork struct {
    Parent *BVLC
    IBVLCDistributeBroadcastToNetwork
}

// The corresponding interface
type IBVLCDistributeBroadcastToNetwork interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BVLCDistributeBroadcastToNetwork) BvlcFunction() uint8 {
    return 0x09
}


func (m *BVLCDistributeBroadcastToNetwork) InitializeParent(parent *BVLC) {
}

func NewBVLCDistributeBroadcastToNetwork() *BVLC {
    child := &BVLCDistributeBroadcastToNetwork{
        Parent: NewBVLC(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBVLCDistributeBroadcastToNetwork(structType interface{}) BVLCDistributeBroadcastToNetwork {
    castFunc := func(typ interface{}) BVLCDistributeBroadcastToNetwork {
        if casted, ok := typ.(BVLCDistributeBroadcastToNetwork); ok {
            return casted
        }
        if casted, ok := typ.(*BVLCDistributeBroadcastToNetwork); ok {
            return *casted
        }
        if casted, ok := typ.(BVLC); ok {
            return CastBVLCDistributeBroadcastToNetwork(casted.Child)
        }
        if casted, ok := typ.(*BVLC); ok {
            return CastBVLCDistributeBroadcastToNetwork(casted.Child)
        }
        return BVLCDistributeBroadcastToNetwork{}
    }
    return castFunc(structType)
}

func (m *BVLCDistributeBroadcastToNetwork) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BVLCDistributeBroadcastToNetwork) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BVLCDistributeBroadcastToNetworkParse(io *utils.ReadBuffer) (*BVLC, error) {

    // Create a partially initialized instance
    _child := &BVLCDistributeBroadcastToNetwork{
        Parent: &BVLC{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BVLCDistributeBroadcastToNetwork) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BVLCDistributeBroadcastToNetwork) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BVLCDistributeBroadcastToNetwork) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.BVLCDistributeBroadcastToNetwork"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

