package client_proxy

import (
	"GoLang/jsonrpc_test/handler"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}

func NewHelloServiceClient(protocol, address string) HelloServiceStub {
	conn, err := rpc.Dial(protocol, address)
	if err != nil {
		panic(err)
	}
	return HelloServiceStub{conn}
}

func (c *HelloServiceStub) Hello(request string, reply *string) error {
	err := c.Call(handler.HelloServiceName+".Hello", request, reply)
	if err != nil {
		return err
	}
	return nil
}
