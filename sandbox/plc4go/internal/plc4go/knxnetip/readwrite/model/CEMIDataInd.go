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
type CEMIDataInd struct {
    AdditionalInformationLength uint8
    AdditionalInformation []ICEMIAdditionalInformation
    CemiDataFrame ICEMIDataFrame
    CEMI
}

// The corresponding interface
type ICEMIDataInd interface {
    ICEMI
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m CEMIDataInd) MessageCode() uint8 {
    return 0x29
}

func (m CEMIDataInd) initialize() spi.Message {
    return m
}

func NewCEMIDataInd(additionalInformationLength uint8, additionalInformation []ICEMIAdditionalInformation, cemiDataFrame ICEMIDataFrame) CEMIInitializer {
    return &CEMIDataInd{AdditionalInformationLength: additionalInformationLength, AdditionalInformation: additionalInformation, CemiDataFrame: cemiDataFrame}
}

func CastICEMIDataInd(structType interface{}) ICEMIDataInd {
    castFunc := func(typ interface{}) ICEMIDataInd {
        if iCEMIDataInd, ok := typ.(ICEMIDataInd); ok {
            return iCEMIDataInd
        }
        return nil
    }
    return castFunc(structType)
}

func CastCEMIDataInd(structType interface{}) CEMIDataInd {
    castFunc := func(typ interface{}) CEMIDataInd {
        if sCEMIDataInd, ok := typ.(CEMIDataInd); ok {
            return sCEMIDataInd
        }
        if sCEMIDataInd, ok := typ.(*CEMIDataInd); ok {
            return *sCEMIDataInd
        }
        return CEMIDataInd{}
    }
    return castFunc(structType)
}

func (m CEMIDataInd) LengthInBits() uint16 {
    var lengthInBits uint16 = m.CEMI.LengthInBits()

    // Simple field (additionalInformationLength)
    lengthInBits += 8

    // Array field
    if len(m.AdditionalInformation) > 0 {
        for _, element := range m.AdditionalInformation {
            lengthInBits += element.LengthInBits()
        }
    }

    // Simple field (cemiDataFrame)
    lengthInBits += m.CemiDataFrame.LengthInBits()

    return lengthInBits
}

func (m CEMIDataInd) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIDataIndParse(io *utils.ReadBuffer) (CEMIInitializer, error) {

    // Simple Field (additionalInformationLength)
    additionalInformationLength, _additionalInformationLengthErr := io.ReadUint8(8)
    if _additionalInformationLengthErr != nil {
        return nil, errors.New("Error parsing 'additionalInformationLength' field " + _additionalInformationLengthErr.Error())
    }

    // Array field (additionalInformation)
    // Length array
    additionalInformation := make([]ICEMIAdditionalInformation, 0)
    _additionalInformationLength := additionalInformationLength
    _additionalInformationEndPos := io.GetPos() + uint16(_additionalInformationLength)
    for ;io.GetPos() < _additionalInformationEndPos; {
        _message, _err := CEMIAdditionalInformationParse(io)
        if _err != nil {
            return nil, errors.New("Error parsing 'additionalInformation' field " + _err.Error())
        }
        var _item ICEMIAdditionalInformation
        _item, _ok := _message.(ICEMIAdditionalInformation)
        if !_ok {
            return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_item).Name() + " to CEMIAdditionalInformation")
        }
        additionalInformation = append(additionalInformation, _item)
    }

    // Simple Field (cemiDataFrame)
    _cemiDataFrameMessage, _err := CEMIDataFrameParse(io)
    if _err != nil {
        return nil, errors.New("Error parsing simple field 'cemiDataFrame'. " + _err.Error())
    }
    var cemiDataFrame ICEMIDataFrame
    cemiDataFrame, _cemiDataFrameOk := _cemiDataFrameMessage.(ICEMIDataFrame)
    if !_cemiDataFrameOk {
        return nil, errors.New("Couldn't cast message of type " + reflect.TypeOf(_cemiDataFrameMessage).Name() + " to ICEMIDataFrame")
    }

    // Create the instance
    return NewCEMIDataInd(additionalInformationLength, additionalInformation, cemiDataFrame), nil
}

func (m CEMIDataInd) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

    // Simple Field (additionalInformationLength)
    additionalInformationLength := uint8(m.AdditionalInformationLength)
    _additionalInformationLengthErr := io.WriteUint8(8, (additionalInformationLength))
    if _additionalInformationLengthErr != nil {
        return errors.New("Error serializing 'additionalInformationLength' field " + _additionalInformationLengthErr.Error())
    }

    // Array Field (additionalInformation)
    if m.AdditionalInformation != nil {
        for _, _element := range m.AdditionalInformation {
            _elementErr := _element.Serialize(io)
            if _elementErr != nil {
                return errors.New("Error serializing 'additionalInformation' field " + _elementErr.Error())
            }
        }
    }

    // Simple Field (cemiDataFrame)
    cemiDataFrame := CastICEMIDataFrame(m.CemiDataFrame)
    _cemiDataFrameErr := cemiDataFrame.Serialize(io)
    if _cemiDataFrameErr != nil {
        return errors.New("Error serializing 'cemiDataFrame' field " + _cemiDataFrameErr.Error())
    }

        return nil
    }
    return CEMISerialize(io, m.CEMI, CastICEMI(m), ser)
}

func (m *CEMIDataInd) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
            case "additionalInformationLength":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.AdditionalInformationLength = data
            case "additionalInformation":
                var _values []ICEMIAdditionalInformation
                switch tok.Attr[0].Value {
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIAdditionalInformationBusmonitorInfo":
                        var dt *CEMIAdditionalInformationBusmonitorInfo
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        _values = append(_values, dt)
                    case "org.apache.plc4x.java.knxnetip.readwrite.CEMIAdditionalInformationRelativeTimestamp":
                        var dt *CEMIAdditionalInformationRelativeTimestamp
                        if err := d.DecodeElement(&dt, &tok); err != nil {
                            return err
                        }
                        _values = append(_values, dt)
                    }
                    m.AdditionalInformation = _values
            case "cemiDataFrame":
                var data *CEMIDataFrame
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.CemiDataFrame = CastICEMIDataFrame(data)
            }
        }
    }
}

func (m CEMIDataInd) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.CEMIDataInd"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AdditionalInformationLength, xml.StartElement{Name: xml.Name{Local: "additionalInformationLength"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.AdditionalInformation, xml.StartElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "additionalInformation"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.CemiDataFrame, xml.StartElement{Name: xml.Name{Local: "cemiDataFrame"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

