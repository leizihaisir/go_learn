package main

import "fmt"

const (
	a1 = 'B'
	b1 = iota
	c
	d
)

func main() {
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(^4)
}
