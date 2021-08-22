package main

import "fmt"

func main() {
	//var name string = "leo"
	//var age int = 30
	//fmt.Printf("%v is %v years old\n", name, age)
	//fmt.Printf("%s is %d years old\n", name, age)
	//fmt.Printf("%s is %-2dyears old\n", name, age)
	var name1 string
	var age1 int

	//fmt.Scan(&name1)
	println("Please input your name:")
	//fmt.Scanln(&name1)
	fmt.Scanf("%s %d", &name1, &age1)
	fmt.Println(name1, age1)

}
