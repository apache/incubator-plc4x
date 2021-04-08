/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.spi.messages;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonIgnore;
import com.fasterxml.jackson.annotation.JsonProperty;
import com.fasterxml.jackson.annotation.JsonTypeInfo;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.messages.PlcWriteResponse;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.spi.messages.utils.ResponseItem;
import org.apache.plc4x.java.spi.utils.XmlSerializable;
import org.w3c.dom.Document;
import org.w3c.dom.Element;

import java.util.Collection;
import java.util.Map;

@JsonTypeInfo(use = JsonTypeInfo.Id.CLASS, property = "className")
public class DefaultPlcWriteResponse implements PlcWriteResponse, XmlSerializable {

    private final PlcWriteRequest request;
    private final Map<String, PlcResponseCode> values;

    @JsonCreator(mode = JsonCreator.Mode.PROPERTIES)
    public DefaultPlcWriteResponse(@JsonProperty("request") PlcWriteRequest request,
                                   @JsonProperty("values") Map<String, PlcResponseCode> values) {
        this.request = request;
        this.values = values;
    }

    @Override
    public PlcWriteRequest getRequest() {
        return request;
    }

    @Override
    @JsonIgnore
    public Collection<String> getFieldNames() {
        return request.getFieldNames();
    }

    @Override
    @JsonIgnore
    public PlcField getField(String name) {
        return request.getField(name);
    }

    @Override
    @JsonIgnore
    public PlcResponseCode getResponseCode(String name) {
        return values.get(name);
    }

    @Override
    public void xmlSerialize(Element parent) {
        Document doc = parent.getOwnerDocument();
        Element messageElement = doc.createElement("PlcWriteResponse");
        if(request instanceof XmlSerializable) {
            ((XmlSerializable) request).xmlSerialize(messageElement);
        }
        Element fieldsElement = doc.createElement("fields");
        messageElement.appendChild(fieldsElement);
        for (Map.Entry<String, PlcResponseCode> fieldEntry : values.entrySet()) {
            String fieldName = fieldEntry.getKey();
            final PlcResponseCode fieldResponseCode = fieldEntry.getValue();
            Element fieldNameElement = doc.createElement(fieldName);
            fieldNameElement.setAttribute("result", fieldResponseCode.name());
            fieldsElement.appendChild(fieldNameElement);
        }
        parent.appendChild(messageElement);
    }

}
