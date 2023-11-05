## Remote Procedure Call Implementation 

This sub repository consists of simple Proof of concept RPC Call between Java->Go  and Python -> GO

![rpc-model](https://github.com/AjayBadrinath/DistributedComputing/assets/92035508/90b8eb3b-21c2-43db-8146-2a27676f5a1d)

The current implementation uses JSON RPC 

<h2>JSON RPC Structure</h2>

```json
  {
"jsonrpc": "2.0",
"id": 0,
"method": ExposedVar.yourMethod,
"params": {
}
}
```
Required request elements |  Description
--------------------------|--------------
jsonrpc                    |Version, which must be “2.0”. No other JSON RPC versions are currently supported.
id                          |Client-provided integer. The JSON RPC responds with the same ID, which allows the client to match requests to responses when there are concurrent requests.
method                     |Defines the method accessed with the JSON RPC. Supported options are get, set, and cli. See method options for more information.
params                      |Defines a container for any parameters related to the request. The type of parameter is dependent on the method used. See params options for more information.

<h3>References:</h3>


https://infocenter.nokia.com/public/SRLINUX216R1A/index.jsp?topic=%2Fcom.srlinux.sysmgmt%2Fhtml%2Frbc-interface.html
