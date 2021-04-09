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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type DeviceConfigurationAck struct {
    DeviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock
    Parent *KNXNetIPMessage
    IDeviceConfigurationAck
}

// The corresponding interface
type IDeviceConfigurationAck interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *DeviceConfigurationAck) MsgType() uint16 {
    return 0x0311
}


func (m *DeviceConfigurationAck) InitializeParent(parent *KNXNetIPMessage) {
}

func NewDeviceConfigurationAck(deviceConfigurationAckDataBlock *DeviceConfigurationAckDataBlock, ) *KNXNetIPMessage {
    child := &DeviceConfigurationAck{
        DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
        Parent: NewKNXNetIPMessage(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastDeviceConfigurationAck(structType interface{}) DeviceConfigurationAck {
    castFunc := func(typ interface{}) DeviceConfigurationAck {
        if casted, ok := typ.(DeviceConfigurationAck); ok {
            return casted
        }
        if casted, ok := typ.(*DeviceConfigurationAck); ok {
            return *casted
        }
        if casted, ok := typ.(KNXNetIPMessage); ok {
            return CastDeviceConfigurationAck(casted.Child)
        }
        if casted, ok := typ.(*KNXNetIPMessage); ok {
            return CastDeviceConfigurationAck(casted.Child)
        }
        return DeviceConfigurationAck{}
    }
    return castFunc(structType)
}

func (m *DeviceConfigurationAck) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Simple field (deviceConfigurationAckDataBlock)
    lengthInBits += m.DeviceConfigurationAckDataBlock.LengthInBits()

    return lengthInBits
}

func (m *DeviceConfigurationAck) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func DeviceConfigurationAckParse(io *utils.ReadBuffer) (*KNXNetIPMessage, error) {

    // Simple Field (deviceConfigurationAckDataBlock)
    deviceConfigurationAckDataBlock, _deviceConfigurationAckDataBlockErr := DeviceConfigurationAckDataBlockParse(io)
    if _deviceConfigurationAckDataBlockErr != nil {
        return nil, errors.New("Error parsing 'deviceConfigurationAckDataBlock' field " + _deviceConfigurationAckDataBlockErr.Error())
    }

    // Create a partially initialized instance
    _child := &DeviceConfigurationAck{
        DeviceConfigurationAckDataBlock: deviceConfigurationAckDataBlock,
        Parent: &KNXNetIPMessage{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *DeviceConfigurationAck) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (deviceConfigurationAckDataBlock)
    _deviceConfigurationAckDataBlockErr := m.DeviceConfigurationAckDataBlock.Serialize(io)
    if _deviceConfigurationAckDataBlockErr != nil {
        return errors.New("Error serializing 'deviceConfigurationAckDataBlock' field " + _deviceConfigurationAckDataBlockErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *DeviceConfigurationAck) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "deviceConfigurationAckDataBlock":
                var data *DeviceConfigurationAckDataBlock
                if err := d.DecodeElement(data, &tok); err != nil {
                    return err
                }
                m.DeviceConfigurationAckDataBlock = data
            }
        }
    }
}

func (m *DeviceConfigurationAck) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.DeviceConfigurationAck"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.DeviceConfigurationAckDataBlock, xml.StartElement{Name: xml.Name{Local: "deviceConfigurationAckDataBlock"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

