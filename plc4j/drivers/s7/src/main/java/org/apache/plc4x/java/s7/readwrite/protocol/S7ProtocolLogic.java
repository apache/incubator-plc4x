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
package org.apache.plc4x.java.s7.readwrite.protocol;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.Unpooled;
import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.s7.readwrite.*;
import org.apache.plc4x.java.s7.readwrite.context.S7DriverContext;
import org.apache.plc4x.java.s7.readwrite.field.S7Field;
import org.apache.plc4x.java.s7.readwrite.io.DataItemIO;
import org.apache.plc4x.java.s7.readwrite.types.*;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.spi.context.DriverContext;
import org.apache.plc4x.java.spi.generation.ParseException;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.generation.WriteBuffer;
import org.apache.plc4x.java.spi.messages.*;
import org.apache.plc4x.java.spi.transaction.RequestTransactionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * The S7 Protocol states that there can not be more then {min(maxAmqCaller, maxAmqCallee} "ongoing" requests.
 * So we need to limit those.
 * Thus, each request goes to a Work Queue and this Queue ensures, that only 3 are open at the same time.
 */
public class S7ProtocolLogic extends Plc4xProtocolBase<TPKTPacket>{

    private static final Logger logger = LoggerFactory.getLogger(S7ProtocolLogic.class);
    public static final Duration REQUEST_TIMEOUT = Duration.ofMillis(10000);

    private S7DriverContext s7DriverContext;
    private static final AtomicInteger tpduGenerator = new AtomicInteger(10);
    private RequestTransactionManager tm;

    @Override
    public void setDriverContext(DriverContext driverContext) {
        super.setDriverContext(driverContext);
        this.s7DriverContext = (S7DriverContext) driverContext;

        // Initialize Transaction Manager.
        // Until the number of concurrent requests is successfully negotiated we set it to a
        // maximum of only one request being able to be sent at a time. During the login process
        // No concurrent requests can be sent anyway. It will be updated when receiving the
        // S7ParameterSetupCommunication response.
        this.tm = new RequestTransactionManager(1);
    }


    @Override
    public void onConnect(ConversationContext<TPKTPacket> context) {
        logger.debug("Sending COTP Connection Request");
        // Open the session on ISO Transport Protocol first.
        TPKTPacket packet = new TPKTPacket(createCOTPConnectionRequest(
            s7DriverContext.getCalledTsapId(), s7DriverContext.getCallingTsapId(), s7DriverContext.getCotpTpduSize()));

        context.sendRequest(packet)
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .check(p -> p.getPayload() instanceof COTPPacketConnectionResponse)
            .unwrap(p -> (COTPPacketConnectionResponse) p.getPayload())
            .handle(cotpPacketConnectionResponse -> {
                logger.debug("Got COTP Connection Response");
                logger.debug("Sending S7 Connection Request");
                context.sendRequest(createS7ConnectionRequest(cotpPacketConnectionResponse))
                    .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
                    .unwrap(TPKTPacket::getPayload)
                    .only(COTPPacketData.class)
                    .unwrap(COTPPacket::getPayload)
                    .only(S7MessageResponseData.class)
                    .unwrap(S7Message::getParameter)
                    .only(S7ParameterSetupCommunication.class)
                    .handle(setupCommunication -> {
                        logger.debug("Got S7 Connection Response");
                        // Save some data from the response.
                        s7DriverContext.setMaxAmqCaller(setupCommunication.getMaxAmqCaller());
                        s7DriverContext.setMaxAmqCallee(setupCommunication.getMaxAmqCallee());
                        s7DriverContext.setPduSize(setupCommunication.getPduLength());

                        // Update the number of concurrent requests to the negotiated number.
                        // I have never seen anything else than equal values for caller and
                        // callee, but if they were different, we're only limiting the outgoing
                        // requests.
                        tm.setNumberOfConcurrentRequests(s7DriverContext.getMaxAmqCallee());

                        // If the controller type is explicitly set, were finished with the login
                        // process. If it's set to ANY, we have to query the serial number information
                        // in order to detect the type of PLC.
                        if (s7DriverContext.getControllerType() != S7ControllerType.ANY) {
                            // Send an event that connection setup is complete.
                            context.fireConnected();
                            return;
                        }

                        // Prepare a message to request the remote to identify itself.
                        logger.debug("Sending S7 Identification Request");
                        TPKTPacket tpktPacket = createIdentifyRemoteMessage();
                        context.sendRequest(tpktPacket)
                            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
                            .check(p -> p.getPayload() instanceof COTPPacketData)
                            .unwrap(p -> ((COTPPacketData) p.getPayload()))
                            .check(p -> p.getPayload() instanceof S7MessageUserData)
                            .unwrap(p -> ((S7MessageUserData) p.getPayload()))
                            .check(p -> p.getPayload() instanceof S7PayloadUserData)
                            .handle(messageUserData -> {
                                logger.debug("Got S7 Identification Response");
                                S7PayloadUserData payloadUserData = (S7PayloadUserData) messageUserData.getPayload();
                                extractControllerTypeAndFireConnected(context, payloadUserData);
                            });
                    });
            });
    }

    @Override
    public CompletableFuture<PlcReadResponse> read(PlcReadRequest readRequest) {
        DefaultPlcReadRequest request = (DefaultPlcReadRequest) readRequest;
        List<S7VarRequestParameterItem> requestItems = new ArrayList<>(request.getNumberOfFields());
        for (PlcField field : request.getFields()) {
            requestItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(field)));
        }

        // Create a read request template.
        // tpuId will be inserted before sending in #readInternal so we insert -1 as dummy here
        final S7MessageRequest s7MessageRequest = new S7MessageRequest(-1,
            new S7ParameterReadVarRequest(requestItems.toArray(new S7VarRequestParameterItem[0])),
            null);

        // Just send a single response and chain it as Response
        return toPlcReadResponse((InternalPlcReadRequest) readRequest, readInternal(s7MessageRequest));
    }

    /** Maps the S7ReadResponse of a PlcReadRequest to a PlcReadRespoonse */
    private CompletableFuture<PlcReadResponse> toPlcReadResponse(InternalPlcReadRequest readRequest, CompletableFuture<S7MessageResponseData> response) {
        return response
            .thenApply(p -> {
                try {
                    return ((PlcReadResponse) decodeReadResponse(p, readRequest));
                } catch (PlcProtocolException e) {
                    throw new PlcRuntimeException("Unable to decode Response", e);
                }
            });
    }

    /**
     * Sends one Read over the Wire and internally returns the Response
     * Do sending of normally sized single-message request.
     *
     * Assumes that the {@link S7MessageRequest} and its expected {@link S7MessageResponseData}
     * and does not further check that!
     */
    private CompletableFuture<S7MessageResponseData> readInternal(S7MessageRequest request) {
        CompletableFuture<S7MessageResponseData> future = new CompletableFuture<>();

        //To stay synced, the tpudGenerator must not go over the Short.MAX_VALUE
        //as it is cqsted to short later on to resync request and response
        //tpudRef 0-10 is reserved and if it exceeds 65535 it will be casted to 0
        //here we reset it to its initial value of 10
        if(tpduGenerator.get()>= (Short.MAX_VALUE-1)){
            tpduGenerator.set(10);
        }
        int tpduId = tpduGenerator.getAndIncrement();

        // Create a new Request with correct tpuId (is not known before)
        S7MessageRequest s7MessageRequest = new S7MessageRequest(tpduId, request.getParameter(), request.getPayload());

        TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null, s7MessageRequest, true, (short) tpduId));
        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(tpktPacket)
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> (COTPPacketData) p.getPayload())
            .check(p -> p.getPayload() instanceof S7MessageResponseData)
            .unwrap(p -> (S7MessageResponseData) p.getPayload())
            .check(p -> p.getTpduReference() == tpduId)
            .check(p -> p.getParameter() instanceof S7ParameterReadVarResponse)
            .handle(p -> {
                future.complete(p);
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    @Override
    public CompletableFuture<PlcWriteResponse> write(PlcWriteRequest writeRequest) {
        CompletableFuture<PlcWriteResponse> future = new CompletableFuture<>();
        DefaultPlcWriteRequest request = (DefaultPlcWriteRequest) writeRequest;
        List<S7VarRequestParameterItem> parameterItems = new ArrayList<>(request.getNumberOfFields());
        List<S7VarPayloadDataItem> payloadItems = new ArrayList<>(request.getNumberOfFields());
        for (String fieldName : request.getFieldNames()) {
            final S7Field field = (S7Field) request.getField(fieldName);
            final PlcValue plcValue = request.getPlcValue(fieldName);
            parameterItems.add(new S7VarRequestParameterItemAddress(encodeS7Address(field)));
            payloadItems.add(serializePlcValue(field, plcValue));
        }
        if(tpduGenerator.get()>= (Short.MAX_VALUE-1)){
            tpduGenerator.set(10);
        }
        final int tpduId = tpduGenerator.getAndIncrement();
        TPKTPacket tpktPacket = new TPKTPacket(new COTPPacketData(null,
            new S7MessageRequest(tpduId,
                new S7ParameterWriteVarRequest(parameterItems.toArray(new S7VarRequestParameterItem[0])),
                new S7PayloadWriteVarRequest(payloadItems.toArray(new S7VarPayloadDataItem[0]))),
            true, (short) tpduId));

        // Start a new request-transaction (Is ended in the response-handler)
        RequestTransactionManager.RequestTransaction transaction = tm.startRequest();
        transaction.submit(() -> context.sendRequest(tpktPacket)
            .expectResponse(TPKTPacket.class, REQUEST_TIMEOUT)
            .onTimeout(future::completeExceptionally)
            .onError((p, e) -> future.completeExceptionally(e))
            .check(p -> p.getPayload() instanceof COTPPacketData)
            .unwrap(p -> ((COTPPacketData) p.getPayload()))
            .check(p -> p.getPayload() instanceof S7MessageResponseData)
            .unwrap(p -> ((S7MessageResponseData) p.getPayload()))
            .check(p -> p.getTpduReference() == tpduId)
            .check(p -> p.getParameter() instanceof S7ParameterWriteVarResponse)
            .handle(p -> {
                try {
                    future.complete(((PlcWriteResponse) decodeWriteResponse(p, ((InternalPlcWriteRequest) writeRequest))));
                } catch (PlcProtocolException e) {
                    logger.warn(String.format("Error sending 'write' message: '%s'", e.getMessage()), e);
                }
                // Finish the request-transaction.
                transaction.endRequest();
            }));
        return future;
    }

    @Override
    public void close(ConversationContext<TPKTPacket> context) {
        // TODO Implement Closing on Protocol Level
    }


    private void extractControllerTypeAndFireConnected(ConversationContext<TPKTPacket> context, S7PayloadUserData payloadUserData) {
        for (S7PayloadUserDataItem item : payloadUserData.getItems()) {
            if (!(item instanceof S7PayloadUserDataItemCpuFunctionReadSzlResponse)) {
                continue;
            }
            S7PayloadUserDataItemCpuFunctionReadSzlResponse readSzlResponseItem =
                (S7PayloadUserDataItemCpuFunctionReadSzlResponse) item;
            for (SzlDataTreeItem readSzlResponseItemItem : readSzlResponseItem.getItems()) {
                if (readSzlResponseItemItem.getItemIndex() != 0x0001) {
                    continue;
                }
                final String articleNumber = new String(readSzlResponseItemItem.getMlfb());
                s7DriverContext.setControllerType(decodeControllerType(articleNumber));

                // Send an event that connection setup is complete.
                context.fireConnected();
            }
        }
    }

    private TPKTPacket createIdentifyRemoteMessage() {
        S7MessageUserData identifyRemoteMessage = new S7MessageUserData(1, new S7ParameterUserData(new S7ParameterUserDataItem[]{
            new S7ParameterUserDataItemCPUFunctions((short) 0x11, (byte) 0x4, (byte) 0x4, (short) 0x01, (short) 0x00, null, null, null)
        }), new S7PayloadUserData(new S7PayloadUserDataItem[]{
            new S7PayloadUserDataItemCpuFunctionReadSzlRequest(DataTransportErrorCode.OK, DataTransportSize.OCTET_STRING, new SzlId(SzlModuleTypeClass.CPU, (byte) 0x00, SzlSublist.MODULE_IDENTIFICATION), 0x0000)
        }));
        COTPPacketData cotpPacketData = new COTPPacketData(null, identifyRemoteMessage, true, (short) 2);
        return new TPKTPacket(cotpPacketData);
    }

    private TPKTPacket createS7ConnectionRequest(COTPPacketConnectionResponse cotpPacketConnectionResponse) {
        for (COTPParameter parameter : cotpPacketConnectionResponse.getParameters()) {
            if (parameter instanceof COTPParameterCalledTsap) {
                COTPParameterCalledTsap cotpParameterCalledTsap = (COTPParameterCalledTsap) parameter;
                s7DriverContext.setCalledTsapId(cotpParameterCalledTsap.getTsapId());
            } else if (parameter instanceof COTPParameterCallingTsap) {
                COTPParameterCallingTsap cotpParameterCallingTsap = (COTPParameterCallingTsap) parameter;
                if(cotpParameterCallingTsap.getTsapId() != s7DriverContext.getCallingTsapId()) {
                    s7DriverContext.setCallingTsapId(cotpParameterCallingTsap.getTsapId());
                    logger.warn(String.format("Switching calling TSAP id to '%s'", s7DriverContext.getCallingTsapId()));
                }
            } else if (parameter instanceof COTPParameterTpduSize) {
                COTPParameterTpduSize cotpParameterTpduSize = (COTPParameterTpduSize) parameter;
                s7DriverContext.setCotpTpduSize(cotpParameterTpduSize.getTpduSize());
            } else {
                logger.warn(String.format("Got unknown parameter type '%s'", parameter.getClass().getName()));
            }
        }

        // Send an S7 login message.
        S7ParameterSetupCommunication s7ParameterSetupCommunication =
            new S7ParameterSetupCommunication(
                s7DriverContext.getMaxAmqCaller(), s7DriverContext.getMaxAmqCallee(), s7DriverContext.getPduSize());
        S7Message s7Message = new S7MessageRequest(0, s7ParameterSetupCommunication,
            null);
        COTPPacketData cotpPacketData = new COTPPacketData(null, s7Message, true, (short) 1);
        return new TPKTPacket(cotpPacketData);
    }

    private COTPPacketConnectionRequest createCOTPConnectionRequest(int calledTsapId, int callingTsapId, COTPTpduSize cotpTpduSize) {
        return new COTPPacketConnectionRequest(
            new COTPParameter[]{
                new COTPParameterCalledTsap(calledTsapId),
                new COTPParameterCallingTsap(callingTsapId),
                new COTPParameterTpduSize(cotpTpduSize)
            }, null, (short) 0x0000, (short) 0x000F, COTPProtocolClass.CLASS_0);
    }


    private PlcResponse decodeReadResponse(S7MessageResponseData responseMessage, InternalPlcReadRequest plcReadRequest) throws PlcProtocolException {
        S7PayloadReadVarResponse payload = (S7PayloadReadVarResponse) responseMessage.getPayload();

        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcReadRequest.getNumberOfFields() != payload.getItems().length) {
            throw new PlcProtocolException(
                "The number of requested items doesn't match the number of returned items");
        }

        Map<String, Pair<PlcResponseCode, PlcValue>> values = new HashMap<>();
        S7VarPayloadDataItem[] payloadItems = payload.getItems();
        int index = 0;
        for (String fieldName : plcReadRequest.getFieldNames()) {
            S7Field field = (S7Field) plcReadRequest.getField(fieldName);
            S7VarPayloadDataItem payloadItem = payloadItems[index];

            PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
            PlcValue plcValue = null;
            ByteBuf data = Unpooled.wrappedBuffer(payloadItem.getData());
            if (responseCode == PlcResponseCode.OK) {
                plcValue = parsePlcValue(field, data);
            }
            Pair<PlcResponseCode, PlcValue> result = new ImmutablePair<>(responseCode, plcValue);
            values.put(fieldName, result);
            index++;
        }

        return new DefaultPlcReadResponse(plcReadRequest, values);
    }

    private PlcResponse decodeWriteResponse(S7MessageResponseData responseMessage, InternalPlcWriteRequest plcWriteRequest) throws PlcProtocolException {
        S7PayloadWriteVarResponse payload = (S7PayloadWriteVarResponse) responseMessage.getPayload();

        // If the numbers of items don't match, we're in big trouble as the only
        // way to know how to interpret the responses is by aligning them with the
        // items from the request as this information is not returned by the PLC.
        if (plcWriteRequest.getNumberOfFields() != payload.getItems().length) {
            throw new PlcProtocolException(
                "The number of requested items doesn't match the number of returned items");
        }

        Map<String, PlcResponseCode> responses = new HashMap<>();
        S7VarPayloadStatusItem[] payloadItems = payload.getItems();
        int index = 0;
        for (String fieldName : plcWriteRequest.getFieldNames()) {
            S7VarPayloadStatusItem payloadItem = payloadItems[index];

            PlcResponseCode responseCode = decodeResponseCode(payloadItem.getReturnCode());
            responses.put(fieldName, responseCode);
            index++;
        }

        return new DefaultPlcWriteResponse(plcWriteRequest, responses);
    }

    private S7VarPayloadDataItem serializePlcValue(S7Field field, PlcValue plcValue) {
        try {
            DataTransportSize transportSize = (field.getDataType().getDataProtocolId() == 1) ?
                DataTransportSize.BIT : DataTransportSize.BYTE_WORD_DWORD;
            WriteBuffer writeBuffer = DataItemIO.staticSerialize(plcValue, field.getDataType().getDataProtocolId());
            if(writeBuffer != null) {
                byte[] data = writeBuffer.getData();
                return new S7VarPayloadDataItem(DataTransportErrorCode.OK, transportSize, data.length, data);
            }
        } catch (ParseException e) {
            logger.warn(String.format("Error serializing field item of type: '%s'", field.getDataType().name()), e);
        }
        return null;
    }

    private PlcValue parsePlcValue(S7Field field, ByteBuf data) {
        ReadBuffer readBuffer = new ReadBuffer(data.array());
        try {
            return DataItemIO.staticParse(readBuffer, field.getDataType().getDataProtocolId());
        } catch (ParseException e) {
            logger.warn(String.format("Error parsing field item of type: '%s'", field.getDataType().name()), e);
        }
        return null;
    }

    /**
     * Helper to convert the return codes returned from the S7 into one of our standard
     * PLC4X return codes
     * @param dataTransportErrorCode S7 return code
     * @return PLC4X return code.
     */
    private PlcResponseCode decodeResponseCode(DataTransportErrorCode dataTransportErrorCode) {
        if (dataTransportErrorCode == null) {
            return PlcResponseCode.INTERNAL_ERROR;
        }
        switch (dataTransportErrorCode) {
            case OK:
                return PlcResponseCode.OK;
            case NOT_FOUND:
                return PlcResponseCode.NOT_FOUND;
            case INVALID_ADDRESS:
                return PlcResponseCode.INVALID_ADDRESS;
            case DATA_TYPE_NOT_SUPPORTED:
                return PlcResponseCode.INVALID_DATATYPE;
            default:
                return PlcResponseCode.INTERNAL_ERROR;
        }
    }

    /**
     * Little helper method to parse Siemens article numbers and extract the type of controller.
     *
     * @param articleNumber article number string.
     * @return type of controller.
     */
    private S7ControllerType decodeControllerType(String articleNumber) {
        if (!articleNumber.startsWith("6ES7 ")) {
            return S7ControllerType.ANY;
        }
        String model = articleNumber.substring(articleNumber.indexOf(' ') + 1, articleNumber.indexOf(' ') + 2);
        switch (model) {
            case "2":
                return S7ControllerType.S7_1200;
            case "5":
                return S7ControllerType.S7_1500;
            case "3":
                return S7ControllerType.S7_300;
            case "4":
                return S7ControllerType.S7_400;
            default:
                if (logger.isInfoEnabled()) {
                    logger.info(String.format("Looking up unknown article number %s", articleNumber));
                }
                return S7ControllerType.ANY;
        }
    }

    /**
     * Currently we only support the S7 Any type of addresses. This helper simply converts the S7Field
     * from PLC4X into S7Address objects.
     * @param field S7Field instance we need to convert into an S7Address
     * @return the S7Address
     */
    protected S7Address encodeS7Address(PlcField field) {
        if (!(field instanceof S7Field)) {
            throw new PlcRuntimeException("Unsupported address type " + field.getClass().getName());
        }
        S7Field s7Field = (S7Field) field;
        return new S7AddressAny(s7Field.getDataType(), s7Field.getNumElements(), s7Field.getBlockNumber(),
            s7Field.getMemoryArea(), s7Field.getByteOffset(), s7Field.getBitOffset());
    }
}
