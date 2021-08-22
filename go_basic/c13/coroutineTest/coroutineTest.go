package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000000; i++ {
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 2)
		}(i)
	}
	time.Sleep(time.Second * 10000000)
}
