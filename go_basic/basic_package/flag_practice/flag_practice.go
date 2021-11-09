package main

import (
	"flag"
	"fmt"
)

var (
	para1 string
	para2 string
)

//usage:
//go run flag_practise.go -para1=123 -para2=456
//go run flag_practise.go -h
func main() {
	flag.StringVar(&para1, "para1", "test1", "help information")
	flag.StringVar(&para2, "para2", "test2", "help information")
	flag.Parse()

	if para1 != "" {
		fmt.Println(para1)
	}

	if para2 != "" {
		fmt.Println(para2)
	}
}
