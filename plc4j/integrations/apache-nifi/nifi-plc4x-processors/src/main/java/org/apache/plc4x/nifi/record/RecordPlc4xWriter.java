package org.apache.plc4x.nifi.record;

import java.io.IOException;
import java.io.OutputStream;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.atomic.AtomicReference;

import org.apache.nifi.avro.AvroTypeUtil;
import org.apache.nifi.flowfile.attributes.CoreAttributes;
import org.apache.nifi.logging.ComponentLog;
import org.apache.nifi.processor.ProcessSession;
import org.apache.nifi.serialization.RecordSetWriter;
import org.apache.nifi.serialization.RecordSetWriterFactory;
import org.apache.nifi.serialization.WriteResult;
import org.apache.nifi.serialization.record.Record;
import org.apache.nifi.serialization.record.RecordSchema;
import org.apache.nifi.serialization.record.RecordSet;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.nifi.util.PLC4X_PROTOCOL;
import org.apache.plc4x.nifi.util.Plc4xCommon;
import org.apache.avro.Schema;

//https://github.com/apache/nifi/blob/main/nifi-nar-bundles/nifi-standard-bundle/nifi-standard-processors/src/main/java/org/apache/nifi/processors/standard/sql/RecordSqlWriter.java
public class RecordPlc4xWriter implements Plc4xWriter {

	private final RecordSetWriterFactory recordSetWriterFactory;
	private final AtomicReference<WriteResult> writeResultRef;
	private final Map<String, String> originalAttributes;
    private String mimeType;
	
	private RecordSet fullRecordSet;
	private RecordSchema writeSchema;
	
	public RecordPlc4xWriter(RecordSetWriterFactory recordSetWriterFactory, Map<String, String> originalAttributes) {
		this.recordSetWriterFactory = recordSetWriterFactory;
        this.writeResultRef = new AtomicReference<>();
        this.originalAttributes = originalAttributes;
	}

	@Override
	public long writeResultSet(PlcReadResponse response, Map<String, String> plcAddressMap, OutputStream outputStream, ComponentLog logger, Plc4xReadResponseRowCallback callback) throws Exception {
		if (fullRecordSet == null) {	
			//TODO: infer Protocol form connection string
			
			//TODO version 1
            //final Schema avroSchema = Plc4xCommon.createSchema(plcAddressMap, PLC4X_PROTOCOL.S7);
                        
			//TODO version 2: is this mapping based on "PlcDatatype" classes valid for aall protocols?
			Map<String, ? extends PlcValue> responseDataStructure = response.getAsPlcValue().getStruct();
			
			final Schema avroSchema = Plc4xCommon.createSchema(responseDataStructure);
            
            final RecordSchema recordAvroSchema = AvroTypeUtil.createSchema(avroSchema);
            fullRecordSet = new Plc4xReadResponseRecordSetWithCallback(plcAddressMap, response, recordAvroSchema, callback);
            writeSchema = recordSetWriterFactory.getSchema(originalAttributes, fullRecordSet.getSchema());
        }
                
        try (final RecordSetWriter resultSetWriter = recordSetWriterFactory.createWriter(logger, writeSchema, outputStream, Collections.emptyMap())) {
            writeResultRef.set(resultSetWriter.write(fullRecordSet));
            if (mimeType == null) {
                mimeType = resultSetWriter.getMimeType();
            }
            return writeResultRef.get().getRecordCount();
        } catch (final Exception e) {
            throw new IOException(e);
        }
	}

	@Override
	public void writeEmptyResultSet(OutputStream outputStream, ComponentLog logger) throws IOException {
		try (final RecordSetWriter resultSetWriter = recordSetWriterFactory.createWriter(logger, writeSchema, outputStream, Collections.emptyMap())) {
            mimeType = resultSetWriter.getMimeType();
            resultSetWriter.beginRecordSet();
            resultSetWriter.finishRecordSet();
        } catch (final Exception e) {
            throw new IOException(e);
        }
	}

	@Override
	public String getMimeType() {
		return mimeType;
	}
	
	@Override
    public Map<String, String> getAttributesToAdd() {
        Map<String, String> attributesToAdd = new HashMap<>();
        attributesToAdd.put(CoreAttributes.MIME_TYPE.key(), mimeType);
        // Add any attributes from the record writer (if present)
        final WriteResult result = writeResultRef.get();
        if (result != null) {
            if (result.getAttributes() != null) {
                attributesToAdd.putAll(result.getAttributes());
            }
            attributesToAdd.put("record.count", String.valueOf(result.getRecordCount()));
        }
        return attributesToAdd;
    }

	@Override
    public void updateCounters(ProcessSession session) {
        final WriteResult result = writeResultRef.get();
        if (result != null) {
            session.adjustCounter("Records Written", result.getRecordCount(), false);
        }
    }
	
	private static class Plc4xReadResponseRecordSetWithCallback extends Plc4xReadResponseRecordSet {
        private final Plc4xReadResponseRowCallback callback;
        public Plc4xReadResponseRecordSetWithCallback(final Map<String, String> plcAddressMap, final PlcReadResponse readResponse, final RecordSchema readerSchema, Plc4xReadResponseRowCallback callback) throws IOException {
            super(plcAddressMap, readResponse, readerSchema);
            this.callback = callback;
        }
        @Override
        public Record next() throws IOException {
                if (hasMoreRows()) {
                	PlcReadResponse response = getReadResponse();
                    final Record record = createRecord(response);
                    setMoreRows(false);
                    if (callback != null) {
                        callback.processRow(response);
                    }
                    return record;
                } else {
                    return null;
                }
        }
	}

}
