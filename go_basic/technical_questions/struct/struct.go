package main

import "fmt"

type User struct {
	Age  int
	Name string
}

func main() {
	a1 := User{
		Age:  1,
		Name: "a1",
	}
	fmt.Sprintf("%p\n", &a1)
	a2 := User{
		Age:  2,
		Name: "a2",
	}
	fmt.Printf("%p\n", &a2)
	if a1 == a2 {
		fmt.Println("a1和a2相等")
	} else {
		fmt.Println("a1和a2不相等")
	}
}
