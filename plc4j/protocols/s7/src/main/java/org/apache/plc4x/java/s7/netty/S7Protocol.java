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
package org.apache.plc4x.java.s7.netty;

import io.netty.buffer.ByteBuf;
import io.netty.buffer.ByteBufUtil;
import io.netty.buffer.Unpooled;
import io.netty.channel.*;
import io.netty.handler.codec.MessageToMessageDecoder;
import io.netty.util.AttributeKey;
import io.netty.util.concurrent.Future;
import io.netty.util.concurrent.PromiseCombiner;
import java.lang.reflect.Field;
import java.nio.charset.Charset;
import java.time.LocalDate;
import java.time.LocalDateTime;
import java.time.LocalTime;
import java.time.temporal.ChronoUnit;
import java.util.*;
import org.apache.commons.lang3.reflect.FieldUtils;
import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.exceptions.PlcProtocolException;
import org.apache.plc4x.java.api.exceptions.PlcProtocolPayloadTooBigException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.isotp.protocol.IsoTPProtocol;
import org.apache.plc4x.java.isotp.protocol.events.IsoTPConnectedEvent;
import org.apache.plc4x.java.isotp.protocol.model.IsoTPMessage;
import org.apache.plc4x.java.isotp.protocol.model.tpdus.DataTpdu;
import org.apache.plc4x.java.s7.netty.events.S7ConnectedEvent;
import org.apache.plc4x.java.s7.netty.model.messages.S7Message;
import org.apache.plc4x.java.s7.netty.model.messages.S7RequestMessage;
import org.apache.plc4x.java.s7.netty.model.messages.S7ResponseMessage;
import org.apache.plc4x.java.s7.netty.model.messages.SetupCommunicationRequestMessage;
import org.apache.plc4x.java.s7.netty.model.params.*;
import org.apache.plc4x.java.s7.netty.model.params.items.S7AnyVarParameterItem;
import org.apache.plc4x.java.s7.netty.model.params.items.VarParameterItem;
import org.apache.plc4x.java.s7.netty.model.payloads.AlarmMessagePayload;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuCyclicServicesRequestPayload;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuCyclicServicesResponsePayload;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuDiagnosticMessagePayload;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuMessageSubscriptionServicePayload;
import org.apache.plc4x.java.s7.netty.model.payloads.CpuServicesPayload;
import org.apache.plc4x.java.s7.netty.model.payloads.S7Payload;
import org.apache.plc4x.java.s7.netty.model.payloads.VarPayload;
import org.apache.plc4x.java.s7.netty.model.payloads.items.AlarmMessageItem;
import org.apache.plc4x.java.s7.netty.model.payloads.items.AssociatedValueItem;
import org.apache.plc4x.java.s7.netty.model.payloads.items.CpuDiagnosticMessageItem;
import org.apache.plc4x.java.s7.netty.model.payloads.items.MessageObjectItem;
import org.apache.plc4x.java.s7.netty.model.payloads.items.S7AnyVarPayloadItem;
import org.apache.plc4x.java.s7.netty.model.payloads.items.VarPayloadItem;
import org.apache.plc4x.java.s7.netty.model.types.*;
import org.apache.plc4x.java.s7.netty.strategies.S7MessageProcessor;
import org.apache.plc4x.java.s7.netty.util.S7SizeHelper;
import org.apache.plc4x.java.s7.types.S7ControllerType;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * Communication Layer between the Application level ({@link Plc4XS7Protocol}) and the lower level (tcp) that sends and receives {@link S7Message}s.
 * This layer also handles the control over the "wire", i.e., the queues of incoming and outgoing messages.
 * Furthermore, here {@link S7Message}s are marshalled and unmarshalled to {@link ByteBuf}s to be send over wire.
 *
 * Before messages are send to the wire an optional {@link S7MessageProcessor} can be applied.
 *
 * @see S7MessageProcessor
 */
public class S7Protocol extends ChannelDuplexHandler {

    private static final byte S7_PROTOCOL_MAGIC_NUMBER = 0x32;

    private static final Logger logger = LoggerFactory.getLogger(S7Protocol.class);

    private final MessageToMessageDecoder<Object> decoder = new MessageToMessageDecoder<Object>() {

        @Override
        public boolean acceptInboundMessage(Object msg) {
            return msg instanceof IsoTPMessage;
        }

        @Override
        @SuppressWarnings("unchecked")
        protected void decode(ChannelHandlerContext ctx, Object msg, List<Object> out) {
            S7Protocol.this.decode(ctx, (IsoTPMessage) msg, out);
        }
    };

    private short maxAmqCaller;
    private short maxAmqCallee;
    private short pduSize;
    private final AttributeKey<Short> pduSizeKey = AttributeKey.valueOf("pduSizeKey");    
    private S7ControllerType controllerType;

    // For detecting the lower layers.
    private ChannelHandler prevChannelHandler;
    private S7MessageProcessor messageProcessor;

    // For being able to respect the max AMQ restrictions.
    private PendingWriteQueue queue;
    private Map<Short, DataTpdu> sentButUnacknowledgedTpdus;
    
    /*
     * All fragmentsof data send by PLC are stored here,
     * until the PLC indicates that it is the last one,
     * then a single message is assembled.
    */    
    private Map<Short, Pair<LocalDateTime, Queue<ByteBuf>>> fragmentedData;

    public S7Protocol(short requestedMaxAmqCaller, short requestedMaxAmqCallee, short requestedPduSize,
                      S7ControllerType controllerType, S7MessageProcessor messageProcessor) {
        this.maxAmqCaller = requestedMaxAmqCaller;
        this.maxAmqCallee = requestedMaxAmqCallee;
        this.pduSize = requestedPduSize;
        this.controllerType = controllerType;
        this.messageProcessor = messageProcessor;
        sentButUnacknowledgedTpdus = new HashMap<>();
        fragmentedData = new HashMap<>();
    }

    @Override
    public void channelRegistered(ChannelHandlerContext ctx) {
        this.queue = new PendingWriteQueue(ctx);
        ctx.channel().attr(pduSizeKey).set(pduSize);
        
        try {
            Field prevField = FieldUtils.getField(ctx.getClass(), "prev", true);
            if(prevField != null) {
                ChannelHandlerContext prevContext = (ChannelHandlerContext) prevField.get(ctx);
                prevChannelHandler = prevContext.handler();
            }
        } catch(Exception e) {
            logger.error("Error accessing field 'prev'", e);
        }
    }

    @Override
    public void channelUnregistered(ChannelHandlerContext ctx) throws Exception {
        this.queue.removeAndWriteAll();
        super.channelUnregistered(ctx);
    }

    @Override
    public void channelInactive(ChannelHandlerContext ctx) throws Exception {
        // Send everything so we get a proper failure for those pending writes
        this.queue.removeAndWriteAll();
        super.channelInactive(ctx);
    }

    /**
     * If the S7 protocol layer is used over Iso TP, then after receiving a 
     * {@link IsoTPConnectedEvent} the corresponding S7 setup communication 
     * message has to be sent in order to negotiate the S7 protocol layer.
     *
     * For high availability handling, the proxy must be incorporated.
     * This will replicate the event to the communication channels.
     * 
     * @param ctx the current protocol layers context
     * @param evt the event
     * @throws Exception throws an exception if something goes wrong internally
     */
    @Override
    public void userEventTriggered(ChannelHandlerContext ctx, Object evt) throws Exception {
        // If we are using S7 inside of IsoTP or we have a proxy, then we need 
        // to intercept IsoTPs connected events.
        if (((prevChannelHandler instanceof IsoTPProtocol) || 
            (prevChannelHandler instanceof S7HProxy)) && 
            (evt instanceof IsoTPConnectedEvent)) {
            // Setup Communication
            SetupCommunicationRequestMessage setupCommunicationRequest =
                new SetupCommunicationRequestMessage((short) 0, maxAmqCaller, maxAmqCallee, pduSize, null);

            ctx.channel().writeAndFlush(setupCommunicationRequest);
        }

        else {
            super.userEventTriggered(ctx, evt);
        }
    }

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Encoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    public void write(ChannelHandlerContext ctx, Object msg, ChannelPromise promise) {
        try {
            if(msg instanceof S7Message) {
                S7Message in = (S7Message) msg;

                // Give message processors to process the incoming message.
                //TODO: messageProcessor move to netty next layer, no a plugin code
                Collection<? extends S7Message> messages;
                if ((messageProcessor != null) && (in instanceof S7RequestMessage)) {
                    messages = messageProcessor.processRequest((S7RequestMessage) in, pduSize);
                } else {
                    messages = Collections.singleton(in);
                }

                // Create a promise that has to be called multiple times.
                PromiseCombiner promiseCombiner = new PromiseCombiner();
                for (S7Message message : messages) {
                    ByteBuf buf = Unpooled.buffer();
                    writeS7Message(promise.channel(), promiseCombiner, message, buf);
                }
                promiseCombiner.finish(promise);

                // Start sending the queue content.
                trySendingMessages(ctx);
            }
            // Especially during the phase of connection establishment, we might be sending
            // messages of a lower level protocol, so if it's not S7, we forward it to the next
            // in the pipeline and hope it can handle it. If no layer can handle it Netty will
            // exceptionally complete the future.
            else {
                ctx.write(msg, promise);
            }
        } catch (Exception e) {
            promise.setFailure(e);
        }
    }

    private void writeS7Message(Channel channel, PromiseCombiner promiseCombiner,
                                S7Message message, ByteBuf buf) throws PlcProtocolException {
        encodeHeader(message, buf);
        encodeParameters(message, buf);
        encodePayloads(message, buf);

        // Check if the message doesn't exceed the negotiated maximum size.
        if (buf.writerIndex() > pduSize) {
            throw new PlcProtocolPayloadTooBigException("s7:"+message.getTpduReference(), pduSize, buf.writerIndex(), message);
        } else {
            ChannelPromise subPromise = new DefaultChannelPromise(channel);
            // The tpduRef was 0x01 but had to be changed to 0x00 in order to support Siemens LOGO devices.
            System.out.println("S7Message; " + message.getTpduReference());
            queue.add(new DataTpdu(true, (byte) 0x00, Collections.emptyList(), buf, message), subPromise);
            
            promiseCombiner.add((Future) subPromise);
            logger.debug("writeS7Message: S7 Message with id {} queued", message.getTpduReference());
        }
    }

