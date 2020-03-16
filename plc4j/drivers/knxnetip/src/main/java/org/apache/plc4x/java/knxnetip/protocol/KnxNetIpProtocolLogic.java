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
package org.apache.plc4x.java.knxnetip.protocol;

import io.netty.channel.socket.DatagramChannel;
import org.apache.commons.codec.binary.Hex;
import org.apache.commons.lang3.tuple.ImmutablePair;
import org.apache.commons.lang3.tuple.Pair;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.PlcSubscriptionEvent;
import org.apache.plc4x.java.api.messages.PlcSubscriptionRequest;
import org.apache.plc4x.java.api.messages.PlcSubscriptionResponse;
import org.apache.plc4x.java.api.model.PlcConsumerRegistration;
import org.apache.plc4x.java.api.model.PlcField;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.types.PlcResponseCode;
import org.apache.plc4x.java.api.value.PlcString;
import org.apache.plc4x.java.api.value.PlcStruct;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.knxnetip.configuration.KnxNetIpConfiguration;
import org.apache.plc4x.java.knxnetip.ets5.Ets5Parser;
import org.apache.plc4x.java.knxnetip.ets5.model.Ets5Model;
import org.apache.plc4x.java.knxnetip.ets5.model.GroupAddress;
import org.apache.plc4x.java.knxnetip.field.KnxNetIpField;
import org.apache.plc4x.java.knxnetip.model.KnxNetIpSubscriptionHandle;
import org.apache.plc4x.java.knxnetip.readwrite.KNXGroupAddress;
import org.apache.plc4x.java.knxnetip.readwrite.KNXGroupAddress2Level;
import org.apache.plc4x.java.knxnetip.readwrite.KNXGroupAddress3Level;
import org.apache.plc4x.java.knxnetip.readwrite.KNXGroupAddressFreeLevel;
import org.apache.plc4x.java.knxnetip.readwrite.io.KNXGroupAddressIO;
import org.apache.plc4x.java.knxnetip.readwrite.io.KnxDatapointIO;
import org.apache.plc4x.java.spi.ConversationContext;
import org.apache.plc4x.java.spi.Plc4xProtocolBase;
import org.apache.plc4x.java.knxnetip.readwrite.*;
import org.apache.plc4x.java.knxnetip.readwrite.types.HostProtocolCode;
import org.apache.plc4x.java.knxnetip.readwrite.types.KnxLayer;
import org.apache.plc4x.java.knxnetip.readwrite.types.Status;
import org.apache.plc4x.java.spi.configuration.HasConfiguration;
import org.apache.plc4x.java.spi.generation.ReadBuffer;
import org.apache.plc4x.java.spi.messages.DefaultPlcSubscriptionEvent;
import org.apache.plc4x.java.spi.messages.DefaultPlcSubscriptionResponse;
import org.apache.plc4x.java.spi.messages.InternalPlcSubscriptionRequest;
import org.apache.plc4x.java.spi.messages.PlcSubscriber;
import org.apache.plc4x.java.spi.model.DefaultPlcConsumerRegistration;
import org.apache.plc4x.java.spi.model.InternalPlcSubscriptionHandle;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.File;
import java.net.InetSocketAddress;
import java.time.Duration;
import java.time.Instant;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Consumer;

public class KnxNetIpProtocolLogic extends Plc4xProtocolBase<KNXNetIPMessage> implements HasConfiguration<KnxNetIpConfiguration>, PlcSubscriber {

    private static final Logger LOGGER = LoggerFactory.getLogger(KnxNetIpProtocolLogic.class);

    private boolean passiveMode = false;
    private KNXAddress gatewayAddress;
    private String gatewayName;
    private IPAddress localIPAddress;
    private int localPort;
    private short communicationChannelId;

    private Timer connectionStateTimer;

    private byte groupAddressType;
    private Ets5Model ets5Model;

    private Map<DefaultPlcConsumerRegistration, Consumer<PlcSubscriptionEvent>> consumers = new ConcurrentHashMap<>();

