package main

import "fmt"

func test(para ...string) {
	for _, v := range para {
		fmt.Println(v)
	}
}

func test1(slice []string) {
	fmt.Println(slice)
}

func filter(slice []int, f func(score int) bool) {
	for _, v := range slice {
		if f(v) {
			fmt.Println(v)
		}
	}
}

type sub func(a, b int) int

//type sex struct {
//	man int,
//	female int,
//	unknown
//}

func main() {
	a := []string{"1", "2", "3"}
	s := append(a, "1", "2")
	fmt.Println(s)
	test(a...)
	test1(a)

	myFunc := test
	fmt.Printf("%T", myFunc)
	myFunc(a...)

	newFunc := func(a, b int) int {
		return a + b
	}

	fmt.Println(newFunc(10, 2))

	newFunc1 := func(a, b int) int {
		return a + b
	}(12, 12)
	fmt.Println(newFunc1)

	var mySub sub = func(a, b int) int {
		return a + b
	}
	fmt.Println(mySub(22, 22))

	scores := []int{20, 30, 60, 70, 80, 90}
	filter(scores, func(score int) bool {
		if score >= 70 {
			return true
		} else {
			return false
		}
	})

	func() {
		fmt.Println("anonymous function!")
	}()
}
