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

import org.apache.commons.lang3.NotImplementedException;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcUnsupportedOperationException;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcUnsubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcWriteRequest;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;
import org.apache.plc4x.java.base.messages.InternalPlcMessage;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.net.InetSocketAddress;
import java.net.Socket;
import java.util.Objects;
import java.util.Optional;

/**
 * Base class for implementing connections.
 * Per default, all operations (read, write, subscribe) are unsupported.
 * Concrete implementations should override the methods indicating connection capabilities
 * and for obtaining respective request builders.
 */
public abstract class AbstractPlcConnection implements PlcConnection, PlcConnectionMetadata {

    private static final Logger LOGGER = LoggerFactory.getLogger(AbstractPlcConnection.class);

    private static final int PING_TIMEOUT_MS = 1_000;

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

    /**
     * The default implementation uses the {@link #ping(int)} method.
     * Drivers that want to implement a more specific version have to overide it.
     */
    @Override
    public boolean isConnected() {
        return ping(PING_TIMEOUT_MS);
    }

    /**
     * Method that sends a test-request or ping to the PLC to check if the PLC is still there.
     * In most cases this method should only be used from the {@link #isConnected()} method.
     * This method can only be used if ghe {@link #getInetSocketAddress()} returns a Socket Adress.
     * Otherwise it throws a {@link NotImplementedException} to inform the user about that.
     */
    protected boolean ping(int timeout) {
        Optional<InetSocketAddress> address = getInetSocketAddress();
        if (!address.isPresent()) {
            throw new NotImplementedException("Tries to check the connection with ping, " +
                "but no Socket Address is given!");
        }
        Socket s = null;
        try {
            s = new Socket();
            s.connect(address.get(), timeout);
            return true;
        } catch (Exception e) {
            LOGGER.debug("Unable to ping PLC", e);
            return false;
        } finally {
            if (s != null) {
                try {
                    s.close();
                } catch (Exception e) {
                }
            }
        }
    }

    /**
     * Strategy Pattern method for the {@link #ping(int)} method.
     * If a driver has no Inet Socket adress, it has to return an Empty Optional, never null.
     */
    protected abstract Optional<InetSocketAddress> getInetSocketAddress();

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

    /**
     * Can be used to check and cast a parameter to its required internal type (can be used for general type checking too).
     *
     * @param o     the object to be checked against target {@code clazz}.
     * @param clazz the expected {@code clazz}.
     * @param <T>   the type of the expected {@code clazz}.
     * @return the cast type of {@code clazz}.
     */
    protected <T extends InternalPlcMessage> T checkInternal(Object o, Class<T> clazz) {
        Objects.requireNonNull(o);
        Objects.requireNonNull(clazz);
        if (!clazz.isInstance(o)) {
            throw new IllegalArgumentException("illegal type " + o.getClass() + ". Expected " + clazz);
        }
        return clazz.cast(o);
    }

}