    @Override
    public void setConfiguration(KnxNetIpConfiguration configuration) {
        if (configuration.knxprojFilePath != null) {
            File knxprojFile = new File(configuration.knxprojFilePath);
            if (knxprojFile.exists() && knxprojFile.isFile()) {
                ets5Model = new Ets5Parser().parse(knxprojFile);
                groupAddressType = ets5Model.getGroupAddressType();
            } else {
                throw new PlcRuntimeException(String.format(
                    "File specified with 'knxproj-file-path' does not exist or is not a file: '%s'",
                    configuration.knxprojFilePath));
            }
        } else {
            groupAddressType = (byte) configuration.groupAddressType;
        }
    }

    @Override
    public void onConnect(ConversationContext<KNXNetIPMessage> context) {
        // Only the UDP transport supports login.
        if(context.getChannel() instanceof DatagramChannel) {
            LOGGER.info("KNX Driver running in ACTIVE mode.");
            passiveMode = false;

            DatagramChannel channel = (DatagramChannel) context.getChannel();
            final InetSocketAddress localSocketAddress = channel.localAddress();
            localIPAddress = new IPAddress(localSocketAddress.getAddress().getAddress());
            localPort = localSocketAddress.getPort();

            // First send out a search request
            // REMARK: This might be optional ... usually we would send a search request to ip 224.0.23.12
            // Any KNX Gateway will respond with a search response. We're currently directly sending to the
            // known gateway address, so it's sort of pointless, but at least only one device will respond.
            LOGGER.info("Sending KNXnet/IP Search Request.");
            SearchRequest searchRequest = new SearchRequest(
                new HPAIDiscoveryEndpoint(HostProtocolCode.IPV4_UDP, localIPAddress, localPort));
            context.sendRequest(searchRequest)
                .expectResponse(KNXNetIPMessage.class, Duration.ofMillis(1000))
                .check(p -> p instanceof SearchResponse)
                .unwrap(p -> (SearchResponse) p)
                .handle(searchResponse -> {
                    LOGGER.info("Got KNXnet/IP Search Response.");
                    // Check if this device supports tunneling services.
                    final ServiceId tunnelingService = Arrays.stream(searchResponse.getDibSuppSvcFamilies().getServiceIds()).filter(serviceId -> serviceId instanceof KnxNetIpTunneling).findFirst().orElse(null);

                    // If this device supports this type of service, tell the driver, we found a suitable device.
                    if (tunnelingService != null) {
                        // Extract the required information form the search request.
                        gatewayAddress = searchResponse.getDibDeviceInfo().getKnxAddress();
                        gatewayName = new String(searchResponse.getDibDeviceInfo().getDeviceFriendlyName()).trim();

                        LOGGER.info(String.format("Found KNXnet/IP Gateway '%s' with KNX address '%d.%d.%d'", gatewayName,
                            gatewayAddress.getMainGroup(), gatewayAddress.getMiddleGroup(), gatewayAddress.getSubGroup()));

                        // Next send a connection request to the gateway.
                        ConnectionRequest connectionRequest = new ConnectionRequest(
                            new HPAIDiscoveryEndpoint(HostProtocolCode.IPV4_UDP, localIPAddress, localPort),
                            new HPAIDataEndpoint(HostProtocolCode.IPV4_UDP, localIPAddress, localPort),
                            new ConnectionRequestInformationTunnelConnection(KnxLayer.TUNNEL_BUSMONITOR));
                        LOGGER.info("Sending KNXnet/IP Connection Request.");
                        context.sendRequest(connectionRequest)
                            .expectResponse(KNXNetIPMessage.class, Duration.ofMillis(1000))
                            .check(p -> p instanceof ConnectionResponse)
                            .unwrap(p -> (ConnectionResponse) p)
                            .handle(connectionResponse -> {
                                // Remember the communication channel id.
                                communicationChannelId = connectionResponse.getCommunicationChannelId();

                                LOGGER.info(String.format("Received KNXnet/IP Connection Response (Connection Id %s)",
                                    communicationChannelId));

                                // Check if everything went well.
                                Status status = connectionResponse.getStatus();
                                if (status == Status.NO_ERROR) {
                                    LOGGER.info(String.format("Successfully connected to KNXnet/IP Gateway '%s' with KNX address '%d.%d.%d'", gatewayName,
                                        gatewayAddress.getMainGroup(), gatewayAddress.getMiddleGroup(), gatewayAddress.getSubGroup()));

                                    // Send an event that connection setup is complete.
                                    context.fireConnected();

                                    // Start a timer to check the connection state every 60 seconds.
                                    // This keeps the connection open if no data is transported.
                                    // Otherwise the gateway will terminate the connection.
                                    connectionStateTimer = new Timer();
                                    connectionStateTimer.scheduleAtFixedRate(new TimerTask() {
                                        @Override
                                        public void run() {
                                            ConnectionStateRequest connectionStateRequest =
                                                new ConnectionStateRequest(communicationChannelId,
                                                    new HPAIControlEndpoint(HostProtocolCode.IPV4_UDP, localIPAddress, localPort));
                                            context.sendRequest(connectionStateRequest)
                                                .expectResponse(KNXNetIPMessage.class, Duration.ofMillis(1000))
                                                .check(p -> p instanceof ConnectionStateResponse)
                                                .unwrap(p -> (ConnectionStateResponse) p)
                                                .handle(connectionStateResponse -> {
                                                    if (connectionStateResponse.getStatus() != Status.NO_ERROR) {
                                                        if (connectionStateResponse.getStatus() != null) {
                                                            LOGGER.error(String.format("Connection state problems. Got %s",
                                                                connectionStateResponse.getStatus().name()));
                                                        } else {
                                                            LOGGER.error("Connection state problems. Got no status information.");
                                                        }
                                                    }
                                                });
                                        }
                                    }, 60000, 60000);
                                } else {
                                    // The connection request wasn't successful.
                                    LOGGER.error(String.format(
                                        "Not connected to KNXnet/IP Gateway '%s' with KNX address '%d.%d.%d' got status: '%s'",
                                        gatewayName, gatewayAddress.getMainGroup(), gatewayAddress.getMiddleGroup(),
                                        gatewayAddress.getSubGroup(), status.toString()));
                                    // TODO: Actively disconnect
                                }
                            });
                    } else {
                        // This device doesn't support tunneling ... do some error handling.
                        LOGGER.error("Not connected to KNCnet/IP Gateway. The device doesn't support Tunneling.");
                        // TODO: Actively disconnect
                    }
                });
        }
        // This usually when we're running a passive mode river.
        else {
            LOGGER.info("KNX Driver running in PASSIVE mode.");
            passiveMode = true;

            // No login required, just confirm that we're connected.
            context.fireConnected();
        }
    }

