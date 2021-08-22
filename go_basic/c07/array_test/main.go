package main

import "fmt"

func replace(array [3]string) {
	array[0] = "hello"
	fmt.Println(array)
}

func main() {
	//var arr [4]int = [4]int{}
	arr := [4]int{}
	fmt.Println(arr)

	arr1 := [...]int{1, 2, 3, 4}
	arr1[0] = 1
	fmt.Println(arr1)

	//Array is value transfer
	array := [3]string{"1", "2", "3"}
	replace(array)
	fmt.Println(array)

}
