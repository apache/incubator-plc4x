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
package drivers

import (
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"net"
	"os"
	"github.com/apache/plc4x/plc4go/internal/plc4go/modbus"
	"github.com/apache/plc4x/plc4go/internal/plc4go/modbus/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/transports/tcp"
	"github.com/apache/plc4x/plc4go/internal/plc4go/utils"
	"github.com/apache/plc4x/plc4go/pkg/plc4go"
	"strings"
	"testing"
)

func TestModbus(t *testing.T) {
	test(t, "000000000006ff0408d20002", false)
	test(t, "7cfe000000c9ff04c600000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000100000004000000000000000000000000000001db000001d600004a380000000000000000000000000000000000000000000000000000000000006461696d006e0000000000000000000000000000303100300000000000000000000000000000000000000000000000000000000000000000000000000000", true)
	test(t, "000a0000001101140e060003270e000206000400000008", false)
	test(t, "000a0000001b011418050600000000110600000000000000000000000000000000", true)
	test(t, "000a0000000c011509060002000000010008", false)
	test(t, "000a00000015011512060001270F00010000060002000000010000", false)
}

func test(t *testing.T, rawMessage string, response bool) {
	// Cr9edz47r4eate the input data
	// "000a 0000 0006 01 03 00 00 00 04"
	request, err := hex.DecodeString(rawMessage)
	if err != nil {
		t.Errorf("Error decoding test input")
	}
	rb := utils.NewReadBuffer(request)
	adu, err := model.ModbusTcpADUParse(rb, response)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}
	if adu != nil {
		serialized, err := xml.Marshal(adu)
		if err != nil {
			fmt.Println("Hurz!" + err.Error())
			return
		}
		fmt.Println(string(serialized))
		var deserializedAdu *model.ModbusTcpADU
		xml.Unmarshal(serialized, &deserializedAdu)

		wb := utils.NewWriteBuffer()
		val := model.CastModbusTcpADU(deserializedAdu)
		val.Serialize(*wb)
		serializedMessage := hex.EncodeToString(wb.GetBytes())
		if strings.ToUpper(serializedMessage) != strings.ToUpper(rawMessage) {
			t.Errorf("The serilized result doesn't match the input")
		}
	}
}

//
// Test that actually sends a read-request to a remote Modbus Slave
//
func Connection(t *testing.T) {

	pdu := model.NewModbusPDUReadInputRegistersRequest(1, 1)
    adu := model.NewModbusTcpADU( 0,255, pdu)

	wb := utils.NewWriteBuffer()
	adu.Serialize(*wb)

	servAddr := "192.168.23.30:502?unit-identifier=1"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write(wb.GetBytes())
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	buffer := make([]byte, 1024)

	numBytes, err := conn.Read(buffer)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	rb := utils.NewReadBuffer(buffer[0:numBytes])
	response, err := model.ModbusTcpADUParse(rb, true)
	if err != nil {
		println("Parsing response failed:", err.Error())
		os.Exit(1)
	}

	fmt.Println(response)

	conn.Close()
}

func TestPlc4goDriver(t *testing.T) {
	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(modbus.NewModbusDriver())
	driverManager.RegisterTransport(tcp.NewTcpTransport())

	// Get a connection to a remote PLC
	crc := driverManager.GetConnection("modbus://192.168.23.30?unit-identifier=1")

	// Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		t.Errorf("error connecting to PLC: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}
	connection := connectionResult.Connection

	if !connection.GetMetadata().CanRead() {
	    fmt.Printf("This connection doesn't support read operations")
        return
    }

	// Try to ping the remote device
	pingResultChannel := connection.Ping()
	pingResult := <-pingResultChannel
	if pingResult.Err != nil {
		t.Errorf("couldn't ping device: %s", pingResult.Err.Error())
		t.Fail()
		return
	}

	// Make sure the connection is closed at the end
	defer connection.Close()

	// Prepare a read-request
	rrb := connection.ReadRequestBuilder()
	rrb.AddItem("field1", "holding-register:1:REAL")
	rrb.AddItem("field2", "holding-register:3:REAL")
	readRequest, err := rrb.Build()
	if err != nil {
		t.Errorf("error preparing read-request: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}

	// Execute a read-request
	rrc := readRequest.Execute()

	// Wait for the response to finish
	rrr := <-rrc
	if rrr.Err != nil {
		t.Errorf("error executing read-request: %s", rrr.Err.Error())
		t.Fail()
		return
	}

	// Do something with the response
	value1 := rrr.Response.GetValue("field1")
	value2 := rrr.Response.GetValue("field2")
	fmt.Printf("\n\nResult field1: %f\n", value1.GetFloat32())
	fmt.Printf("\n\nResult field1: %f\n", value2.GetFloat32())

	// Prepare a write-request
	wrb := connection.WriteRequestBuilder()
	wrb.AddItem("field1", "holding-register:1:REAL", 1.2345)
	wrb.AddItem("field2", "holding-register:3:REAL", 2.3456)
	writeRequest, err := rrb.Build()
	if err != nil {
		t.Errorf("error preparing read-request: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}

	// Execute a write-request
	wrc := writeRequest.Execute()

	// Wait for the response to finish
	wrr := <-wrc
	if wrr.Err != nil {
		t.Errorf("error executing read-request: %s", rrr.Err.Error())
		t.Fail()
		return
	}

	fmt.Printf("\n\nResult field1: %d\n", wrr.Response.GetResponseCode("field1"))
	fmt.Printf("\n\nResult field2: %d\n", wrr.Response.GetResponseCode("field2"))
}
