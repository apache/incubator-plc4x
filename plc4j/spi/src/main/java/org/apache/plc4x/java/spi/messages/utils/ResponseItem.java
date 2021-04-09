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
package org.apache.plc4x.java.spi.messages.utils;

import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.spi.utils.XmlSerializable;
import org.w3c.dom.Element;

public class ResponseItem<T> implements XmlSerializable {

    private final PlcResponseCode code;
    private final T value;

    public ResponseItem(PlcResponseCode code, T value) {
        this.code = code;
        this.value = value;
    }

    public PlcResponseCode getCode() {
        return code;
    }

    public T getValue() {
        return value;
    }

    @Override
    public void xmlSerialize(Element parent) {
        parent.setAttribute("result", code.name());
        if(value != null) {
            if (!(value instanceof XmlSerializable)) {
                throw new RuntimeException("Error serializing. Field value doesn't implement XmlSerializable");
            }
            ((XmlSerializable) value).xmlSerialize(parent);
        }
    }

}
