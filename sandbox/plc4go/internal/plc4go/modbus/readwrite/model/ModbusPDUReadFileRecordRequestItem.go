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
    "io"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
    "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/utils"
)

// The data-structure of this message
type ModbusPDUReadFileRecordRequestItem struct {
    ReferenceType uint8
    FileNumber uint16
    RecordNumber uint16
    RecordLength uint16

}

// The corresponding interface
type IModbusPDUReadFileRecordRequestItem interface {
    spi.Message
    Serialize(io utils.WriteBuffer) error
}


func NewModbusPDUReadFileRecordRequestItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordLength uint16) spi.Message {
    return &ModbusPDUReadFileRecordRequestItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordLength: recordLength}
}

func CastIModbusPDUReadFileRecordRequestItem(structType interface{}) IModbusPDUReadFileRecordRequestItem {
    castFunc := func(typ interface{}) IModbusPDUReadFileRecordRequestItem {
        if iModbusPDUReadFileRecordRequestItem, ok := typ.(IModbusPDUReadFileRecordRequestItem); ok {
            return iModbusPDUReadFileRecordRequestItem
        }
        return nil
    }
    return castFunc(structType)
}

func CastModbusPDUReadFileRecordRequestItem(structType interface{}) ModbusPDUReadFileRecordRequestItem {
    castFunc := func(typ interface{}) ModbusPDUReadFileRecordRequestItem {
        if sModbusPDUReadFileRecordRequestItem, ok := typ.(ModbusPDUReadFileRecordRequestItem); ok {
            return sModbusPDUReadFileRecordRequestItem
        }
        if sModbusPDUReadFileRecordRequestItem, ok := typ.(*ModbusPDUReadFileRecordRequestItem); ok {
            return *sModbusPDUReadFileRecordRequestItem
        }
        return ModbusPDUReadFileRecordRequestItem{}
    }
    return castFunc(structType)
}

func (m ModbusPDUReadFileRecordRequestItem) LengthInBits() uint16 {
    var lengthInBits uint16 = 0

    // Simple field (referenceType)
    lengthInBits += 8

    // Simple field (fileNumber)
    lengthInBits += 16

    // Simple field (recordNumber)
    lengthInBits += 16

    // Simple field (recordLength)
    lengthInBits += 16

    return lengthInBits
}

func (m ModbusPDUReadFileRecordRequestItem) LengthInBytes() uint16 {
    return m.LengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestItemParse(io *utils.ReadBuffer) (spi.Message, error) {

    // Simple Field (referenceType)
    referenceType, _referenceTypeErr := io.ReadUint8(8)
    if _referenceTypeErr != nil {
        return nil, errors.New("Error parsing 'referenceType' field " + _referenceTypeErr.Error())
    }

    // Simple Field (fileNumber)
    fileNumber, _fileNumberErr := io.ReadUint16(16)
    if _fileNumberErr != nil {
        return nil, errors.New("Error parsing 'fileNumber' field " + _fileNumberErr.Error())
    }

    // Simple Field (recordNumber)
    recordNumber, _recordNumberErr := io.ReadUint16(16)
    if _recordNumberErr != nil {
        return nil, errors.New("Error parsing 'recordNumber' field " + _recordNumberErr.Error())
    }

    // Simple Field (recordLength)
    recordLength, _recordLengthErr := io.ReadUint16(16)
    if _recordLengthErr != nil {
        return nil, errors.New("Error parsing 'recordLength' field " + _recordLengthErr.Error())
    }

    // Create the instance
    return NewModbusPDUReadFileRecordRequestItem(referenceType, fileNumber, recordNumber, recordLength), nil
}

func (m ModbusPDUReadFileRecordRequestItem) Serialize(io utils.WriteBuffer) error {

    // Simple Field (referenceType)
    referenceType := uint8(m.ReferenceType)
    _referenceTypeErr := io.WriteUint8(8, (referenceType))
    if _referenceTypeErr != nil {
        return errors.New("Error serializing 'referenceType' field " + _referenceTypeErr.Error())
    }

    // Simple Field (fileNumber)
    fileNumber := uint16(m.FileNumber)
    _fileNumberErr := io.WriteUint16(16, (fileNumber))
    if _fileNumberErr != nil {
        return errors.New("Error serializing 'fileNumber' field " + _fileNumberErr.Error())
    }

    // Simple Field (recordNumber)
    recordNumber := uint16(m.RecordNumber)
    _recordNumberErr := io.WriteUint16(16, (recordNumber))
    if _recordNumberErr != nil {
        return errors.New("Error serializing 'recordNumber' field " + _recordNumberErr.Error())
    }

    // Simple Field (recordLength)
    recordLength := uint16(m.RecordLength)
    _recordLengthErr := io.WriteUint16(16, (recordLength))
    if _recordLengthErr != nil {
        return errors.New("Error serializing 'recordLength' field " + _recordLengthErr.Error())
    }

    return nil
}

func (m *ModbusPDUReadFileRecordRequestItem) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
    for {
        token, err := d.Token()
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
            case "referenceType":
                var data uint8
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.ReferenceType = data
            case "fileNumber":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.FileNumber = data
            case "recordNumber":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.RecordNumber = data
            case "recordLength":
                var data uint16
                if err := d.DecodeElement(&data, &tok); err != nil {
                    return err
                }
                m.RecordLength = data
            }
        }
    }
}

func (m ModbusPDUReadFileRecordRequestItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    if err := e.EncodeToken(xml.StartElement{Name: start.Name, Attr: []xml.Attr{
            {Name: xml.Name{Local: "className"}, Value: "org.apache.plc4x.java.modbus.readwrite.ModbusPDUReadFileRecordRequestItem"},
        }}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.ReferenceType, xml.StartElement{Name: xml.Name{Local: "referenceType"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.FileNumber, xml.StartElement{Name: xml.Name{Local: "fileNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.RecordNumber, xml.StartElement{Name: xml.Name{Local: "recordNumber"}}); err != nil {
        return err
    }
    if err := e.EncodeElement(m.RecordLength, xml.StartElement{Name: xml.Name{Local: "recordLength"}}); err != nil {
        return err
    }
    if err := e.EncodeToken(xml.EndElement{Name: start.Name}); err != nil {
        return err
    }
    return nil
}

