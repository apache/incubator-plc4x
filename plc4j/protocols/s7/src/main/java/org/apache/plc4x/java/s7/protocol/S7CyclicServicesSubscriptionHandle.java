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

import java.util.HashSet;
import java.util.Set;
import java.util.function.Consumer;
import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.base.model.InternalPlcSubscriptionHandle;
import org.apache.plc4x.java.s7.netty.model.payloads.items.AssociatedValueItem;

/**
 *
 * @author cgarcia
 */
public class S7CyclicServicesSubscriptionHandle implements InternalPlcSubscriptionHandle {
    
    Set<Consumer<PlcSubscriptionEvent>> consumers = new HashSet<>();
        
    //CyClic Services information hanler
    private final String fieldName; //Subscription id from the request
    private final byte jobId; //Job-Id from cyclic subscriotion
    private final short error; //Register the error from the suscription.
    private final AssociatedValueItem item;

    public S7CyclicServicesSubscriptionHandle(String fieldName, 
            byte jobId, 
            short error, 
            AssociatedValueItem item) {
        this.fieldName = fieldName;
        this.jobId = jobId;
        this.error = error;
        this.item = item;
    }

    public Set<Consumer<PlcSubscriptionEvent>> getConsumers() {
        return consumers;
    }

    public String getFieldName() {
        return fieldName;
    }
    
    public byte getJobId() {
        return jobId;
    }

    public short getError() {
        return error;
    }
    
    public AssociatedValueItem getValueItem(){
        return item;
    }
                           
    @Override
    public PlcConsumerRegistration register(Consumer<PlcSubscriptionEvent> consumer) {
        consumers.add(consumer);
        return () -> {consumers.remove(consumer);};
    }
    
}
