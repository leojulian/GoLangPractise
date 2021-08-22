package main

import "fmt"

func main1() {
	//f := os.Open()
	fmt.Println("test1")
	defer fmt.Println("defer test")
	defer fmt.Println("defer test1")
	fmt.Println("test2")
}

func test() int {
	a := 1
	defer func() { a++ }()
	return a
}

func test1() *int {
	a := 1
	b := &a
	defer func() {
		*b++
	}()
	return b
}

func main() {

	//fmt.Println(test())
	fmt.Println(*test1())
}
