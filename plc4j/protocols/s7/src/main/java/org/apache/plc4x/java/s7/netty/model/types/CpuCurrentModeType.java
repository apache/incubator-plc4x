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

package org.apache.plc4x.java.s7.netty.model.types;

import java.util.HashMap;
import java.util.Map;
import org.apache.plc4x.java.api.model.PlcField;

/**
 *
 * @author cgarcia
 */
public enum CpuCurrentModeType implements PlcField {
    
    STOP("STOP", (byte) 0x00),
    WARM_RESTART("WARM_RESTART", (byte) 0x01),
    RUN("RUN", (byte) 0x02),
    HOT_RESTART("HOT_RESTART", (byte) 0x03),
    HOLD("HOLD", (byte) 0x04),
    COLD_RESTART("COLD_RESTART", (byte) 0x06),
    RUN_R("RUN_R", (byte) 0x09),
    LINK_UP("LINK_UP", (byte) 0x0b),
    UPDATE("UPDATE", (byte) 0x0c),
    ;
    
    
    private static final Map<Byte, CpuCurrentModeType> map;
    
    static {
        map = new HashMap<>();
        for (CpuCurrentModeType subevent : CpuCurrentModeType.values()) {
            map.put(subevent.code, subevent);
        }
    }    
    
    private final String event;
    private final byte code;
    
    CpuCurrentModeType(String event, byte code){
        this.event = event;
        this.code = code;
    }
    
    public String getEvent(){
        return event;
    }    
    
    public byte getCode() {
        return code;
    }    
    
    public static CpuCurrentModeType valueOfEvent(String event) {
        for (CpuCurrentModeType value : CpuCurrentModeType.values()) {
            if(value.getEvent().equals(event)) {
                return value;
            }
        }
        return null;
    }

    public static CpuCurrentModeType valueOf(byte code) {
        return map.get(code);
    }
    
    
}
