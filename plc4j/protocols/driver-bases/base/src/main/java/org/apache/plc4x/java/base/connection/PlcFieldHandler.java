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
package org.apache.plc4x.java.base.connection;

import org.apache.plc4x.java.api.exceptions.PlcInvalidFieldException;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.base.messages.items.FieldItem;

public interface PlcFieldHandler {

    PlcField createField(String fieldQuery) throws PlcInvalidFieldException;

    FieldItem encodeBoolean(PlcField field, Object[] values);

    FieldItem encodeByte(PlcField field, Object[] values);

    FieldItem encodeShort(PlcField field, Object[] values);

    FieldItem encodeInteger(PlcField field, Object[] values);

    FieldItem encodeLong(PlcField field, Object[] values);

    FieldItem encodeFloat(PlcField field, Object[] values);

    FieldItem encodeDouble(PlcField field, Object[] values);

    FieldItem encodeString(PlcField field, Object[] values);

    FieldItem encodeTime(PlcField field, Object[] values);

    FieldItem encodeDate(PlcField field, Object[] values);

    FieldItem encodeDateTime(PlcField field, Object[] values);

}
