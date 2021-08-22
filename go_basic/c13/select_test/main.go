package main

import (
	"fmt"
	"time"
)

func main() {
	// 1, 随机公平的选择一个case语句执行
	// 2, 查看channel是否驻留
	timeOut := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 5)
		timeOut <- true
	}()

	timeOut1 := make(chan bool, 1)
	go func() {
		time.Sleep(time.Second * 1)
		timeOut1 <- true
	}()

	select {
	case <-timeOut:
		fmt.Println("goroutine 0 finish")
		//break
	case <-timeOut1:
		fmt.Println("goroutine 1 finish")
	default:
		fmt.Println("default finished")
	}
}
