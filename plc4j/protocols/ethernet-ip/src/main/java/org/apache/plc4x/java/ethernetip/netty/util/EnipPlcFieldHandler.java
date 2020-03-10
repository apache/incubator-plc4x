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
package org.apache.plc4x.java.ethernetip.netty.util;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.connection.DefaultPlcFieldHandler;
import org.apache.plc4x.java.ethernetip.model.EtherNetIpField;

public class EnipPlcFieldHandler extends DefaultPlcFieldHandler {

    @Override
    public PlcField createField(String fieldQuery) throws PlcInvalidFieldException {
        if (EtherNetIpField.matches(fieldQuery)) {
            return EtherNetIpField.of(fieldQuery);
        }
        throw new PlcInvalidFieldException(fieldQuery);
    }

    @Override
    public PlcValue encodeBoolean(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeByte(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeShort(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeInteger(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeBigInteger(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeLong(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeFloat(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeDouble(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeString(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeTime(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeDate(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

    @Override
    public PlcValue encodeDateTime(PlcField field, Object[] values) {
        EtherNetIpField enipField = (EtherNetIpField) field;
        throw new PlcRuntimeException("Invalid encoder for type " + enipField);
    }

}
