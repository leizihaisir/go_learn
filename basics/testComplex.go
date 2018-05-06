package main

import "fmt"

func main() {
	// go可直接定义复数
	var x complex64 = 10 + 9i
	fmt.Print(x + x)
}
