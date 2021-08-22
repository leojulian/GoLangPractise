package server_proxy

import (
	"GoLang/rpc_test/handler"
	"net/rpc"
)

type HelloServicer interface {
	Hello(request string, reply *string) error
}

func RegisterHelloService(srv HelloServicer) error {
	return rpc.RegisterName(handler.HelloServiceName, srv)
}
