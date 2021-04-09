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
    "fmt"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/knxnetip"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/knxnetip/readwrite/model"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/transports/udp"
    "plc4x.apache.org/plc4go/v0/internal/plc4go/utils"
    "plc4x.apache.org/plc4go/v0/pkg/plc4go"
    apiModel "plc4x.apache.org/plc4go/v0/pkg/plc4go/model"
    "testing"
    "time"
)

func KnxNetIp(t *testing.T) {
	t.Skip()
	request, err := hex.DecodeString("000a00000006010300000004")
	if err != nil {
		t.Errorf("Error decoding test input")
	}
	rb := utils.NewReadBuffer(request)
	adu, err := model.KnxNetIpMessageParse(rb)
	if err != nil {
		t.Errorf("Error parsing: %s", err)
	}
	if adu != nil {
		// Output success ...
	}

}

func TestKnxNetIpPlc4goDriver(t *testing.T) {
    driverManager := plc4go.NewPlcDriverManager()
    driverManager.RegisterDriver(knxnetip.NewKnxNetIpDriver())
    driverManager.RegisterTransport(udp.NewUdpTransport())

    // Get a connection to a remote PLC
    crc := driverManager.GetConnection("knxnet-ip://192.168.42.11")
    //crc := driverManager.GetConnection("knxnet-ip://-discover-")

    // Wait for the driver to connect (or not)
    connectionResult := <-crc
    if connectionResult.Err != nil {
        t.Errorf("error connecting to PLC: %s", connectionResult.Err.Error())
        t.Fail()
        return
    }
    connection := connectionResult.Connection

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
    pollingInterval, err := time.ParseDuration("5s")
    if err != nil {
        t.Errorf("invalid format")
        t.Fail()
        return
    }
    srb := connection.SubscriptionRequestBuilder()
    srb.AddChangeOfStateItem("field1", "*/*/*")
    srb.AddCyclicItem("field2", "holding-register:3:REAL", pollingInterval)
    srb.AddItemHandler(knxEventHandler)
    subscriptionRequest, err := srb.Build()
    if err != nil {
        t.Errorf("error preparing subscription-request: %s", connectionResult.Err.Error())
        t.Fail()
        return
    }

    // Execute a subscription-request
    rrc := subscriptionRequest.Execute()

    // Wait for the response to finish
    rrr := <-rrc
    if rrr.Err != nil {
        t.Errorf("error executing read-request: %s", rrr.Err.Error())
        t.Fail()
        return
    }

    // Do something with the response
    /*value1 := rrr.Response.GetValue("field1")
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
    fmt.Printf("\n\nResult field2: %d\n", wrr.Response.GetResponseCode("field2"))*/
}

func knxEventHandler(event apiModel.PlcSubscriptionEvent) {
    for _, fieldName := range event.GetFieldNames() {
        if event.GetResponseCode(fieldName) == apiModel.PlcResponseCode_OK {
            fmt.Printf("Got update for field %s with value %s", fieldName, event.GetValue(fieldName).GetString())
        }
    }
}