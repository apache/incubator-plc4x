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

import org.apache.plc4x.java.isotp.protocol.model.types.DeviceGroup;
import org.junit.jupiter.api.Test;

import java.lang.reflect.Constructor;

import static org.hamcrest.core.IsEqual.equalTo;
import static org.junit.Assert.assertThat;

class S7TsapIdEncoderTest {

    @Test
    void testConstructorPrivacy() {
        // Check that every declared constructor is inaccessible.
        for (Constructor<?> declaredConstructor : S7TsapIdEncoder.class.getDeclaredConstructors()) {
            assertThat(declaredConstructor.isAccessible(), equalTo(false));
        }
    }

    @Test
    void encodeS7TsapId() {
        short tsapId = S7TsapIdEncoder.encodeS7TsapId(DeviceGroup.PG_OR_PC, 1, 2);

        assertThat(tsapId, equalTo((short) 0x112));
    }

    @Test
    void decodeDeviceGroup() {
        DeviceGroup deviceGroup = S7TsapIdEncoder.decodeDeviceGroup((short) 0x112);

        assertThat(deviceGroup, equalTo(DeviceGroup.PG_OR_PC));
    }

    @Test
    void decodeRack() {
        int rack = S7TsapIdEncoder.decodeRack((short) 0x112);

        assertThat(rack, equalTo(1));
    }

    @Test
    void decodeSlot() {
        int slot = S7TsapIdEncoder.decodeSlot((short) 0x112);

        assertThat(slot, equalTo(2));
    }

}