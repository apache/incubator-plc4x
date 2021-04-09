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
    "reflect"
)

// The data-structure of this message
type ConnectionResponseDataBlockTunnelConnection struct {
    KnxAddress IKNXAddress
    ConnectionResponseDataBlock
}

// The corresponding interface
type IConnectionResponseDataBlockTunnelConnection interface {
    IConnectionResponseDataBlock
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m ConnectionResponseDataBlockTunnelConnection) ConnectionType() uint8 {
    return 0x04
}

func (m ConnectionResponseDataBlockTunnelConnection) initialize() spi.Message {
    return m
}

func NewConnectionResponseDataBlockTunnelConnection(knxAddress IKNXAddress) ConnectionResponseDataBlockInitializer {
    return &ConnectionResponseDataBlockTunnelConnection{KnxAddress: knxAddress}
}

func CastIConnectionResponseDataBlockTunnelConnection(structType interface{}) IConnectionResponseDataBlockTunnelConnection {
    castFunc := func(typ interface{}) IConnectionResponseDataBlockTunnelConnection {
        if iConnectionResponseDataBlockTunnelConnection, ok := typ.(IConnectionResponseDataBlockTunnelConnection); ok {
            return iConnectionResponseDataBlockTunnelConnection
        }
        return nil
    }
    return castFunc(structType)
}

func CastConnectionResponseDataBlockTunnelConnection(structType interface{}) ConnectionResponseDataBlockTunnelConnection {
    castFunc := func(typ interface{}) ConnectionResponseDataBlockTunnelConnection {
        if sConnectionResponseDataBlockTunnelConnection, ok := typ.(ConnectionResponseDataBlockTunnelConnection); ok {
            return sConnectionResponseDataBlockTunnelConnection
        }
        if sConnectionResponseDataBlockTunnelConnection, ok := typ.(*ConnectionResponseDataBlockTunnelConnection); ok {
            return *sConnectionResponseDataBlockTunnelConnection
        }
        return ConnectionResponseDataBlockTunnelConnection{}
    }
    return castFunc(structType)
}

func (m ConnectionResponseDataBlockTunnelConnection) LengthInBits() uint16 {
    var lengthInBits uint16 = m.ConnectionResponseDataBlock.LengthInBits()

    // Simple field (knxAddress)
    lengthInBits += m.KnxAddress.LengthInBits()

    return lengthInBits
}

func (m ConnectionResponseDataBlockTunnelConnection) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ConnectionResponseDataBlockTunnelConnectionParse(io *utils.ReadBuffer) (ConnectionResponseDataBlockInitializer, error) {

    // Simple Field (knxAddress)
    _knxAddressMessage, _err := KNXAddressParse(io)
    if _err != nil {
        return nil, errors.New("Error parsing simple field 'knxAddress'. " + _err.Error())
    }
    var knxAddress IKNXAddress
    knxAddress, _knxAddressOk := _knxAddressMessage.(IKNXAddress)
    if !_knxAddressOk {
        return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_knxAddressMessage).Name() + " to IKNXAddress")
    }

    // Create the instance
    return NewConnectionResponseDataBlockTunnelConnection(knxAddress), nil
}

func (m ConnectionResponseDataBlockTunnelConnection) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (knxAddress)
    knxAddress := CastIKNXAddress(m.KnxAddress)
    _knxAddressErr := knxAddress.Serialize(io)
    if _knxAddressErr != nil {
        return errors.New("Error serializing 'knxAddress' field " + _knxAddressErr.Error())
    }

        return nil
    }
    return ConnectionResponseDataBlockSerialize(io, m.ConnectionResponseDataBlock, CastIConnectionResponseDataBlock(m), ser)
}

func (m *ConnectionResponseDataBlockTunnelConnection) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "knxAddress":
                var data *KNXAddress
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.KnxAddress = CastIKNXAddress(data)
            }
        }
    }
}

func (m ConnectionResponseDataBlockTunnelConnection) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.ConnectionResponseDataBlockTunnelConnection"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.KnxAddress, xml.StartElement{Name: xml.Name{Local: "knxAddress"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

