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

package org.apache.plc4x.java.simulated.field;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.connection.DefaultPlcFieldHandler;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Arrays;

public class SimulatedFieldHandler extends DefaultPlcFieldHandler {

    @Override
    public PlcField createField(String fieldQuery) {
        if (SimulatedField.matches(fieldQuery)) {
            return SimulatedField.of(fieldQuery);
        }
        throw new PlcInvalidFieldException(fieldQuery);
    }

    @Override
    public PlcValue encodeString(PlcField field, Object[] values) {
        SimulatedField testField = (SimulatedField) field;
        if (testField.getDataType() == String.class) {
            if(values.length == 1) {
                return new PlcSTRING((String) values[0]);
            } else {
                return new PlcList(Arrays.asList(values));
            }
        }
        throw new PlcRuntimeException("Invalid encoder for type " + testField.getDataType().getName());
    }

    @Override
    public PlcValue encodeTime(PlcField field, Object[] values) {
        SimulatedField testField = (SimulatedField) field;
        if (testField.getDataType() == LocalTime.class) {
            if(values.length == 1) {
                return new PlcTime((LocalTime) values[0]);
            } else {
                return new PlcList(Arrays.asList(values));
            }
        }
        throw new PlcRuntimeException("Invalid encoder for type " + testField.getDataType().getName());
    }

    @Override
    public PlcValue encodeDate(PlcField field, Object[] values) {
        SimulatedField testField = (SimulatedField) field;
        if (testField.getDataType() == LocalDate.class) {
            if(values.length == 1) {
                return new PlcDate((LocalDate) values[0]);
            } else {
                return new PlcList(Arrays.asList(values));
            }
        }
        throw new PlcRuntimeException("Invalid encoder for type " + testField.getDataType().getName());
    }

    @Override
    public PlcValue encodeDateTime(PlcField field, Object[] values) {
        SimulatedField testField = (SimulatedField) field;
        if (testField.getDataType() == LocalDateTime.class) {
            if(values.length == 1) {
                return new PlcDateTime((LocalDateTime) values[0]);
            } else {
                return new PlcList(Arrays.asList(values));
            }
        }
        throw new PlcRuntimeException("Invalid encoder for type " + testField.getDataType().getName());
    }

}
