package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test() {
	c := make(chan int, 5)
	c <- 5
	c <- 6
	close(c)
	//c <- 7 //will panic here when try to send msg to a closed signal
	fmt.Println(<-c) //but still can fetch from the closed channel
}

func consumer(queue chan int) {
	defer wg.Done()
	for {
		data, ok := <-queue
		if !ok {
			break
		}
		fmt.Println(data)
		time.Sleep(time.Second)
	}
}

func main() {
	// 1,
	//var msg chan int
	// 2,
	//msg := make(chan int)
	// 3,
	msg := make(chan int, 1)
	msg <- 1
	//data := <-msg
	//fmt.Print(data)
	wg.Add(1)
	go consumer(msg)
	msg <- 2
	close(msg)
	wg.Wait()
}
