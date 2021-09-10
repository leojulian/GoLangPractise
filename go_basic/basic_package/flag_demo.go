package main

import (
	"flag"
	"fmt"
)

var (
	para1 string
	para2 string
)

func main() {
	flag.StringVar(&para1, "para1", "test1", "")
	flag.StringVar(&para2, "para2", "test2", "")
	flag.Parse()

	if para1 != "" {
		fmt.Println(para1)
	}

	if para2 != "" {
		fmt.Println(para2)
	}
}
