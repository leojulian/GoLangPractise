package main

import "fmt"

func replace(slce []int) {
	slce[0] = 21
	fmt.Println(slce)
}

func main() {

	//slice
	var nameArr = [4]string{"a", "b", "c", "d"}
	var nameArrSlice = nameArr[0:3]
	fmt.Printf("nameArr is %v\n", nameArr)
	fmt.Println(nameArrSlice)
	var nameSlice = make([]string, 4) //[]string{}
	copy(nameSlice, nameArrSlice)
	fmt.Println(nameSlice)
	//delete 'b'
	deleteNameSlice := append(nameArrSlice[:1], nameArrSlice[2:]...)
	fmt.Println(deleteNameSlice)

	//slice is reference transfer
	b := []int{1, 2, 3}
	fmt.Printf("Orignal is %v\n", b)
	c := b[:]
	c[0] = 2
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	//Modify the transferred slice will make change for the original slice too
	d := []int{1, 2, 3}
	replace(d)
	fmt.Println(d)

	slice := make([]int, 5, 16)
	fmt.Println(slice)

	data := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%T\n", data)
	dataSlice := data[2:4]
	dataSlice[0] = 6
	fmt.Println(data)
	fmt.Println(dataSlice)
	fmt.Printf("%T\n", data)
	fmt.Printf("%T\n", dataSlice)

	//切片有不同的初始化方式
	//1, make
	slice1 := make([]int, 5)
	fmt.Printf("length is %d capacity is %d \n", len(slice1), cap(slice1))

	//2, slice from array
	array1 := [5]int{1, 2, 3, 4, 5}
	slice2 := array1[:4]
	//delete index:2 value:3
	slice2 = append(slice2[0:2], slice2[3:]...)
	fmt.Printf("length is %d capacity is %d \n", len(slice2), cap(slice2))

	//3,
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Printf("length is %d capacity is %d \n", len(slice3), cap(slice3))

	//slice4 := make([]int, 3)
	var slice4 []int
	slice4 = append(slice4, 5)
	fmt.Println(slice4)
}
