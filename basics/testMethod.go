package main

import "fmt"

type rectangle struct {
	Width  float64
	Length float64
}

// 结构体rectangle的method，可由结构体的实现直接调用
func (r rectangle) area() {
	fmt.Println(r.Length * r.Width)
}

func testMethod() {
	r := rectangle{
		Width:  88.8,
		Length: 100,
	}
	r.area()

}

func main() {
	testMethod()
}
