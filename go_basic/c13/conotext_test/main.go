package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var stop = make(chan bool)

func test() {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Sprintln("goroutine stopped")
			return
		default:
			time.Sleep(time.Second * 2)
			fmt.Println("Running")
		}
		//time.Sleep(time.Second * 2)
		//fmt.Println("Running")
	}
}

func test1(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Test1 Stopped")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("Running")
		}
	}
}

func main1() {
	wg.Add(1)
	go test()
	time.Sleep(time.Second * 6)
	stop <- true
	wg.Wait()
	fmt.Printf("Finished")
}

func main2() {
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go test1(ctx)

	time.Sleep(time.Second * 6)
	cancel()

	wg.Wait()
	fmt.Printf("Finished")
}

func main() {
	wg.Add(1)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*6)
	//ctx, _ := context.WithDeadline(context.Background(), )

	go test1(ctx)

	wg.Wait()
	fmt.Printf("Finished")
}
