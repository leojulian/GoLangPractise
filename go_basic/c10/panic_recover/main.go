package main

import (
	"fmt"
	"net/http"
	"time"
)

func div(a, b int) int {
	return a / b
}

func main1() {
	a, b := 1, 0
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("catch error")
	}()
	fmt.Println(div(a, b))
}

func main2() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

//panic 在协程里
func main3() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		panic("error")
		w.Write([]byte("hello world!"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

//panic 在子协程里 会导致主线程抛异常 同时其他协程也会一起挂掉
func main4() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			panic("error")
		}()

		w.Write([]byte("hello world!"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

//panic 在子协程里 然后加入recover
func main5() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		go func() {
			defer func() {
				err := recover()
				fmt.Println(err)
			}()
			panic("error")
		}()

		w.Write([]byte("hello world!"))
	})
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func f1() {
	//a, b := 2, 0
	//fmt.Println(a / b)
	defer func() {
		err := recover()
		fmt.Printf("target error! %v", err)
	}()

	//fmt.Println(2 / 0)

	//go func() {
	//	panic("error")
	//}()
	panic("error")
	time.Sleep(10 * time.Second)
}

func main() {
	f1()
}
