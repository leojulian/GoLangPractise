package main

import (
	"GoLang/rpc_test/client_proxy"
	"fmt"
)

func main() {
	//client, err := rpc.Dial("tcp", "localhost:1234")
	//if err != nil {
	//	fmt.Println("failed")
	//}
	//var reply *string = new(string)
	//// var reply string
	////&reply
	//err = client.Call(handler.HelloServiceName, "leo", reply)
	//
	var reply string
	client := client_proxy.NewHelloServiceClient("tcp", "localhost:1234")
	err := client.Hello("leo", &reply)
	if err != nil {
		fmt.Println("Call failed")
	} else {
		fmt.Println(reply)
	}

}
