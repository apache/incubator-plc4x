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
    "github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
    "io"
    "strconv"
    "reflect"
    "strings"
)

// Constant values.
const BVLC_BACNETTYPE uint8 = 0x81

// The data-structure of this message
type BVLC struct {
    Child IBVLCChild
    IBVLC
    IBVLCParent
}

// The corresponding interface
type IBVLC interface {
    BvlcFunction() uint8
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
    xml.Marshaler
}

type IBVLCParent interface {
    SerializeParent(io utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error
    GetTypeName() string
}

type IBVLCChild interface {
    Serialize(io utils.WriteBuffer) error
    InitializeParent(parent *BVLC)
    GetTypeName() string
    IBVLC
}

func NewBVLC() *BVLC {
    return &BVLC{}
}

func CastBVLC(structType interface{}) *BVLC {
    castFunc := func(typ interface{}) *BVLC {
        if casted, ok := typ.(BVLC); ok {
            return &casted
        }
        if casted, ok := typ.(*BVLC); ok {
            return casted
        }
        return nil
    }
    return castFunc(structType)
}

func (m *BVLC) GetTypeName() string {
    return "BVLC"
}

func (m *BVLC) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Const Field (bacnetType)
    lengthInBits += 8

    // Discriminator Field (bvlcFunction)
    lengthInBits += 8

    // Implicit Field (bvlcLength)
    lengthInBits += 16

    // Length of sub-type elements will be added by sub-type...
    lengthInBits += m.Child.LengthInBits()

    return lengthInBits
}

func (m *BVLC) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func BVLCParse(io *utils.ReadBuffer) (*BVLC, error) {

    // Const Field (bacnetType)
    bacnetType, _bacnetTypeErr := io.ReadUint8(8)
    if _bacnetTypeErr != nil {
        return nil, errors.New("Error parsing 'bacnetType' field " + _bacnetTypeErr.Error())
    }
    if bacnetType != BVLC_BACNETTYPE {
        return nil, errors.New("Expected constant value " + strconv.Itoa(int(BVLC_BACNETTYPE)) + " but got " + strconv.Itoa(int(bacnetType)))
    }

    // Discriminator Field (bvlcFunction) (Used as input to a switch field)
    bvlcFunction, _bvlcFunctionErr := io.ReadUint8(8)
    if _bvlcFunctionErr != nil {
        return nil, errors.New("Error parsing 'bvlcFunction' field " + _bvlcFunctionErr.Error())
    }

    // Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    bvlcLength, _bvlcLengthErr := io.ReadUint16(16)
    if _bvlcLengthErr != nil {
        return nil, errors.New("Error parsing 'bvlcLength' field " + _bvlcLengthErr.Error())
    }

    // Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
    var _parent *BVLC
    var typeSwitchError error
    switch {
    case bvlcFunction == 0x00:
        _parent, typeSwitchError = BVLCResultParse(io)
    case bvlcFunction == 0x01:
        _parent, typeSwitchError = BVLCWideBroadcastDistributionTableParse(io)
    case bvlcFunction == 0x02:
        _parent, typeSwitchError = BVLCReadBroadcastDistributionTableParse(io)
    case bvlcFunction == 0x03:
        _parent, typeSwitchError = BVLCReadBroadcastDistributionTableAckParse(io)
    case bvlcFunction == 0x04:
        _parent, typeSwitchError = BVLCForwardedNPDUParse(io, bvlcLength)
    case bvlcFunction == 0x05:
        _parent, typeSwitchError = BVLCRegisterForeignDeviceParse(io)
    case bvlcFunction == 0x06:
        _parent, typeSwitchError = BVLCReadForeignDeviceTableParse(io)
    case bvlcFunction == 0x07:
        _parent, typeSwitchError = BVLCReadForeignDeviceTableAckParse(io)
    case bvlcFunction == 0x08:
        _parent, typeSwitchError = BVLCDeleteForeignDeviceTableEntryParse(io)
    case bvlcFunction == 0x09:
        _parent, typeSwitchError = BVLCDistributeBroadcastToNetworkParse(io)
    case bvlcFunction == 0x0A:
        _parent, typeSwitchError = BVLCOriginalUnicastNPDUParse(io, bvlcLength)
    case bvlcFunction == 0x0B:
        _parent, typeSwitchError = BVLCOriginalBroadcastNPDUParse(io, bvlcLength)
    case bvlcFunction == 0x0C:
        _parent, typeSwitchError = BVLCSecureBVLLParse(io)
    }
    if typeSwitchError != nil {
        return nil, errors.New("Error parsing sub-type for type-switch. " + typeSwitchError.Error())
    }

    // Finish initializing
    _parent.Child.InitializeParent(_parent)
    return _parent, nil
}

