package main

func main() {
	//for loop
	//for i := 0; i < 10; i++ {
	//	println(i)
	//}

	//while
	//var i int = 0
	//for i < 10 {
	//	println(i)
	//	i++
	//}

	//loop string
	//var str string = "name"
	// method 1
	//for i := 0; i < len(str); i++ {
	//	println(str[i])
	//}
	// method 2
	//for index, value := range str {
	//	//fmt.Printf("%d %d \n", index, value)
	//	//println(index, value)
	//	fmt.Printf("%d %c\n", index, value)
	//}

	//Chinese
	var name = "leo利奥"
	//nameArr := []rune(name)
	//for index := 0; index < len(nameArr); index++ {
	//	fmt.Printf("%d %c\n", index, nameArr[index])
	//}

	//e:
	//
	//	goto e

	switch name {
	case "leo", "leo利奥1":
		println(name)
	default:
	case "leo利奥":
		println(name)
	}

}
