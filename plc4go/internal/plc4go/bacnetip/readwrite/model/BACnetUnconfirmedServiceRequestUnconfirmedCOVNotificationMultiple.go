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
type BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple struct {
    Parent *BACnetUnconfirmedServiceRequest
    IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) ServiceChoice() uint8 {
    return 0x0B
}


func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple() *BACnetUnconfirmedServiceRequest {
    child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
        Parent: NewBACnetUnconfirmedServiceRequest(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(structType interface{}) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
    castFunc := func(typ interface{}) BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple {
        if casted, ok := typ.(BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
            return casted
        }
        if casted, ok := typ.(*BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple); ok {
            return *casted
        }
        if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
            return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
        }
        if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
            return CastBACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple(casted.Child)
        }
        return BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{}
    }
    return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultipleParse(io *utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {

    // Create a partially initialized instance
    _child := &BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple{
        Parent: &BACnetUnconfirmedServiceRequest{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m *BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    return nil
}

