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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduControlAck struct {
	Parent *ApduControl
}

// The corresponding interface
type IApduControlAck interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduControlAck) ControlType() uint8 {
	return 0x2
}

func (m *ApduControlAck) InitializeParent(parent *ApduControl) {
}

func NewApduControlAck() *ApduControl {
	child := &ApduControlAck{
		Parent: NewApduControl(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduControlAck(structType interface{}) *ApduControlAck {
	castFunc := func(typ interface{}) *ApduControlAck {
		if casted, ok := typ.(ApduControlAck); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduControlAck); ok {
			return casted
		}
		if casted, ok := typ.(ApduControl); ok {
			return CastApduControlAck(casted.Child)
		}
		if casted, ok := typ.(*ApduControl); ok {
			return CastApduControlAck(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduControlAck) GetTypeName() string {
	return "ApduControlAck"
}

func (m *ApduControlAck) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduControlAck) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlAck) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduControlAckParse(readBuffer utils.ReadBuffer) (*ApduControl, error) {
	if pullErr := readBuffer.PullContext("ApduControlAck"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduControlAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduControlAck{
		Parent: &ApduControl{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduControlAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlAck"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduControlAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
