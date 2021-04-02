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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type COTPPacketData struct {
	Eot     bool
	TpduRef uint8
	Parent  *COTPPacket
}

// The corresponding interface
type ICOTPPacketData interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *COTPPacketData) TpduCode() uint8 {
	return 0xF0
}

func (m *COTPPacketData) InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message) {
	m.Parent.Parameters = parameters
	m.Parent.Payload = payload
}

func NewCOTPPacketData(eot bool, tpduRef uint8, parameters []*COTPParameter, payload *S7Message) *COTPPacket {
	child := &COTPPacketData{
		Eot:     eot,
		TpduRef: tpduRef,
		Parent:  NewCOTPPacket(parameters, payload),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastCOTPPacketData(structType interface{}) *COTPPacketData {
	castFunc := func(typ interface{}) *COTPPacketData {
		if casted, ok := typ.(COTPPacketData); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacketData); ok {
			return casted
		}
		if casted, ok := typ.(COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return CastCOTPPacketData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacketData) GetTypeName() string {
	return "COTPPacketData"
}

func (m *COTPPacketData) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (eot)
	lengthInBits += 1

	// Simple field (tpduRef)
	lengthInBits += 7

	return lengthInBits
}

func (m *COTPPacketData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketDataParse(io *utils.ReadBuffer) (*COTPPacket, error) {

	// Simple Field (eot)
	eot, _eotErr := io.ReadBit()
	if _eotErr != nil {
		return nil, errors.Wrap(_eotErr, "Error parsing 'eot' field")
	}

	// Simple Field (tpduRef)
	tpduRef, _tpduRefErr := io.ReadUint8(7)
	if _tpduRefErr != nil {
		return nil, errors.Wrap(_tpduRefErr, "Error parsing 'tpduRef' field")
	}

	// Create a partially initialized instance
	_child := &COTPPacketData{
		Eot:     eot,
		TpduRef: tpduRef,
		Parent:  &COTPPacket{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *COTPPacketData) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (eot)
		eot := bool(m.Eot)
		_eotErr := io.WriteBit((eot))
		if _eotErr != nil {
			return errors.Wrap(_eotErr, "Error serializing 'eot' field")
		}

		// Simple Field (tpduRef)
		tpduRef := uint8(m.TpduRef)
		_tpduRefErr := io.WriteUint8(7, (tpduRef))
		if _tpduRefErr != nil {
			return errors.Wrap(_tpduRefErr, "Error serializing 'tpduRef' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *COTPPacketData) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "eot":
				var data bool
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Eot = data
			case "tpduRef":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TpduRef = data
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

func (m *COTPPacketData) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Eot, xml.StartElement{Name: xml.Name{Local: "eot"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TpduRef, xml.StartElement{Name: xml.Name{Local: "tpduRef"}}); err != nil {
		return err
	}
	return nil
}

func (m COTPPacketData) String() string {
	return string(m.Box("COTPPacketData", utils.DefaultWidth*2))
}

func (m COTPPacketData) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "COTPPacketData"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("Eot", m.Eot, width-2))
	boxes = append(boxes, utils.BoxAnything("TpduRef", m.TpduRef, width-2))
	return utils.BoxString(name, string(utils.AlignBoxes(boxes, width-2)), width)
}
