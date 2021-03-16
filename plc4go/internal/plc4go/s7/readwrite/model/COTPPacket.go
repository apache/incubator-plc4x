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
	"reflect"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type COTPPacket struct {
	Parameters []*COTPParameter
	Payload    *S7Message
	Child      ICOTPPacketChild
}

// The corresponding interface
type ICOTPPacket interface {
	TpduCode() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

type ICOTPPacketParent interface {
	SerializeParent(io utils.WriteBuffer, child ICOTPPacket, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICOTPPacketChild interface {
	Serialize(io utils.WriteBuffer) error
	InitializeParent(parent *COTPPacket, parameters []*COTPParameter, payload *S7Message)
	GetTypeName() string
	ICOTPPacket
}

func NewCOTPPacket(parameters []*COTPParameter, payload *S7Message) *COTPPacket {
	return &COTPPacket{Parameters: parameters, Payload: payload}
}

func CastCOTPPacket(structType interface{}) *COTPPacket {
	castFunc := func(typ interface{}) *COTPPacket {
		if casted, ok := typ.(COTPPacket); ok {
			return &casted
		}
		if casted, ok := typ.(*COTPPacket); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *COTPPacket) GetTypeName() string {
	return "COTPPacket"
}

func (m *COTPPacket) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (headerLength)
	lengthInBits += 8
	// Discriminator Field (tpduCode)
	lengthInBits += 8

	// Length of sub-type elements will be added by sub-type...
	lengthInBits += m.Child.LengthInBits()

	// Array field
	if len(m.Parameters) > 0 {
		for _, element := range m.Parameters {
			lengthInBits += element.LengthInBits()
		}
	}

	// Optional Field (payload)
	if m.Payload != nil {
		lengthInBits += (*m.Payload).LengthInBits()
	}

	return lengthInBits
}

func (m *COTPPacket) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func COTPPacketParse(io *utils.ReadBuffer, cotpLen uint16) (*COTPPacket, error) {
	var startPos = io.GetPos()
	var curPos uint16

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength, _headerLengthErr := io.ReadUint8(8)
	if _headerLengthErr != nil {
		return nil, errors.Wrap(_headerLengthErr, "Error parsing 'headerLength' field")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode, _tpduCodeErr := io.ReadUint8(8)
	if _tpduCodeErr != nil {
		return nil, errors.Wrap(_tpduCodeErr, "Error parsing 'tpduCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *COTPPacket
	var typeSwitchError error
	switch {
	case tpduCode == 0xF0:
		_parent, typeSwitchError = COTPPacketDataParse(io)
	case tpduCode == 0xE0:
		_parent, typeSwitchError = COTPPacketConnectionRequestParse(io)
	case tpduCode == 0xD0:
		_parent, typeSwitchError = COTPPacketConnectionResponseParse(io)
	case tpduCode == 0x80:
		_parent, typeSwitchError = COTPPacketDisconnectRequestParse(io)
	case tpduCode == 0xC0:
		_parent, typeSwitchError = COTPPacketDisconnectResponseParse(io)
	case tpduCode == 0x70:
		_parent, typeSwitchError = COTPPacketTpduErrorParse(io)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Array field (parameters)
	curPos = io.GetPos() - startPos
	// Length array
	parameters := make([]*COTPParameter, 0)
	_parametersLength := uint16(uint16(uint16(headerLength)+uint16(uint16(1)))) - uint16(curPos)
	_parametersEndPos := io.GetPos() + uint16(_parametersLength)
	for io.GetPos() < _parametersEndPos {
		_item, _err := COTPParameterParse(io, uint8(uint8(uint8(headerLength)+uint8(uint8(1))))-uint8(curPos))
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'parameters' field")
		}
		parameters = append(parameters, _item)
		curPos = io.GetPos() - startPos
	}

	// Optional Field (payload) (Can be skipped, if a given expression evaluates to false)
	curPos = io.GetPos() - startPos
	var payload *S7Message = nil
	if bool((curPos) < (cotpLen)) {
		_val, _err := S7MessageParse(io)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'payload' field")
		}
		payload = _val
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, parameters, payload)
	return _parent, nil
}

func (m *COTPPacket) Serialize(io utils.WriteBuffer) error {
	return m.Child.Serialize(io)
}

func (m *COTPPacket) SerializeParent(io utils.WriteBuffer, child ICOTPPacket, serializeChildFunction func() error) error {

	// Implicit Field (headerLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	headerLength := uint8(uint8(uint8(m.LengthInBytes())) - uint8(uint8(uint8(uint8(utils.InlineIf(bool(bool((m.Payload) != (nil))), uint16(m.Payload.LengthInBytes()), uint16(uint8(0)))))+uint8(uint8(1)))))
	_headerLengthErr := io.WriteUint8(8, (headerLength))
	if _headerLengthErr != nil {
		return errors.Wrap(_headerLengthErr, "Error serializing 'headerLength' field")
	}

	// Discriminator Field (tpduCode) (Used as input to a switch field)
	tpduCode := uint8(child.TpduCode())
	_tpduCodeErr := io.WriteUint8(8, (tpduCode))
	if _tpduCodeErr != nil {
		return errors.Wrap(_tpduCodeErr, "Error serializing 'tpduCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Array Field (parameters)
	if m.Parameters != nil {
		for _, _element := range m.Parameters {
			_elementErr := _element.Serialize(io)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'parameters' field")
			}
		}
	}

	// Optional Field (payload) (Can be skipped, if the value is null)
	var payload *S7Message = nil
	if m.Payload != nil {
		payload = m.Payload
		_payloadErr := payload.Serialize(io)
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}
	}

	return nil
}

func (m *COTPPacket) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "parameters":
				var _values []*COTPParameter
				var dt *COTPParameter
				if err := d.DecodeElement(&dt, &tok); err != nil {
					return err
				}
				_values = append(_values, dt)
				m.Parameters = _values
			case "payload":
				var dt *S7Message
				if err := d.DecodeElement(&dt, &tok); err != nil {
					return err
				}
				m.Payload = dt
			default:
				attr := start.Attr
				if attr == nil || len(attr) <= 0 {
					// TODO: workaround for bug with nested lists
					attr = tok.Attr
				}
				if attr == nil || len(attr) <= 0 {
					panic("Couldn't determine class type for childs of COTPPacket")
				}
				switch attr[0].Value {
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketData":
					var dt *COTPPacketData
					if m.Child != nil {
						dt = m.Child.(*COTPPacketData)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketConnectionRequest":
					var dt *COTPPacketConnectionRequest
					if m.Child != nil {
						dt = m.Child.(*COTPPacketConnectionRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketConnectionResponse":
					var dt *COTPPacketConnectionResponse
					if m.Child != nil {
						dt = m.Child.(*COTPPacketConnectionResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketDisconnectRequest":
					var dt *COTPPacketDisconnectRequest
					if m.Child != nil {
						dt = m.Child.(*COTPPacketDisconnectRequest)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketDisconnectResponse":
					var dt *COTPPacketDisconnectResponse
					if m.Child != nil {
						dt = m.Child.(*COTPPacketDisconnectResponse)
					}
					if err := d.DecodeElement(&dt, &tok); err != nil {
						return err
					}
					if m.Child == nil {
						dt.Parent = m
						m.Child = dt
					}
				case "org.apache.plc4x.java.s7.readwrite.COTPPacketTpduError":
					var dt *COTPPacketTpduError
					if m.Child != nil {
						dt = m.Child.(*COTPPacketTpduError)
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

func (m *COTPPacket) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := reflect.TypeOf(m.Child).String()
	className = "org.apache.plc4x.java.s7.readwrite." + className[strings.LastIndex(className, ".")+1:]
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	marshaller, ok := m.Child.(xml.Marshaler)
	if !ok {
		return errors.Errorf("child is not castable to Marshaler. Actual type %T", m.Child)
	}
	if err := marshaller.MarshalXML(e, start); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "parameters"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Parameters, xml.StartElement{Name: xml.Name{Local: "parameters"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "parameters"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Payload, xml.StartElement{Name: xml.Name{Local: "payload"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}
