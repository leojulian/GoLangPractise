package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var pointer *int
	fmt.Println(pointer)
	v := 2
	pointer = &v
	fmt.Println(pointer)

	var p *int = new(int)
	*p = 10
	fmt.Println(*p)

	var info map[string]string
	if info == nil {
		fmt.Println("map default value is nil")
	}

	var info2 []int
	if info2 == nil {
		fmt.Println("slice default value is nil")
	}

	var err error
	if err == nil {
		fmt.Println("error default value is nil")
	}

	fmt.Println(unsafe.Sizeof(info), unsafe.Sizeof(info2), unsafe.Sizeof(err))
}
