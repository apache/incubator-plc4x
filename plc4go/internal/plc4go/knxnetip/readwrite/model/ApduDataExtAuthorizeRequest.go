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
)

// The data-structure of this message
type ApduDataExtAuthorizeRequest struct {
	Level  uint8
	Data   []uint8
	Parent *ApduDataExt
	IApduDataExtAuthorizeRequest
}

// The corresponding interface
type IApduDataExtAuthorizeRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtAuthorizeRequest) ExtApciType() uint8 {
	return 0x11
}

func (m *ApduDataExtAuthorizeRequest) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtAuthorizeRequest(level uint8, data []uint8) *ApduDataExt {
	child := &ApduDataExtAuthorizeRequest{
		Level:  level,
		Data:   data,
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtAuthorizeRequest(structType interface{}) *ApduDataExtAuthorizeRequest {
	castFunc := func(typ interface{}) *ApduDataExtAuthorizeRequest {
		if casted, ok := typ.(ApduDataExtAuthorizeRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtAuthorizeRequest); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtAuthorizeRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtAuthorizeRequest) GetTypeName() string {
	return "ApduDataExtAuthorizeRequest"
}

func (m *ApduDataExtAuthorizeRequest) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (level)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *ApduDataExtAuthorizeRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtAuthorizeRequestParse(io *utils.ReadBuffer) (*ApduDataExt, error) {

	// Simple Field (level)
	level, _levelErr := io.ReadUint8(8)
	if _levelErr != nil {
		return nil, errors.New("Error parsing 'level' field " + _levelErr.Error())
	}

	// Array field (data)
	// Count array
	data := make([]uint8, uint16(4))
	for curItem := uint16(0); curItem < uint16(uint16(4)); curItem++ {
		_item, _err := io.ReadUint8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'data' field " + _err.Error())
		}
		data[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &ApduDataExtAuthorizeRequest{
		Level:  level,
		Data:   data,
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtAuthorizeRequest) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (level)
		level := uint8(m.Level)
		_levelErr := io.WriteUint8(8, (level))
		if _levelErr != nil {
			return errors.New("Error serializing 'level' field " + _levelErr.Error())
		}

		// Array Field (data)
		if m.Data != nil {
			for _, _element := range m.Data {
				_elementErr := io.WriteUint8(8, _element)
				if _elementErr != nil {
					return errors.New("Error serializing 'data' field " + _elementErr.Error())
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *ApduDataExtAuthorizeRequest) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "level":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Level = data
			case "data":
				var data []uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Data = data
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

func (m *ApduDataExtAuthorizeRequest) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Level, xml.StartElement{Name: xml.Name{Local: "level"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Data, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	return nil
}
