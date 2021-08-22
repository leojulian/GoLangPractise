package main

import (
	"GoLang/rpc_test/handler"
	"net"
	"net/rpc"

	"GoLang/rpc_test/server_proxy"
)

func main() {
	listener, _ := net.Listen("tcp", ":1234")

	//_ = rpc.RegisterName(handler.HelloServiceName, handler.HelloService{})

	_ = server_proxy.RegisterHelloService(&handler.HelloService{})

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
