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

package org.apache.plc4x.java.s7.protocol;

import io.netty.buffer.ByteBuf;
import java.math.BigDecimal;
import java.math.BigInteger;
import java.time.Duration;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.util.Collection;
import java.util.Map;
import java.util.Objects;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.base.messages.InternalPlcReadRequest;
import org.apache.plc4x.java.base.messages.InternalPlcReadResponse;
import org.apache.plc4x.java.base.messages.items.BaseDefaultFieldItem;

/**
 *
 * @author cgarcia
 */
public class S7ByteReadResponse implements InternalPlcReadResponse{
    
    private final InternalPlcReadRequest request;
    private final Map<String, Pair<PlcResponseCode, ByteBuf>> values;        

    public S7ByteReadResponse(InternalPlcReadRequest request, Map<String, Pair<PlcResponseCode, ByteBuf>> values) {
        this.request = request;
        this.values = values;
    }

    /**
     *
     * @return
     */
    @Override
    public Map<String, Pair<PlcResponseCode, BaseDefaultFieldItem>> getValues() {
        return null;
    }
   
    public Map<String, Pair<PlcResponseCode, ByteBuf>> getByteBufValues() {
        return values;
    }   
    
    public ByteBuf getByteBufValues(String name) {
     ByteBuf bytebuf = getByteBufInternal(name);         
    return bytebuf;
    }  

    @Override
    public InternalPlcReadRequest  getRequest() {
        return request;
    }

    @Override
    public int getNumberOfValues(String name) {
        ByteBuf bytebuf = getByteBufInternal(name);
        return bytebuf.capacity();
    }

    @Override
    public Object getObject(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Object getObject(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Object> getAllObjects(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBoolean(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBoolean(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Boolean getBoolean(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Boolean getBoolean(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Boolean> getAllBooleans(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidByte(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidByte(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Byte getByte(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Byte getByte(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Byte> getAllBytes(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidShort(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidShort(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Short getShort(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Short getShort(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Short> getAllShorts(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidInteger(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidInteger(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Integer getInteger(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Integer getInteger(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Integer> getAllIntegers(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBigInteger(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBigInteger(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public BigInteger getBigInteger(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public BigInteger getBigInteger(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<BigInteger> getAllBigIntegers(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidLong(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidLong(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Long getLong(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Long getLong(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Long> getAllLongs(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidFloat(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidFloat(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Float getFloat(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Float getFloat(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Float> getAllFloats(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDouble(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDouble(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Double getDouble(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Double getDouble(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Double> getAllDoubles(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBigDecimal(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidBigDecimal(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public BigDecimal getBigDecimal(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public BigDecimal getBigDecimal(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<BigDecimal> getAllBigDecimals(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidString(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidString(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public String getString(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public String getString(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<String> getAllStrings(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidTime(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidTime(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalTime getTime(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalTime getTime(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<LocalTime> getAllTimes(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDate(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDate(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalDate getDate(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalDate getDate(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<LocalDate> getAllDates(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDateTime(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidDateTime(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalDateTime getDateTime(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public LocalDateTime getDateTime(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<LocalDateTime> getAllDateTimes(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidByteArray(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public boolean isValidByteArray(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Byte[] getByteArray(String name) {
        ByteBuf bytebuf = getByteBufInternal(name);
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Byte[] getByteArray(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<Byte[]> getAllByteArrays(String name) {
        throw new UnsupportedOperationException("Not supported yet.");  
    }

    @Override
    public Collection<String> getFieldNames() {
        return request.getFieldNames();
    }

    @Override
    public PlcField getField(String name) {
        return request.getField(name);
    }

    @Override
    public PlcResponseCode getResponseCode(String name) {
        if (values.get(name) == null) {
            throw new PlcInvalidFieldException(name);
        }
        return values.get(name).getKey();
    }

    @Override
    public boolean isValidDuration(String name) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public boolean isValidDuration(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Duration getDuration(String name) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Duration getDuration(String name, int index) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }

    @Override
    public Collection<Duration> getAllDuration(String name) {
        throw new UnsupportedOperationException("Not supported yet."); //To change body of generated methods, choose Tools | Templates.
    }            
    
    protected ByteBuf getByteBufInternal(String name) {
        Objects.requireNonNull(name, "Name argument required");
        // If this field doesn't exist, ignore it.
        if (values.get(name) == null) {
            throw new PlcInvalidFieldException(name);
        }
        //TODO: Check for error respont
        /*
        if (values.get(name).getKey() != PlcResponseCode.OK) {
            throw new PlcRuntimeException("Field '" + name + "' could not be fetched, response was " + values.get(name).getKey());
        }
        */
        // No need to check for "null" as this is already captured by the constructors.
        return values.get(name).getValue();
    }    
    
}
