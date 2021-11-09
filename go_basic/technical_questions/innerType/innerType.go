package main

import "fmt"

func main() {
	deferCall()
}

//考察defer函数，先进后出
func deferCall() {
	defer func() {
		fmt.Println("打印前")
	}()
	defer func() {
		fmt.Println("打印中")
	}()
	defer func() {
		fmt.Println("打印后")
	}()
	panic("触发异常!")
	//fmt.Println("打印")
}
