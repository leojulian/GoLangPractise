package main

import "fmt"

func test(i interface{}) {
	// 1,
	//v, ok := i.(int)
	//if ok {
	//	// v is int
	//	fmt.Printf("%d\n", v)
	//}
	//s, ok := i.(string)
	//if ok {
	//	// v is string
	//	fmt.Printf("%s\n", s)
	//}

	//2,
	//if v, ok := i.(int); ok {
	//	// v is int
	//	fmt.Printf("%d\n", v)
	//}
	//
	//if s, ok := i.(string); ok {
	//	// v is string
	//	fmt.Printf("%s\n", s)
	//}

	//3,
	switch v := i.(type) {
	case int:
		fmt.Printf("%d\n", v)
	case string:
		fmt.Printf("%s\n", v)
	}
}

func main() {
	//1, empty interface可以接任何类型的值
	var i interface{}
	i = []int{}
	fmt.Println(i)
	//2, 作为参数传递
	test("test")
	test(111)
	//作为map的value的值
	var teacher = make(map[string]interface{})
	teacher["name"] = "leo"
	teacher["sex"] = "man"
	teacher["age"] = 30
	teacher["courses"] = []string{"go", "python"}
	fmt.Println(teacher)

}
