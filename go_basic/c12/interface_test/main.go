package main

import (
	"errors"
	"fmt"
)

type UIDesigner interface {
	Design()
}

type Programmer interface {
	//UIDesigner
	Coding() string
	Debugging() string
}

type Pythoner struct {
}

func (p Pythoner) Coding() string {
	return "Coding Python"
}

func (p Pythoner) Debugging() string {
	return "Debugging Pyhon"
}

type Golanger struct {
}

func (g Golanger) Coding() string {
	return "Coding Golang"
}

func (g Golanger) Debugging() string {
	return "Debugging Golang"
}

type MyError struct{}

func (m MyError) Error() string {
	return "My Error happen"
}

func main() {
	var pro Programmer = Pythoner{}
	fmt.Println(pro.Coding(), pro.Debugging())

	var pros = []Programmer{}
	pros = append(pros, Pythoner{}, Golanger{})
	fmt.Println(pros)
	var err error = MyError{}
	fmt.Print(err)

	var err1 error = errors.New("")
	fmt.Println(err1)

	var err2 = fmt.Errorf("%s", "Not found")
	fmt.Println(err2)
}