    @Override
    public void onDisconnect(ConversationContext<KNXNetIPMessage> context) {
        // Cancel the timer for sending connection state requests.
        connectionStateTimer.cancel();

        DisconnectRequest disconnectRequest = new DisconnectRequest(communicationChannelId,
            new HPAIControlEndpoint(HostProtocolCode.IPV4_UDP, localIPAddress, localPort));
        context.sendRequest(disconnectRequest)
            .expectResponse(KNXNetIPMessage.class, Duration.ofMillis(1000))
            .check(p -> p instanceof DisconnectResponse)
            .unwrap(p -> (DisconnectResponse) p)
            .handle(disconnectResponse -> {
                // In general we should probably check if the disconnect was successful, but in
                // the end we couldn't do much if the disconnect would fail.
                LOGGER.info(String.format("Disconnected from KNX Gateway '%s' with KNX address '%d.%d.%d'", gatewayName,
                    gatewayAddress.getMainGroup(), gatewayAddress.getMiddleGroup(), gatewayAddress.getSubGroup()));

                // Send an event that connection disconnect is complete.
                context.fireDisconnected();
            });
    }

    @Override
    protected void decode(ConversationContext<KNXNetIPMessage> context, KNXNetIPMessage msg) throws Exception {
        // Handle a normal tunneling request, which is delivering KNX data.
        if(msg instanceof TunnelingRequest) {
            TunnelingRequest tunnelingRequest = (TunnelingRequest) msg;
            final short curCommunicationChannelId =
                tunnelingRequest.getTunnelingRequestDataBlock().getCommunicationChannelId();

            // Only if the communication channel id match, do anything with the request.
            // In case of a passive-mode driver we'll simply accept all communication ids.
            if(passiveMode || (curCommunicationChannelId == communicationChannelId)) {
                if(tunnelingRequest.getCemi() instanceof CEMIBusmonInd) {
                    CEMIBusmonInd busmonInd = (CEMIBusmonInd) tunnelingRequest.getCemi();
                    if (busmonInd.getCemiFrame() instanceof CEMIFrameData) {
                        CEMIFrameData cemiDataFrame = (CEMIFrameData) busmonInd.getCemiFrame();

                        // The first byte is actually just 6 bit long, but we'll treat it as a full one.
                        // So here we create a byte array containing the first and all the following bytes.
                        byte[] payload = new byte[1 + cemiDataFrame.getData().length];
                        payload[0] = cemiDataFrame.getDataFirstByte();
                        System.arraycopy(cemiDataFrame.getData(), 0, payload, 1, cemiDataFrame.getData().length);

                        final KNXAddress sourceAddress = cemiDataFrame.getSourceAddress();
                        final byte[] destinationGroupAddress = cemiDataFrame.getDestinationAddress();

                        // Decode the group address depending on the project settings.
                        ReadBuffer addressBuffer = new ReadBuffer(destinationGroupAddress);
                        final KNXGroupAddress knxGroupAddress =
                            KNXGroupAddressIO.staticParse(addressBuffer, groupAddressType);
                        final String destinationAddress = toString(knxGroupAddress);

                        // If there is an ETS5 model provided, continue decoding the payload.
                        if (ets5Model != null) {
                            final GroupAddress groupAddress = ets5Model.getGroupAddresses().get(destinationAddress);

                            if ((groupAddress != null) && (groupAddress.getType() != null)) {
                                LOGGER.info(String.format("Message from: '%s' to: '%s'",
                                    toString(sourceAddress), destinationAddress));

                                // Parse the payload depending on the type of the group-address.
                                ReadBuffer rawDataReader = new ReadBuffer(payload);
                                final PlcValue value = KnxDatapointIO.staticParse(rawDataReader,
                                    groupAddress.getType().getMainType(), groupAddress.getType().getSubType());

                                // Assemble the plc4x return data-structure.
                                Map<String, PlcValue> dataPointMap = new HashMap<>();
                                dataPointMap.put("sourceAddress", new PlcString(toString(sourceAddress)));
                                dataPointMap.put("targetAddress", new PlcString(groupAddress.getGroupAddress()));
                                if (groupAddress.getFunction() != null) {
                                    dataPointMap.put("location", new PlcString(groupAddress.getFunction().getSpaceName()));
                                    dataPointMap.put("function", new PlcString(groupAddress.getFunction().getName()));
                                } else {
                                    dataPointMap.put("location", null);
                                    dataPointMap.put("function", null);
                                }
                                dataPointMap.put("description", new PlcString(groupAddress.getName()));
                                dataPointMap.put("unitOfMeasurement", new PlcString(groupAddress.getType().getName()));
                                dataPointMap.put("value", value);
                                final PlcStruct dataPoint = new PlcStruct(dataPointMap);

                                // Send the data-structure.
                                publishEvent(groupAddress, dataPoint);
                            } else {
                                LOGGER.warn(
                                    String.format("Message from: '%s' to unknown group address: '%s'%n payload: '%s'",
                                        toString(sourceAddress), destinationAddress, Hex.encodeHexString(payload)));
                            }
                        }
                        // Else just output the raw payload.
                        else {
                            LOGGER.info(String.format("Raw Message: '%s' to: '%s'%n payload: '%s'",
                                KnxNetIpProtocolLogic.toString(sourceAddress), destinationAddress,
                                Hex.encodeHexString(payload))
                            );
                        }
                    }
                }

                // Confirm receipt of the request.
                final short sequenceCounter = tunnelingRequest.getTunnelingRequestDataBlock().getSequenceCounter();
                TunnelingResponse tunnelingResponse = new TunnelingResponse(
                    new TunnelingResponseDataBlock(communicationChannelId, sequenceCounter, Status.NO_ERROR));
                context.sendToWire(tunnelingResponse);
            }
        }
    }

