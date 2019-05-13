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
package org.apache.plc4x.java.opcua.protocol.model;



import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.base.connection.DefaultPlcFieldHandler;
import org.apache.plc4x.java.base.messages.items.*;
import org.apache.plc4x.java.opcua.protocol.OpcuaField;

import java.math.BigInteger;
import java.util.ArrayList;
import java.util.List;

public class OpcuaPlcFieldHandler extends DefaultPlcFieldHandler {

    @Override
    public PlcField createField(String fieldQuery) throws PlcInvalidFieldException {
        if (OpcuaField.matches(fieldQuery)) {
            return OpcuaField.of(fieldQuery);
        }
        throw new PlcInvalidFieldException(fieldQuery);
    }

    @Override
    public BaseDefaultFieldItem encodeString(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<String> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add(item.toString());
        }
        return new DefaultStringFieldItem(resultSet.toArray(new String[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeBoolean(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Boolean> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Boolean) item);
        }
        return new DefaultBooleanFieldItem(resultSet.toArray(new Boolean[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeByte(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Byte> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Byte) item);
        }
        return new DefaultByteFieldItem(resultSet.toArray(new Byte[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeShort(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Short> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Short) item);
        }
        return new DefaultShortFieldItem(resultSet.toArray(new Short[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeInteger(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Integer> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Integer) item);
        }
        return new DefaultIntegerFieldItem(resultSet.toArray(new Integer[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeBigInteger(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<BigInteger> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((BigInteger) item);
        }
        return new DefaultBigIntegerFieldItem(resultSet.toArray(new BigInteger[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeLong(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Long> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Long) item);
        }
        return new DefaultLongFieldItem(resultSet.toArray(new Long[0]));
    }

    @Override
    public BaseDefaultFieldItem encodeFloat(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Float> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Float) item);
        }
        return new DefaultFloatFieldItem(resultSet.toArray(new Float[0]));
    }



    @Override
    public BaseDefaultFieldItem encodeDouble(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Double> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Double) item);
        }
        return new DefaultDoubleFieldItem(resultSet.toArray(new Double[0]));
    }


    @Override
    public BaseDefaultFieldItem encodeByteArray(PlcField field, Object[] values) {
        OpcuaField adsField = (OpcuaField) field;
        ArrayList<Byte[]> resultSet = new ArrayList<>();
        for(Object item : values){
            resultSet.add((Byte[]) item);
        }
        return new DefaultByteArrayFieldItem(resultSet.toArray(new Byte[0][0]));
    }
}
