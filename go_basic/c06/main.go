package main

import "strings"

func main() {
	var name string = "leo李华东"
	println(len(name))
	//nameArr := []rune(name)
	//println(len(nameArr))
	//date := `2020\06\08`
	//print(date)

	println(strings.Contains(name, "leo"))
	println(strings.Index(name, "q"))
	println(strings.Count(name, "l"))
	println(strings.HasPrefix(name, "leo"))
	println(strings.HasSuffix(name, "leo"))

	var name1 string = " 123 "
	name1 = strings.TrimSpace(name1)
	println(name1)

	println(strings.Replace(name, "l", "L", 1))
}
