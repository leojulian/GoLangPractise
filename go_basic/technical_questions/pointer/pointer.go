package main

import "fmt"

func main() {
	var a *int = nil
	f(a)
}

func f(a *int) {
	if a == nil {
		fmt.Println("a")
	} else {
		fmt.Println("b")
	}
}
