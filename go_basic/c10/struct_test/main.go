package main

import "fmt"

//封装

type Course struct {
	Name  string
	Price int
	Url   string
}

func (c Course) printInfo() {
	fmt.Println(c.Name, c.Price, c.Url)
}

func NewCourse() *Course {
	return &Course{}
}

type GoCourse struct {
	Course
	Name  string
	Price int
	Url   string
}

func (c GoCourse) printInfo() {
	fmt.Println(c.Name, c.Price, c.Url)
}

func main() {
	// 1
	var course Course = Course{
		Name:  "go",
		Price: 100,
		Url:   "www.baidu.com",
	}
	fmt.Println(course)

	// 2
	course1 := Course{"go", 100, "www.baidu.com"}
	fmt.Println(course1)

	// 3 指针
	c2 := &Course{"go", 100, "www.baidu.com"}
	fmt.Printf("%T\n", c2)
	fmt.Println(c2.Name, c2.Price, c2.Url) // struct 语法糖 (*c2).Name => c2.Name

	// 4
	c3 := Course{}
	fmt.Println(c3.Price)

	// 5
	var c4 = &Course{} //new(Course)
	fmt.Printf("%T\n", c4)

	// 6
	c5 := Course{"leo", 180, "leo.com"}
	c6 := c5
	fmt.Printf("%p %p\n", &c5, &c6)
	//fmt.Println(&c6)
	s := "124"
	s1 := s
	fmt.Printf("%p %p\n", &s, &s1)
	//fmt.Println(s, s1)
	//s = "456"
	//fmt.Println(s, s1)

	a := [...]int{1, 2, 3}
	a1 := a
	fmt.Printf("%p %p\n", &a, &a1)
	a[0] = 111
	fmt.Println(a, a1)

	sl := []int{1, 2, 3}
	sl1 := sl
	fmt.Printf("%p %p\n", &sl, &sl1)
	sl[0] = 10
	fmt.Println(sl, sl1)

	//struct method
	c7 := Course{"leo", 180, "leo.com"}
	c7.printInfo()

	//constructor
	NewCourse()

}
