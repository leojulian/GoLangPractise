package main

import (
	"GoLang/stream_grpc_test/proto"
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = ":50052"

type server struct{}

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		err := res.Send(&proto.StreamResData{Data: fmt.Sprintf("%v", time.Now().Unix())})
		if err != nil {
			return err
		}
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}
	return nil
}

func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}
	}
	return nil
}

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data, err := allStr.Recv()
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("Receive data from client: " + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := allStr.Send(&proto.StreamResData{Data: "Server Response: Leo"})
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return nil
}

//GetStream(*StreamReqData, Greeter_GetStreamServer) error
//PutStream(Greeter_PutStreamServer) error
//AllStream(Greeter_AllStreamServer) error

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &server{})

	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		panic(err)
	}

	g.Serve(lis)

}
