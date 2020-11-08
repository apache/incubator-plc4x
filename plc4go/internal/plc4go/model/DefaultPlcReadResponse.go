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
	"plc4x.apache.org/plc4go/v0/pkg/plc4go/model"
	"plc4x.apache.org/plc4go/v0/pkg/plc4go/values"
)

type DefaultPlcReadResponse struct {
	request       model.PlcReadRequest
	responseCodes map[string]model.PlcResponseCode
	values        map[string]values.PlcValue
	model.PlcReadResponse
}

func NewDefaultPlcReadResponse(request model.PlcReadRequest, responseCodes map[string]model.PlcResponseCode, values map[string]values.PlcValue) DefaultPlcReadResponse {
	return DefaultPlcReadResponse{
		request:       request,
		responseCodes: responseCodes,
		values:        values,
	}
}

func (m DefaultPlcReadResponse) GetFieldNames() []string {
	var fieldNames []string
	for fieldName, _ := range m.values {
		fieldNames = append(fieldNames, fieldName)
	}
	return fieldNames
}

func (m DefaultPlcReadResponse) GetRequest() model.PlcReadRequest {
	return m.request
}

func (m DefaultPlcReadResponse) GetResponseCode(name string) model.PlcResponseCode {
	return m.responseCodes[name]
}

func (m DefaultPlcReadResponse) GetValue(name string) values.PlcValue {
	return m.values[name]
}

func (m DefaultPlcReadResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "PlcReadResponse"}}); err != nil {
		return err
	}

	if err := e.EncodeElement(m.request, xml.StartElement{Name: xml.Name{Local: "PlcReadRequest"}}); err != nil {
		return err
	}

	if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: "values"}}); err != nil {
		return err
	}
	for _, fieldName := range m.GetFieldNames() {
		if err := e.EncodeToken(xml.StartElement{Name: xml.Name{Local: fieldName},
			Attr: []xml.Attr{
				{Name: xml.Name{Local: "result"}, Value: m.GetResponseCode(fieldName).GetName()},
			}}); err != nil {
			return err
		}
		if err := e.EncodeElement(m.GetValue(fieldName), xml.StartElement{Name: xml.Name{Local: "field"}}); err != nil {
			return err
		}
		if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: fieldName}}); err != nil {
			return err
		}
	}
	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "values"}}); err != nil {
		return err
	}

	if err := e.EncodeToken(xml.EndElement{Name: xml.Name{Local: "PlcReadResponse"}}); err != nil {
		return err
	}
	return nil
}