    @Override
    public void close(ConversationContext<KNXNetIPMessage> context) {
        // TODO Implement Closing on Protocol Level
    }

    @Override
    public CompletableFuture<PlcSubscriptionResponse> subscribe(PlcSubscriptionRequest subscriptionRequest) {
        Map<String, Pair<PlcResponseCode, PlcSubscriptionHandle>> values = new HashMap<>();
        for (String fieldName : subscriptionRequest.getFieldNames()) {
            final PlcField field = subscriptionRequest.getField(fieldName);
            if(!(field instanceof KnxNetIpField)) {
                values.put(fieldName, new ImmutablePair<>(PlcResponseCode.INVALID_ADDRESS, null));
            } else {
                values.put(fieldName, new ImmutablePair<>(PlcResponseCode.OK,
                    new KnxNetIpSubscriptionHandle(this, (KnxNetIpField) field)));
            }
        }
        return CompletableFuture.completedFuture(
            new DefaultPlcSubscriptionResponse((InternalPlcSubscriptionRequest) subscriptionRequest, values));
    }

    @Override
    public PlcConsumerRegistration register(Consumer<PlcSubscriptionEvent> consumer, Collection<PlcSubscriptionHandle> collection) {
        final DefaultPlcConsumerRegistration consumerRegistration =
            new DefaultPlcConsumerRegistration(this, consumer, collection.toArray(new InternalPlcSubscriptionHandle[0]));
        consumers.put(consumerRegistration, consumer);
        return consumerRegistration;
    }

