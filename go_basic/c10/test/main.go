package main

import "fmt"

type Teacher struct {
	name  string
	age   int
	title string
}

type Car interface {
	Drive()
	Wheel()
}

type BMW struct {
}

func (b BMW) Drive() {}
func (b BMW) Wheel() {}

func (t Teacher) PrintInfo() {

}

type Course struct {
	Teacher
	price int
	name  string
	url   string
}

//func (c Course) PrintInfo() {
//
//}

func (c *Course) NewCourse(name string) {
	c.name = name
	getInfo(*c)
}

func getInfo(c Course) {
	fmt.Println(c.name, c.age)
}

func main() {
	var c Course = Course{
		Teacher: Teacher{ //还可以这样声明一些属性值,因为Teacher是结构体,匿名,所以需要这样声明
			"bobby", 18, "",
		},
		price: 100,
		name:  "scrapy分布式爬虫",
		url:   "", // 注意这里的逗号不能少
	}
	c.PrintInfo()
	c.NewCourse("leo")
	getInfo(c)
}
