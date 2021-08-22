package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func consumer(queue <-chan int) {
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

func consumer1(msg chan int) {
	defer wg.Done()
	data := <-msg
	fmt.Println(data)
}

// 1, init space for channel and use close channel
func main1() {
	//var msg chan<- int
	//var msg <-chan int
	var msg chan int
	msg = make(chan int, 1)
	msg <- 1
	//data := <-msg
	//fmt.Print(data)
	wg.Add(1)
	go consumer(msg)
	//msg <- 2
	close(msg)
	wg.Wait()
}

// 2,
func main() {
	var msg chan int
	msg = make(chan int)
	wg.Add(1)
	go consumer1(msg)
	msg <- 10
	wg.Wait()
}
