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
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type ApduDataExt struct {
	Child IApduDataExtChild
}

// The corresponding interface
type IApduDataExt interface {
	ExtApciType() uint8
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IApduDataExtParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error
	GetTypeName() string
}

type IApduDataExtChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ApduDataExt)
	GetTypeName() string
	IApduDataExt
}

func NewApduDataExt() *ApduDataExt {
	return &ApduDataExt{}
}

func CastApduDataExt(structType interface{}) *ApduDataExt {
	castFunc := func(typ interface{}) *ApduDataExt {
		if casted, ok := typ.(ApduDataExt); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExt) GetTypeName() string {
	return "ApduDataExt"
}

func (m *ApduDataExt) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExt) LengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.LengthInBits()
}

func (m *ApduDataExt) ParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (extApciType)
	lengthInBits += 6

	return lengthInBits
}

func (m *ApduDataExt) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExt"); pullErr != nil {
		return nil, pullErr
	}

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType, _extApciTypeErr := readBuffer.ReadUint8("extApciType", 6)
	if _extApciTypeErr != nil {
		return nil, errors.Wrap(_extApciTypeErr, "Error parsing 'extApciType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ApduDataExt
	var typeSwitchError error
	switch {
	case extApciType == 0x00: // ApduDataExtOpenRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtOpenRoutingTableRequestParse(readBuffer)
	case extApciType == 0x01: // ApduDataExtReadRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtReadRoutingTableRequestParse(readBuffer)
	case extApciType == 0x02: // ApduDataExtReadRoutingTableResponse
		_parent, typeSwitchError = ApduDataExtReadRoutingTableResponseParse(readBuffer)
	case extApciType == 0x03: // ApduDataExtWriteRoutingTableRequest
		_parent, typeSwitchError = ApduDataExtWriteRoutingTableRequestParse(readBuffer)
	case extApciType == 0x08: // ApduDataExtReadRouterMemoryRequest
		_parent, typeSwitchError = ApduDataExtReadRouterMemoryRequestParse(readBuffer)
	case extApciType == 0x09: // ApduDataExtReadRouterMemoryResponse
		_parent, typeSwitchError = ApduDataExtReadRouterMemoryResponseParse(readBuffer)
	case extApciType == 0x0A: // ApduDataExtWriteRouterMemoryRequest
		_parent, typeSwitchError = ApduDataExtWriteRouterMemoryRequestParse(readBuffer)
	case extApciType == 0x0D: // ApduDataExtReadRouterStatusRequest
		_parent, typeSwitchError = ApduDataExtReadRouterStatusRequestParse(readBuffer)
	case extApciType == 0x0E: // ApduDataExtReadRouterStatusResponse
		_parent, typeSwitchError = ApduDataExtReadRouterStatusResponseParse(readBuffer)
	case extApciType == 0x0F: // ApduDataExtWriteRouterStatusRequest
		_parent, typeSwitchError = ApduDataExtWriteRouterStatusRequestParse(readBuffer)
	case extApciType == 0x10: // ApduDataExtMemoryBitWrite
		_parent, typeSwitchError = ApduDataExtMemoryBitWriteParse(readBuffer)
	case extApciType == 0x11: // ApduDataExtAuthorizeRequest
		_parent, typeSwitchError = ApduDataExtAuthorizeRequestParse(readBuffer)
	case extApciType == 0x12: // ApduDataExtAuthorizeResponse
		_parent, typeSwitchError = ApduDataExtAuthorizeResponseParse(readBuffer)
	case extApciType == 0x13: // ApduDataExtKeyWrite
		_parent, typeSwitchError = ApduDataExtKeyWriteParse(readBuffer)
	case extApciType == 0x14: // ApduDataExtKeyResponse
		_parent, typeSwitchError = ApduDataExtKeyResponseParse(readBuffer)
	case extApciType == 0x15: // ApduDataExtPropertyValueRead
		_parent, typeSwitchError = ApduDataExtPropertyValueReadParse(readBuffer)
	case extApciType == 0x16: // ApduDataExtPropertyValueResponse
		_parent, typeSwitchError = ApduDataExtPropertyValueResponseParse(readBuffer, length)
	case extApciType == 0x17: // ApduDataExtPropertyValueWrite
		_parent, typeSwitchError = ApduDataExtPropertyValueWriteParse(readBuffer, length)
	case extApciType == 0x18: // ApduDataExtPropertyDescriptionRead
		_parent, typeSwitchError = ApduDataExtPropertyDescriptionReadParse(readBuffer)
	case extApciType == 0x19: // ApduDataExtPropertyDescriptionResponse
		_parent, typeSwitchError = ApduDataExtPropertyDescriptionResponseParse(readBuffer)
	case extApciType == 0x1A: // ApduDataExtNetworkParameterRead
		_parent, typeSwitchError = ApduDataExtNetworkParameterReadParse(readBuffer)
	case extApciType == 0x1B: // ApduDataExtNetworkParameterResponse
		_parent, typeSwitchError = ApduDataExtNetworkParameterResponseParse(readBuffer)
	case extApciType == 0x1C: // ApduDataExtIndividualAddressSerialNumberRead
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberReadParse(readBuffer)
	case extApciType == 0x1D: // ApduDataExtIndividualAddressSerialNumberResponse
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberResponseParse(readBuffer)
	case extApciType == 0x1E: // ApduDataExtIndividualAddressSerialNumberWrite
		_parent, typeSwitchError = ApduDataExtIndividualAddressSerialNumberWriteParse(readBuffer)
	case extApciType == 0x20: // ApduDataExtDomainAddressWrite
		_parent, typeSwitchError = ApduDataExtDomainAddressWriteParse(readBuffer)
	case extApciType == 0x21: // ApduDataExtDomainAddressRead
		_parent, typeSwitchError = ApduDataExtDomainAddressReadParse(readBuffer)
	case extApciType == 0x22: // ApduDataExtDomainAddressResponse
		_parent, typeSwitchError = ApduDataExtDomainAddressResponseParse(readBuffer)
	case extApciType == 0x23: // ApduDataExtDomainAddressSelectiveRead
		_parent, typeSwitchError = ApduDataExtDomainAddressSelectiveReadParse(readBuffer)
	case extApciType == 0x24: // ApduDataExtNetworkParameterWrite
		_parent, typeSwitchError = ApduDataExtNetworkParameterWriteParse(readBuffer)
	case extApciType == 0x25: // ApduDataExtLinkRead
		_parent, typeSwitchError = ApduDataExtLinkReadParse(readBuffer)
	case extApciType == 0x26: // ApduDataExtLinkResponse
		_parent, typeSwitchError = ApduDataExtLinkResponseParse(readBuffer)
	case extApciType == 0x27: // ApduDataExtLinkWrite
		_parent, typeSwitchError = ApduDataExtLinkWriteParse(readBuffer)
	case extApciType == 0x28: // ApduDataExtGroupPropertyValueRead
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueReadParse(readBuffer)
	case extApciType == 0x29: // ApduDataExtGroupPropertyValueResponse
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueResponseParse(readBuffer)
	case extApciType == 0x2A: // ApduDataExtGroupPropertyValueWrite
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueWriteParse(readBuffer)
	case extApciType == 0x2B: // ApduDataExtGroupPropertyValueInfoReport
		_parent, typeSwitchError = ApduDataExtGroupPropertyValueInfoReportParse(readBuffer)
	case extApciType == 0x2C: // ApduDataExtDomainAddressSerialNumberRead
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberReadParse(readBuffer)
	case extApciType == 0x2D: // ApduDataExtDomainAddressSerialNumberResponse
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberResponseParse(readBuffer)
	case extApciType == 0x2E: // ApduDataExtDomainAddressSerialNumberWrite
		_parent, typeSwitchError = ApduDataExtDomainAddressSerialNumberWriteParse(readBuffer)
	case extApciType == 0x30: // ApduDataExtFileStreamInfoReport
		_parent, typeSwitchError = ApduDataExtFileStreamInfoReportParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ApduDataExt"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ApduDataExt) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ApduDataExt) SerializeParent(writeBuffer utils.WriteBuffer, child IApduDataExt, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ApduDataExt"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType := uint8(child.ExtApciType())
	_extApciTypeErr := writeBuffer.WriteUint8("extApciType", 6, (extApciType))

	if _extApciTypeErr != nil {
		return errors.Wrap(_extApciTypeErr, "Error serializing 'extApciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	_typeSwitchErr := serializeChildFunction()
	if _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduDataExt"); popErr != nil {
		return popErr
	}
	return nil
}
