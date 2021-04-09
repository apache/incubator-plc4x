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
	"github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/transports/udp"
	"github.com/apache/plc4x/plc4go/pkg/plc4go"
	apiModel "github.com/apache/plc4x/plc4go/pkg/plc4go/model"
	"testing"
	"time"
)

func TestKnxNetIpPlc4goBrowse(t *testing.T) {
	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(knxnetip.NewKnxNetIpDriver())
	driverManager.RegisterTransport(udp.NewUdpTransport())

	// Create a connection string from the discovery result.
	connectionString := "knxnet-ip:udp://192.168.42.11:3671"
	crc := driverManager.GetConnection(connectionString)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		t.Errorf("Got error connecting: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}
	connection := connectionResult.Connection
	defer connection.Close()

	// Build a browse request, to scan the KNX network for KNX devices
	// (Limiting the range to only the actually used range of addresses)
	browseRequestBuilder := connection.BrowseRequestBuilder()
	browseRequestBuilder.AddItem("findAllKnxDevices", "[1-3].[1-6].[1-60]")
	browseRequest, err := browseRequestBuilder.Build()
	if err != nil {
		t.Errorf("Got error preparing browse-request: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}

	// Execute the browse-request
	brr := browseRequest.Execute()
	browseRequestResults := <-brr
	if browseRequestResults.Err != nil {
		t.Errorf("Got error executing browse-request: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}

	// Output the addresses found
	for _, queryName := range browseRequestResults.Response.GetQueryNames() {
		results := browseRequestResults.Response.GetQueryResults(queryName)
		for _, result := range results {
			fmt.Printf("Found KNX device at address: %v querying device information: \n", result.Address)

			// Create a read-request to read the manufacturer and hardware ids
			readRequestBuilder := connection.ReadRequestBuilder()
			readRequestBuilder.AddItem("manufacturerId", result.Address+"/0/12")
			readRequestBuilder.AddItem("programVersion", result.Address+"/0/13")
			readRequest, err := readRequestBuilder.Build()
			if err != nil {
				t.Errorf("Error creating read-request. %s", err.Error())
				t.Fail()
				return
			}

			// Execute the read-requests
			rrr := readRequest.Execute()
			readResult := <-rrr
			if readResult.Err != nil {
				t.Errorf("Error executing read-request. %s", readResult.Err.Error())
				t.Fail()
				return
			}

			// Check the response
			readResponse := readResult.Response
			if readResponse.GetResponseCode("manufacturerId") != apiModel.PlcResponseCode_OK {
				t.Errorf("Got response code %d for 'manufacturerId'", readResponse.GetResponseCode("manufacturerId"))
				t.Fail()
				return
			}
			if readResponse.GetResponseCode("programVersion") != apiModel.PlcResponseCode_OK {
				t.Errorf("Got response code %d for 'hardwareType'", readResponse.GetResponseCode("programVersion"))
				t.Fail()
				return
			}

			manufacturerId := readResponse.GetValue("manufacturerId").GetUint16()
			programVersion := readResponse.GetValue("programVersion").GetRaw()
			fmt.Printf(" - Manufacturer Id: %d, Program Version: %s\n", manufacturerId, hex.EncodeToString(programVersion))
		}
	}
}

func TestKnxNetIpPlc4goDiscovery(t *testing.T) {
	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(knxnetip.NewKnxNetIpDriver())
	driverManager.RegisterTransport(udp.NewUdpTransport())

	found := make(chan bool)
	err := driverManager.Discover(func(event apiModel.PlcDiscoveryEvent) {
		fmt.Printf("Found device: %s:%s://%s\n - Name: %s\n", event.ProtocolCode, event.TransportCode,
			event.TransportUrl.Host, event.Name)
		//go func() {
		// Create a connection string from the discovery result.
		connectionString := fmt.Sprintf("%s:%s://%s", event.ProtocolCode, event.TransportCode,
			event.TransportUrl.Host)
		crc := driverManager.GetConnection(connectionString)
		connectionResult := <-crc
		if connectionResult.Err != nil {
			fmt.Printf("Got error connecting: %s", connectionResult.Err.Error())
			return
		}
		connection := connectionResult.Connection
		defer connection.Close()

		// Build a browse request, to scan the KNX network for KNX devices
		// (Limiting the range to only the actually used range of addresses)
		browseRequestBuilder := connection.BrowseRequestBuilder()
		browseRequestBuilder.AddItem("findAllKnxDevices", "[1-3].[1-6].[1-60]")
		browseRequest, err := browseRequestBuilder.Build()
		if err != nil {
			fmt.Printf("Got error preparing browse-request: %s", connectionResult.Err.Error())
			return
		}

		// Execute the browse-request
		brr := browseRequest.Execute()
		browseRequestResults := <-brr
		if browseRequestResults.Err != nil {
			fmt.Printf("Got error executing browse-request: %s", connectionResult.Err.Error())
			return
		}

		// Output the addresses found
		for _, queryName := range browseRequestResults.Response.GetQueryNames() {
			results := browseRequestResults.Response.GetQueryResults(queryName)
			for _, result := range results {
				fmt.Printf("Found KNX device at address: %v", result.Address)
			}
		}
		//}()
		found <- true
	})
	if err != nil {
		fmt.Printf("got error %s", err.Error())
		return
	}
	for {
		select {
		case _ = <-found:
			time.Sleep(time.Second * 2)
			fmt.Print("Found devices")
			return
		case <-time.After(time.Second * 60):
			t.Error("Couldn't find device in the last 60 seconds")
			t.Fail()
			return
		}
	}
}

func TestKnxNetIpPlc4goGroupAddressRead(t *testing.T) {
	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(knxnetip.NewKnxNetIpDriver())
	driverManager.RegisterTransport(udp.NewUdpTransport())

	// Get a connection to a remote PLC
	crc := driverManager.GetConnection("knxnet-ip://192.168.42.11")

	// Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		t.Errorf("error connecting to PLC: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}
	connection := connectionResult.Connection
	defer connection.Close()

	attributes := connection.GetMetadata().GetConnectionAttributes()
	fmt.Printf("Successfully connected to KNXnet/IP Gateway '%s' with KNX address '%s' got assigned client KNX address '%s'\n",
		attributes["GatewayName"],
		attributes["GatewayKnxAddress"],
		attributes["ClientKnxAddress"])

	// TODO: Find out why a connection-state request breaks everything ...
	// Try to ping the remote device
	pingResultChannel := connection.Ping()
	pingResult := <-pingResultChannel
	fmt.Println("Ping Received")
	if pingResult.Err != nil {
		t.Errorf("couldn't ping device: %s", pingResult.Err.Error())
		t.Fail()
		return
	}

	srb := connection.SubscriptionRequestBuilder()
	srb.AddChangeOfStateItem("heating-actual-temperature", "*/*/10:DPT_Value_Temp")
	srb.AddChangeOfStateItem("heating-target-temperature", "*/*/11:DPT_Value_Temp")
	srb.AddChangeOfStateItem("heating-valve-open", "*/*/12:DPT_OpenClose")
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

	// Wait 2 minutes
	time.Sleep(120 * time.Second)

	// Execute a read request
	rrb := connection.ReadRequestBuilder()
	rrb.AddItem("energy-consumption", "1/1/211:DPT_Value_Power")
	rrb.AddItem("actual-temperatures", "*/*/10:DPT_Value_Temp")
	rrb.AddItem("set-temperatures", "*/*/11:DPT_Value_Temp")
	rrb.AddItem("window-status", "*/*/[60,64]:DPT_Value_Temp")
	rrb.AddItem("power-consumption", "*/*/[111,121,131,141]:DPT_Value_Temp")
	readRequest, err := rrb.Build()
	if err == nil {
		rrr := readRequest.Execute()
		readRequestResult := <-rrr
		if readRequestResult.Err == nil {
			for _, fieldName := range readRequestResult.Response.GetFieldNames() {
				if readRequestResult.Response.GetResponseCode(fieldName) == apiModel.PlcResponseCode_OK {
					fmt.Printf(" - Field %s Value %s\n", fieldName, readRequestResult.Response.GetValue(fieldName).GetString())
				}
			}
		}
	}
}

func TestKnxNetIpPlc4goPropertyRead(t *testing.T) {
	driverManager := plc4go.NewPlcDriverManager()
	driverManager.RegisterDriver(knxnetip.NewKnxNetIpDriver())
	driverManager.RegisterTransport(udp.NewUdpTransport())

	// Get a connection to a remote PLC
	crc := driverManager.GetConnection("knxnet-ip://192.168.42.11")

	// Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.Err != nil {
		t.Errorf("error connecting to PLC: %s", connectionResult.Err.Error())
		t.Fail()
		return
	}
	connection := connectionResult.Connection
	defer connection.Close()

	readRequestBuilder := connection.ReadRequestBuilder()
	readRequestBuilder.AddItem("manufacturerId", "1.1.10/0/12")
	readRequestBuilder.AddItem("programVersion", "1.1.10/3/13")
	//readRequestBuilder.AddItem("hardwareType", "1.1.10/0/78")
	readRequest, _ := readRequestBuilder.Build()

	rrr := readRequest.Execute()
	readResult := <-rrr

	fmt.Printf("Got result %v", readResult)
}

func knxEventHandler(event apiModel.PlcSubscriptionEvent) {
	for _, fieldName := range event.GetFieldNames() {
		if event.GetResponseCode(fieldName) == apiModel.PlcResponseCode_OK {
			groupAddress := event.GetAddress(fieldName)
			fmt.Printf("Got update for field %s with address %s. Value changed to: %s\n",
				fieldName, groupAddress, event.GetValue(fieldName).GetString())
		}
	}
}
