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
package org.apache.plc4x.java.s7.utils;

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;

import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.temporal.ChronoUnit;

public class StaticHelper {

    public static LocalTime parseTiaTime(ReadBuffer io) {
        try {
            int millisSinceMidnight = io.readInt(32);
            return LocalTime.now().withHour(0).withMinute(0).withSecond(0).withNano(0).plus(
                millisSinceMidnight, ChronoUnit.MILLIS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static String parseS7String(ReadBuffer io){
        try {
            byte total = io.readByte(8);
            short stringLength = io.readShort(8);
            byte[] string = new byte[stringLength];
            for(int i = 0; i< stringLength ; i++){
                string[i]=io.readByte(8);
            }
            return new String(string);
        } catch (ParseException e) {
            e.printStackTrace();
            return null;
        }
    }

    public static void serializeS7String(WriteBuffer io, PlcValue value) {

    }

    public static void serializeTiaTime(WriteBuffer io, PlcValue value) {

    }

    public static LocalTime parseS5Time(ReadBuffer io) {
        try {
            int stuff = io.readInt(16);
            // TODO: Implement this correctly.
            throw new NotImplementedException("S5TIME not implemented");
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeS5Time(WriteBuffer io, PlcValue value) {

    }

    public static LocalTime parseTiaLTime(ReadBuffer io) {
        throw new NotImplementedException("LTime not implemented");
    }

    public static void serializeTiaLTime(WriteBuffer io, PlcValue value) {

    }

    public static LocalTime parseTiaTimeOfDay(ReadBuffer io) {
        try {
            int millisSinceMidnight = io.readUnsignedInt(32);
            return LocalTime.now().withHour(0).withMinute(0).withSecond(0).withNano(0).plus(
                millisSinceMidnight, ChronoUnit.MILLIS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaTimeOfDay(WriteBuffer io, PlcValue value) {

    }

    public static LocalDate parseTiaDate(ReadBuffer io) {
        try {
            int daysSince1990 = io.readUnsignedShort(16);
            return LocalDate.now().withYear(1990).withDayOfMonth(1).withMonth(1).plus(daysSince1990, ChronoUnit.DAYS);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaDate(WriteBuffer io, PlcValue value) {

    }

    public static LocalDateTime parseTiaDateTime(ReadBuffer io) {
        try {
            int year = convertByteToBcd(io.readByte(8));
            int month = convertByteToBcd(io.readByte(8));
            int day = convertByteToBcd(io.readByte(8));
            int hour = convertByteToBcd(io.readByte(8));
            int minute = convertByteToBcd(io.readByte(8));
            int second = convertByteToBcd(io.readByte(8));
            //skip the last 2 bytes no information present
            io.readByte(8);
            io.readByte(8);

            //data-type ranges from 1990 up to 2089
            if (year >= 90) {
                year += 1900;
            } else {
                year += 2000;
            }

            return LocalDateTime.of(year, month, day, hour, minute, second);
        } catch (ParseException e) {
            return null;
        }
    }

    public static void serializeTiaDateTime(WriteBuffer io, PlcValue value) {

    }

    /**
     * converts incoming byte to an integer regarding used BCD format
     *
     * @param incomingByte
     * @return converted BCD number
     */
    private static int convertByteToBcd(byte incomingByte) {
        int dec = (incomingByte >> 4) * 10;
        return dec + (incomingByte & 0x0f);
    }

}
