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
package org.apache.plc4x.java.s7.readwrite.field;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.connection.DefaultPlcFieldHandler;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.InvocationTargetException;
import java.math.BigInteger;
import java.nio.charset.StandardCharsets;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Arrays;
import java.util.BitSet;
import java.util.LinkedList;
import java.util.List;

public class S7PlcFieldHandler extends DefaultPlcFieldHandler {
    private static final Logger LOGGER = LoggerFactory.getLogger(S7PlcFieldHandler.class);
    @Override
    public PlcField createField(String fieldQuery) {
        if (S7Field.matches(fieldQuery)) {
            return S7Field.of(fieldQuery);
        }
        throw new PlcInvalidFieldException(fieldQuery);
    }

    @Override
    public PlcValue encodeBoolean(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        // All of these types are declared as Bit or Bit-String types.
        switch (s7Field.getDataType()) {
            case BOOL:
            case BYTE:
            case WORD:
            case DWORD:
            case LWORD:
                return internalEncodeBoolean(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeByte(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        // All of these types are declared as Bit or Bit-String types.
        switch (s7Field.getDataType()) {
            case BYTE:
            case SINT:
            case USINT:
            case CHAR:
                return internalEncodeInteger(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeShort(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case WORD:
            case INT:
            case UINT:
                return internalEncodeInteger(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeInteger(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case DWORD:
            case DINT:
            case UDINT:
                return internalEncodeInteger(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeBigInteger(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case DWORD:
            case DINT:
            case UDINT:
                return internalEncodeInteger(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeLong(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case LWORD:
            case LINT:
            case ULINT:
                return internalEncodeInteger(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeFloat(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case REAL:
                return internalEncodeFloatingPoint(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeDouble(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case LREAL:
                return internalEncodeFloatingPoint(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeString(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case CHAR:
            case WCHAR:
            case STRING:
            case WSTRING:
                return internalEncodeString(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeTime(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case TIME:
                return internalEncodeTemporal(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeDate(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case DATE:
                return internalEncodeTemporal(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    @Override
    public PlcValue encodeDateTime(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case DATE_AND_TIME:
                return internalEncodeTemporal(field, values);
            default:
                throw new PlcRuntimeException("Invalid encoder for type " + s7Field.getDataType().name());
        }
    }

    private PlcValue internalEncodeBoolean(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case BOOL:
            case BYTE:
            case WORD:
            case DWORD:
            case LWORD:
                break;
            default:
                throw new IllegalArgumentException(
                    "Cannot assign boolean values to " + s7Field.getDataType().name() + " fields.");
        }
        List<Boolean> booleanValues = new LinkedList<>();
        for (Object value : values) {
            if (value instanceof Boolean) {
                Boolean booleanValue = (Boolean) value;
                booleanValues.add(booleanValue);
            } else if (value instanceof Byte) {
                Byte byteValue = (Byte) value;
                BitSet bitSet = BitSet.valueOf(new byte[]{byteValue});
                for (int i = 0; i < 8; i++) {
                    booleanValues.add(bitSet.get(i));
                }
            } else if (value instanceof Short) {
                Short shortValue = (Short) value;
                BitSet bitSet = BitSet.valueOf(new long[]{shortValue});
                for (int i = 0; i < 16; i++) {
                    booleanValues.add(bitSet.get(i));
                }
            } else if (value instanceof Integer) {
                Integer integerValue = (Integer) value;
                BitSet bitSet = BitSet.valueOf(new long[]{integerValue});
                for (int i = 0; i < 32; i++) {
                    booleanValues.add(bitSet.get(i));
                }
            } else if (value instanceof Long) {
                long longValue = (Long) value;
                BitSet bitSet = BitSet.valueOf(new long[]{longValue});
                for (int i = 0; i < 64; i++) {
                    booleanValues.add(bitSet.get(i));
                }
            } else {
                throw new IllegalArgumentException(
                    "Value of type " + value.getClass().getName() +
                        " is not assignable to " + s7Field.getDataType().name() + " fields.");
            }
        }
        if(booleanValues.size() == 1) {
            return new PlcBoolean(booleanValues.get(0));
        } else {
            return new PlcList(booleanValues);
        }
    }

    private PlcValue internalEncodeInteger(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;
        // Initialize the constraints.
        BigInteger minValue;
        BigInteger maxValue;
        Class<? extends PlcValue> fieldType;
        Class<?> valueType;
        Object[] castedValues;
        //Here the App threw a noSuchMethodError because on line 398 we're looking for
        // a constructor depending on a type and we pass him an Array
        // so I change the <T>[].class to <T>.class
        switch (s7Field.getDataType()) {
            case BYTE:
                minValue = BigInteger.valueOf(Byte.MIN_VALUE);
                maxValue = BigInteger.valueOf(Byte.MAX_VALUE);
                fieldType = PlcByte.class;
                valueType = Byte.class;
                castedValues = new Byte[values.length];
                break;
            case WORD:
                minValue = BigInteger.valueOf(Short.MIN_VALUE);
                maxValue = BigInteger.valueOf(Short.MAX_VALUE);
                fieldType = PlcInteger.class;
                valueType = Short.class;
                castedValues = new Short[values.length];
                break;
            case DWORD:
                minValue = BigInteger.valueOf(Integer.MIN_VALUE);
                maxValue = BigInteger.valueOf(Integer.MAX_VALUE);
                fieldType = PlcInteger.class;
                valueType= Integer.class;
                castedValues = new Integer[values.length];
                break;
            case LWORD:
                minValue = BigInteger.valueOf(Long.MIN_VALUE);
                maxValue = BigInteger.valueOf(Long.MAX_VALUE);
                fieldType = PlcLong.class;
                valueType = Long.class;
                castedValues = new Long[values.length];
                break;
            case SINT:
                minValue = BigInteger.valueOf(Byte.MIN_VALUE);
                maxValue = BigInteger.valueOf(Byte.MAX_VALUE);
                fieldType = PlcInteger.class;
                valueType= Byte.class;
                castedValues = new Byte[values.length];
                break;
            case USINT:
                minValue = BigInteger.valueOf(0);
                maxValue = BigInteger.valueOf((long) Byte.MAX_VALUE * 2);
                fieldType = PlcInteger.class;
                valueType = Short.class;
                castedValues = new Short[values.length];
                break;
            case INT:
                minValue = BigInteger.valueOf(Short.MIN_VALUE);
                maxValue = BigInteger.valueOf(Short.MAX_VALUE);
                fieldType = PlcInteger.class;
                valueType = Short.class;
                castedValues = new Short[values.length];
                break;
            case UINT:
                minValue = BigInteger.valueOf(0);
                maxValue = BigInteger.valueOf(((long) Short.MAX_VALUE) * 2);
                fieldType = PlcInteger.class;
                valueType = Integer.class;
                castedValues = new Integer[values.length];
                break;
            case DINT:
                minValue = BigInteger.valueOf(Integer.MIN_VALUE);
                maxValue = BigInteger.valueOf(Integer.MAX_VALUE);
                fieldType = PlcInteger.class;
                //------------------------------------------------------------------------------------
                valueType = Integer.class;
                castedValues = new Integer[values.length];
                break;
            case UDINT:
                minValue = BigInteger.valueOf(0);
                maxValue = BigInteger.valueOf(((long) Integer.MAX_VALUE) * 2);
                fieldType = PlcLong.class;
                valueType= Long.class;
                castedValues = new Long[values.length];
                break;
            case LINT:
                minValue = BigInteger.valueOf(Long.MIN_VALUE);
                maxValue = BigInteger.valueOf(Long.MAX_VALUE);
                fieldType = PlcLong.class;
                valueType = Long.class;
                castedValues = new Long[values.length];
                break;
            case ULINT:
                minValue = BigInteger.valueOf(0);
                maxValue = BigInteger.valueOf(Long.MAX_VALUE).multiply(BigInteger.valueOf(2));
                fieldType = PlcBigInteger.class;
                valueType = BigInteger.class;
                castedValues = new BigInteger[values.length];
                break;
            default:
                throw new IllegalArgumentException(
                    "Cannot assign integer values to " + s7Field.getDataType().name() + " fields.");
        }

        // Check the constraints
        for (int i = 0; i < values.length; i++) {
            BigInteger value;
            if (values[i] instanceof BigInteger) {
                value = (BigInteger) values[i];
            } else if ((values[i] instanceof Byte) || (values[i] instanceof Short) ||
                (values[i] instanceof Integer) || (values[i] instanceof Long)) {
                value = BigInteger.valueOf(((Number) values[i]).longValue());
            } else {
                throw new IllegalArgumentException(
                    "Value of type " + values[i].getClass().getName() +
                        " is not assignable to " + s7Field.getDataType().name() + " fields.");
            }
            if (minValue.compareTo(value) > 0) {
                throw new IllegalArgumentException(
                    "Value of " + value.toString() + " exceeds allowed minimum for type "
                        + s7Field.getDataType().name() + " (min " + minValue.toString() + ")");
            }
            if (maxValue.compareTo(value) < 0) {
                throw new IllegalArgumentException(
                    "Value of " + value.toString() + " exceeds allowed maximum for type "
                        + s7Field.getDataType().name() + " (max " + maxValue.toString() + ")");
            }
            if (valueType == Byte.class) {
                castedValues[i] = value.byteValue();
            } else if (valueType== Short.class) {
                castedValues[i] = value.shortValue();
                //---------------------------------------
            } else if (valueType == Integer.class) {
                castedValues[i] = value.intValue();
            } else if (valueType == Long.class) {
                castedValues[i] = value.longValue();
            } else {
                castedValues[i] = value;
            }
        }
        // Create the field item.
        try {
            return fieldType.getDeclaredConstructor(valueType).newInstance(castedValues);
        } catch (InstantiationException | IllegalAccessException | InvocationTargetException | NoSuchMethodException e) {
            throw new PlcRuntimeException("Error initializing field class " + fieldType.getSimpleName(), e);
        }
    }

    private PlcValue internalEncodeFloatingPoint(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;

        // Initialize the constraints.
        Double minValue;
        Double maxValue;
        Class<? extends PlcValue> fieldType;
        Class<?> valueType;
        Object[] castedValues;
        switch (s7Field.getDataType()) {
            case REAL:
                minValue = (double) -Float.MAX_VALUE;
                maxValue = (double) Float.MAX_VALUE;
                fieldType = PlcFloat.class;
                valueType = Float[].class;
                castedValues = new Float[values.length];
                break;
            case LREAL:
                minValue = -Double.MAX_VALUE;
                maxValue = Double.MAX_VALUE;
                fieldType = PlcDouble.class;
                valueType = Double[].class;
                castedValues = new Double[values.length];
                break;
            default:
                throw new IllegalArgumentException(
                    "Cannot assign floating point values to " + s7Field.getDataType().name() + " fields.");
        }

        // Check the constraints
        for (int i = 0; i < values.length; i++) {
            Double value;
            if (values[i] instanceof Float) {
                value = ((Float) values[i]).doubleValue();
            } else if (values[i] instanceof Double) {
                value = (Double) values[i];
            } else {
                throw new IllegalArgumentException(
                    "Value of type " + values[i].getClass().getName() +
                        " is not assignable to " + s7Field.getDataType().name() + " fields.");
            }
            if (value < minValue) {
                throw new IllegalArgumentException(
                    "Value of " + value + " exceeds allowed minimum for type "
                        + s7Field.getDataType().name() + " (min " + minValue.toString() + ")");
            }
            if (value > maxValue) {
                throw new IllegalArgumentException(
                    "Value of " + value + " exceeds allowed maximum for type "
                        + s7Field.getDataType().name() + " (max " + maxValue.toString() + ")");
            }
            if (valueType == Float[].class) {
                castedValues[i] = value.floatValue();
            } else {
                castedValues[i] = value;
            }
        }

        // Create the field item.
        try {
            return fieldType.getDeclaredConstructor(valueType).newInstance(castedValues);
        } catch (InstantiationException | IllegalAccessException | InvocationTargetException | NoSuchMethodException e) {
            throw new PlcRuntimeException("Error initializing field class " + fieldType.getSimpleName(), e);
        }
    }

    private PlcValue internalEncodeString(PlcField field, Object[] values) {
        S7Field s7Field = (S7Field) field;

        // Initialize the constraints.
        int maxLength;
        boolean encoding16Bit;
        switch (s7Field.getDataType()) {
            case CHAR:
                maxLength = 1;
                encoding16Bit = false;
                break;
            case WCHAR:
                maxLength = 1;
                encoding16Bit = true;
                break;
            case STRING:
                maxLength = 254;
                encoding16Bit = false;
                break;
            case WSTRING:
                maxLength = 254;
                encoding16Bit = true;
                break;
            default:
                throw new IllegalArgumentException(
                    "Cannot assign string values to " + s7Field.getDataType().name() + " fields.");
        }

        // Check the constraints and create the strings.
        List<String> stringValues = new LinkedList<>();
        for (Object value : values) {
            if (value instanceof String) {
                String stringValue = (String) value;
                if (stringValue.length() > maxLength) {
                    throw new IllegalArgumentException(
                        "String length " + stringValue.length() + " exceeds allowed maximum for type "
                            + s7Field.getDataType().name() + " (max " + maxLength + ")");
                }
                stringValues.add(stringValue);
            }
            // All other types just translate to max one String character.
            else if (value instanceof Byte) {
                Byte byteValue = (Byte) value;
                byte[] stringBytes = new byte[]{byteValue};
                if (encoding16Bit) {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_16));
                } else {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_8));
                }
            } else if (value instanceof Short) {
                Short shortValue = (Short) value;
                byte[] stringBytes = new byte[2];
                stringBytes[0] = (byte) (shortValue >> 8);
                stringBytes[1] = (byte) (shortValue & 0xFF);
                if (encoding16Bit) {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_16));
                } else {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_8));
                }
            } else if (value instanceof Integer) {
                Integer integerValue = (Integer) value;
                byte[] stringBytes = new byte[4];
                stringBytes[0] = (byte) ((integerValue >> 24) & 0xFF);
                stringBytes[1] = (byte) ((integerValue >> 16) & 0xFF);
                stringBytes[2] = (byte) ((integerValue >> 8) & 0xFF);
                stringBytes[3] = (byte) (integerValue & 0xFF);
                if (encoding16Bit) {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_16));
                } else {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_8));
                }
            } else if (value instanceof Long) {
                Long longValue = (Long) value;
                byte[] stringBytes = new byte[8];
                stringBytes[0] = (byte) ((longValue >> 56) & 0xFF);
                stringBytes[1] = (byte) ((longValue >> 48) & 0xFF);
                stringBytes[2] = (byte) ((longValue >> 40) & 0xFF);
                stringBytes[3] = (byte) ((longValue >> 32) & 0xFF);
                stringBytes[4] = (byte) ((longValue >> 24) & 0xFF);
                stringBytes[5] = (byte) ((longValue >> 16) & 0xFF);
                stringBytes[6] = (byte) ((longValue >> 8) & 0xFF);
                stringBytes[7] = (byte) (longValue & 0xFF);
                if (encoding16Bit) {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_16));
                } else {
                    stringValues.add(new String(stringBytes, StandardCharsets.UTF_8));
                }
            } else {
                throw new IllegalArgumentException(
                    "Value of type " + value.getClass().getName() +
                        " is not assignable to " + s7Field.getDataType().name() + " fields.");
            }
        }
        // Create the field item.
        if(stringValues.size() == 1) {
            return new PlcString(stringValues.get(0));
        } else {
            return new PlcList(stringValues);
        }
    }

    private PlcValue internalEncodeTemporal(PlcField field, Object[] values) {
        if(values.length > 1) {
            return new PlcList(Arrays.asList(values));
        }
        S7Field s7Field = (S7Field) field;
        switch (s7Field.getDataType()) {
            case TIME:
                return new PlcTime((LocalTime) values[0]);
            case DATE:
                return new PlcDate((LocalDate) values[0]);
            case DATE_AND_TIME:
                return new PlcDateTime((LocalDateTime) values[0]);
            default:
                throw new IllegalArgumentException(
                    "Cannot assign temporal values to " + s7Field.getDataType().name() + " fields.");
        }
    }

}
