package main

import (
	"GoLang/grpc_test/proto"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type Server struct{}

func (s *Server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{
		Message: "Hello " + request.Name,
	}, nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	ls, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
	}
	err = g.Serve(ls)
	if err != nil {
		panic(err)
	}
}
