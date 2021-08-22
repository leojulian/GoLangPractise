package main

import "fmt"

func main() {
	var m11 map[string]string = map[string]string{"name": "leo"}
	fmt.Println(m11)
	//1,
	//var m1 map[string]string = map[string]string{}
	m1 := map[string]string{
		"key1": "value1",
		"key2": "value2",
	}
	fmt.Println(m1)
	//2, make
	m2 := make(map[string]string)
	m2["key3"] = "value3"
	fmt.Println(m2)
	//3
	var m3 = map[string]string{}
	m3["key4"] = "value4"
	fmt.Println(m3)

	var m12 map[string]string
	m12["name"] = "leo"
	fmt.Println(m12)
}
