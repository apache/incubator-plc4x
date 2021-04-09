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
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type CEMIFramePollingDataExt struct {
    CEMIFrame
}

// The corresponding interface
type ICEMIFramePollingDataExt interface {
    ICEMIFrame
    Serialize(io utils.WriteBuffer) error
}

// Accessors for discriminator values.
func (m CEMIFramePollingDataExt) NotAckFrame() bool {
    return true
}

func (m CEMIFramePollingDataExt) StandardFrame() bool {
    return false
}

func (m CEMIFramePollingDataExt) Polling() bool {
    return true
}

func (m CEMIFramePollingDataExt) initialize(repeated bool, priority ICEMIPriority, acknowledgeRequested bool, errorFlag bool) spi.Message {
    m.Repeated = repeated
    m.Priority = priority
    m.AcknowledgeRequested = acknowledgeRequested
    m.ErrorFlag = errorFlag
    return m
}

func NewCEMIFramePollingDataExt() CEMIFrameInitializer {
    return &CEMIFramePollingDataExt{}
}

func CastICEMIFramePollingDataExt(structType interface{}) ICEMIFramePollingDataExt {
    castFunc := func(typ interface{}) ICEMIFramePollingDataExt {
        if iCEMIFramePollingDataExt, ok := typ.(ICEMIFramePollingDataExt); ok {
            return iCEMIFramePollingDataExt
        }
        return nil
    }
    return castFunc(structType)
}

func CastCEMIFramePollingDataExt(structType interface{}) CEMIFramePollingDataExt {
    castFunc := func(typ interface{}) CEMIFramePollingDataExt {
        if sCEMIFramePollingDataExt, ok := typ.(CEMIFramePollingDataExt); ok {
            return sCEMIFramePollingDataExt
        }
        if sCEMIFramePollingDataExt, ok := typ.(*CEMIFramePollingDataExt); ok {
            return *sCEMIFramePollingDataExt
        }
        return CEMIFramePollingDataExt{}
    }
    return castFunc(structType)
}

func (m CEMIFramePollingDataExt) LengthInBits() uint16 {
    var lengthInBits uint16 = m.CEMIFrame.LengthInBits()

    return lengthInBits
}

func (m CEMIFramePollingDataExt) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func CEMIFramePollingDataExtParse(io *utils.ReadBuffer) (CEMIFrameInitializer, error) {

    // Create the instance
    return NewCEMIFramePollingDataExt(), nil
}

func (m CEMIFramePollingDataExt) Serialize(io utils.WriteBuffer) error {
    ser := func() error {

        return nil
    }
    return CEMIFrameSerialize(io, m.CEMIFrame, CastICEMIFrame(m), ser)
}

func (m *CEMIFramePollingDataExt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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

func (m CEMIFramePollingDataExt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.knxnetip.readwrite.CEMIFramePollingDataExt"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

