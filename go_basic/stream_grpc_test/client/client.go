package main

import (
	"GoLang/stream_grpc_test/proto"
	"context"
	"fmt"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	c := proto.NewGreeterClient(conn)

	//服务端流模式
	res, err := c.GetStream(context.Background(), &proto.StreamReqData{
		Data: "leo",
	})

	for {
		a, err := res.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}

	//客户端流模式
	putS, _ := c.PutStream(context.Background())
	for i := 0; i < 10; i++ {
		_ = putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("leo %v", i),
		})
		time.Sleep(time.Second)
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
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
			fmt.Println("Recieve from server " + data.Data)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			err := allStr.Send(&proto.StreamReqData{Data: "I am Leo"})
			if err != nil {
				fmt.Println(err)
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	var a map[string]string = make(map[string]string)
	a["a"] = "a"

}
