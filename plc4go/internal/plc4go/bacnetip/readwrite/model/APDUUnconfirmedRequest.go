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
    log "github.com/sirupsen/logrus"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type APDUUnconfirmedRequest struct {
    ServiceRequest *BACnetUnconfirmedServiceRequest
    Parent *APDU
    IAPDUUnconfirmedRequest
}

// The corresponding interface
type IAPDUUnconfirmedRequest interface {
    LengthInBytes() uint16
    LengthInBits() uint16
    Serialize(io utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUUnconfirmedRequest) ApduType() uint8 {
    return 0x1
}


func (m *APDUUnconfirmedRequest) InitializeParent(parent *APDU) {
}

func NewAPDUUnconfirmedRequest(serviceRequest *BACnetUnconfirmedServiceRequest, ) *APDU {
    child := &APDUUnconfirmedRequest{
        ServiceRequest: serviceRequest,
        Parent: NewAPDU(),
    }
    child.Parent.Child = child
    return child.Parent
}

func CastAPDUUnconfirmedRequest(structType interface{}) APDUUnconfirmedRequest {
    castFunc := func(typ interface{}) APDUUnconfirmedRequest {
        if casted, ok := typ.(APDUUnconfirmedRequest); ok {
            return casted
        }
        if casted, ok := typ.(*APDUUnconfirmedRequest); ok {
            return *casted
        }
        if casted, ok := typ.(APDU); ok {
            return CastAPDUUnconfirmedRequest(casted.Child)
        }
        if casted, ok := typ.(*APDU); ok {
            return CastAPDUUnconfirmedRequest(casted.Child)
        }
        return APDUUnconfirmedRequest{}
    }
    return castFunc(structType)
}

func (m *APDUUnconfirmedRequest) LengthInBits() uint16 {
    lengthInBits := uint16(0)

    // Reserved Field (reserved)
    lengthInBits += 4

    // Simple field (serviceRequest)
    lengthInBits += m.ServiceRequest.LengthInBits()

    return lengthInBits
}

func (m *APDUUnconfirmedRequest) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func APDUUnconfirmedRequestParse(io *utils.ReadBuffer, apduLength uint16) (*APDU, error) {

    // Reserved Field (Compartmentalized so the "reserved" variable can't leak)
    {
        reserved, _err := io.ReadUint8(4)
        if _err != nil {
            return nil, errors.New("Error parsing 'reserved' field " + _err.Error())
        }
        if reserved != uint8(0) {
            log.WithFields(log.Fields{
                "expected value": uint8(0),
                "got value": reserved,
            }).Info("Got unexpected response.")
        }
    }

    // Simple Field (serviceRequest)
    serviceRequest, _serviceRequestErr := BACnetUnconfirmedServiceRequestParse(io, uint16(apduLength) - uint16(uint16(1)))
    if _serviceRequestErr != nil {
        return nil, errors.New("Error parsing 'serviceRequest' field " + _serviceRequestErr.Error())
    }

    // Create a partially initialized instance
    _child := &APDUUnconfirmedRequest{
        ServiceRequest: serviceRequest,
        Parent: &APDU{},
    }
    _child.Parent.Child = _child
    return _child.Parent, nil
}

func (m *APDUUnconfirmedRequest) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Reserved Field (reserved)
    {
        _err := io.WriteUint8(4, uint8(0))
        if _err != nil {
            return errors.New("Error serializing 'reserved' field " + _err.Error())
        }
    }

    // Simple Field (serviceRequest)
    _serviceRequestErr := m.ServiceRequest.Serialize(io)
    if _serviceRequestErr != nil {
        return errors.New("Error serializing 'serviceRequest' field " + _serviceRequestErr.Error())
    }

        return nil
    }
    return m.Parent.SerializeParent(io, m, ser)
}

func (m *APDUUnconfirmedRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "serviceRequest":
                switch tok.Attr[0].Value {
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestIAm":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestIHave":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotification":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedEventNotification":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedTextMessage":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestTimeSynchronization":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWhoHas":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWhoIs":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUTCTimeSynchronization":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestWriteGroup":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    case "org.apache.plc4x.java.bacnetip.readwrite.BACnetUnconfirmedServiceRequestUnconfirmedCOVNotificationMultiple":
                        var dt *BACnetUnconfirmedServiceRequest
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        m.ServiceRequest = dt
                    }
            }
        }
    }
}

func (m *APDUUnconfirmedRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.bacnetip.readwrite.APDUUnconfirmedRequest"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ServiceRequest, xml.StartElement{Name: xml.Name{Local: "serviceRequest"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

