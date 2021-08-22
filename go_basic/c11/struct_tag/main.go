package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type info struct {
	Name  string `json:"name""`
	AgeT  int    `json:"age,omitempty"`
	Title string `json:"-"`
}

type ormInfo struct {
	Name  string `orm:"name""`
	AgeT  int    `orm:"age,omitempty"`
	Title string `orm:"-"`
}

func main() {
	//json tag
	i := info{
		Name:  "Leo",
		Title: "Dev",
	}
	v, _ := json.Marshal(i)
	fmt.Println(string(v))

	//反射 取tag
	ormi := ormInfo{
		Name:  "Leo",
		AgeT:  12,
		Title: "Dev",
	}
	t := reflect.TypeOf(ormi)
	fmt.Println(t.Name())
	fmt.Println(t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("orm")
		fmt.Println(i+1, field.Name, tag)
	}
}
