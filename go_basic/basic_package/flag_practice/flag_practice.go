package main

import (
	"flag"
	"fmt"
)

var (
	name  string
	title string
)

//usage:
//go run flag_practice.go -name=leo -title=UFT-RnD
//go run flag_practice.go -h
func main() {
	flag.StringVar(&name, "name", "leo", "help information")
	flag.StringVar(&title, "title", "UFT-RnD", "help information")
	flag.Parse()

	if name != "" {
		fmt.Println(name)
	}

	if title != "" {
		fmt.Println(title)
	}
}
