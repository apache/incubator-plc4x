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
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"io"
)

// The data-structure of this message
type AdsReadDeviceInfoResponse struct {
	Result       ReturnCode
	MajorVersion uint8
	MinorVersion uint8
	Version      uint16
	Device       []int8
	Parent       *AdsData
	IAdsReadDeviceInfoResponse
}

// The corresponding interface
type IAdsReadDeviceInfoResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadDeviceInfoResponse) CommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoResponse) Response() bool {
	return true
}

func (m *AdsReadDeviceInfoResponse) InitializeParent(parent *AdsData) {
}

func NewAdsReadDeviceInfoResponse(result ReturnCode, majorVersion uint8, minorVersion uint8, version uint16, device []int8) *AdsData {
	child := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		Parent:       NewAdsData(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastAdsReadDeviceInfoResponse(structType interface{}) *AdsReadDeviceInfoResponse {
	castFunc := func(typ interface{}) *AdsReadDeviceInfoResponse {
		if casted, ok := typ.(AdsReadDeviceInfoResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadDeviceInfoResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadDeviceInfoResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadDeviceInfoResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadDeviceInfoResponse) GetTypeName() string {
	return "AdsReadDeviceInfoResponse"
}

func (m *AdsReadDeviceInfoResponse) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (result)
	lengthInBits += 32

	// Simple field (majorVersion)
	lengthInBits += 8

	// Simple field (minorVersion)
	lengthInBits += 8

	// Simple field (version)
	lengthInBits += 16

	// Array field
	if len(m.Device) > 0 {
		lengthInBits += 8 * uint16(len(m.Device))
	}

	return lengthInBits
}

func (m *AdsReadDeviceInfoResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsReadDeviceInfoResponseParse(io *utils.ReadBuffer) (*AdsData, error) {

	// Simple Field (result)
	result, _resultErr := ReturnCodeParse(io)
	if _resultErr != nil {
		return nil, errors.New("Error parsing 'result' field " + _resultErr.Error())
	}

	// Simple Field (majorVersion)
	majorVersion, _majorVersionErr := io.ReadUint8(8)
	if _majorVersionErr != nil {
		return nil, errors.New("Error parsing 'majorVersion' field " + _majorVersionErr.Error())
	}

	// Simple Field (minorVersion)
	minorVersion, _minorVersionErr := io.ReadUint8(8)
	if _minorVersionErr != nil {
		return nil, errors.New("Error parsing 'minorVersion' field " + _minorVersionErr.Error())
	}

	// Simple Field (version)
	version, _versionErr := io.ReadUint16(16)
	if _versionErr != nil {
		return nil, errors.New("Error parsing 'version' field " + _versionErr.Error())
	}

	// Array field (device)
	// Count array
	device := make([]int8, uint16(16))
	for curItem := uint16(0); curItem < uint16(uint16(16)); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.New("Error parsing 'device' field " + _err.Error())
		}
		device[curItem] = _item
	}

	// Create a partially initialized instance
	_child := &AdsReadDeviceInfoResponse{
		Result:       result,
		MajorVersion: majorVersion,
		MinorVersion: minorVersion,
		Version:      version,
		Device:       device,
		Parent:       &AdsData{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *AdsReadDeviceInfoResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (result)
		_resultErr := m.Result.Serialize(io)
		if _resultErr != nil {
			return errors.New("Error serializing 'result' field " + _resultErr.Error())
		}

		// Simple Field (majorVersion)
		majorVersion := uint8(m.MajorVersion)
		_majorVersionErr := io.WriteUint8(8, (majorVersion))
		if _majorVersionErr != nil {
			return errors.New("Error serializing 'majorVersion' field " + _majorVersionErr.Error())
		}

		// Simple Field (minorVersion)
		minorVersion := uint8(m.MinorVersion)
		_minorVersionErr := io.WriteUint8(8, (minorVersion))
		if _minorVersionErr != nil {
			return errors.New("Error serializing 'minorVersion' field " + _minorVersionErr.Error())
		}

		// Simple Field (version)
		version := uint16(m.Version)
		_versionErr := io.WriteUint16(16, (version))
		if _versionErr != nil {
			return errors.New("Error serializing 'version' field " + _versionErr.Error())
		}

		// Array Field (device)
		if m.Device != nil {
			for _, _element := range m.Device {
				_elementErr := io.WriteInt8(8, _element)
				if _elementErr != nil {
					return errors.New("Error serializing 'device' field " + _elementErr.Error())
				}
			}
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *AdsReadDeviceInfoResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "version":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Version = data
			case "device":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded := make([]byte, base64.StdEncoding.DecodedLen(len(_encoded)))
				_len, err := base64.StdEncoding.Decode(_decoded, []byte(_encoded))
				if err != nil {
					return err
				}
				m.Device = utils.ByteArrayToInt8Array(_decoded[0:_len])
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

func (m *AdsReadDeviceInfoResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.Result, xml.StartElement{Name: xml.Name{Local: "result"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MajorVersion, xml.StartElement{Name: xml.Name{Local: "majorVersion"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MinorVersion, xml.StartElement{Name: xml.Name{Local: "minorVersion"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Version, xml.StartElement{Name: xml.Name{Local: "version"}}); err != nil {
		return err
	}
	_encodedDevice := make([]byte, base64.StdEncoding.EncodedLen(len(m.Device)))
	base64.StdEncoding.Encode(_encodedDevice, utils.Int8ArrayToByteArray(m.Device))
	if err := e.EncodeElement(_encodedDevice, xml.StartElement{Name: xml.Name{Local: "device"}}); err != nil {
		return err
	}
	return nil
}
