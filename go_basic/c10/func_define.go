package main

import (
	"errors"
	"fmt"
)

func test() int {
	return 1
}

func add(a, b int) int {
	return a + b
}

func add1(a, b int) (int, int) {
	return a + b, a + b
}

func add2(a, b int) (sum int) {
	sum = a + b
	return
}

func add3(a, b int) (int, error) {
	var result int
	var err error
	if a < 0 || b < 0 {
		err = errors.New("can not less than o")
	} else {
		result = a + b
	}
	return result, err
}

func add4(a, b int) (result int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("can not less than o")
	} else {
		result = a + b
	}
	return
}

func main() {
	a := test()
	fmt.Println(a)
	fmt.Println(add(1, 2))

	//b, err := add3(3, 0)
	b, err := add4(3, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(b)
	}
}
