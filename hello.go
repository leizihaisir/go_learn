package main

import (
	"fmt"
)

//func main() {
//	var balance = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
//	var newArray []int
//	newArray = mysort.BubbleSort(balance)
//	fmt.Println(newArray)
//}

func main() {
	array := [5]int{1: 2, 3: 4}
	modify(array)
	fmt.Println(array)
}
func modify(a [5]int) {
	a[1] = 3
	fmt.Println(a)
}
