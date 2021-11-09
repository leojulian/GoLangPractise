package main

import "fmt"

type user struct{}

func main() {
	a := &user{}
	b := &user{}
	//fmt.Printf("%p\n", &a)
	//fmt.Printf("%p\n", &b)
	//fmt.Println(a, b, a == b)
	println(a, b, a == b) //之所以是false，go在编译器代码优化阶段对其（a==b）直接优化，返回false

	c := &user{}
	d := &user{}
	//fmt.Println(c, d, c == d)
	fmt.Println(c, d) //fmt.Println会导致指针逃逸
	println(c, d, c == d)
}

//go run -gcflags="-m -l" memoryEscape.go
