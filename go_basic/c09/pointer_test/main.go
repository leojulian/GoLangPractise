package main

import "fmt"

func main() {
	var ip *int
	a := 10
	//var ip *int = &a
	ip = &a
	fmt.Printf("内存地址是%p \n", ip)
	*ip = 20
	fmt.Printf("内存地址是%p 值是%d\n", ip, *ip)

	arr := [3]int{1, 2, 3}
	var pointer *[3]int = &arr
	fmt.Println(*pointer)

	var n1 *int
	var n2 *int
	var n3 *int
	var arrp []*int = []*int{n1, n2, n3}
	fmt.Println(arrp)
	fmt.Println(len(arrp))
	//
	var p = new(string)
	*p = "124"
	fmt.Println(*p)

	var p1 = new(int)

	fmt.Println(*p1)
}