    @Override
    public void unregister(PlcConsumerRegistration plcConsumerRegistration) {
        DefaultPlcConsumerRegistration consumerRegistration = (DefaultPlcConsumerRegistration) plcConsumerRegistration;
        consumers.remove(consumerRegistration);
    }

    protected void publishEvent(GroupAddress groupAddress, PlcValue plcValue) {
        // Create a subscription event from the input.
        // TODO: Check this ... this is sort of not really right ...
        final PlcSubscriptionEvent event = new DefaultPlcSubscriptionEvent(Instant.now(),
            Collections.singletonMap("knxData", Pair.of(PlcResponseCode.OK, plcValue)));

        // Try sending the subscription event to all listeners.
        for (Map.Entry<DefaultPlcConsumerRegistration, Consumer<PlcSubscriptionEvent>> entry : consumers.entrySet()) {
            final DefaultPlcConsumerRegistration registration = entry.getKey();
            final Consumer<PlcSubscriptionEvent> consumer = entry.getValue();
            // Only if the current data point matches the subscription, publish the event to it.
            for (InternalPlcSubscriptionHandle handle : registration.getAssociatedHandles()) {
                if(handle instanceof KnxNetIpSubscriptionHandle) {
                    KnxNetIpSubscriptionHandle subscriptionHandle = (KnxNetIpSubscriptionHandle) handle;
                    // Check if the subscription matches this current event.
                    if (subscriptionHandle.getField().matchesGroupAddress(groupAddress)) {
                        consumer.accept(event);
                    }
                }
            }
        }
    }

    protected static String toString(KNXAddress knxAddress) {
        return knxAddress.getMainGroup() + "." + knxAddress.getMiddleGroup() + "." + knxAddress.getSubGroup();
    }

    protected static String toString(KNXGroupAddress groupAddress) {
        if (groupAddress instanceof KNXGroupAddress3Level) {
            KNXGroupAddress3Level level3 = (KNXGroupAddress3Level) groupAddress;
            return level3.getMainGroup() + "/" + level3.getMiddleGroup() + "/" + level3.getSubGroup();
        } else if (groupAddress instanceof KNXGroupAddress2Level) {
            KNXGroupAddress2Level level2 = (KNXGroupAddress2Level) groupAddress;
            return level2.getMainGroup() + "/" + level2.getSubGroup();
        } else if(groupAddress instanceof KNXGroupAddressFreeLevel) {
            KNXGroupAddressFreeLevel free = (KNXGroupAddressFreeLevel) groupAddress;
            return free.getSubGroup() + "";
        }
        throw new PlcRuntimeException("Unsupported Group Address Type " + groupAddress.getClass().getName());
    }

}
