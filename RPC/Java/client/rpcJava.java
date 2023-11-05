package testss;
/*
 * @author: AjayBadrinath
 * @date: 04-11-23
 * 											                      RPC Client Implementation
 * 
 * 
 * This is a demo for Executing Remote Function Call running over a Go Server . The Example is a Proof Of Concept Calculator Application .
 * This can however be extended to other practical purposes.
 * Server : Servercal.go
 * Client : rpcJava.java
 * Communication between both server and client is established over TCP socket and The Request and Response is sent/recieved via JSON 
 * 
 * I have made the code as Modular and reusable as possible with the thought of instantiating each class for each new request similar to API Calls 
 * wherein each request is a new request. 
 *
 * 
 * */
/*
 *                                                             Imports
 * */
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.Socket;
import java.net.SocketException;
import java.net.UnknownHostException;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import org.json.simple.JSONObject;
/*
 * Class TCPConnection: Base Class for Establishing connection to remote server via Socket
 * @functions:send,recieve
 * */
class TCPConnection{
	
	String Address;
	int port;
	Socket c;
	PrintWriter write;
	BufferedReader readres;
	/*
	 * Constructor Establishes connection to RPC Server 
	 * */
	TCPConnection(String Address,int port) throws UnknownHostException, IOException,SocketException{
		this.Address=Address;
		this.port=port;
		c = new Socket(Address,port);
		
		
	}
	/*
	 * Send a JSON Object to Server using PrintWriter
	 * */
	 void send(JSONObject object) throws IOException {
		write=new PrintWriter(c.getOutputStream(),true);
		write.println(object);
		
		
	}
	 /*
	  * Recieve a JSON Object  using InputStream and Pipe to Buffered Reader 
	  * */
	 String recieve() throws IOException {
		readres= new BufferedReader(new InputStreamReader(c.getInputStream()));
		return readres.readLine();
	 }
	
	
	
}
/*
 * Class JSONConstruct: This class is for construction of JSON Payload 
 * @functions: constructJSON()->Implicit call ->void param
 * */
class JSONConstruct{
	String Method;
	int num1,num2;
	HashMap<String,Object> request ,reqparam;
	List <HashMap<String,Object>>k1;
	public JSONObject RPCRequest;
	/*
	 * Construct JSONObject with id,method,param 
	 * for the param format [{'A':x,'B':t}] List <HashMap<String,Object>> is used by passing the previously constructed param object
	 * Access the JSON by accessing the public JSONObject variable .
	 * */
	JSONConstruct(String Method,int num1,int num2){
		this.Method=Method;
		this.num1=num1;
		this.num2=num2;
		request= new HashMap<String,Object>();
		reqparam= new HashMap<String,Object>();
		
		k1=new ArrayList<HashMap<String,Object>>();
		constructJSON();
		RPCRequest=new JSONObject(request);
	}
	void constructJSON() {
		reqparam.put("A", num1);
		reqparam.put("B", num2);
		k1.add(reqparam);
		
		request.put("method", "ExportVariable."+Method);
		request.put("params", k1);
		request.put("id", 1);
		
	}
	
	
	
}
class rpcJava{
 public static void main(String[]args) throws UnknownHostException, IOException  {
	 
	/*
	 * Perform Single RPC Request to Server.
	 * */
	JSONConstruct x1=new JSONConstruct("Divide",100,2);
	TCPConnection s1=new TCPConnection("localhost",1234);
	s1.send(x1.RPCRequest);
	System.out.println(s1.recieve());
	
 }
}