    private void encodePayloads(S7Message in, ByteBuf buf) throws PlcProtocolException {
        if(in.getPayloads() != null) {
            Iterator<S7Payload> payloadIterator = in.getPayloads().iterator();
            while(payloadIterator.hasNext()) {
                S7Payload payload = payloadIterator.next();
                switch (payload.getType()) {
                    case WRITE_VAR:
                        encodeWriteVarPayload((VarPayload) payload, buf, payloadIterator.hasNext());
                        break;
                    case CPU_SERVICES: 
                        if (payload instanceof CpuServicesPayload) {
                            encodeCpuServicesPayload((CpuServicesPayload) payload, buf);
                        } else if (payload instanceof CpuMessageSubscriptionServicePayload) {
                            encodeCpuMessageSubcriptionPayload((CpuMessageSubscriptionServicePayload) payload, buf);
                        } else if (payload instanceof AlarmMessagePayload) {
                             encodeAlarmMessagePayload((AlarmMessagePayload) payload, buf);
                        } else if (payload instanceof CpuCyclicServicesRequestPayload){
                            encodeCpuCyclicSubscriptionPayload((CpuCyclicServicesRequestPayload) payload, buf);
                        }
                        break;
                    default:
                        throw new PlcProtocolException("Writing payloads of type " +
                            payload.getType().name() + " not implemented.");
                }
            }
        }
    }

    /*
    * 
    */
    private void encodeWriteVarPayload(VarPayload varPayload, ByteBuf buf, boolean lastItem) {
        
        //All payloads for the request come here
        int i=1;
        for (VarPayloadItem payloadItem : varPayload.getItems()) {            
            buf.writeByte(payloadItem.getReturnCode().getCode());
            buf.writeByte(payloadItem.getDataTransportSize().getCode());
            //TODO: Check if this is correct?!?! Might be problems with sizeInBits = true/false
            //TODO: This field is in BCD
            switch (payloadItem.getDataTransportSize()){
                case BIT: 
                    //Check for bits length
                    buf.writeShort(payloadItem.getData().length);                    
                    break;
                case BYTE_WORD_DWORD: 
                case INTEGER:
                    //Length in bits
                    buf.writeShort(payloadItem.getData().length * 8);                      
                    break;
                case  NULL:
                    break;
                case DINTEGER:                    
                case OCTET_STRING:
                case REAL:
                    //Length in bytes
                    buf.writeShort(payloadItem.getData().length);                     
                    break;
            }           
            buf.writeBytes(payloadItem.getData());  

            // if this is not the last item and it's payload is exactly one byte, we need to output a fill-byte.
            //if((payloadItem.getData().length == 1) && (!lastItem)) {
          
            if((payloadItem.getData().length == 1) && (varPayload.getItems().size() > i) || (payloadItem.getData().length % 2 != 0)) {
                buf.writeByte(0x00); 
            }  
            
            i++;            
        }
    }
    
    private void encodeCpuServicesPayload(CpuServicesPayload cpuServicesPayload, ByteBuf buf)
            throws PlcProtocolException {

        buf.writeByte(cpuServicesPayload.getReturnCode().getCode());

        if (cpuServicesPayload.getReturnCode() == DataTransportErrorCode.NOT_FOUND) {
            buf.writeByte(DataTransportSize.NULL.getCode()); 
            buf.writeShort(0x0000);
        // This seems to be constantly set to this.
        } else {
            buf.writeByte(DataTransportSize.OCTET_STRING.getCode());

            // A request payload is simple.
            if (cpuServicesPayload.getSslDataRecords().isEmpty()) {
                buf.writeShort(4);
                buf.writeShort(cpuServicesPayload.getSslId().getCode());
                buf.writeShort(cpuServicesPayload.getSslIndex());
            }
            // The response payload contains a lot more information.
            else {
                throw new PlcProtocolException("Unexpected SZL Data Records");
                /*short length = 8;
                short sizeOfDataItem = 0;
                for (SslDataRecord sslDataRecord : cpuServicesPayload.getSslDataRecords()) {
                    sizeOfDataItem = (short) (sslDataRecord.getLengthInWords() * (short) 2);
                    length += sizeOfDataItem;
                }
                buf.writeShort(length);
                buf.writeShort(cpuServicesPayload.getSslId().getCode());
                buf.writeShort(cpuServicesPayload.getSslIndex());
                buf.writeShort(sizeOfDataItem);
                buf.writeShort(cpuServicesPayload.getSslDataRecords().size());
                // Output any sort of ssl list items, if there are any.
                for (SslDataRecord sslDataRecord : cpuServicesPayload.getSslDataRecords()) {
                    if(sslDataRecord instanceof SslModuleIdentificationDataRecord) {
                        SslModuleIdentificationDataRecord midr = (SslModuleIdentificationDataRecord) sslDataRecord;
                        buf.writeShort(midr.getIndex());
                        byte[] articleNumberBytes = midr.getArticleNumber().getBytes(StandardCharsets.UTF_8);
                        // An array full of 20 spaces.
                        byte[] data = new byte[]{0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
                            0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20};
                        // Copy max 20 bytes from the article number into the dest array.
                        System.arraycopy(articleNumberBytes, 0, data, 0, 20);
                        buf.writeBytes(data);
                        buf.writeShort(midr.getModuleOrOsVersion());
                        buf.writeShort(midr.getPgDescriptionFileVersion());
                    }
                }*/
            }
        }
    }    
    
    private void encodeCpuMessageSubcriptionPayload(CpuMessageSubscriptionServicePayload cpuServicesPayload, ByteBuf buf)
        throws PlcProtocolException {
        buf.writeByte(cpuServicesPayload.getReturnCode().getCode());
        buf.writeByte(cpuServicesPayload.getDataTransportSize().getCode());
        if ((cpuServicesPayload.getSubscribedEvents() & 0x80) == 0){
            buf.writeShort(0x000A);
        } else {
            buf.writeShort(0x000C);  
        };
        buf.writeByte(cpuServicesPayload.getSubscribedEvents());
        buf.writeByte(0x00);
        buf.writeBytes(cpuServicesPayload.getId().getBytes());
        if ((cpuServicesPayload.getSubscribedEvents() & 0x80) == 0x80){
            buf.writeByte(cpuServicesPayload.getAlarm().getCode());
            buf.writeByte(0x00);
        }
    }  
    
    //TODO: Agregar ALARM_ACK
    private void encodeAlarmMessagePayload(AlarmMessagePayload payload, ByteBuf buf) {

       AlarmMessageItem alarmitem = payload.getMsg();
       
       buf.writeByte(payload.getReturnCode().getCode());
       buf.writeByte(payload.getDataTransportSize().getCode());
       buf.writeShort(payload.getLength());
       
       //Response for fragmented PLC response
       //Check Plc4XS7Protocol. decode method
       if (alarmitem == null) { return; };
       List<Object> items = alarmitem.getMsgItems();
       
       buf.writeByte(alarmitem.getFunction());
       buf.writeByte(alarmitem.getObjects());
       
       for(Object thisitem:items){
           MessageObjectItem item = (MessageObjectItem) thisitem;
           switch (item.getSyntaxID()) {
                case S7ANY:;
                 logger.debug("encodeAlarmMessagePayload: S7ANY no supportesd");
                break;
                case PBC_R_ID:;
                 logger.debug("encodeAlarmMessagePayload: PBC_R_ID no supportesd");                
                break;                        
                case ALARM_LOCKFREE:;
                 logger.debug("encodeAlarmMessagePayload: ALARM_LOCKFREE no supportesd");                
                break;                        
                case ALARM_IND:;
                 logger.debug("encodeAlarmMessagePayload: ALARM no supportesd");                
                break;                        
                case ALARM_ACK:{
                   logger.debug("encodeAlarmMessagePayload: ALARM_ACK");
                   buf.writeByte(item.getVariableSpecification());
                   buf.writeByte(item.getLength());
                   buf.writeByte(item.getSyntaxID().getCode());
                   buf.writeByte(item.getNumberOfValues());
                   buf.writeInt(item.getEventID());
                   buf.writeByte(item.getAckStateGoing());
                   buf.writeByte(item.getAckStateComming());                    
                };
                break;                        
                case ALARM_QUERYREQ:{
                   logger.debug("encodeAlarmMessagePayload: ALARM_QUERYREQ");                    
                   buf.writeByte(item.getVariableSpecification());
                   buf.writeByte(item.getLength());
                   buf.writeByte(item.getSyntaxID().getCode());
                   buf.writeByte(0x00); //Unknown/Reserved
                   buf.writeByte(item.getQuerytype().getCode());
                   buf.writeByte(0x34); //Unknown/Reserved
                   buf.writeInt(item.getAlarmtype().getCode());                     
                };
                break;                        
                case NOTIFY_IND:;
                 logger.debug("encodeAlarmMessagePayload: NOTIFY_IND no supportesd");                
                break;                        
                case NIC:;
                 logger.debug("encodeAlarmMessagePayload: NIC no supportesd");                
                break;                        
                case DRIVEESANY:;
                 logger.debug("encodeAlarmMessagePayload: DRIVEESANY no supportesd");                
                break;                        
                case DBREAD:;
                 logger.debug("encodeAlarmMessagePayload: DBREAD no supportesd");                
                break;                        
                case SYM1200:;
                 logger.debug("encodeAlarmMessagePayload: SYM1200 no supportesd");                
                break;                        
               default:;               
           }
       }         
    }    

