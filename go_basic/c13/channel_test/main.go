package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

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
