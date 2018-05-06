package main

import (
	"fmt"
)

func main() {
	var a animal
	var c cat
	a = &c
	a.printInfo()
	var d dog
	a = d
	a.printInfo()

	var m cat
	invoke(&m)
}

type animal interface {
	printInfo()
}
type cat int
type dog int

func (c *cat) printInfo() {
	fmt.Println("c cat")
}
func (d dog) printInfo() {
	fmt.Println("d dog")
}

// 实体类型以值接收者实现接口的时候，不管是实体类型的值，还是实体类型值的指针，都实现了该接口。
// 如果是值接收者，实体类型的值和指针都可以实现对应的接口；如果是指针接收者，那么只有类型的指针能够实现对应的接口。
func invoke(a animal) {
	a.printInfo()
}
