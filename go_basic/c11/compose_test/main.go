package main

import "fmt"

//继承
//内嵌结构体

type Teacher struct {
	Name  string
	Age   int
	Title string
}

func (t Teacher) teacherInfo() {
	fmt.Println(t.Name, t.Age, t.Title)
}

type Course struct {
	Teacher
	Name  string
	Price int
	Url   string
}

func (c Course) courseInfo() {
	fmt.Println(c.Name, c.Price, c.Url, c.Teacher.Name, c.Age, c.Title)
}

func main() {
	c := Course{
		Teacher{"leo", 30, "Dev"},
		"Go",
		100,
		"www.baidu.com",
	}
	c.courseInfo()
}
