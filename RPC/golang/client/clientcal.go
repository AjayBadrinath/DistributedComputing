/*
					Simple RPC Communication using Golang  to Golang Environment
Author: AjayBadrinath
Date: 23-09-2023

*/

package main // GoPath default Package Unless your own package main  is the default entry path for go Execution

//imports for logging and rpc
import (
	"fmt"
	"log"
	"net/rpc"
)

// Go Specifies that  only exported objects are visible to  server

//Struct Similar to one defined in the server with same data types .

type SomethingElse struct {
	A int64
	B int64
}

func main() {
	//DialHTTP perform Raw network connection
	//Then creates client object which has access to methods to call and Go
	//This Example is a Synchronous call to the server while it can also be done asynchronously

	client, errorHandle := rpc.DialHTTP("tcp", ":4200")
	var reply int64 // Reply From RPC Server of the same type
	if errorHandle != nil {
		log.Print(errorHandle)
	}
	var a, b int64 // Some Example variable
	var mode string
	var rpcCall error
	for {

		fmt.Scanln(&mode, &a, &b)
		args := SomethingElse{a, b}
		/*
			client.Call -> should call the Exposed method/methodtype along with function name .
		*/
		if mode == "add" {
			rpcCall = client.Call("ExportVariable.Add", args, &reply)
		} else if mode == "subtract" {
			rpcCall = client.Call("ExportVariable.Subtract", args, &reply)

		} else if mode == "multiply" {
			rpcCall = client.Call("ExportVariable.Multiply", args, &reply)

		} else {
			rpcCall = client.Call("ExportVariable.Divide", args, &reply)

		}
		if rpcCall != nil {
			log.Print(rpcCall)
		}
		// Printthe result .
		fmt.Printf("%d\n", reply)
	}
}
