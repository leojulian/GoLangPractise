package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var rwlock sync.RWMutex

var total int

func read() {
	defer wg.Done()
	rwlock.RLock()

	fmt.Println("Read data...")
	time.Sleep(time.Second)
	fmt.Println("Read finished...")

	rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	rwlock.Lock()

	fmt.Println("Write data...")
	time.Sleep(time.Second * 10)
	fmt.Println("Write finished...")

	rwlock.Unlock()
}

func main() {
	//wg.Add(5)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
	fmt.Println(total)
}
