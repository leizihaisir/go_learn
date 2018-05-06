package main

import "fmt"

func main() {

	c := closure(10)
	fmt.Println(c(2))
	fmt.Println(c(3))
}

// 闭包的使用，返回的是一个函数
func closure(x int) func(int) int {

	return func(y int) int {
		return x + y
	}
}
