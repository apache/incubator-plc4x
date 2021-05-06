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

package s7

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/s7/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"reflect"
	"strings"
	"testing"
)

func TestS7MessageBytes(t *testing.T) {
	type debuggable interface {
		utils.Serializable
	}
	type args struct {
		debuggable debuggable
	}
	tests := []struct {
		name                        string
		args                        args
		wantString                  string
		wantStringSerialized        string
		wantStringSerializedCompact string
		wantStringXml               string
		wantStringJson              string
		wantDump                    string
	}{
		{
			name: "TPKT Packet with Read var response data",
			args: args{
				debuggable: model.NewTPKTPacket(
					model.NewCOTPPacketData(
						false,
						13,
						[]*model.COTPParameter{model.NewCOTPParameterTpduSize(model.COTPTpduSize_SIZE_4096)},
						model.NewS7MessageResponseData(
							0,
							0,
							11,
							model.NewS7ParameterReadVarResponse(1),
							model.NewS7PayloadReadVarResponse(
								[]*model.S7VarPayloadDataItem{
									model.NewS7VarPayloadDataItem(
										model.DataTransportErrorCode_OK,
										model.DataTransportSize_BIT,
										[]byte{1},
									),
								},
							),
						),
					),
				),
			},
			wantString: `
╔═TPKTPacket══════════════════════════════════════════════════════════════════════════════════════════════════════════╗
║╔═ProtocolId╗╔═reserved╗╔═Len═════╗                                                                                  ║
║║  0x03 3   ║║ 0x00 0  ║║0x001d 29║                                                                                  ║
║╚═══════════╝╚═════════╝╚═════════╝                                                                                  ║
║╔═COTPPacket/COTPPacketData/payload═════════════════════════════════════════════════════════════════════════════════╗║
║║╔═HeaderLength╗╔═TpduCode╗╔═Eot════╗╔═TpduRef╗╔═Parameters═════════════════════════════════════════════════╗       ║║
║║║   0x05 5    ║║0xf0 240 ║║b0 false║║0x0d 13 ║║╔═COTPParameter/COTPParameterTpduSize══════════════════════╗║       ║║
║║╚═════════════╝╚═════════╝╚════════╝╚════════╝║║╔═ParameterType╗╔═ParameterLength╗╔═COTPTpduSize/tpduSize╗║║       ║║
║║                                              ║║║   0xc0 192   ║║     0x01 1     ║║    0x0c SIZE_4096    ║║║       ║║
║║                                              ║║╚══════════════╝╚════════════════╝╚══════════════════════╝║║       ║║
║║                                              ║╚══════════════════════════════════════════════════════════╝║       ║║
║║                                              ╚════════════════════════════════════════════════════════════╝       ║║
║║╔═S7Message/S7MessageResponseData/payload═════════════════════════════════════════════════════════════════════════╗║║
║║║╔═ProtocolId╗╔═MessageType╗╔═reserved╗╔═TpduReference╗╔═ParameterLength╗╔═PayloadLength╗╔═ErrorClass╗╔═ErrorCode╗║║║
║║║║  0x32 50  ║║   0x03 3   ║║0x0000 0 ║║  0x000b 11   ║║    0x0002 2    ║║   0x0005 5   ║║  0x00 0   ║║  0x00 0  ║║║║
║║║╚═══════════╝╚════════════╝╚═════════╝╚══════════════╝╚════════════════╝╚══════════════╝╚═══════════╝╚══════════╝║║║
║║║╔═S7Parameter/S7ParameterReadVarResponse/parameter╗                                                              ║║║
║║║║           ╔═ParameterType╗╔═NumItems╗           ║                                                              ║║║
║║║║           ║    0x04 4    ║║ 0x01 1  ║           ║                                                              ║║║
║║║║           ╚══════════════╝╚═════════╝           ║                                                              ║║║
║║║╚═════════════════════════════════════════════════╝                                                              ║║║
║║║╔═S7Payload/S7PayloadReadVarResponse/payload════════════════════════════════════════════════════════════════╗    ║║║
║║║║╔═Items═══════════════════════════════════════════════════════════════════════════════════════════════════╗║    ║║║
║║║║║╔═S7VarPayloadDataItem══════════════════════════════════════════════════════════════════════════════════╗║║    ║║║
║║║║║║╔═DataTransportErrorCode/returnCode╗╔═DataTransportSize/transportSize╗╔═DataLength╗╔═Data═══╗╔═padding╗║║║    ║║║
║║║║║║║             0xff OK              ║║            0x03 BIT            ║║ 0x0001 1  ║║╔══════╗║║ 0x00 0 ║║║║    ║║║
║║║║║║╚══════════════════════════════════╝╚════════════════════════════════╝╚═══════════╝║║0x01 1║║╚════════╝║║║    ║║║
║║║║║║                                                                                   ║╚══════╝║          ║║║    ║║║
║║║║║║                                                                                   ╚════════╝          ║║║    ║║║
║║║║║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════╝║║    ║║║
║║║║╚═════════════════════════════════════════════════════════════════════════════════════════════════════════╝║    ║║║
║║║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════╝    ║║║
║║╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║║
║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║
╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
`,
			wantStringSerialized: `
╔═TPKTPacket═══════════════════════════════════════════════════════════════════════════════════════════════════════╗
║╔═protocolId╗╔═reserved╗╔═len═════╗                                                                               ║
║║  0x03 3   ║║ 0x00 0  ║║0x001d 29║                                                                               ║
║╚═══════════╝╚═════════╝╚═════════╝                                                                               ║
║╔═payload════════════════════════════════════════════════════════════════════════════════════════════════════════╗║
║║╔═COTPPacket═══════════════════════════════════════════════════════════════════════════════════════════════════╗║║
║║║╔═headerLength╗╔═tpduCode╗╔═COTPPacketData═════╗╔═parameters═════════════════════════════════════════════════╗║║║
║║║║   0x05 5    ║║0xf0 240 ║║╔═eot════╗╔═tpduRef╗║║╔═COTPParameter════════════════════════════════════════════╗║║║║
║║║╚═════════════╝╚═════════╝║║b0 false║║ 0xd 13 ║║║║╔═parameterType╗╔═parameterLength╗╔═COTPParameterTpduSize╗║║║║║
║║║                          ║╚════════╝╚════════╝║║║║   0xc0 192   ║║     0x01 1     ║║╔═tpduSize══════════╗ ║║║║║║
║║║                          ╚════════════════════╝║║╚══════════════╝╚════════════════╝║║╔═COTPTpduSize════╗║ ║║║║║║
║║║                                                ║║                                  ║║║0x0c 12 SIZE_4096║║ ║║║║║║
║║║                                                ║║                                  ║║╚═════════════════╝║ ║║║║║║
║║║                                                ║║                                  ║╚═══════════════════╝ ║║║║║║
║║║                                                ║║                                  ╚══════════════════════╝║║║║║
║║║                                                ║╚══════════════════════════════════════════════════════════╝║║║║
║║║                                                ╚════════════════════════════════════════════════════════════╝║║║
║║║╔═S7Message══════════════════════════════════════════════════════════════════════════════╗                    ║║║
║║║║╔═protocolId╗╔═messageType╗╔═reserved╗╔═tpduReference╗╔═parameterLength╗╔═payloadLength╗║                    ║║║
║║║║║  0x32 50  ║║   0x03 3   ║║0x0000 0 ║║  0x000b 11   ║║    0x0002 2    ║║   0x0005 5   ║║                    ║║║
║║║║╚═══════════╝╚════════════╝╚═════════╝╚══════════════╝╚════════════════╝╚══════════════╝║                    ║║║
║║║║╔═S7MessageResponseData═══╗╔═S7Parameter═════════════════════════════════╗              ║                    ║║║
║║║║║╔═errorClass╗╔═errorCode╗║║╔═parameterType╗╔═S7ParameterReadVarResponse╗║              ║                    ║║║
║║║║║║  0x00 0   ║║  0x00 0  ║║║║    0x04 4    ║║        ╔═numItems╗        ║║              ║                    ║║║
║║║║║╚═══════════╝╚══════════╝║║╚══════════════╝║        ║ 0x01 1  ║        ║║              ║                    ║║║
║║║║╚═════════════════════════╝║                ║        ╚═════════╝        ║║              ║                    ║║║
║║║║                           ║                ╚═══════════════════════════╝║              ║                    ║║║
║║║║                           ╚═════════════════════════════════════════════╝              ║                    ║║║
║║║║╔═S7Payload══════════════════════════════════════════════════════════╗                  ║                    ║║║
║║║║║╔═S7PayloadReadVarResponse═════════════════════════════════════════╗║                  ║                    ║║║
║║║║║║╔═items══════════════════════════════════════════════════════════╗║║                  ║                    ║║║
║║║║║║║╔═S7VarPayloadDataItem═════════════════════════════════════════╗║║║                  ║                    ║║║
║║║║║║║║╔═returnCode══════════════╗╔═transportSize══════╗╔═dataLength╗║║║║                  ║                    ║║║
║║║║║║║║║╔═DataTransportErrorCode╗║║╔═DataTransportSize╗║║ 0x0001 1  ║║║║║                  ║                    ║║║
║║║║║║║║║║      0xff 255 OK      ║║║║    0x03 3 BIT    ║║╚═══════════╝║║║║                  ║                    ║║║
║║║║║║║║║╚═══════════════════════╝║║╚══════════════════╝║             ║║║║                  ║                    ║║║
║║║║║║║║╚═════════════════════════╝╚════════════════════╝             ║║║║                  ║                    ║║║
║║║║║║║║╔═data═══════════════════════════════════════╗╔═padding╗      ║║║║                  ║                    ║║║
║║║║║║║║║0|01                            '.         '║║        ║      ║║║║                  ║                    ║║║
║║║║║║║║╚════════════════════════════════════════════╝╚════════╝      ║║║║                  ║                    ║║║
║║║║║║║╚══════════════════════════════════════════════════════════════╝║║║                  ║                    ║║║
║║║║║║╚════════════════════════════════════════════════════════════════╝║║                  ║                    ║║║
║║║║║╚══════════════════════════════════════════════════════════════════╝║                  ║                    ║║║
║║║║╚════════════════════════════════════════════════════════════════════╝                  ║                    ║║║
║║║╚════════════════════════════════════════════════════════════════════════════════════════╝                    ║║║
║║╚══════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║║
║╚════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝║
╚══════════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
`,
			wantStringSerializedCompact: `
╔═TPKTPacket══════════════════════════════════════════════════════════════════════════════════════════════════╗
║╔═protocolId╗╔═reserved╗╔═len═════╗                                                                          ║
║║  0x03 3   ║║ 0x00 0  ║║0x001d 29║                                                                          ║
║╚═══════════╝╚═════════╝╚═════════╝                                                                          ║
║╔═payload/COTPPacket════════════════════════════════════════════════════════════════════════════════════════╗║
║║╔═headerLength╗╔═tpduCode╗╔═COTPPacketData═════╗                                                           ║║
║║║   0x05 5    ║║0xf0 240 ║║╔═eot════╗╔═tpduRef╗║                                                           ║║
║║╚═════════════╝╚═════════╝║║b0 false║║ 0xd 13 ║║                                                           ║║
║║                          ║╚════════╝╚════════╝║                                                           ║║
║║                          ╚════════════════════╝                                                           ║║
║║╔═parameters/COTPParameter════════════════════════════════════════════════════════════╗                    ║║
║║║╔═parameterType╗╔═parameterLength╗╔═COTPParameterTpduSize/tpduSize/COTPTpduSize═════╗║                    ║║
║║║║   0xc0 192   ║║     0x01 1     ║║                0x0c 12 SIZE_4096                ║║                    ║║
║║║╚══════════════╝╚════════════════╝╚═════════════════════════════════════════════════╝║                    ║║
║║╚═════════════════════════════════════════════════════════════════════════════════════╝                    ║║
║║╔═S7Message═══════════════════════════════════════════════════════════════════════════════════════════════╗║║
║║║╔═protocolId╗╔═messageType╗╔═reserved╗╔═tpduReference╗╔═parameterLength╗╔═payloadLength╗                 ║║║
║║║║  0x32 50  ║║   0x03 3   ║║0x0000 0 ║║  0x000b 11   ║║    0x0002 2    ║║   0x0005 5   ║                 ║║║
║║║╚═══════════╝╚════════════╝╚═════════╝╚══════════════╝╚════════════════╝╚══════════════╝                 ║║║
║║║╔═S7MessageResponseData═══╗╔═S7Parameter═════════════════════════════════════════════╗                   ║║║
║║║║╔═errorClass╗╔═errorCode╗║║╔═parameterType╗╔═S7ParameterReadVarResponse/numItems═══╗║                   ║║║
║║║║║  0x00 0   ║║  0x00 0  ║║║║    0x04 4    ║║                0x01 1                 ║║                   ║║║
║║║║╚═══════════╝╚══════════╝║║╚══════════════╝╚═══════════════════════════════════════╝║                   ║║║
║║║╚═════════════════════════╝╚═════════════════════════════════════════════════════════╝                   ║║║
║║║╔═S7Payload/S7PayloadReadVarResponse/items/S7VarPayloadDataItem═════════════════════════════════════════╗║║║
║║║║╔═returnCode/DataTransportErrorCode════════════╗╔═transportSize/DataTransportSize════════╗╔═dataLength╗║║║║
║║║║║                 0xff 255 OK                  ║║               0x03 3 BIT               ║║ 0x0001 1  ║║║║║
║║║║╚══════════════════════════════════════════════╝╚════════════════════════════════════════╝╚═══════════╝║║║║
║║║║╔═data═══════════════════════════════════════╗                                                         ║║║║
║║║║║0|01                            '.         '║                                                         ║║║║
║║║║╚════════════════════════════════════════════╝                                                         ║║║║
║║║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════╝║║║
║║╚═════════════════════════════════════════════════════════════════════════════════════════════════════════╝║║
║╚═══════════════════════════════════════════════════════════════════════════════════════════════════════════╝║
╚═════════════════════════════════════════════════════════════════════════════════════════════════════════════╝
`,
			wantStringXml: `
<TPKTPacket>
  <protocolId dataType="uint" bitLength="8">3</protocolId>
  <reserved dataType="uint" bitLength="8">0</reserved>
  <len dataType="uint" bitLength="16">29</len>
  <payload>
    <COTPPacket>
      <headerLength dataType="uint" bitLength="8">5</headerLength>
      <tpduCode dataType="uint" bitLength="8">240</tpduCode>
      <COTPPacketData>
        <eot dataType="bit" bitLength="1">false</eot>
        <tpduRef dataType="uint" bitLength="7">13</tpduRef>
      </COTPPacketData>
      <parameters isList="true">
        <COTPParameter>
          <parameterType dataType="uint" bitLength="8">192</parameterType>
          <parameterLength dataType="uint" bitLength="8">1</parameterLength>
          <COTPParameterTpduSize>
            <tpduSize>
              <COTPTpduSize dataType="int" bitLength="8" stringRepresentation="SIZE_4096">12</COTPTpduSize>
            </tpduSize>
          </COTPParameterTpduSize>
        </COTPParameter>
      </parameters>
      <S7Message>
        <protocolId dataType="uint" bitLength="8">50</protocolId>
        <messageType dataType="uint" bitLength="8">3</messageType>
        <reserved dataType="uint" bitLength="16">0</reserved>
        <tpduReference dataType="uint" bitLength="16">11</tpduReference>
        <parameterLength dataType="uint" bitLength="16">2</parameterLength>
        <payloadLength dataType="uint" bitLength="16">5</payloadLength>
        <S7MessageResponseData>
          <errorClass dataType="uint" bitLength="8">0</errorClass>
          <errorCode dataType="uint" bitLength="8">0</errorCode>
        </S7MessageResponseData>
        <S7Parameter>
          <parameterType dataType="uint" bitLength="8">4</parameterType>
          <S7ParameterReadVarResponse>
            <numItems dataType="uint" bitLength="8">1</numItems>
          </S7ParameterReadVarResponse>
        </S7Parameter>
        <S7Payload>
          <S7PayloadReadVarResponse>
            <items isList="true">
              <S7VarPayloadDataItem>
                <returnCode>
                  <DataTransportErrorCode dataType="uint" bitLength="8" stringRepresentation="OK">255</DataTransportErrorCode>
                </returnCode>
                <transportSize>
                  <DataTransportSize dataType="uint" bitLength="8" stringRepresentation="BIT">3</DataTransportSize>
                </transportSize>
                <dataLength dataType="uint" bitLength="16">1</dataLength>
                <data dataType="byte" bitLength="8">0x01</data>
                <padding isList="true"></padding>
              </S7VarPayloadDataItem>
            </items>
          </S7PayloadReadVarResponse>
        </S7Payload>
      </S7Message>
    </COTPPacket>
  </payload>
</TPKTPacket>
`,
			wantStringJson: `
{
  "TPKTPacket": {
    "len": 29,
    "len__plc4x_bitLength": 16,
    "len__plc4x_dataType": "uint",
    "payload": {
      "COTPPacket": {
        "COTPPacketData": {
          "eot": false,
          "eot__plc4x_bitLength": 1,
          "eot__plc4x_dataType": "bit",
          "tpduRef": 13,
          "tpduRef__plc4x_bitLength": 7,
          "tpduRef__plc4x_dataType": "uint"
        },
        "S7Message": {
          "S7MessageResponseData": {
            "errorClass": 0,
            "errorClass__plc4x_bitLength": 8,
            "errorClass__plc4x_dataType": "uint",
            "errorCode": 0,
            "errorCode__plc4x_bitLength": 8,
            "errorCode__plc4x_dataType": "uint"
          },
          "S7Parameter": {
            "S7ParameterReadVarResponse": {
              "numItems": 1,
              "numItems__plc4x_bitLength": 8,
              "numItems__plc4x_dataType": "uint"
            },
            "parameterType": 4,
            "parameterType__plc4x_bitLength": 8,
            "parameterType__plc4x_dataType": "uint"
          },
          "S7Payload": {
            "S7PayloadReadVarResponse": {
              "items": [
                {
                  "S7VarPayloadDataItem": {
                    "data": "0x01",
                    "dataLength": 1,
                    "dataLength__plc4x_bitLength": 16,
                    "dataLength__plc4x_dataType": "uint",
                    "data__plc4x_bitLength": 8,
                    "data__plc4x_dataType": "byte",
                    "padding": [],
                    "returnCode": {
                      "DataTransportErrorCode": 255,
                      "DataTransportErrorCode__plc4x_bitLength": 8,
                      "DataTransportErrorCode__plc4x_dataType": "uint",
                      "DataTransportErrorCode__plc4x_stringRepresentation": "OK"
                    },
                    "transportSize": {
                      "DataTransportSize": 3,
                      "DataTransportSize__plc4x_bitLength": 8,
                      "DataTransportSize__plc4x_dataType": "uint",
                      "DataTransportSize__plc4x_stringRepresentation": "BIT"
                    }
                  }
                }
              ]
            }
          },
          "messageType": 3,
          "messageType__plc4x_bitLength": 8,
          "messageType__plc4x_dataType": "uint",
          "parameterLength": 2,
          "parameterLength__plc4x_bitLength": 16,
          "parameterLength__plc4x_dataType": "uint",
          "payloadLength": 5,
          "payloadLength__plc4x_bitLength": 16,
          "payloadLength__plc4x_dataType": "uint",
          "protocolId": 50,
          "protocolId__plc4x_bitLength": 8,
          "protocolId__plc4x_dataType": "uint",
          "reserved": 0,
          "reserved__plc4x_bitLength": 16,
          "reserved__plc4x_dataType": "uint",
          "tpduReference": 11,
          "tpduReference__plc4x_bitLength": 16,
          "tpduReference__plc4x_dataType": "uint"
        },
        "headerLength": 5,
        "headerLength__plc4x_bitLength": 8,
        "headerLength__plc4x_dataType": "uint",
        "parameters": [
          {
            "COTPParameter": {
              "COTPParameterTpduSize": {
                "tpduSize": {
                  "COTPTpduSize": 12,
                  "COTPTpduSize__plc4x_bitLength": 8,
                  "COTPTpduSize__plc4x_dataType": "int",
                  "COTPTpduSize__plc4x_stringRepresentation": "SIZE_4096"
                }
              },
              "parameterLength": 1,
              "parameterLength__plc4x_bitLength": 8,
              "parameterLength__plc4x_dataType": "uint",
              "parameterType": 192,
              "parameterType__plc4x_bitLength": 8,
              "parameterType__plc4x_dataType": "uint"
            }
          }
        ],
        "tpduCode": 240,
        "tpduCode__plc4x_bitLength": 8,
        "tpduCode__plc4x_dataType": "uint"
      }
    },
    "protocolId": 3,
    "protocolId__plc4x_bitLength": 8,
    "protocolId__plc4x_dataType": "uint",
    "reserved": 0,
    "reserved__plc4x_bitLength": 8,
    "reserved__plc4x_dataType": "uint"
  }
}
`,
			wantDump: `
00|03 00 00 1d 05 f0 0d c0 01 0c '..........'
10|32 03 00 00 00 0b 00 02 00 05 '2.........'
20|00 00 04 01 ff 03 00 01 01    '......... '
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// TODO: reimplement String() method with serializer and box based
			/*t.Run("Simple 2 String", func(t *testing.T) {
				tt.wantString = strings.Trim(tt.wantString, "\n")
				if got := tt.args.debuggable.String(); got != tt.wantString {
					t.Errorf("String() = '\n%v\n', want '\n%v\n'", got, tt.wantString)
				}
			})*/
			t.Run("Simple 2 Box", func(t *testing.T) {
				boxWriter := utils.NewBoxedWriteBuffer()
				if err := tt.args.debuggable.Serialize(boxWriter); err != nil {
					t.Error(err)
				}
				tt.wantStringSerialized = strings.Trim(tt.wantStringSerialized, "\n")
				if got := string(boxWriter.GetBox()); got != tt.wantStringSerialized {
					t.Errorf("Serialize Boxed() = '\n%v\n', want '\n%v\n'", got, tt.wantStringSerialized)
				}
			})
			t.Run("Simple 2 Compact Box", func(t *testing.T) {
				boxWriter := utils.NewBoxedWriteBufferWithOptions(true, true)
				if err := tt.args.debuggable.Serialize(boxWriter); err != nil {
					t.Error(err)
				}
				tt.wantStringSerializedCompact = strings.Trim(tt.wantStringSerializedCompact, "\n")
				if got := string(boxWriter.GetBox()); got != tt.wantStringSerializedCompact {
					t.Errorf("Serialize BoxedCompact() = '\n%v\n', want '\n%v\n'", got, tt.wantStringSerializedCompact)
				}
			})
			t.Run("Simple 2 Xml", func(t *testing.T) {
				xmlWriteBuffer := utils.NewXmlWriteBuffer()
				if err := tt.args.debuggable.Serialize(xmlWriteBuffer); err != nil {
					t.Error(err)
				}
				tt.wantStringXml = strings.Trim(tt.wantStringXml, "\n")
				if got := xmlWriteBuffer.GetXmlString(); got != tt.wantStringXml {
					t.Errorf("Serialize Xml() = '\n%v\n', want '\n%v\n'", got, tt.wantStringXml)
				}
			})
			t.Run("Simple 2 Json", func(t *testing.T) {
				jsonWriteBuffer := utils.NewJsonWriteBuffer()
				if err := tt.args.debuggable.Serialize(jsonWriteBuffer); err != nil {
					t.Error(err)
				}
				tt.wantStringJson = strings.Trim(tt.wantStringJson, "\n")
				if got, err := jsonWriteBuffer.GetJsonString(); err != nil || strings.Trim(got, "\n") != tt.wantStringJson {
					if err != nil {
						t.Error(err)
					} else {
						t.Errorf("Serialize Json() = '\n%v\n', want '\n%v\n'", got, tt.wantStringJson)
					}
				}
			})
			t.Run("Simple Binary Serialize", func(t *testing.T) {
				buffer := utils.NewWriteBuffer()
				if err := tt.args.debuggable.Serialize(buffer); err != nil {
					t.Error(err)
				}
				tt.wantDump = strings.Trim(tt.wantDump, "\n")
				if got := utils.Dump(buffer.GetBytes()); !reflect.DeepEqual(got, tt.wantDump) {
					t.Errorf("Serialize() = '\n%v\n', want '\n%v\n'", got, tt.wantDump)
				}
			})
			t.Run("xml roundtip", func(t *testing.T) {
				reader := strings.NewReader(tt.wantStringXml)
				readBuffer := utils.NewXmlReadBuffer(reader)
				if got, err := model.TPKTPacketParse(readBuffer); err != nil || !reflect.DeepEqual(got, tt.args.debuggable) {
					if err != nil {
						t.Error(err)
					} else {
						t.Errorf("Roundtrip(xml) = '\n%v\n', want '\n%v\n'", got, tt.wantDump)
					}
				}
			})
			t.Run("json roundtip", func(t *testing.T) {
				reader := strings.NewReader(tt.wantStringJson)
				readBuffer := utils.NewJsonReadBuffer(reader)
				if got, err := model.TPKTPacketParse(readBuffer); err != nil || !reflect.DeepEqual(got, tt.args.debuggable) {
					if err != nil {
						t.Error(err)
					} else {
						t.Errorf("Roundtrip(json) = '\n%v\n', want '\n%v\n'", got, tt.wantDump)
					}
				}
			})
		})
	}
}
