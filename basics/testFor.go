package main

import "fmt"

func main() {
	var b = 15
	var a int
	for i := 0; i < 10; i++ {
		a++
		fmt.Printf("i得值为：%d\n", a)
	}

	for a < b {
		a++
		fmt.Printf("i二次for得值为：%d\n", a)
	}

	numbers := [...]int{1, 2, 3, 4}
	for i, x := range numbers {
		fmt.Printf("numbers[%d]得值为：%d\n", i, x)
	}
}
