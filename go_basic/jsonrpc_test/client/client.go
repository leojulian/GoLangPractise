package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"GoLang/jsonrpc_test/handler"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		fmt.Println("failed")
	}
	var reply *string = new(string)
	// var reply string
	//&reply
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))
	err = client.Call(handler.HelloServiceName+".Hello", "leo", reply)
	if err != nil {
		fmt.Println("Call failed")
	} else {
		fmt.Println(*reply)
	}
}
