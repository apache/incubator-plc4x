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
package modbus

import (
	"errors"
	model2 "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/modbus/readwrite/model"
	internalModel "plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/model"
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi"
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/spi/interceptors"
	"plc4x.apache.org/plc4go-modbus-driver/v0/internal/plc4go/transports"
	"plc4x.apache.org/plc4go-modbus-driver/v0/pkg/plc4go"
	"plc4x.apache.org/plc4go-modbus-driver/v0/pkg/plc4go/model"
)

type ConnectionMetadata struct {
	model.PlcConnectionMetadata
}

func (m ConnectionMetadata) CanRead() bool {
	return true
}

func (m ConnectionMetadata) CanWrite() bool {
	return true
}

func (m ConnectionMetadata) CanSubscribe() bool {
	return false
}

type ModbusConnection struct {
	unitIdentifier     uint8
	messageCodec       spi.MessageCodec
	options            map[string][]string
	fieldHandler       spi.PlcFieldHandler
	valueHandler       spi.PlcValueHandler
	requestInterceptor internalModel.RequestInterceptor
	plc4go.PlcConnection
}

func NewModbusConnection(unitIdentifier uint8, messageCodec spi.MessageCodec, options map[string][]string, fieldHandler spi.PlcFieldHandler) ModbusConnection {
	return ModbusConnection{
		unitIdentifier:     unitIdentifier,
		messageCodec:       messageCodec,
		options:            options,
		fieldHandler:       fieldHandler,
		valueHandler:       NewValueHandler(),
		requestInterceptor: interceptors.NewSingleItemRequestInterceptor(),
	}
}

func (m ModbusConnection) Connect() <-chan plc4go.PlcConnectionConnectResult {
	ch := make(chan plc4go.PlcConnectionConnectResult)
	go func() {
		err := m.messageCodec.Connect()
		ch <- plc4go.NewPlcConnectionConnectResult(m, err)
	}()
	return ch
}

func (m ModbusConnection) Close() <-chan plc4go.PlcConnectionCloseResult {
	// TODO: Implement ...
	ch := make(chan plc4go.PlcConnectionCloseResult)
	go func() {
		ch <- plc4go.NewPlcConnectionCloseResult(m, nil)
	}()
	return ch
}

func (m ModbusConnection) IsConnected() bool {
	panic("implement me")
}

func (m ModbusConnection) Ping() <-chan plc4go.PlcConnectionPingResult {
	result := make(chan plc4go.PlcConnectionPingResult)
	diagnosticRequestPdu := model2.CastIModbusPDU(model2.NewModbusPDUDiagnosticRequest(0, 0x42))
	go func() {
		pingRequest := model2.NewModbusTcpADU(1, m.unitIdentifier, diagnosticRequestPdu)
		err := m.messageCodec.Send(pingRequest)
		if err != nil {
			result <- plc4go.NewPlcConnectionPingResult(err)
			return
		}
		// Register an expected response
		check := func(response interface{}) bool {
			responseAdu := model2.CastModbusTcpADU(response)
			return responseAdu.TransactionIdentifier == 1 &&
				responseAdu.UnitIdentifier == m.unitIdentifier
		}
		// Register a callback to handle the response
		pingResponseChanel := m.messageCodec.Expect(check)
		if pingResponseChanel == nil {
			result <- plc4go.NewPlcConnectionPingResult(errors.New("no response channel"))
			return
		}
		// Wait for the response to come in
		pingResponse := <-pingResponseChanel
		if pingResponse == nil {
			result <- plc4go.NewPlcConnectionPingResult(errors.New("no response"))
			return
		}
		// If we got a valid response (even if it will probably contain an error, we know the remote is available)
		result <- plc4go.NewPlcConnectionPingResult(nil)
	}()

	return result
}

func (m ModbusConnection) GetMetadata() model.PlcConnectionMetadata {
	return ConnectionMetadata{}
}

func (m ModbusConnection) ReadRequestBuilder() model.PlcReadRequestBuilder {
	return internalModel.NewDefaultPlcReadRequestBuilderWithInterceptor(m.fieldHandler,
		NewModbusReader(m.unitIdentifier, m.messageCodec), m.requestInterceptor)
}

func (m ModbusConnection) WriteRequestBuilder() model.PlcWriteRequestBuilder {
	return internalModel.NewDefaultPlcWriteRequestBuilder(
		m.fieldHandler, m.valueHandler, NewModbusWriter(m.unitIdentifier, m.messageCodec))
}

func (m ModbusConnection) SubscriptionRequestBuilder() model.PlcSubscriptionRequestBuilder {
	panic("implement me")
}

func (m ModbusConnection) UnsubscriptionRequestBuilder() model.PlcUnsubscriptionRequestBuilder {
	panic("implement me")
}

func (m ModbusConnection) GetTransportInstance() transports.TransportInstance {
	if mc, ok := m.messageCodec.(spi.TransportInstanceExposer); ok {
		return mc.GetTransportInstance()
	}
	return nil
}

func (m ModbusConnection) GetPlcFieldHandler() spi.PlcFieldHandler {
	return m.fieldHandler
}

func (m ModbusConnection) GetPlcValueHandler() spi.PlcValueHandler {
	return m.valueHandler
}
