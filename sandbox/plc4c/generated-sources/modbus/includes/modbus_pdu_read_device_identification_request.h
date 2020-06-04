/*
  Licensed to the Apache Software Foundation (ASF) under one
  or more contributor license agreements.  See the NOTICE file
  distributed with this work for additional information
  regarding copyright ownership.  The ASF licenses this file
  to you under the Apache License, Version 2.0 (the
  "License"); you may not use this file except in compliance
  with the License.  You may obtain a copy of the License at

      http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing,
  software distributed under the License is distributed on an
  "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
  KIND, either express or implied.  See the License for the
  specific language governing permissions and limitations
  under the License.
*/
#ifndef PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_READ_DEVICE_IDENTIFICATION_REQUEST_H_
#define PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_READ_DEVICE_IDENTIFICATION_REQUEST_H_
#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <stdint.h>
#include <plc4c/utils/list.h>
#include "modbus_pdu.h"

struct plc4c_modbus_read_write_modbus_pdu_read_device_identification_request {
  plc4c_modbus_read_write_modbus_pdu_type _type;
};
typedef struct plc4c_modbus_read_write_modbus_pdu_read_device_identification_request plc4c_modbus_read_write_modbus_pdu_read_device_identification_request;

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_read_device_identification_request_parse(plc4c_spi_read_buffer* buf, bool response, plc4c_modbus_read_write_modbus_pdu_read_device_identification_request** message);

plc4c_return_code plc4c_modbus_read_write_modbus_pdu_read_device_identification_request_serialize(plc4c_spi_write_buffer* buf, plc4c_modbus_read_write_modbus_pdu_read_device_identification_request* message);

#ifdef __cplusplus
}
#endif
#endif  // PLC4C_MODBUS_READ_WRITE_MODBUS_PDU_READ_DEVICE_IDENTIFICATION_REQUEST_H_