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
type AmsSerialResetFrame struct {
	MagicCookie        uint16
	TransmitterAddress int8
	ReceiverAddress    int8
	FragmentNumber     int8
	Length             int8
	Crc                uint16
	IAmsSerialResetFrame
}

// The corresponding interface
type IAmsSerialResetFrame interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
}

func NewAmsSerialResetFrame(magicCookie uint16, transmitterAddress int8, receiverAddress int8, fragmentNumber int8, length int8, crc uint16) *AmsSerialResetFrame {
	return &AmsSerialResetFrame{MagicCookie: magicCookie, TransmitterAddress: transmitterAddress, ReceiverAddress: receiverAddress, FragmentNumber: fragmentNumber, Length: length, Crc: crc}
}

func CastAmsSerialResetFrame(structType interface{}) *AmsSerialResetFrame {
	castFunc := func(typ interface{}) *AmsSerialResetFrame {
		if casted, ok := typ.(AmsSerialResetFrame); ok {
			return &casted
		}
		if casted, ok := typ.(*AmsSerialResetFrame); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AmsSerialResetFrame) GetTypeName() string {
	return "AmsSerialResetFrame"
}

func (m *AmsSerialResetFrame) LengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (magicCookie)
	lengthInBits += 16

	// Simple field (transmitterAddress)
	lengthInBits += 8

	// Simple field (receiverAddress)
	lengthInBits += 8

	// Simple field (fragmentNumber)
	lengthInBits += 8

	// Simple field (length)
	lengthInBits += 8

	// Simple field (crc)
	lengthInBits += 16

	return lengthInBits
}

func (m *AmsSerialResetFrame) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AmsSerialResetFrameParse(io *utils.ReadBuffer) (*AmsSerialResetFrame, error) {

	// Simple Field (magicCookie)
	magicCookie, _magicCookieErr := io.ReadUint16(16)
	if _magicCookieErr != nil {
		return nil, errors.New("Error parsing 'magicCookie' field " + _magicCookieErr.Error())
	}

	// Simple Field (transmitterAddress)
	transmitterAddress, _transmitterAddressErr := io.ReadInt8(8)
	if _transmitterAddressErr != nil {
		return nil, errors.New("Error parsing 'transmitterAddress' field " + _transmitterAddressErr.Error())
	}

	// Simple Field (receiverAddress)
	receiverAddress, _receiverAddressErr := io.ReadInt8(8)
	if _receiverAddressErr != nil {
		return nil, errors.New("Error parsing 'receiverAddress' field " + _receiverAddressErr.Error())
	}

	// Simple Field (fragmentNumber)
	fragmentNumber, _fragmentNumberErr := io.ReadInt8(8)
	if _fragmentNumberErr != nil {
		return nil, errors.New("Error parsing 'fragmentNumber' field " + _fragmentNumberErr.Error())
	}

	// Simple Field (length)
	length, _lengthErr := io.ReadInt8(8)
	if _lengthErr != nil {
		return nil, errors.New("Error parsing 'length' field " + _lengthErr.Error())
	}

	// Simple Field (crc)
	crc, _crcErr := io.ReadUint16(16)
	if _crcErr != nil {
		return nil, errors.New("Error parsing 'crc' field " + _crcErr.Error())
	}

	// Create the instance
	return NewAmsSerialResetFrame(magicCookie, transmitterAddress, receiverAddress, fragmentNumber, length, crc), nil
}

func (m *AmsSerialResetFrame) Serialize(io utils.WriteBuffer) error {

	// Simple Field (magicCookie)
	magicCookie := uint16(m.MagicCookie)
	_magicCookieErr := io.WriteUint16(16, (magicCookie))
	if _magicCookieErr != nil {
		return errors.New("Error serializing 'magicCookie' field " + _magicCookieErr.Error())
	}

	// Simple Field (transmitterAddress)
	transmitterAddress := int8(m.TransmitterAddress)
	_transmitterAddressErr := io.WriteInt8(8, (transmitterAddress))
	if _transmitterAddressErr != nil {
		return errors.New("Error serializing 'transmitterAddress' field " + _transmitterAddressErr.Error())
	}

	// Simple Field (receiverAddress)
	receiverAddress := int8(m.ReceiverAddress)
	_receiverAddressErr := io.WriteInt8(8, (receiverAddress))
	if _receiverAddressErr != nil {
		return errors.New("Error serializing 'receiverAddress' field " + _receiverAddressErr.Error())
	}

	// Simple Field (fragmentNumber)
	fragmentNumber := int8(m.FragmentNumber)
	_fragmentNumberErr := io.WriteInt8(8, (fragmentNumber))
	if _fragmentNumberErr != nil {
		return errors.New("Error serializing 'fragmentNumber' field " + _fragmentNumberErr.Error())
	}

	// Simple Field (length)
	length := int8(m.Length)
	_lengthErr := io.WriteInt8(8, (length))
	if _lengthErr != nil {
		return errors.New("Error serializing 'length' field " + _lengthErr.Error())
	}

	// Simple Field (crc)
	crc := uint16(m.Crc)
	_crcErr := io.WriteUint16(16, (crc))
	if _crcErr != nil {
		return errors.New("Error serializing 'crc' field " + _crcErr.Error())
	}

	return nil
}

func (m *AmsSerialResetFrame) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			case "magicCookie":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.MagicCookie = data
			case "transmitterAddress":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.TransmitterAddress = data
			case "receiverAddress":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.ReceiverAddress = data
			case "fragmentNumber":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.FragmentNumber = data
			case "length":
				var data int8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Length = data
			case "crc":
				var data uint16
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.Crc = data
			}
		}
	}
}

func (m *AmsSerialResetFrame) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	className := "org.apache.plc4x.java.ads.readwrite.AmsSerialResetFrame"
	if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
		{Name: xml.Name{Local: "className"}, Value: className},
	}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.MagicCookie, xml.StartElement{Name: xml.Name{Local: "magicCookie"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.TransmitterAddress, xml.StartElement{Name: xml.Name{Local: "transmitterAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.ReceiverAddress, xml.StartElement{Name: xml.Name{Local: "receiverAddress"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.FragmentNumber, xml.StartElement{Name: xml.Name{Local: "fragmentNumber"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Length, xml.StartElement{Name: xml.Name{Local: "length"}}); err != nil {
		return err
	}
	if err := e.EncodeElement(m.Crc, xml.StartElement{Name: xml.Name{Local: "crc"}}); err != nil {
		return err
	}
	if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
		return err
	}
	return nil
}