func (m *BVLC) Serialize(io utils.WriteBuffer) error {
    return m.Child.Serialize(io)
}

func (m *BVLC) SerializeParent(io utils.WriteBuffer, child IBVLC, serializeChildFunction func() error) error {

    // Const Field (bacnetType)
    _bacnetTypeErr := io.WriteUint8(8, 0x81)
    if _bacnetTypeErr != nil {
        return errors.New("Error serializing 'bacnetType' field " + _bacnetTypeErr.Error())
    }

    // Discriminator Field (bvlcFunction) (Used as input to a switch field)
    bvlcFunction := uint8(child.BvlcFunction())
    _bvlcFunctionErr := io.WriteUint8(8, (bvlcFunction))
    if _bvlcFunctionErr != nil {
        return errors.New("Error serializing 'bvlcFunction' field " + _bvlcFunctionErr.Error())
    }

    // Implicit Field (bvlcLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
    bvlcLength := uint16(uint16(m.LengthInBytes()))
    _bvlcLengthErr := io.WriteUint16(16, (bvlcLength))
    if _bvlcLengthErr != nil {
        return errors.New("Error serializing 'bvlcLength' field " + _bvlcLengthErr.Error())
    }

    // Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
    _typeSwitchErr := serializeChildFunction()
    if _typeSwitchErr != nil {
        return errors.New("Error serializing sub-type field " + _typeSwitchErr.Error())
    }

    return nil
}

func (m *BVLC) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCResult":
                        var dt *BVLCResult
                        if m.Child != nil {
                            dt = m.Child.(*BVLCResult)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCWideBroadcastDistributionTable":
                        var dt *BVLCWideBroadcastDistributionTable
                        if m.Child != nil {
                            dt = m.Child.(*BVLCWideBroadcastDistributionTable)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTable":
                        var dt *BVLCReadBroadcastDistributionTable
                        if m.Child != nil {
                            dt = m.Child.(*BVLCReadBroadcastDistributionTable)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadBroadcastDistributionTableAck":
                        var dt *BVLCReadBroadcastDistributionTableAck
                        if m.Child != nil {
                            dt = m.Child.(*BVLCReadBroadcastDistributionTableAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCForwardedNPDU":
                        var dt *BVLCForwardedNPDU
                        if m.Child != nil {
                            dt = m.Child.(*BVLCForwardedNPDU)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCRegisterForeignDevice":
                        var dt *BVLCRegisterForeignDevice
                        if m.Child != nil {
                            dt = m.Child.(*BVLCRegisterForeignDevice)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTable":
                        var dt *BVLCReadForeignDeviceTable
                        if m.Child != nil {
                            dt = m.Child.(*BVLCReadForeignDeviceTable)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCReadForeignDeviceTableAck":
                        var dt *BVLCReadForeignDeviceTableAck
                        if m.Child != nil {
                            dt = m.Child.(*BVLCReadForeignDeviceTableAck)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDeleteForeignDeviceTableEntry":
                        var dt *BVLCDeleteForeignDeviceTableEntry
                        if m.Child != nil {
                            dt = m.Child.(*BVLCDeleteForeignDeviceTableEntry)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCDistributeBroadcastToNetwork":
                        var dt *BVLCDistributeBroadcastToNetwork
                        if m.Child != nil {
                            dt = m.Child.(*BVLCDistributeBroadcastToNetwork)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCOriginalUnicastNPDU":
                        var dt *BVLCOriginalUnicastNPDU
                        if m.Child != nil {
                            dt = m.Child.(*BVLCOriginalUnicastNPDU)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCOriginalBroadcastNPDU":
                        var dt *BVLCOriginalBroadcastNPDU
                        if m.Child != nil {
                            dt = m.Child.(*BVLCOriginalBroadcastNPDU)
                        }
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        if m.Child == nil {
                            dt.Parent = m
                            m.Child = dt
                        }
                    case "org.apache.plc4x.java.bacnetip.readwrite.BVLCSecureBVLL":
                        var dt *BVLCSecureBVLL
                        if m.Child != nil {
                            dt = m.Child.(*BVLCSecureBVLL)
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

func (m *BVLC) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
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

