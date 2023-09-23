/*
					Simple RPC Communication using Golang  to Golang Environment
Author: AjayBadrinath
Date: 23-09-2023

*/

package main //Default entry point for go
// Required imports
import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)

// Cal holds two numbers on which we have to perform operation.

type Cal struct {
	A, B int64
}

/*
For remote Access to Work
The below Criteria has to be met;
Method/Method type has to be exposed (Start with Capital Letter in the beginning)
Method Second arg has to be a pointer(result.) then write back the result by dereferencing the pointer.
Method must have return type Error.

*/

type ExportVariable int64

// Some Sample functions to perform operations.

func (f *ExportVariable) Add(args *Cal, response *int64) error {
	log.Printf("Add %d with %d\n", args.A, args.B)
	*response = args.A + args.B
	return nil
}
func (f *ExportVariable) Multiply(args *Cal, response *int64) error {
	log.Printf("Multiply %d with %d\n", args.A, args.B)
	*response = args.A * args.B
	return nil
}
func (f *ExportVariable) Subtract(args *Cal, response *int64) error {
	log.Printf("Subtract %d with %d\n", args.A, args.B)
	*response = args.A - args.B
	return nil
}
func (f *ExportVariable) Divide(args *Cal, response *int64) error {
	log.Printf("Divide %d with %d\n", args.A, args.B)
	*response = args.A / args.B
	return nil
}

/*Serving for remote access*/
func main() {
	a := new(ExportVariable)
	rpc.Register(a) //publish to the server for clients to see the methods.
	rpc.HandleHTTP()
	listener, errors := net.Listen("tcp", ":4200") //listen over tcp on port 4200
	if errors != nil {
		log.Print(errors)
	}
	http.Serve(listener, nil)
}
