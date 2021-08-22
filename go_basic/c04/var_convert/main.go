package main

import "strconv"

func main() {
	var a string = "s"
	data, err := strconv.Atoi(a)
	if err != nil {
		println(data)
	} else {
		println(err)
	}

	var b int = 120
	println(strconv.Itoa(b))
}
