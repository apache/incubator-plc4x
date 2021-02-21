package org.apache.plc4x.processors.plc4x4nifi;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.TimeUnit;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.utils.connectionpool.PooledPlcDriverManager;

public class SandBox {
	
	static PooledPlcDriverManager pool = new PooledPlcDriverManager();
	static Boolean recreateOnError = true;
	
	public static void main(String[] args) throws Exception {
		for (int i = 0; i<12; i++) {
			try (PlcConnection connection = pool.getConnection("s7://10.105.143.1:102?remote-rack=0&remote-slot=0&controller-type=S7_300")){
					System.out.println(i);
					System.out.println(connection.getMetadata().canRead());
					PlcReadRequest.Builder builder = connection.readRequestBuilder();
					//testRunner.setProperty(Plc4xSourceRecordProcessor.PLC_ADDRESS_STRING, "miboolean=%DB20:DBX6.0:BOOL;miint=%DB20:DBW06:INT");
	                builder.addItem("miboolean", "%DB20:DBX6.0:BOOL");
	                builder.addItem("miint", "%DB20:DBW06:INT");
	                
	                PlcReadRequest readRequest = builder.build();
	                PlcReadResponse response = readRequest.execute().get(10, TimeUnit.SECONDS);
	                Map<String, String> attributes = new HashMap<>();
	                for (String fieldName : response.getFieldNames()) {
	                	System.out.println("fieldName: "+fieldName);
	                    for (int k = 0; k < response.getNumberOfValues(fieldName); k++) {
	                    	System.out.println("fieldName number of values: "+response.getNumberOfValues(fieldName));
	                        Object value = response.getObject(fieldName, k);
	                        attributes.put(fieldName, String.valueOf(value));
	                        System.out.println("fieldName value: "+value);
	                    }
	                }
			} catch (Exception e) {
				if(recreateOnError)
					pool = new PooledPlcDriverManager();
				e.printStackTrace();
			}
		}
    	
	}
}
