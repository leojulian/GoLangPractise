package main

import (
	"Example/grpc_proto_test/proto"
	"context"
	"time"

	"google.golang.org/grpc"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	c.SayHello(context.Background(), &proto.HelloRequest{
		Name:    "123",
		G:       proto.Gender_Male,
		Mp:      map[string]string{"name": "leo"},
		AddTime: timestamppb.New(time.Now()),
	})

}
