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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "reflect"
    "strings"
)

// The data-structure of this message
type BACnetError struct {
    Child IBACnetErrorChild
    IBACnetError
    IBACnetErrorParent
}

// The corresponding interface
type IBACnetError interface {
    ServiceChoice() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IBACnetErrorParent interface {
    SerializeParent(io utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error
    GetTypeName() string
}

type IBACnetErrorChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *BACnetError)
    GetTypeName() string
    IBACnetError
}

func NewBACnetError() *BACnetError {
    return &BACnetError{}
}

func CastBACnetError(structType interface{}) *BACnetError {
    castFunc := func(typ interface{}) *BACnetError {
        if casted, ok := typ.(BACnetError); ok {
            return &casted
        }
        if casted, ok := typ.(*BACnetError); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *BACnetError) GetTypeName() string {
    return "BACnetError"
}

func (m *BACnetError) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Discriminator Field (serviceChoice)
    lengthInBits += 8

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *BACnetError) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BACnetErrorParse(io *utils.ReadBuffer) (*BACnetError, error) {

    // Discriminator Field (serviceChoice) (Used as input to a switch field)
    serviceChoice, _serviceChoiceErr := io.ReadUint8(8)
    if _serviceChoiceErr != nil {
        return nil, errors.New("Error parsing 'serviceChoice' field " + _serviceChoiceErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *BACnetError
    var typeSwitchError error
    switch {
    case serviceChoice == 0x03:
        _parent, typeSwitchError = BACnetErrorGetAlarmSummaryParse(io)
    case serviceChoice == 0x04:
        _parent, typeSwitchError = BACnetErrorGetEnrollmentSummaryParse(io)
    case serviceChoice == 0x1D:
        _parent, typeSwitchError = BACnetErrorGetEventInformationParse(io)
    case serviceChoice == 0x06:
        _parent, typeSwitchError = BACnetErrorAtomicReadFileParse(io)
    case serviceChoice == 0x07:
        _parent, typeSwitchError = BACnetErrorAtomicWriteFileParse(io)
    case serviceChoice == 0x0A:
        _parent, typeSwitchError = BACnetErrorCreateObjectParse(io)
    case serviceChoice == 0x0C:
        _parent, typeSwitchError = BACnetErrorReadPropertyParse(io)
    case serviceChoice == 0x0E:
        _parent, typeSwitchError = BACnetErrorReadPropertyMultipleParse(io)
    case serviceChoice == 0x1A:
        _parent, typeSwitchError = BACnetErrorReadRangeParse(io)
    case serviceChoice == 0x12:
        _parent, typeSwitchError = BACnetErrorConfirmedPrivateTransferParse(io)
    case serviceChoice == 0x15:
        _parent, typeSwitchError = BACnetErrorVTOpenParse(io)
    case serviceChoice == 0x17:
        _parent, typeSwitchError = BACnetErrorVTDataParse(io)
    case serviceChoice == 0x18:
        _parent, typeSwitchError = BACnetErrorRemovedAuthenticateParse(io)
    case serviceChoice == 0x0D:
        _parent, typeSwitchError = BACnetErrorRemovedReadPropertyConditionalParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *BACnetError) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *BACnetError) SerializeParent(io utils.WriteBuffer, child IBACnetError, serializeChildFunction func() error) error {

    // Discriminator Field (serviceChoice) (Used as input to a switch field)
    serviceChoice := uint8(child.ServiceChoice())
    _serviceChoiceErr := io.WriteUint8(8, (serviceChoice))
    if _serviceChoiceErr != nil {
        return errors.New("Error serializing 'serviceChoice' field " + _serviceChoiceErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *BACnetError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    var token xml.Token
    var err error
    for {
        token, err = d.Token()
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
            default:
                switch start.Attr[0].Value {
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetAlarmSummary":
                        var dt *BACnetErrorGetAlarmSummary
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorGetAlarmSummary)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEnrollmentSummary":
                        var dt *BACnetErrorGetEnrollmentSummary
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorGetEnrollmentSummary)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorGetEventInformation":
                        var dt *BACnetErrorGetEventInformation
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorGetEventInformation)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicReadFile":
                        var dt *BACnetErrorAtomicReadFile
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorAtomicReadFile)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorAtomicWriteFile":
                        var dt *BACnetErrorAtomicWriteFile
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorAtomicWriteFile)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorCreateObject":
                        var dt *BACnetErrorCreateObject
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorCreateObject)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadProperty":
                        var dt *BACnetErrorReadProperty
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorReadProperty)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadPropertyMultiple":
                        var dt *BACnetErrorReadPropertyMultiple
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorReadPropertyMultiple)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorReadRange":
                        var dt *BACnetErrorReadRange
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorReadRange)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorConfirmedPrivateTransfer":
                        var dt *BACnetErrorConfirmedPrivateTransfer
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorConfirmedPrivateTransfer)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTOpen":
                        var dt *BACnetErrorVTOpen
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorVTOpen)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorVTData":
                        var dt *BACnetErrorVTData
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorVTData)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedAuthenticate":
                        var dt *BACnetErrorRemovedAuthenticate
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorRemovedAuthenticate)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetErrorRemovedReadPropertyConditional":
                        var dt *BACnetErrorRemovedReadPropertyConditional
                        if m.Child != nil {
                            dt = m.Child.(*BACnetErrorRemovedReadPropertyConditional)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                }
            }
        }
    }
}

func (m *BACnetError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    className := reflect.TypeOf(m.Child).String()
    className = "org.apache.plc4x.java.bacnetip.readwrite." + className[strings.LastIndex(className, ".") + 1:]
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: className},
        }}); err != nil {
        return err
    }
    marshaller, ok := m.Child.(xml.Marshaler)
    if !ok {
        return errors.New("child is not castable to Marshaler")
    }
    marshaller.MarshalXML(e, start)
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

