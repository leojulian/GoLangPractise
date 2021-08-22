// example3.go
package main

import "math/rand"

func fa() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	index := rand.Intn(7)
	_ = s[:index] // 第9行： 需要边界检查
	_ = s[index:] // 第10行： 边界检查消除了！
}

func fb(s []int, index int) {
	_ = s[:index] // 第14行： 需要边界检查
	_ = s[index:] // 第15行： 需要边界检查（不够智能？）
}

func fc() {
	s := []int{0, 1, 2, 3, 4, 5, 6}
	s = s[:4]
	index := rand.Intn(7)
	_ = s[:index] // 第22行： 需要边界检查
	_ = s[index:] // 第23行： 需要边界检查（不够智能？）
}

func main() {}
