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
	"encoding/hex"
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	"math"
	"strings"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7VarPayloadDataItem struct {
	ReturnCode    DataTransportErrorCode
	TransportSize DataTransportSize
	Data          []int8
}

// The corresponding interface
type IS7VarPayloadDataItem interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

func NewS7VarPayloadDataItem(returnCode DataTransportErrorCode, transportSize DataTransportSize, data []int8) *S7VarPayloadDataItem {
	return &S7VarPayloadDataItem{ReturnCode: returnCode, TransportSize: transportSize, Data: data}
}

func CastS7VarPayloadDataItem(structType interface{}) *S7VarPayloadDataItem {
	castFunc := func(typ interface{}) *S7VarPayloadDataItem {
		if casted, ok := typ.(S7VarPayloadDataItem); ok {
			return &casted
		}
		if casted, ok := typ.(*S7VarPayloadDataItem); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7VarPayloadDataItem) GetTypeName() string {
	return "S7VarPayloadDataItem"
}

func (m *S7VarPayloadDataItem) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Enum Field (returnCode)
	lengthInBits += 8

	// Enum Field (transportSize)
	lengthInBits += 8

	// Implicit Field (dataLength)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	// Padding Field (padding)
	_timesPadding := uint8(utils.InlineIf(false, uint16(uint8(0)), uint16(uint8(uint8(len(m.Data)))%uint8(uint8(2)))))
	for ; _timesPadding > 0; _timesPadding-- {
		lengthInBits += 8
	}

	return lengthInBits
}

func (m *S7VarPayloadDataItem) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7VarPayloadDataItemParse(io *utils.ReadBuffer, lastItem bool) (*S7VarPayloadDataItem, error) {

	// Enum field (returnCode)
	returnCode, _returnCodeErr := DataTransportErrorCodeParse(io)
	if _returnCodeErr != nil {
		return nil, errors.Wrap(_returnCodeErr, "Error parsing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize, _transportSizeErr := DataTransportSizeParse(io)
	if _transportSizeErr != nil {
		return nil, errors.Wrap(_transportSizeErr, "Error parsing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := io.ReadUint16(16)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Array field (data)
	// Count array
	data := make([]int8, utils.InlineIf(transportSize.SizeInBits(), uint16(math.Ceil(float64(dataLength)/float64(float64(8.0)))), uint16(dataLength)))
	for curItem := uint16(0); curItem < uint16(utils.InlineIf(transportSize.SizeInBits(), uint16(math.Ceil(float64(dataLength)/float64(float64(8.0)))), uint16(dataLength))); curItem++ {
		_item, _err := io.ReadInt8(8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'data' field")
		}
		data[curItem] = _item
	}

	// Padding Field (padding)
	{
		_timesPadding := (utils.InlineIf(lastItem, uint16(uint8(0)), uint16(uint8(uint8(len(data)))%uint8(uint8(2)))))
		for ; (io.HasMore(8)) && (_timesPadding > 0); _timesPadding-- {
			// Just read the padding data and ignore it
			_, _err := io.ReadUint8(8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'padding' field")
			}
		}
	}

	// Create the instance
	return NewS7VarPayloadDataItem(returnCode, transportSize, data), nil
}

func (m *S7VarPayloadDataItem) Serialize(io utils.WriteBuffer, lastItem bool) error {

	// Enum field (returnCode)
	returnCode := CastDataTransportErrorCode(m.ReturnCode)
	_returnCodeErr := returnCode.Serialize(io)
	if _returnCodeErr != nil {
		return errors.Wrap(_returnCodeErr, "Error serializing 'returnCode' field")
	}

	// Enum field (transportSize)
	transportSize := CastDataTransportSize(m.TransportSize)
	_transportSizeErr := transportSize.Serialize(io)
	if _transportSizeErr != nil {
		return errors.Wrap(_transportSizeErr, "Error serializing 'transportSize' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	dataLength := uint16(uint16(uint16(len(m.Data))) * uint16(uint16(utils.InlineIf(bool(bool((m.TransportSize) == (DataTransportSize_BIT))), uint16(uint16(1)), uint16(uint16(utils.InlineIf(transportSize.SizeInBits(), uint16(uint16(8)), uint16(uint16(1)))))))))
	_dataLengthErr := io.WriteUint16(16, (dataLength))
	if _dataLengthErr != nil {
		return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
	}

	// Array Field (data)
	if m.Data != nil {
		for _, _element := range m.Data {
			_elementErr := io.WriteInt8(8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
	}

	// Padding Field (padding)
	{
		_timesPadding := uint8(utils.InlineIf(lastItem, uint16(uint8(0)), uint16(uint8(uint8(len(m.Data)))%uint8(uint8(2)))))
		for ; _timesPadding > 0; _timesPadding-- {
			_paddingValue := uint8(uint8(0))
			_paddingErr := io.WriteUint8(8, (_paddingValue))
			if _paddingErr != nil {
				return errors.Wrap(_paddingErr, "Error serializing 'padding' field")
			}
		}
	}

	return nil
}

func (m *S7VarPayloadDataItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "returnCode":
				var data DataTransportErrorCode
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReturnCode = data
			case "transportSize":
				var data DataTransportSize
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransportSize = data
			case "data":
				var _encoded string
				if err := d.DecodeElement(&_encoded, &tok); err != nil {
					return err
				}
				_decoded, err := hex.DecodeString(_encoded)
				_len := len(_decoded)
				if err != nil {
					return err
				}
				m.Data = utils.ByteArrayToInt8Array(_decoded[0:_len])
			}
		}
	}
}

func (m *S7VarPayloadDataItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.s7.readwrite.S7VarPayloadDataItem"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReturnCode, xml.StartElement{Name: xml.Name{Local: "returnCode"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransportSize, xml.StartElement{Name: xml.Name{Local: "transportSize"}}); err != nil {
		return err
	}
	_encodedData := hex.EncodeToString(utils.Int8ArrayToByteArray(m.Data))
	_encodedData = strings.ToUpper(_encodedData)
	if err := e.EncodeElement(_encodedData, xml.StartElement{Name: xml.Name{Local: "data"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}

func (m S7VarPayloadDataItem) String() string {
	return string(m.Box("S7VarPayloadDataItem", utils.DefaultWidth*2))
}

func (m S7VarPayloadDataItem) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7VarPayloadDataItem"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("ReturnCode", m.ReturnCode, width-2))
	boxes = append(boxes, utils.BoxAnything("TransportSize", m.TransportSize, width-2))
	boxes = append(boxes, utils.BoxAnything("Data", m.Data, width-2))
	return utils.BoxString(name, string(utils.AlignBoxes(boxes, width-2)), width)
}
