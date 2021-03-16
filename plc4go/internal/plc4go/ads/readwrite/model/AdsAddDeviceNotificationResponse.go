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
type AdsAddDeviceNotificationResponse struct {
	Result             ReturnCode
	NotificationHandle uint32
	Parent             *AdsData
}

// The corresponding interface
type IAdsAddDeviceNotificationResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsAddDeviceNotificationResponse) CommandId() CommandId {
	return ADS_ADD_DEVICE_NOTIFICATION
}

func (m *AdsAddDeviceNotificationResponse) Response() bool {
	return true
}

func (m *AdsAddDeviceNotificationResponse) InitializeParent(parent *AdsData) {
}

func NewAdsAddDeviceNotificationResponse(result ReturnCode, notificationHandle uint32) *AdsData {
	child := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		Parent:             NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsAddDeviceNotificationResponse(structType interface{}) *AdsAddDeviceNotificationResponse {
	castFunc := func(typ interface{}) *AdsAddDeviceNotificationResponse {
		if casted, ok := typ.(AdsAddDeviceNotificationResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsAddDeviceNotificationResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsAddDeviceNotificationResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsAddDeviceNotificationResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsAddDeviceNotificationResponse) GetTypeName() string {
	return "AdsAddDeviceNotificationResponse"
}

func (m *AdsAddDeviceNotificationResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (result)
	lengthInBits += 32

	// Simple field (notificationHandle)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsAddDeviceNotificationResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsAddDeviceNotificationResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}

	// Simple Field (notificationHandle)
	notificationHandle, _notificationHandleErr := io.ReadUint32(32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}

	// Create a partially initialized instance
	_child := &AdsAddDeviceNotificationResponse{
		Result:             result,
		NotificationHandle: notificationHandle,
		Parent:             &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsAddDeviceNotificationResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (result)
		_resultErr := m.Result.Serialize(io)
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Simple Field (notificationHandle)
		notificationHandle := uint32(m.NotificationHandle)
		_notificationHandleErr := io.WriteUint32(32, (notificationHandle))
		if _notificationHandleErr != nil {
			return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsAddDeviceNotificationResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "result":
				var data ReturnCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Result = data
			case "notificationHandle":
				var data uint32
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NotificationHandle = data
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

func (m *AdsAddDeviceNotificationResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.NotificationHandle, xml.StartElement{Name: xml.Name{Local: "notificationHandle"}}); err != nil {
		return err
	}
	return nil
}
