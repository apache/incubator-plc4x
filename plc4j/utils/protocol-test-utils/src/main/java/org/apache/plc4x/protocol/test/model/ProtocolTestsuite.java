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

package org.apache.plc4x.protocol.test.model;

import java.util.List;

public class ProtocolTestsuite {

    private final String name;
    private final List<Testcase> testcases;
    private final boolean littleEndian;

    public ProtocolTestsuite(String name, List<Testcase> testcases, boolean littleEndian) {
        this.name = name;
        this.testcases = testcases;
        this.littleEndian = littleEndian;
    }

    public String getName() {
        return name;
    }

    public List<Testcase> getTestcases() {
        return testcases;
    }

    public boolean isLittleEndian() {
        return littleEndian;
    }

}