    private void encodeCpuCyclicSubscriptionPayload(CpuCyclicServicesRequestPayload payload, ByteBuf buf) {
        buf.writeByte(payload.getReturnCode().getCode());
        buf.writeByte(payload.getDataTransportSize().getCode());
        buf.writeShort(payload.getLength());
        //Unsuscription
        if (payload.getLength() == 2) {
            buf.writeByte(payload.getFunction());
            buf.writeByte(payload.getJobId());
        } else {
           //Suscription message
           buf.writeShort(payload.getItemCount());
           buf.writeByte(payload.getTimeBase().getCode());
           buf.writeByte(payload.getTimeFactor());
           payload.getItems().forEach((anyvar)->{
               encodeS7AnyPayloadItem(anyvar,buf);
           });     
       };
   }    
    
    private void encodeParameters(S7Message in, ByteBuf buf) throws PlcProtocolException {
        for (S7Parameter s7Parameter : in.getParameters()) {
            buf.writeByte(s7Parameter.getType().getCode());
            switch (s7Parameter.getType()) {
                case READ_VAR:
                case WRITE_VAR:
                    encodeParameterReadWriteVar(buf, (VarParameter) s7Parameter);
                    break;
                case SETUP_COMMUNICATION:
                    encodeParameterSetupCommunication(buf, (SetupCommunicationParameter) s7Parameter);
                    break;
                case CPU_SERVICES:
                    encodeCpuServicesParameter(buf, (CpuServicesParameter) s7Parameter);
                    break;
                default:
                    throw new PlcProtocolException("Writing parameters of type " +
                        s7Parameter.getType().name() + " not implemented.");
            }
        }
    }

    private void encodeHeader(S7Message in, ByteBuf buf) {
        buf.writeByte(S7_PROTOCOL_MAGIC_NUMBER);
        buf.writeByte(in.getMessageType().getCode());
        // Reserved (is always constant 0x0000)
        buf.writeShort((short) 0x0000);
        // PDU Reference (Request Id, generated by the initiating node)
        buf.writeShort(in.getTpduReference());
        // S7 message parameters length
        buf.writeShort(S7SizeHelper.getParametersLength(in.getParameters()));
        // Data field length
        buf.writeShort(S7SizeHelper.getPayloadsLength(in.getPayloads()));
        // Not sure why this is implemented, we should never be sending out responses.
        /*if (in instanceof S7ResponseMessage) {
            S7ResponseMessage s7ResponseMessage = (S7ResponseMessage) in;
            buf.writeByte(s7ResponseMessage.getErrorClass());
            buf.writeByte(s7ResponseMessage.getErrorCode());
        }*/
    }

    private void encodeParameterSetupCommunication(ByteBuf buf, SetupCommunicationParameter s7Parameter) {
        // Reserved (is always constant 0x00)
        buf.writeByte((byte) 0x00);
        buf.writeShort(s7Parameter.getMaxAmqCaller());
        buf.writeShort(s7Parameter.getMaxAmqCallee());
        buf.writeShort(s7Parameter.getPduLength());        
    }

    private void encodeParameterReadWriteVar(ByteBuf buf, VarParameter s7Parameter) throws PlcProtocolException {
        List<VarParameterItem> items = s7Parameter.getItems();
        // PlcReadRequestItem count (Read one variable at a time)
        buf.writeByte((byte) items.size());
        for (VarParameterItem item : items) {
            VariableAddressingMode addressMode = item.getAddressingMode();
            if (addressMode == VariableAddressingMode.S7ANY) {
                encodeS7AnyParameterItem(buf, (S7AnyVarParameterItem) item);
            } else {
                throw new PlcProtocolException("Writing VarParameterItems with addressing mode " +
                    addressMode.name() + " not implemented");
            }
        }
    }

    private void encodeCpuServicesParameter(ByteBuf buf, CpuServicesParameter parameter) {
        // Output the header for a CPU Services parameter.
        buf.writeByte(0x01);
        buf.writeByte(0x12);
        // Length of the parameter.
        
        buf.writeByte((parameter instanceof CpuServicesRequestParameter) ? 0x04 : 0x08);
        // Is this a request or a response?
        buf.writeByte((parameter instanceof CpuServicesRequestParameter) ? 0x11 : 0x12);
        // This is a mixture of request/response and function group .
        byte nextByte = (byte) (((parameter instanceof CpuServicesRequestParameter) ?
            (byte) 0x40 : (byte) 0x80) | parameter.getFunctionGroup().getCode());
        
        //For fragmented response
        if ((parameter instanceof CpuServicesResponseParameter) && 
                (parameter.getSequenceNumber() != 0)){
            nextByte = (byte) (0x40 + parameter.getFunctionGroup().getCode());
        }
        
        buf.writeByte(nextByte);
        if (parameter instanceof CpuCyclicServicesRequestParameter) {
            buf.writeByte(((CpuCyclicServicesRequestParameter) parameter).getCyclicSubFunction().getCode());
        } else {
            buf.writeByte(parameter.getSubFunctionGroup().getCode());
        }
        buf.writeByte(parameter.getSequenceNumber());

        // A response parameter has some more fields.
        // Not sure why this is implemented, we should never be sending out responses.
        /*if(parameter instanceof CpuServicesResponseParameter) {
            CpuServicesResponseParameter responseParameter = (CpuServicesResponseParameter) parameter;
            buf.writeByte(responseParameter.getDataUnitReferenceNumber());
            buf.writeByte(responseParameter.isLastDataUnit() ? 0x00 : 0x01);
            buf.writeShort(responseParameter.getError().getCode());
        }*/
        //R: if parameter.getSequenceNumber() == 2 then is fragment message query
        //   from Plc4XS7Protocol decode method.
        if (parameter.getSequenceNumber() != 0x00) {
            buf.writeByte(0x00); //DataUnitReferenceNumber
            buf.writeByte(0x00); //LastDataUnit
            buf.writeShort(0x0000); //Error code
        }
        
    }

    //TODO: Refactor S7AnyVarParameterItem to S7AnyVarItem
    private void encodeS7AnyParameterItem(ByteBuf buf, S7AnyVarParameterItem s7AnyRequestItem) {
        buf.writeByte(s7AnyRequestItem.getSpecificationType().getCode());
        // Length of this item (excluding spec type and length)
        buf.writeByte((byte) 0x0a);
        buf.writeByte(s7AnyRequestItem.getAddressingMode().getCode());
        //buf.writeByte(s7AnyRequestItem.getDataType().getTypeCode());
        
        buf.writeByte(s7AnyRequestItem.getDataType().getBaseType().getTypeCode());

        buf.writeShort(encodeNumElements(s7AnyRequestItem));
        buf.writeShort(s7AnyRequestItem.getDataBlockNumber());
        buf.writeByte(s7AnyRequestItem.getMemoryArea().getCode());
        // A S7 address is 3 bytes long. Unfortunately the byte-offset is NOT located in
        // byte 1 and byte 2 and the bit offset in byte 3. Siemens used the last 3 bits of
        // byte 3 for the bit-offset and the remaining 5 bits of byte 3 to contain the lowest
        // 5 bits of the byte-offset. The highest 5 bits of byte 1 are probably left unused
        // for future extensions.
        buf.writeShort((short) (s7AnyRequestItem.getByteOffset() >> 5));
        buf.writeByte((byte) ((
                (s7AnyRequestItem.getByteOffset() & 0x1F) << 3)
                | (s7AnyRequestItem.getBitOffset() & 0x07)));
    }
    
    //TODO: Refactor S7AnyVarPayloadItem to S7AnyVarItem
    private void encodeS7AnyPayloadItem(S7AnyVarPayloadItem s7AnyRequestItem, ByteBuf buf) {
        buf.writeByte(s7AnyRequestItem.getSpecificationType().getCode());
        // Length of this item (excluding spec type and length)
        buf.writeByte((byte) 0x0a);
        buf.writeByte(s7AnyRequestItem.getAddressingMode().getCode());
        //buf.writeByte(s7AnyRequestItem.getDataType().getTypeCode());
        buf.writeByte(s7AnyRequestItem.getDataType().getBaseType().getTypeCode());

        buf.writeShort(encodeNumElements(s7AnyRequestItem));
        buf.writeShort(s7AnyRequestItem.getDataBlockNumber());
        buf.writeByte(s7AnyRequestItem.getMemoryArea().getCode());
        // A S7 address is 3 bytes long. Unfortunately the byte-offset is NOT located in
        // byte 1 and byte 2 and the bit offset in byte 3. Siemens used the last 3 bits of
        // byte 3 for the bit-offset and the remaining 5 bits of byte 3 to contain the lowest
        // 5 bits of the byte-offset. The highest 5 bits of byte 1 are probably left unused
        // for future extensions.
        buf.writeShort((short) (s7AnyRequestItem.getByteOffset() >> 5));
        buf.writeByte((byte) ((
                (s7AnyRequestItem.getByteOffset() & 0x1F) << 3)
                | (s7AnyRequestItem.getBitOffset() & 0x07)));
    }
        
    /**
     * this is a workaround for the date and time types, as native requests with the datatypes are
     * @return
     */
        //TODO: Refactor S7AnyVarParameterItem to S7AnyVarItem
    private short encodeNumElements(S7AnyVarParameterItem s7AnyVarParameterItem){
        switch (s7AnyVarParameterItem.getDataType()){
            case DT:
            case DATE_AND_TIME:
            case TOD:
            case TIME_OF_DAY:
            case DATE:
                return (short) (s7AnyVarParameterItem.getNumElements()*s7AnyVarParameterItem.getDataType().getSizeInBytes());
            default:
                return (short) s7AnyVarParameterItem.getNumElements();
        }

    }    
    
