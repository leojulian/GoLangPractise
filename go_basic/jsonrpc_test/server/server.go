package main

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"GoLang/jsonrpc_test/handler"
)

type HelloService struct {
}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "Hello " + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":1234")

	_ = rpc.RegisterName(handler.HelloServiceName, &HelloService{})
	for {
		conn, _ := listener.Accept()
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}
