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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by build-utils. DO NOT EDIT.

// Constant values.
const BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER uint8 = 0x0B
const BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER uint8 = 0x1B
const BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER uint8 = 0x3D

// The data-structure of this message
type BACnetUnconfirmedServiceRequestWhoHas struct {
	DeviceInstanceLowLimit  uint32
	DeviceInstanceHighLimit uint32
	ObjectNameCharacterSet  uint8
	ObjectName              []int8
	Parent                  *BACnetUnconfirmedServiceRequest
}

// The corresponding interface
type IBACnetUnconfirmedServiceRequestWhoHas interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetUnconfirmedServiceRequestWhoHas) ServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func NewBACnetUnconfirmedServiceRequestWhoHas(deviceInstanceLowLimit uint32, deviceInstanceHighLimit uint32, objectNameCharacterSet uint8, objectName []int8) *BACnetUnconfirmedServiceRequest {
	child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceLowLimit:  deviceInstanceLowLimit,
		DeviceInstanceHighLimit: deviceInstanceHighLimit,
		ObjectNameCharacterSet:  objectNameCharacterSet,
		ObjectName:              objectName,
		Parent:                  NewBACnetUnconfirmedServiceRequest(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastBACnetUnconfirmedServiceRequestWhoHas(structType interface{}) *BACnetUnconfirmedServiceRequestWhoHas {
	castFunc := func(typ interface{}) *BACnetUnconfirmedServiceRequestWhoHas {
		if casted, ok := typ.(BACnetUnconfirmedServiceRequestWhoHas); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequestWhoHas); ok {
			return casted
		}
		if casted, ok := typ.(BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
		}
		if casted, ok := typ.(*BACnetUnconfirmedServiceRequest); ok {
			return CastBACnetUnconfirmedServiceRequestWhoHas(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestWhoHas"
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Const Field (deviceInstanceLowLimitHeader)
	lengthInBits += 8

	// Simple field (deviceInstanceLowLimit)
	lengthInBits += 24

	// Const Field (deviceInstanceHighLimitHeader)
	lengthInBits += 8

	// Simple field (deviceInstanceHighLimit)
	lengthInBits += 24

	// Const Field (objectNameHeader)
	lengthInBits += 8

	// Implicit Field (objectNameLength)
	lengthInBits += 8

	// Simple field (objectNameCharacterSet)
	lengthInBits += 8

	// Array field
	if len(m.ObjectName) > 0 {
		lengthInBits += 8 * uint16(len(m.ObjectName))
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestWhoHasParse(readBuffer utils.ReadBuffer) (*BACnetUnconfirmedServiceRequest, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestWhoHas"); pullErr != nil {
		return nil, pullErr
	}

	// Const Field (deviceInstanceLowLimitHeader)
	deviceInstanceLowLimitHeader, _deviceInstanceLowLimitHeaderErr := readBuffer.ReadUint8("deviceInstanceLowLimitHeader", 8)
	if _deviceInstanceLowLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceLowLimitHeaderErr, "Error parsing 'deviceInstanceLowLimitHeader' field")
	}
	if deviceInstanceLowLimitHeader != BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCELOWLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceLowLimitHeader))
	}

	// Simple Field (deviceInstanceLowLimit)
	deviceInstanceLowLimit, _deviceInstanceLowLimitErr := readBuffer.ReadUint32("deviceInstanceLowLimit", 24)
	if _deviceInstanceLowLimitErr != nil {
		return nil, errors.Wrap(_deviceInstanceLowLimitErr, "Error parsing 'deviceInstanceLowLimit' field")
	}

	// Const Field (deviceInstanceHighLimitHeader)
	deviceInstanceHighLimitHeader, _deviceInstanceHighLimitHeaderErr := readBuffer.ReadUint8("deviceInstanceHighLimitHeader", 8)
	if _deviceInstanceHighLimitHeaderErr != nil {
		return nil, errors.Wrap(_deviceInstanceHighLimitHeaderErr, "Error parsing 'deviceInstanceHighLimitHeader' field")
	}
	if deviceInstanceHighLimitHeader != BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_DEVICEINSTANCEHIGHLIMITHEADER) + " but got " + fmt.Sprintf("%d", deviceInstanceHighLimitHeader))
	}

	// Simple Field (deviceInstanceHighLimit)
	deviceInstanceHighLimit, _deviceInstanceHighLimitErr := readBuffer.ReadUint32("deviceInstanceHighLimit", 24)
	if _deviceInstanceHighLimitErr != nil {
		return nil, errors.Wrap(_deviceInstanceHighLimitErr, "Error parsing 'deviceInstanceHighLimit' field")
	}

	// Const Field (objectNameHeader)
	objectNameHeader, _objectNameHeaderErr := readBuffer.ReadUint8("objectNameHeader", 8)
	if _objectNameHeaderErr != nil {
		return nil, errors.Wrap(_objectNameHeaderErr, "Error parsing 'objectNameHeader' field")
	}
	if objectNameHeader != BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BACnetUnconfirmedServiceRequestWhoHas_OBJECTNAMEHEADER) + " but got " + fmt.Sprintf("%d", objectNameHeader))
	}

	// Implicit Field (objectNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	objectNameLength, _objectNameLengthErr := readBuffer.ReadUint8("objectNameLength", 8)
	_ = objectNameLength
	if _objectNameLengthErr != nil {
		return nil, errors.Wrap(_objectNameLengthErr, "Error parsing 'objectNameLength' field")
	}

	// Simple Field (objectNameCharacterSet)
	objectNameCharacterSet, _objectNameCharacterSetErr := readBuffer.ReadUint8("objectNameCharacterSet", 8)
	if _objectNameCharacterSetErr != nil {
		return nil, errors.Wrap(_objectNameCharacterSetErr, "Error parsing 'objectNameCharacterSet' field")
	}

	// Array field (objectName)
	if pullErr := readBuffer.PullContext("objectName", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	objectName := make([]int8, 0)
	_objectNameLength := uint16(objectNameLength) - uint16(uint16(1))
	_objectNameEndPos := readBuffer.GetPos() + uint16(_objectNameLength)
	for readBuffer.GetPos() < _objectNameEndPos {
		_item, _err := readBuffer.ReadInt8("", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'objectName' field")
		}
		objectName = append(objectName, _item)
	}
	if closeErr := readBuffer.CloseContext("objectName", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestWhoHas"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestWhoHas{
		DeviceInstanceLowLimit:  deviceInstanceLowLimit,
		DeviceInstanceHighLimit: deviceInstanceHighLimit,
		ObjectNameCharacterSet:  objectNameCharacterSet,
		ObjectName:              objectName,
		Parent:                  &BACnetUnconfirmedServiceRequest{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *BACnetUnconfirmedServiceRequestWhoHas) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestWhoHas"); pushErr != nil {
			return pushErr
		}

		// Const Field (deviceInstanceLowLimitHeader)
		_deviceInstanceLowLimitHeaderErr := writeBuffer.WriteUint8("deviceInstanceLowLimitHeader", 8, 0x0B)
		if _deviceInstanceLowLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceLowLimitHeaderErr, "Error serializing 'deviceInstanceLowLimitHeader' field")
		}

		// Simple Field (deviceInstanceLowLimit)
		deviceInstanceLowLimit := uint32(m.DeviceInstanceLowLimit)
		_deviceInstanceLowLimitErr := writeBuffer.WriteUint32("deviceInstanceLowLimit", 24, (deviceInstanceLowLimit))
		if _deviceInstanceLowLimitErr != nil {
			return errors.Wrap(_deviceInstanceLowLimitErr, "Error serializing 'deviceInstanceLowLimit' field")
		}

		// Const Field (deviceInstanceHighLimitHeader)
		_deviceInstanceHighLimitHeaderErr := writeBuffer.WriteUint8("deviceInstanceHighLimitHeader", 8, 0x1B)
		if _deviceInstanceHighLimitHeaderErr != nil {
			return errors.Wrap(_deviceInstanceHighLimitHeaderErr, "Error serializing 'deviceInstanceHighLimitHeader' field")
		}

		// Simple Field (deviceInstanceHighLimit)
		deviceInstanceHighLimit := uint32(m.DeviceInstanceHighLimit)
		_deviceInstanceHighLimitErr := writeBuffer.WriteUint32("deviceInstanceHighLimit", 24, (deviceInstanceHighLimit))
		if _deviceInstanceHighLimitErr != nil {
			return errors.Wrap(_deviceInstanceHighLimitErr, "Error serializing 'deviceInstanceHighLimit' field")
		}

		// Const Field (objectNameHeader)
		_objectNameHeaderErr := writeBuffer.WriteUint8("objectNameHeader", 8, 0x3D)
		if _objectNameHeaderErr != nil {
			return errors.Wrap(_objectNameHeaderErr, "Error serializing 'objectNameHeader' field")
		}

		// Implicit Field (objectNameLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		objectNameLength := uint8(uint8(uint8(len(m.ObjectName))) + uint8(uint8(1)))
		_objectNameLengthErr := writeBuffer.WriteUint8("objectNameLength", 8, (objectNameLength))
		if _objectNameLengthErr != nil {
			return errors.Wrap(_objectNameLengthErr, "Error serializing 'objectNameLength' field")
		}

		// Simple Field (objectNameCharacterSet)
		objectNameCharacterSet := uint8(m.ObjectNameCharacterSet)
		_objectNameCharacterSetErr := writeBuffer.WriteUint8("objectNameCharacterSet", 8, (objectNameCharacterSet))
		if _objectNameCharacterSetErr != nil {
			return errors.Wrap(_objectNameCharacterSetErr, "Error serializing 'objectNameCharacterSet' field")
		}

		// Array Field (objectName)
		if m.ObjectName != nil {
			if pushErr := writeBuffer.PushContext("objectName", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ObjectName {
				_elementErr := writeBuffer.WriteInt8("", 8, _element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'objectName' field")
				}
			}
			if popErr := writeBuffer.PopContext("objectName", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestWhoHas"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}