    /**
     * this is a workaround for the date and time types, as native requests with the datatypes are
     * @return
     */
    //TODO: Refactor S7AnyVarPayloadItem to S7AnyVarItem
    private short encodeNumElements(S7AnyVarPayloadItem s7AnyVarParameterItem){
        switch (s7AnyVarParameterItem.getDataType()){
            case DATE_AND_TIME:
            case TOD:
            case TIME_OF_DAY:
            case DATE:
            case DT:
                return (short) (s7AnyVarParameterItem.getNumElements()*s7AnyVarParameterItem.getDataType().getSizeInBytes());               
            case STRING:
               break; 
            default:
                return (short) s7AnyVarParameterItem.getNumElements();
        }
        return 0;
    }  

    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Decoding
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        decoder.channelRead(ctx, msg);
        super.channelRead(ctx, msg);
    }

    protected void decode(ChannelHandlerContext ctx, IsoTPMessage in, List<Object> out) {
        logger.info("DECODE: " + ByteBufUtil.prettyHexDump(in.getUserData()));
        if (logger.isTraceEnabled()) {
            logger.trace("Got Data: {}", ByteBufUtil.hexDump(in.getUserData()));
        }
        ByteBuf userData = in.getUserData();
        if (userData.readableBytes() == 0) {
            return;
        }
                      
        if (userData.readByte() != S7_PROTOCOL_MAGIC_NUMBER) {
            logger.warn("Expecting S7 protocol magic number.");
            if (logger.isDebugEnabled()) {
                logger.debug("Got Data: {}", ByteBufUtil.hexDump(userData));
            }
            return;
        }
                
        MessageType messageType = MessageType.valueOf(userData.readByte());
        //Some response maybe USER_DATA
        //Header: Userdata. Parameter:(Response)->(CPU functions)->(ALARM ack)
           
        //The driver freezes if it does not recognize 
        //the MessageType.ACK type in response
        boolean isResponse = (messageType == MessageType.ACK_DATA) || (messageType == MessageType.ACK);        
        userData.readShort();  // Reserved (is always constant 0x0000)
        short tpduReference = userData.readShort();
        short headerParametersLength = userData.readShort();
        short userDataLength = userData.readShort();
        byte errorClass = 0;
        byte errorCode = 0;

        //If the PLC response qith error then messagetype == MessageType.ACK
        if (isResponse) {            
            errorClass = userData.readByte();
            errorCode = userData.readByte();
        }

        List<S7Parameter> s7Parameters = new LinkedList<>();
        int i = 0;

        while (i < headerParametersLength) {
            S7Parameter parameter = decodeParameter(userData, isResponse);
            s7Parameters.add(parameter);
            if (parameter instanceof SetupCommunicationParameter) {
                handleSetupCommunications(ctx, (SetupCommunicationParameter) parameter);
            }
            i += S7SizeHelper.getParameterLength(parameter);
        }
        
        List<S7Payload> s7Payloads = decodePayloads(userData, isResponse, userDataLength, s7Parameters);
        logger.debug("S7 Message with id {} received", tpduReference);
        
        //Case: USER_DATA CpuCyclicServicesResponseParameter
        if (messageType == MessageType.USER_DATA){
            if (s7Parameters.get(0) instanceof CpuCyclicServicesResponseParameter){
                CpuCyclicServicesResponseParameter parameter = (CpuCyclicServicesResponseParameter) s7Parameters.get(0);
                if (parameter.getSubCycFunctionGroup() == CpuCyclicServicesParameterSubFunctionGroupType.CYCLIC_UNSUBSCRIBE){
                    isResponse = true;
                }
            }
            
            //TODO: Check for other combinations
            if (s7Parameters.get(0) instanceof CpuServicesResponseParameter){                
                CpuServicesResponseParameter parameter = (CpuServicesResponseParameter) s7Parameters.get(0);
                if (parameter.getSubFunctionGroup() == CpuServicesParameterSubFunctionGroup.ALARM_ACK){
                    isResponse = true;
                }
            }
        };

        if (isResponse) {
            S7ResponseMessage responseMessage = new S7ResponseMessage(
                messageType, tpduReference, s7Parameters, s7Payloads, errorClass, errorCode);

            // Remove the current response from the list of unconfirmed messages.
            DataTpdu requestTpdu = sentButUnacknowledgedTpdus.remove(tpduReference);
            logger.debug("decode: S7 Message with id {} remove", tpduReference);

            // Get the corresponding request message.
            S7RequestMessage requestMessage = (requestTpdu != null) ? (S7RequestMessage) requestTpdu.getParent() : null;

            if(requestMessage != null) {
                // Set this individual request to "acknowledged".
                requestMessage.setAcknowledged(true);

                // Give the request and response to a message processor to process the incoming message.
                if(messageProcessor != null) {
                    try {
                        responseMessage = messageProcessor.processResponse(requestMessage, responseMessage);
                    } catch(Exception e) {
                        logger.error("Error processing message", e);
                        ctx.fireExceptionCaught(e);
                        return;
                    }
                }

                if(responseMessage != null) {
                    out.add(responseMessage);
                    // If this is a USER_DATA packet the probability is high that this is
                    // a response to the identification request, we have to handle that.
                    if(responseMessage.getMessageType() == MessageType.USER_DATA) {
                        for (S7Payload payload : responseMessage.getPayloads()) {
                            if(payload instanceof CpuServicesPayload) {
                                handleIdentifyRemote(ctx, (CpuServicesPayload) payload);
                            }
                        }
                    }
                }

                // Try to send the next message (if there is one).
                logger.info("DECODE: Try to send the next message (if there is one).");
                trySendingMessages(ctx);
            }

        } else {
            // !CpuService responses are encoded as requests. 
            // No!, we need check in the next layer Plc4XS7Protocol like response
            for (S7Parameter s7Parameter : s7Parameters) {
                // Only if we have a response parameter, the payload is a response payload.
                if(s7Parameter instanceof CpuServicesResponseParameter) {

                    for (S7Payload s7Payload : s7Payloads) {
                        if((s7Payload instanceof CpuServicesPayload)) {
                            CpuServicesPayload cpuServicesPayload = (CpuServicesPayload) s7Payload;
                            // Remove the current response from the list of unconfirmed messages.
                            sentButUnacknowledgedTpdus.remove(tpduReference);
                            logger.debug("CpuServicesPayload: S7 Message with id {} remove", tpduReference);  
   
                            handleIdentifyRemote(ctx, cpuServicesPayload);
 
                        } else {
                            sentButUnacknowledgedTpdus.remove(tpduReference);
                            logger.debug("CpuCyclicServicesResponsePayload: S7 Message with id {} remove", tpduReference); 
                        }
                    }
                } else if (s7Parameter instanceof CpuServicesPushParameter){
                    //*** MENSAGE PUSH ***"
                    //out.add(new S7ResponseMessage(messageType, tpduReference, s7Parameters, s7Payloads, errorClass, errorCode)); 
                }
            }            
            //out.add(new S7RequestMessage(messageType, tpduReference, s7Parameters, s7Payloads, null));
            out.add(new S7ResponseMessage(messageType, tpduReference, s7Parameters, s7Payloads, errorClass, errorCode)); 
        }
    }

    private void handleSetupCommunications(ChannelHandlerContext ctx, SetupCommunicationParameter setupCommunicationParameter) {
        maxAmqCaller = setupCommunicationParameter.getMaxAmqCaller();
        maxAmqCallee = setupCommunicationParameter.getMaxAmqCallee();
        pduSize = setupCommunicationParameter.getPduLength();
        
        ctx.channel().attr(pduSizeKey).set(pduSize);

        if(logger.isInfoEnabled()) {
            logger.info("S7Connection established pdu-size {}, max-amq-caller {}, " +
                "max-amq-callee {}", pduSize, maxAmqCaller, maxAmqCallee);
        }

        // Only if the controller type is set to "ANY", then try to identify the PLC type.
        if(controllerType == S7ControllerType.ANY) {
            // Prepare a message to request the remote to identify itself.
            S7RequestMessage identifyRemoteMessage = new S7RequestMessage(MessageType.USER_DATA, (short) 2,
                Collections.singletonList(new CpuServicesRequestParameter(
                    CpuServicesParameterFunctionGroup.CPU_FUNCTIONS,
                    CpuServicesParameterSubFunctionGroup.READ_SSL, (byte) 0)),
                Collections.singletonList(new CpuServicesPayload(DataTransportErrorCode.OK, SslId.MODULE_IDENTIFICATION,
                    (short) 0x0000)), null);
            ctx.channel().writeAndFlush(identifyRemoteMessage);
        }
        // If a concrete type was specified, then we're done here.
        else {
            if(logger.isDebugEnabled()) {
                logger.debug("Successfully connected to S7: {}", controllerType.name());
                logger.debug("- max amq caller: {}", maxAmqCaller);
                logger.debug("- max amq callee: {}", maxAmqCallee);
                logger.debug("- pdu size: {}", pduSize);
            }
            if(logger.isInfoEnabled()) {
                logger.info("Successfully connected to S7: {} wit PDU {}", controllerType.name(),pduSize);
            }

            // Send an event that connection setup is complete.
            ctx.channel().pipeline().fireUserEventTriggered(new S7ConnectedEvent());
        }
    }

    private void handleIdentifyRemote(ChannelHandlerContext ctx, CpuServicesPayload cpuServicesPayload) {
        controllerType = S7ControllerType.ANY;
        /*
        for (SslDataRecord sslDataRecord : cpuServicesPayload.getSslDataRecords()) {
            if(sslDataRecord instanceof SslModuleIdentificationDataRecord) {
                SslModuleIdentificationDataRecord sslModuleIdentificationDataRecord =
                    (SslModuleIdentificationDataRecord) sslDataRecord;
                if(sslModuleIdentificationDataRecord.getIndex() == (short) 0x0001) {
                    controllerType = lookupControllerType(sslModuleIdentificationDataRecord.getArticleNumber());
                }
            }
        */
        if (cpuServicesPayload.getSslId() == SslId.MODULE_IDENTIFICATION) {
            if ((cpuServicesPayload.getSslIndex() == 0x0000) || (cpuServicesPayload.getSslIndex() == 0x0001)){
                ByteBuf payload = cpuServicesPayload.getSslPayload();
                CharSequence MlfB = payload.getCharSequence(6, 20, Charset.defaultCharset());
                controllerType = lookupControllerType(MlfB.toString());
            }
        }
        if(logger.isDebugEnabled()) {
            logger.debug("Successfully connected to S7: {}", controllerType.name());
            logger.debug("-  max amq caller: {}", maxAmqCaller);
            logger.debug("-  max amq callee: {}", maxAmqCallee);
            logger.debug("-  pdu size: {}", pduSize);
        }
        if(logger.isInfoEnabled()) {
            logger.info("Successfully connected to S7: {} wit PDU {}", controllerType.name(), pduSize);
        }

        // Send an event that connection setup is complete.
        ctx.channel().pipeline().fireUserEventTriggered(new S7ConnectedEvent());
    }

    private List<S7Payload> decodePayloads(ByteBuf userData, boolean isResponse, short userDataLength, List<S7Parameter> s7Parameters) {
        List<S7Payload> s7Payloads = new LinkedList<>();
        for (S7Parameter s7Parameter : s7Parameters) {
            if(s7Parameter instanceof VarParameter) {
                VarParameter readWriteVarParameter = (VarParameter) s7Parameter;
                VarPayload varPayload = decodeVarPayload(userData, isResponse, userDataLength, readWriteVarParameter);
                s7Payloads.add(varPayload);
            } else if(s7Parameter instanceof CpuServicesParameter) {
                S7Payload cpuServicesPayload = decodeCpuServicesPayload((CpuServicesParameter)s7Parameter, userData);
                s7Payloads.add(cpuServicesPayload);
            } else if (s7Parameter instanceof CpuServicesPushParameter){
                S7Payload cpuServicesPayload = decodeCpuServicesPayload((CpuServicesParameter)s7Parameter, userData);
                s7Payloads.add(cpuServicesPayload); 
            }
        }
        return s7Payloads;
    }

    private VarPayload decodeVarPayload(ByteBuf userData, boolean isResponse, short userDataLength,
                                        VarParameter readWriteVarParameter) {
        //System.out.println("decodeVarPayload:\r\n" + ByteBufUtil.prettyHexDump(userData));
        List<VarPayloadItem> payloadItems = new LinkedList<>();

        // Just keep on reading payloads until the provided length is read.
        int i = 0;
        while (i < userDataLength) {
            DataTransportErrorCode dataTransportErrorCode = DataTransportErrorCode.valueOf(userData.readByte());

            // This is a response to a WRITE_VAR request (It only contains the return code for every sent item.
            if ((readWriteVarParameter.getType() == ParameterType.WRITE_VAR) && isResponse) {
                // Initialize a rudimentary payload (This is updated in the Plc4XS7Protocol class
                VarPayloadItem payload = new VarPayloadItem(dataTransportErrorCode, null, null);
                payloadItems.add(payload);
                i += 1;
            }
            // This is a response to a READ_VAR request.
            else if ((readWriteVarParameter.getType() == ParameterType.READ_VAR) && isResponse) {
                DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
                short length = dataTransportSize.isSizeInBits() ?
                    (short) Math.ceil(userData.readShort() / 8.0) : userData.readShort();
                byte[] data = new byte[length];
                userData.readBytes(data);
                // Initialize a rudimentary payload (This is updated in the Plc4XS7Protocol class
                VarPayloadItem payload = new VarPayloadItem(dataTransportErrorCode, dataTransportSize, data);
                payloadItems.add(payload);
                i += S7SizeHelper.getPayloadLength(payload);

                // It seems that odd-byte payloads require a fill byte, but only if it's not the last item.
                if((length % 2== 1) && (userData.readableBytes() > 0)) {
                    userData.readByte();
                    i++;
                }
            }
        }

        return new VarPayload(readWriteVarParameter.getType(), payloadItems);
    }

    private S7Payload decodeCpuServicesPayload(CpuServicesParameter parameter, ByteBuf userData) {
        if (parameter.getFunctionGroup() == CpuServicesParameterFunctionGroup.CPU_FUNCTIONS) {
            switch(parameter.getSubFunctionGroup()){
                case READ_SSL: {
                    CpuServicesPayload payload = decodeReadSslPayload(parameter, userData);
                    return payload;
                }
                case MESSAGE_SERVICE:{  
                    AlarmMessagePayload payload = decodeMessageServicePayload(parameter, userData); 
                    return payload;            
                }
                case DIAG_MESSAGE:{
                    CpuDiagnosticMessagePayload payload = decodeCpuDiagnosticMessagePayload(parameter, userData);
                    return payload;
                }
                case ALARM8:;
                    logger.info("decode ALARM8.....");
                    
                    break;
                case NOTIFY:;
                    logger.info("decode NOTIFY.....");
                    break;
                case ALARM8_LOCK:;
                    logger.info("decode ALARM8_LOCK.....");
                    break;
                case ALARM8_UNLOCK:;
                    logger.info("decode ALARM8_UNLOCK.....");
                    break;
                case SCAN:;
                    logger.info("decode SCAN.....");
                    break;
                case ALARM_ACK:{  
                    AlarmMessagePayload payload = decodeMessageServiceAckPayload(parameter, userData);                
                    return payload;            
                }
                case ALARM_ACK_IND:{
                    AlarmMessagePayload payload = decodeMessageServicePushPayload(parameter, userData);
                    return payload;
                }
                case ALARM8_LOCK_IND:;
                    logger.info("decode ALARM8_LOCK_IND.....");
                    break;
                case ALARM8_UNLOCK_IND:;
                    logger.info("decode ALARM8_UNLOCK_IND.....");
                    break;
                case ALARM_SQ_IND:{  
                    AlarmMessagePayload payload = decodeMessageServicePushPayload(parameter, userData);                
                    return payload;            
                }
                case ALARM_S_IND: {
                    AlarmMessagePayload payload = decodeMessageServicePushPayload(parameter, userData);
                    return payload;
                }
                case ALARM_QUERY:{
                    AlarmMessagePayload payload = decodeMessageServiceQueryPayload(parameter, userData);
                    return payload;
                }
                case NOTIFY8: {
                    AlarmMessagePayload payload = decodeMessageServicePushPayload(parameter, userData);
                    return payload;
                }
                case ALARM: {
                    AlarmMessagePayload payload = decodeMessageServicePushPayload(parameter, userData);
                    return payload;
                }                
                default:;
                    break;
            } 
            
        } else if (parameter.getFunctionGroup() == CpuServicesParameterFunctionGroup.CYCLIC_SERVICES) {
            if (parameter instanceof CpuCyclicServicesResponseParameter){
                CpuCyclicServicesResponseParameter pushparameter = (CpuCyclicServicesResponseParameter) parameter;
                switch(pushparameter.getSubCycFunctionGroup()){
                    case CYCLIC_CHANGE:;
                        break;
                    case CYCLIC_CHANGE_MODIFY:;
                        break;
                    case CYCLIC_RDREC:;    
                        break;
                    case CYCLIC_TRANSFER: {
                        CpuCyclicServicesResponsePayload payload = decodeCyclicServiceResponsePayload(pushparameter, userData);
                        return payload;
                    }
                    case CYCLIC_UNSUBSCRIBE:{   
                        CpuCyclicServicesResponsePayload payload = decodeCyclicServiceResponsePayload(pushparameter, userData);
                        return payload;
                    }
                    default: logger.info("decodeCpuServicesPayload: No Sub Cyc Function Group.");
                }
            }
        }
        return null;
    }

    private S7Parameter decodeParameter(ByteBuf in, boolean isResponse) {
        ParameterType parameterType = ParameterType.valueOf(in.readByte());
        if (parameterType == null) {
            logger.error("Could not find parameter type");
            return null;
        }
        switch (parameterType) {
            case CPU_SERVICES:
                return decodeCpuServicesParameter(in);
            case MODE_TRANSITION:
                return decodePushModeTransitionParameter(in);
            case READ_VAR:
            case WRITE_VAR:
                List<VarParameterItem> varParameterItems;
                byte numItems = in.readByte();
                if (!isResponse) {
                    varParameterItems = decodeReadWriteVarParameter(in, numItems);
                } else {
                    varParameterItems = Collections.singletonList(
                        new S7AnyVarParameterItem(null, null, null, numItems, (short) 0, (short) 0, (byte) 0));
                }
                return new VarParameter(parameterType, varParameterItems);
            case SETUP_COMMUNICATION:
                // Reserved (is always constant 0x00)
                in.readByte();
                short callingMaxAmq = in.readShort();
                short calledMaxAmq = in.readShort();
                short pduLength = in.readShort();
                return new SetupCommunicationParameter(callingMaxAmq, calledMaxAmq, pduLength);
            default:
                if (logger.isErrorEnabled()) {
                    logger.error("Unimplemented parameter type: {}", parameterType.name());
                }
        }
        return null;
    }

    private CpuServicesParameter decodeCpuServicesParameter(ByteBuf in) {
        //logger.info("decodeParameter... \r\n" + ByteBufUtil.prettyHexDump(in));        
        if(in.readShort() != 0x0112) {
            if (logger.isErrorEnabled()) {
                logger.error("Expecting 0x0112 for CPU_SERVICES parameter");
            }
            return null;
        }
        
        byte parameterLength = in.readByte();
        if((parameterLength != 4) && (parameterLength != 8)) {
            if (logger.isErrorEnabled()) {
                logger.error("Parameter length should be 4 or 8, but was {}", parameterLength);
            }
            return null;
        }
        // Skipping this as it sort of contains redundant information.
        // TODO: We need check the next byte for push message (Request/Response)
        in.readByte();
        byte typeAndFunctionGroup = in.readByte();
        // If bit 7 is set, it's a request (if bit 8 is set it's a response).
        //Must be for request: 0x40 for check X100 0X00
        //For Push message is 0000 XXXX, Request is 0100 XXXX, Response 1000 XXXX       
        boolean pushParameter = (typeAndFunctionGroup & 0xF0) == 0;
        boolean requestParameter = (typeAndFunctionGroup & 0x40) != 0;
        boolean responseParameter = (typeAndFunctionGroup & 0x80) != 0;
        
        // The last 4 bits contain the function group value.
        typeAndFunctionGroup = (byte) (typeAndFunctionGroup & 0x0F);
        
        CpuServicesParameterFunctionGroup functionGroup =
            CpuServicesParameterFunctionGroup.valueOf(typeAndFunctionGroup);

        CpuServicesParameterSubFunctionGroup subFunctionGroup =
            CpuServicesParameterSubFunctionGroup.valueOf(in.readByte());
        
        byte sequenceNumber = in.readByte();
        
        if(pushParameter || (responseParameter && functionGroup == CpuServicesParameterFunctionGroup.CYCLIC_SERVICES)) {
            if (functionGroup == CpuServicesParameterFunctionGroup.CYCLIC_SERVICES ) {
               CpuCyclicServicesParameterSubFunctionGroupType subCycFunctionGroup = 
                       CpuCyclicServicesParameterSubFunctionGroupType.valueOf(subFunctionGroup.getCode());
               byte dataUnitReferenceNumber = in.readByte();
               boolean lastDataUnit = in.readByte() == 0x00;
               ParameterError error = ParameterError.valueOf(in.readShort());
               
               return new CpuCyclicServicesResponseParameter(functionGroup, subCycFunctionGroup, sequenceNumber,
                            dataUnitReferenceNumber, lastDataUnit, error);               
            } else {
                return new CpuServicesPushParameter(functionGroup, subFunctionGroup, sequenceNumber);            
            }
        } else if (requestParameter) {
            return new CpuServicesRequestParameter(functionGroup, subFunctionGroup, sequenceNumber);            
        } else {
            byte dataUnitReferenceNumber = in.readByte();
            boolean lastDataUnit = in.readByte() == 0x00;
            ParameterError error = ParameterError.valueOf(in.readShort());

            return new CpuServicesResponseParameter(functionGroup, subFunctionGroup, sequenceNumber,
                dataUnitReferenceNumber, lastDataUnit, error);
        }
    }

    private CpuDiagnosticPushParameter decodePushModeTransitionParameter(ByteBuf in) {

        if(in.readShort() != 0x0010) {
            if (logger.isErrorEnabled()) {
                logger.error("Expecting 0x0010 for MODE_TRANSITION parameter");
            }
            return null;
        }
        
        byte parameterLength = in.readByte();
        if((parameterLength != 16)) {
            if (logger.isErrorEnabled()) {
                logger.error("Parameter length should be 16, but was {}", parameterLength);
            }
            return null;
        }   
        CpuUserDataMethodType usermethodtype = CpuUserDataMethodType.valueOf(in.readByte());
        byte typeandfunc = in.readByte();
        CpuUserDataParameterType userparamtype = CpuUserDataParameterType.valueOf((byte)(typeandfunc >> 4));
        CpuUserDataParameterFunctionGroupType userfunction =  CpuUserDataParameterFunctionGroupType.valueOf((byte)(typeandfunc & 0x0f));
        CpuCurrentModeType cpumode = CpuCurrentModeType.valueOf(in.readByte());
        byte sequencenumber = in.readByte();
        return new CpuDiagnosticPushParameter(usermethodtype, userparamtype, userfunction, cpumode, sequencenumber);
    }        
    
    private List<VarParameterItem> decodeReadWriteVarParameter(ByteBuf in, byte numItems) {
        List<VarParameterItem> items = new LinkedList<>();
        for (int i = 0; i < numItems; i++) {
            SpecificationType specificationType = SpecificationType.valueOf(in.readByte());
            // Length of the rest of this item.
            byte itemLength = in.readByte();
            if (itemLength != 0x0a) {
                logger.warn("Expecting a length of 10 here.");
                return items;
            }
            VariableAddressingMode variableAddressingMode = VariableAddressingMode.valueOf(in.readByte());
            if (variableAddressingMode == VariableAddressingMode.S7ANY) {
                TransportSize dataType = TransportSize.valueOf(in.readByte());
                short length = in.readShort();
                short dbNumber = in.readShort();
                byte memoryAreaCode = in.readByte();
                MemoryArea memoryArea = MemoryArea.valueOf(memoryAreaCode);
                if(memoryArea == null) {
                    throw new PlcRuntimeException("Unknown memory area '" + memoryAreaCode + "'");
                }
                short byteAddress = (short) (in.readShort() << 5);
                byte tmp = in.readByte();
                // Only the least 3 bits are the bit address, the
                byte bitAddress = (byte) (tmp & 0x07);
                // Bits 4-8 belong to the byte address
                byteAddress = (short) (byteAddress | (tmp >> 3));
                S7AnyVarParameterItem item = new S7AnyVarParameterItem(
                        specificationType, memoryArea, dataType,
                        length, dbNumber, byteAddress, bitAddress);
                items.add(item);
            } else {
                logger.error("Error parsing item type");
                return items;
            }
        }

        return items;
    }
    
    private CpuServicesPayload decodeReadSslPayload(CpuServicesParameter parameter, ByteBuf userData){ 
        //logger.info("decodeReadSslPayload: \r\n" + ByteBufUtil.prettyHexDump(userData));
        CpuServicesResponseParameter thisparameter = null;
        short length;
        short dataUnitReferenceNumber = 0x0000;
        boolean lastdataunit = true;
        
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());

        
        if (parameter instanceof CpuServicesResponseParameter){           
            thisparameter = (CpuServicesResponseParameter) parameter;
            dataUnitReferenceNumber = (short) thisparameter.getDataUnitReferenceNumber();
            lastdataunit = thisparameter.isLastDataUnit();
        }        
        /*
        if(dataTransportSize != DataTransportSize.OCTET_STRING) {
            if(dataTransportSize == DataTransportSize.NULL) {
                length = userData.readShort();
                if (length == 0x0000) {
                    return new CpuServicesPayload(returnCode, SslId.NULL, (short) 0x0000);
                } else {
                // TODO: Output an error.                    
                }
                
            }   
        }
          */      
        length = userData.readShort();         
        //If is fragment message 
        SslId sslId = SslId.NULL;
        short sslIndex = 0x0000;
        
        
        //Check for service not found
        if (returnCode == DataTransportErrorCode.NOT_FOUND && !lastdataunit){
            userData.clear();           
            if (thisparameter != null){               
                userData.writeShort(thisparameter.getError().getCode());
            } else {
                userData.writeShort(ParameterError.PROTOCOL_ERROR.getCode());
            }
           return new CpuServicesPayload(returnCode, sslId, sslIndex, userData);            
        }        
        
        // If the length is 4 there is no `partial list length in bytes` and `partial list count` parameters.
        if(length == 4) {
            //logger.info("decodeReadSslPayload: error response.");
            sslId = SslId.valueOf(userData.readShort());
            sslIndex = userData.readShort();            
            return new CpuServicesPayload(returnCode, sslId, sslIndex);
        }
        // If the length is not 4, then it has to be at least 8.
        else if((length >= 8) || ((length == 0) && lastdataunit)) {                                           
            if (!lastdataunit) {   
                //TODO: Work with two fragments, we need test with 3 and more
                Pair<LocalDateTime, Queue<ByteBuf>> fragments = fragmentedData.get(dataUnitReferenceNumber);
                if (fragments != null) {
                    Queue<ByteBuf> bytebufqueue = fragments.getValue();
                    bytebufqueue.add(userData);
                } else {
                    sslId = SslId.valueOf(userData.readShort());
                    sslIndex = userData.readShort();                       
                    Queue<ByteBuf> bytebufqueue = new ArrayDeque();
                    bytebufqueue.add(userData);
                    Pair<LocalDateTime, Queue<ByteBuf>> firtsfragment = new ImmutablePair(LocalDateTime.now(), bytebufqueue);
                    fragmentedData.put(dataUnitReferenceNumber, firtsfragment);
                } 
                return new CpuServicesPayload(returnCode, sslId, sslIndex);   
            }
            
            //Reemsable the fragments   
            
            if (dataUnitReferenceNumber != 0x0000){
                Pair<LocalDateTime, Queue<ByteBuf>> fragments = fragmentedData.get(dataUnitReferenceNumber);
                if (fragments != null){ 
                    Queue<ByteBuf> bytebufqueue = fragments.getValue();
                    ByteBuf payload = bytebufqueue.remove();
                    for(ByteBuf nextpayload:bytebufqueue){
                        if (nextpayload != null) { 
                            payload.writeBytes(nextpayload);
                            nextpayload.release();                          
                        };
                    }               
                    payload.writeBytes(userData);
                    userData.clear();                 
                    userData.writeBytes(payload);                
                    payload.release();               
                    fragmentedData.remove(dataUnitReferenceNumber);
                    returnCode = DataTransportErrorCode.OK;
                }                  
            } else {
                sslId = SslId.valueOf(userData.readShort());
                sslIndex = userData.readShort();                  
            }

            //TODO: sslId & sslIndex in this point dont care, we need check!
            return new CpuServicesPayload(returnCode, sslId, sslIndex, userData);
            
        }
        // In all other cases, it's probably an error.
        else {
            // TODO: Output an error.
        }   

        return null;   
    }    
    
    private AlarmMessagePayload decodeMessageServicePayload(CpuServicesParameter parameter, ByteBuf userData){
        //logger.info(ByteBufUtil.prettyHexDump(userData));
        AlarmType alarmtype = null;
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        int length = userData.readShort();
        if (length != 0){
            byte result = userData.readByte();
            byte unknown = userData.readByte();

            if (length>2) {
                alarmtype = AlarmType.valueOf(userData.readByte());
                unknown = userData.readByte();
                unknown = userData.readByte();                        
            } else {
                //Free dummy byte

            }
        }
        
       return new AlarmMessagePayload(returnCode,
                        dataTransportSize,
                        alarmtype,
                        length,
                        null);
    }
    
    private CpuDiagnosticMessagePayload decodeCpuDiagnosticMessagePayload(CpuServicesParameter parameter, ByteBuf userData){

        LocalDateTime timestamp;
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        int length = userData.readShort(); //TODO: Validate userData length
        short EventID = userData.readShort();
        byte PriorityClass = userData.readByte();
        byte ObNumber = userData.readByte();
        short DatID = userData.readShort();
        short Info1 =  userData.readShort();
        int Info2 =  userData.readInt();
        
        //It is assumed that you have synchronized the time of your PLC with PC.
        //TODO: Write util function for translate S7 DateTime
        timestamp = readDateAndTime(userData);
        
        CpuDiagnosticMessageItem diagnosticitem = new CpuDiagnosticMessageItem(EventID,
                                                        PriorityClass,
                                                        ObNumber,
                                                        DatID,
                                                        Info1,
                                                        Info2,
                                                        timestamp);
        
        return new CpuDiagnosticMessagePayload(returnCode,
                                    dataTransportSize,
                                    length,
                                    diagnosticitem);
    }    
    
    private AlarmMessagePayload decodeMessageServicePushPayload(CpuServicesParameter parameter, ByteBuf userData){

        List<Object> MessageObjects = new LinkedList<>();
        List<ByteBuf> values = new LinkedList<>();
        int length;
        //Alarm message
        LocalDateTime timestamp;
        Byte FunctionID;
        byte NumberOfMessgaeObjects;
        byte EventGoing = 0x00;
        byte EventComming = 0x00;
        byte EventLastChange = 0x00;
        
        
        //
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        length = userData.readShort();
        
        //It is assumed that you have synchronized the time of your PLC with PC.
        //The most important field.
        timestamp = readDateAndTime(userData);
        
        FunctionID = userData.readByte();
        NumberOfMessgaeObjects = userData.readByte();

        for (int i = 0; i < NumberOfMessgaeObjects; i++){
            {
                byte VariableSpecification = userData.readByte();
                byte Length = userData.readByte();
                VariableAddressingMode SyntaxID = VariableAddressingMode.valueOf(userData.readByte());
                byte NumberOfValues = userData.readByte();
                int EventID = userData.readInt();
                
                byte EventState = 0x00;
                byte State = 0x00;
                if (parameter.getSubFunctionGroup() != CpuServicesParameterSubFunctionGroup.ALARM_ACK_IND) {                                          
                    EventState = userData.readByte();                
                    State = userData.readByte();
                }
                
                byte AckStateGoing = userData.readByte();
                byte AckStateComming = userData.readByte();
                
                switch(parameter.getSubFunctionGroup()){
                    case ALARM8:
                    case NOTIFY8:
                    case NOTIFY:
                        EventGoing  = userData.readByte(); 
                        EventComming  = userData.readByte(); 
                        EventLastChange  = userData.readByte(); 
                        userData.readByte(); //Reserved = 0x00
                        break;                    
                }                
                
                //TODO: If NumberOfValues == 0 then AssociatedValues is null
                List<AssociatedValueItem> AssociatedValues = new LinkedList<>();
                        
                for (int j = 0; j < NumberOfValues; j++){
                    {
                        DataTransportErrorCode valueCode = DataTransportErrorCode.valueOf(userData.readByte());
                        DataTransportSize valueTransportSize = DataTransportSize.valueOf(userData.readByte());
                        int valueLength = userData.readShort();
                        //Max length of value is 12 bytes
                        valueLength = (valueLength >> 4)*2;
                        
                        ByteBuf Data = userData.readBytes(valueLength);
                        
                        AssociatedValues.add(new AssociatedValueItem(valueCode,
                                                valueTransportSize,
                                                valueLength,
                                                Data));
                    }
                }
                
                MessageObjects.add( new MessageObjectItem(VariableSpecification,
                                            Length,
                                            SyntaxID,
                                            NumberOfValues,
                                            EventID,
                                            EventState,
                                            State,
                                            AckStateGoing,
                                            AckStateComming,
                                            EventGoing,
                                            EventComming, 
                                            EventLastChange,                        
                                            AssociatedValues));
                
            }            
        }
        
        return new AlarmMessagePayload(returnCode,
                        dataTransportSize,
                        parameter.getSubFunctionGroup(),
                        length,
                        new AlarmMessageItem(timestamp,
                                FunctionID,
                                NumberOfMessgaeObjects,
                                MessageObjects));  
    };

    private AlarmMessagePayload decodeMessageServiceQueryPayload(CpuServicesParameter parameter, ByteBuf userData){
      
        List<Object> MessageObjects = new LinkedList<>();
        CpuServicesResponseParameter thisparameter = null;
        int length;
        byte FunctionID = 0x00;
        byte NumberOfMessageObjects = 0x00; //Say 1, but I have 2 messages? Why?
        DataTransportErrorCode AlarmReturnCode = DataTransportErrorCode.RESERVED;
        DataTransportSize AlarmTransportSize = DataTransportSize.NULL;
        int CompleteDataLength = 0x0000;
        short dataUnitReferenceNumber = 0x0000;
        boolean lastdataunit = true;
        
        //Data section
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        length = userData.readShort(); //Number of message objects?
        
        if (parameter instanceof CpuServicesResponseParameter){
            thisparameter = (CpuServicesResponseParameter) parameter;
            dataUnitReferenceNumber = (short) thisparameter.getDataUnitReferenceNumber();
            lastdataunit = thisparameter.isLastDataUnit();
        }
        
        //Response to query from S7-1500
        if ((length == 0) && 
            lastdataunit && 
            (thisparameter.getSubFunctionGroup() == CpuServicesParameterSubFunctionGroup.ALARM_QUERY)){
            return new AlarmMessagePayload(returnCode,
                    dataTransportSize,
                    CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                    length,
                    new AlarmMessageItem(FunctionID,
                            NumberOfMessageObjects,
                            AlarmReturnCode,
                            AlarmTransportSize,
                            CompleteDataLength,
                            MessageObjects));              
        }        
        
        //Response to query
        if ((length == 6) && 
            lastdataunit && 
            (thisparameter.getSubFunctionGroup() == CpuServicesParameterSubFunctionGroup.ALARM_QUERY)){
            //Alarm message information
            FunctionID = userData.readByte();
            NumberOfMessageObjects = userData.readByte();
            AlarmReturnCode = DataTransportErrorCode.valueOf(userData.readByte());
            AlarmTransportSize = DataTransportSize.valueOf(userData.readByte());
            CompleteDataLength = userData.readShort();    
            
            return new AlarmMessagePayload(returnCode,
                    dataTransportSize,
                    CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                    length,
                    new AlarmMessageItem(FunctionID,
                            NumberOfMessageObjects,
                            AlarmReturnCode,
                            AlarmTransportSize,
                            CompleteDataLength,
                            MessageObjects));              
            
        }        
        
        //Fragmente code
        //The next Level query again.
                
        if ((length > 2) && (!lastdataunit)){
            
            //Is alway a CpuServicesResponseParameter instance?
            //CpuServicesResponseParameter thisparameter = (CpuServicesResponseParameter) parameter;
            Pair<LocalDateTime, Queue<ByteBuf>> fragments = fragmentedData.get(dataUnitReferenceNumber);
            if (fragments != null) {
                Queue<ByteBuf> bytebufqueue = fragments.getValue();
                bytebufqueue.add(userData);
            } else {
                Queue<ByteBuf> bytebufqueue = new ArrayDeque();
                bytebufqueue.add(userData);
                Pair<LocalDateTime, Queue<ByteBuf>> firtsfragment = new ImmutablePair(LocalDateTime.now(), bytebufqueue);
                fragmentedData.put(dataUnitReferenceNumber, firtsfragment);
            }
            
            return new AlarmMessagePayload(returnCode,
                        dataTransportSize,
                        CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                        length,
                        null);             
        }
        
        //Reemsable the fragments       
        if (dataUnitReferenceNumber != 0x0000){
            Pair<LocalDateTime, Queue<ByteBuf>> fragments = fragmentedData.get(dataUnitReferenceNumber);
            if (fragments != null){ 
                Queue<ByteBuf> bytebufqueue = fragments.getValue();
                ByteBuf payload = bytebufqueue.remove();
                for(ByteBuf nextpayload:bytebufqueue){
                    if (nextpayload != null) { 
                        payload.writeBytes(nextpayload);
                        nextpayload.release();                          
                    };
                }               
                payload.writeBytes(userData);
                userData.clear();                 
                userData.writeBytes(payload);                
                payload.release();               
                fragmentedData.remove(dataUnitReferenceNumber);
            }            
        }        
 
        //Alarm message information
        FunctionID = userData.readByte();
        NumberOfMessageObjects = userData.readByte();
        AlarmReturnCode = DataTransportErrorCode.valueOf(userData.readByte());
        AlarmTransportSize = DataTransportSize.valueOf(userData.readByte());
        CompleteDataLength = userData.readShort();
        
        //Message Object
        if (AlarmReturnCode == DataTransportErrorCode.OK){
            ///length = 1+2+1+4+2+1+1+1+8+1+1+2+8+1+1+2 = 35
            while (userData.isReadable(35)){              
                byte LengthOfDataSet = userData.readByte();
                int Unknown = userData.readShort();
                AlarmQueryType AlarmType = AlarmQueryType.valueOf(userData.readByte());
                int EnentID = userData.readInt();
                Unknown = userData.readByte();
                byte EventState =  userData.readByte();
                byte AckStateGoing = userData.readByte();
                byte AckStateComing = userData.readByte();  
                
                LocalDateTime timestampComing = null;
                LocalDateTime timestampGoing = null;
      
                List<AssociatedValueItem> ComingValues = new LinkedList<>();
                List<AssociatedValueItem> GoingValues = new LinkedList<>();
               
                switch(AlarmType){
                    case ALARM_S:{
                        //Coming data
                        timestampComing = readDateAndTime(userData);                         
                        DataTransportErrorCode valueCode = DataTransportErrorCode.valueOf(userData.readByte());
                        DataTransportSize valueTransportSize = DataTransportSize.valueOf(userData.readByte());
                        
                        int valueLength = userData.readShort() / 8;
                        
                        ByteBuf Data = userData.readBytes(valueLength);
                        ComingValues.add(new AssociatedValueItem(valueCode,
                                                valueTransportSize,
                                                valueLength,
                                                Data));   
                        //Going data
                        timestampGoing = readDateAndTime(userData);
                        
                        valueCode = DataTransportErrorCode.valueOf(userData.readByte());
                        valueTransportSize = DataTransportSize.valueOf(userData.readByte());
                        
                        valueLength = userData.readShort() / 8;;
                        
                        Data = userData.readBytes(valueLength);
                        GoingValues.add(new AssociatedValueItem(valueCode,
                                                valueTransportSize,
                                                valueLength,
                                                Data));                            
                        }
                    break;
                    case ALARM_8: {
                        //TODO: Decode values for ALARM_8
                    }
                    break;
                    default:;                    
                }
              MessageObjects.add(new MessageObjectItem(LengthOfDataSet,
                                        AlarmType,
                                        EnentID,
                                        EventState,
                                        AckStateGoing,
                                        AckStateComing,
                                        timestampComing,
                                        ComingValues,
                                        timestampGoing,
                                        GoingValues));
            }    
        }
        
        return new AlarmMessagePayload(returnCode,
                    dataTransportSize,
                    CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                    length,
                    new AlarmMessageItem(FunctionID,
                            NumberOfMessageObjects,
                            AlarmReturnCode,
                            AlarmTransportSize,
                            CompleteDataLength,
                            MessageObjects));                 

    }    
    
    private AlarmMessagePayload decodeMessageServiceAckPayload(CpuServicesParameter parameter, ByteBuf userData){
        List<Object> MessageObjects = new LinkedList<>();
        //Data section
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        int length = userData.readShort(); //Number of message objects?
        
        if (returnCode == DataTransportErrorCode.OK) {
        
            //Alarm message section
            byte FunctionID = userData.readByte();
            byte NumberOfMessageObjects = userData.readByte();

            //In the next leve if is != null -> Success
            for(int i=0; i< NumberOfMessageObjects; i++) {
                MessageObjects.add(DataTransportErrorCode.valueOf(userData.readByte()));
            };

            return new AlarmMessagePayload(returnCode,
                        dataTransportSize,
                        CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                        length,
                        new AlarmMessageItem(FunctionID,
                                NumberOfMessageObjects,
                                MessageObjects)); 
        }
        
        return new AlarmMessagePayload(returnCode,
                    dataTransportSize,
                    CpuServicesParameterSubFunctionGroup.ALARM_QUERY,
                    length,
                    null);                 
    }    
    
    
    private CpuCyclicServicesResponsePayload decodeCyclicServiceResponsePayload(CpuCyclicServicesResponseParameter parameter, ByteBuf userData){
        //logger.info("decodeCyclicServiceResponsePayload :\r\n" + ByteBufUtil.prettyHexDump(userData));
        DataTransportErrorCode AlarmReturnCode;
        DataTransportSize AlarmTransportSize;
        int length;
        int itemcount = 0;
        List<AssociatedValueItem> Values = new LinkedList<>();
        
        //Data section
        DataTransportErrorCode returnCode = DataTransportErrorCode.valueOf(userData.readByte());
        DataTransportSize dataTransportSize = DataTransportSize.valueOf(userData.readByte());
        length = userData.readShort(); //Number of bytes
        
        //Check for Unsubscribe
        if (length != 0x0000) {
            userData.readByte(); //TODO: Sometimes is 0x00 another 0x01, Blinking when I have TIA running
            itemcount = (length==0?0x0000:userData.readByte());
            try {
                for (int i=0; i < itemcount; i++ ){                         
                    DataTransportErrorCode valueCode = DataTransportErrorCode.valueOf(userData.readByte());
                    DataTransportSize valueTransportSize = DataTransportSize.valueOf(userData.readByte());

                    int valueLength = userData.readShort() / 8;  
                    if ((valueLength % 2 != 0) && (userData.isReadable(valueLength+1))) {

                        userData.readByte(); //Fill byte for odd number of bytes
                    }

                    ByteBuf Data = userData.readBytes(valueLength);

                    Values.add(new AssociatedValueItem(valueCode,
                                    valueTransportSize,
                                    valueLength,
                                    Data));            
                }
            } catch (Exception e) {
                logger.info(e.getMessage());
                return null;
            }
        }

        return new CpuCyclicServicesResponsePayload(
                    returnCode,
                    dataTransportSize,
                    parameter.getSequenceNumber(),
                    length,
                    itemcount,
                    Values);
                

    }    
    
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Helpers
    ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////


    private synchronized void trySendingMessages(ChannelHandlerContext ctx) {
        logger.info("trySendingMessages {} < {}",sentButUnacknowledgedTpdus.size(),maxAmqCaller);
        while(sentButUnacknowledgedTpdus.size() < maxAmqCaller) {
            // Get the TPDU that is up next in the queue.
            DataTpdu curTpdu = (DataTpdu) queue.current();

            if (curTpdu != null) {
                // Send the TPDU.
                try {
                    //logger.info("trySendingMessages: Trata de enviar todos los mensajes...");
                    S7Message message = (S7Message) curTpdu.getParent();
                    System.out.println(">S7Message: " + message.getTpduReference());
                    ChannelFuture channelFuture = queue.removeAndWrite();
                    ctx.flush();
                    if (channelFuture == null) {
                        break;
                    }
                } catch (Exception e) {
                    logger.debug("trySendingMessages: Error sending more queues messages", e);
                    ctx.fireExceptionCaught(e);
                }

                if(curTpdu.getParent() != null) {
                    //logger.info("Tiene un pariente...");
                    // Add it to the list of sentButUnacknowledgedTpdus.
                    // (It seems that the S7 drops the value of the COTP reference id, so we have to use the S7 one)
                    S7RequestMessage s7RequestMessage = (S7RequestMessage) curTpdu.getParent();
                    
                    sentButUnacknowledgedTpdus.put(s7RequestMessage.getTpduReference(), curTpdu);

                    logger.debug("trySendingMessages: S7 Message with id {} sent", s7RequestMessage.getTpduReference());
                }
                // TODO: Perhaps remove this.
                break;
            } else {
                break;
            }
        }
        ctx.flush();
    }

    private S7ControllerType lookupControllerType(String articleNumber) {
        if(!articleNumber.startsWith("6ES7 ")) {
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
                if(logger.isInfoEnabled()) {
                    logger.info(String.format("Looking up unknown article number %s", articleNumber));
                }
                return S7ControllerType.ANY;
        }
    }

    /*
     * Date and time of day (BCD coded).
     *          +----------------+
     * Byte n   | Year   0 to 99 |
     *          +----------------+
     * Byte n+1 | Month  1 to 12 |
     *          +----------------+
     * Byte n+2 | Day    1 to 31 |    
     *          +----------------+
     * Byte n+3 | Hour   0 to 23 |    
     *          +----------------+  
     * Byte n+4 | Minute 0 to 59 |  
     *          +----------------+
     * Byte n+5 | Second 0 to 59 |     
     *          +----------------+
     * Byte n+6 | ms    0 to 999 |
     * Byte n+7 | X X X X X D O W|    
     *          +----------------+    
     * DOW: Day of weed (last 3 bits)
    */
    private LocalDateTime readDateAndTime(ByteBuf data) {
        //from Plc4XS7Protocol
        int year=convertByteToBcd(data.readByte());
        byte themonth = data.readByte();
        int month=convertByteToBcd(themonth==0x00?0x01:themonth);
        byte theday = data.readByte();
        int day=convertByteToBcd(theday==0x00?0x01:theday);
        int hour=convertByteToBcd(data.readByte());
        int minute=convertByteToBcd(data.readByte());
        int second=convertByteToBcd(data.readByte());        
        int milliseconds = (data.readShort() & 0xfff0) >> 4;
        
        int cen = ((milliseconds & 0x0f00) >> 8) * 100;
        int dec = ((milliseconds & 0x00f0) >> 4) * 10;
        milliseconds = cen + dec + (milliseconds & 0x000f);
        int nanoseconds = milliseconds * 1000000;
        
        //data-type ranges from 1990 up to 2089
        if(year>=90){
            year+=1900;
        }
        else{
            year+=2000;
        }
        
        return LocalDateTime.of(year,month,day,hour,minute,second, nanoseconds);
    }
    

    private LocalTime readTimeOfDay(ByteBuf data) {
        //per definition for Date_And_Time only the first 6 bytes are used

        int millisSinsMidnight = data.readInt();

        return LocalTime.now().withHour(0).withMinute(0).withSecond(0).withNano(0).plus(millisSinsMidnight, ChronoUnit.MILLIS);

    }

    private LocalDate readDate(ByteBuf data) {
        //per definition for Date_And_Time only the first 6 bytes are used

        int daysSince1990 = data.readUnsignedShort();

        System.out.println(daysSince1990);
        return LocalDate.now().withYear(1990).withDayOfMonth(1).withMonth(1).plus(daysSince1990, ChronoUnit.DAYS);

    }

    /**
     * converts incoming byte to an integer regarding used BCD format
     * @param incomingByte
     * @return converted BCD number
     */
    private static int convertByteToBcd(byte incomingByte) {
        int dec = (incomingByte >> 4) * 10;
        return dec + (incomingByte & 0x0f);
    }    

    
}
