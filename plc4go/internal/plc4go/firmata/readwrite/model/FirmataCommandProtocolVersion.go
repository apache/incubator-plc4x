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
type FirmataCommandProtocolVersion struct {
	MajorVersion uint8
	MinorVersion uint8
	Parent       *FirmataCommand
}

// The corresponding interface
type IFirmataCommandProtocolVersion interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *FirmataCommandProtocolVersion) CommandCode() uint8 {
	return 0x9
}

func (m *FirmataCommandProtocolVersion) InitializeParent(parent *FirmataCommand) {
}

func NewFirmataCommandProtocolVersion(majorVersion uint8, minorVersion uint8) *FirmataCommand {
	child := &FirmataCommandProtocolVersion{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Parent:       NewFirmataCommand(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastFirmataCommandProtocolVersion(structType interface{}) *FirmataCommandProtocolVersion {
	castFunc := func(typ interface{}) *FirmataCommandProtocolVersion {
		if casted, ok := typ.(FirmataCommandProtocolVersion); ok {
			return &casted
		}
		if casted, ok := typ.(*FirmataCommandProtocolVersion); ok {
			return casted
		}
		if casted, ok := typ.(FirmataCommand); ok {
			return CastFirmataCommandProtocolVersion(casted.Child)
		}
		if casted, ok := typ.(*FirmataCommand); ok {
			return CastFirmataCommandProtocolVersion(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *FirmataCommandProtocolVersion) GetTypeName() string {
	return "FirmataCommandProtocolVersion"
}

func (m *FirmataCommandProtocolVersion) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *FirmataCommandProtocolVersion) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	return lengthInBits
}

func (m *FirmataCommandProtocolVersion) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func FirmataCommandProtocolVersionParse(io utils.ReadBuffer) (*FirmataCommand, error) {
	if pullErr := io.PullContext("FirmataCommandProtocolVersion"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (majorVersion)
	majorVersion, _majorVersionErr := io.ReadUint8("majorVersion", 8)
	if _majorVersionErr != nil {
		return nil, errors.Wrap(_majorVersionErr, "Error parsing 'majorVersion' field")
	}

	// Simple Field (minorVersion)
	minorVersion, _minorVersionErr := io.ReadUint8("minorVersion", 8)
	if _minorVersionErr != nil {
		return nil, errors.Wrap(_minorVersionErr, "Error parsing 'minorVersion' field")
	}

	if closeErr := io.CloseContext("FirmataCommandProtocolVersion"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &FirmataCommandProtocolVersion{
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Parent:       &FirmataCommand{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *FirmataCommandProtocolVersion) Serialize(io utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := io.PushContext("FirmataCommandProtocolVersion"); pushErr != nil {
			return pushErr
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := io.WriteUint8("majorVersion", 8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.Wrap(_majorVersionErr, "Error serializing 'majorVersion' field")
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := io.WriteUint8("minorVersion", 8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.Wrap(_minorVersionErr, "Error serializing 'minorVersion' field")
		}

		if popErr := io.PopContext("FirmataCommandProtocolVersion"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

// Deprecated: the utils.ReadBufferWriteBased should be used instead
func (m *FirmataCommandProtocolVersion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "majorVersion":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MajorVersion = data
			case "minorVersion":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MinorVersion = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

// Deprecated: the utils.WriteBufferReadBased should be used instead
func (m *FirmataCommandProtocolVersion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.MajorVersion, xml.StartElement{Name: xml.Name{Local: "majorVersion"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MinorVersion, xml.StartElement{Name: xml.Name{Local: "minorVersion"}}); err != nil {
		return err
	}
	return nil
}

func (m FirmataCommandProtocolVersion) String() string {
	return string(m.Box("", 120))
}

// Deprecated: the utils.WriteBufferBoxBased should be used instead
func (m FirmataCommandProtocolVersion) Box(name string, width int) utils.AsciiBox {
	boxName := "FirmataCommandProtocolVersion"
	if name != "" {
		boxName += "/" + name
	}
	childBoxer := func() []utils.AsciiBox {
		boxes := make([]utils.AsciiBox, 0)
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("MajorVersion", m.MajorVersion, -1))
		// Simple field (case simple)
		// uint8 can be boxed as anything with the least amount of space
		boxes = append(boxes, utils.BoxAnything("MinorVersion", m.MinorVersion, -1))
		return boxes
	}
	return m.Parent.BoxParent(boxName, width, childBoxer)
}