package main

import (
	"fmt"
	"sort"
)

type Course struct {
	Name  string
	Price int
	Url   string
}

type Courses []Course

func (c Courses) Len() int {
	return len(c)
}

func (c Courses) Less(i, j int) bool {
	return c[i].Price < c[j].Price
}

func (c Courses) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	var courses Courses = Courses{
		Course{"python", 300, "baidu"},
		Course{"go", 200, "baidu"},
		Course{"C#", 400, "baidu"},
	}
	sort.Sort(courses)
	for _, v := range courses {
		fmt.Println(v)
	}
	//fmt.Println(courses)
}
