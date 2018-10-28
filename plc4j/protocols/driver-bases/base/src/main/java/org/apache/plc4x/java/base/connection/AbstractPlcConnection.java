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

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcUnsupportedOperationException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcUnsubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;

/**
 * Base class for implementing connections.
 * Per default, all operations (read, write, subscribe) are unsupported.
 * Concrete implementations should override the methods indicating connection capabilities
 * and for obtaining respective request builders.
 */
public abstract class AbstractPlcConnection implements PlcConnection, PlcConnectionMetadata {

    @Override
    public PlcConnectionMetadata getMetadata() {
        return this;
    }

    @Override
    public boolean canRead() {
        return false;
    }

    @Override
    public boolean canWrite() {
        return false;
    }

    @Override
    public boolean canSubscribe() {
        return false;
    }

    @Override
    public PlcReadRequest.Builder readRequestBuilder() {
        throw new PlcUnsupportedOperationException("The connection does not support reading");
    }

    @Override
    public PlcWriteRequest.Builder writeRequestBuilder() {
        throw new PlcUnsupportedOperationException("The connection does not support writing");
    }

    @Override
    public PlcSubscriptionRequest.Builder subscriptionRequestBuilder() {
        throw new PlcUnsupportedOperationException("The connection does not support subscription");
    }

    @Override
    public PlcUnsubscriptionRequest.Builder unsubscriptionRequestBuilder() {
        throw new PlcUnsupportedOperationException("The connection does not support subscription");
    }

}
