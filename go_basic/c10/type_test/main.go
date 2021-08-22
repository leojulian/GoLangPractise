package main

import "fmt"

type myFunc func(str string)

func main() {

	// 1, alias 为别的类型起别名
	type myByte = byte
	var a myByte
	fmt.Printf("%T\n", a)

	// 2, 基于一个已有的类型定义一个新的类型，类似于继承？
	type myType int
	var b myType
	fmt.Printf("%T\n", b)

	// 3, 定义结构体
	type myStruct struct {
		name string
		sex  string
	}

	var c myStruct = myStruct{name: "123", sex: "man"}
	fmt.Printf("%T\n", c)

	// 4,
	type myInterface interface {
	}
	// 5, 定义函数
	var e myFunc
	fmt.Printf("%T\n", e)
}
