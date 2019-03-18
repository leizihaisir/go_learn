package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	fmt.Printf("高级打印方法：%v \n", array)
	// 现在的z就变为了动态数组slice，类似于指针。切片操作将数组变为切片
	z := array[1:3]
	z[1] = 111
	fmt.Printf("高级打印方法2：%v \n", array)
	// go语言类型打印
	fmt.Println("z type:", reflect.TypeOf(z))
	fmt.Println("array type:", reflect.TypeOf(array))
}

func printArray(a [10]int) {
	for k, v := range a {
		a[k] = v + 1
		fmt.Printf("索引:%d,值:%d\n", k, v)
	}
}
