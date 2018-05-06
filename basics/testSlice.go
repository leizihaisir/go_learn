package main

import "fmt"

func main() {
	s := make([]int, 3, 10)
	fmt.Printf("s切片容量%d\n", cap(s))
	fmt.Printf("%p\n", s)
	fmt.Println(s)
	for k := range s {
		s[k] = k + 1
	}

	// 没有超过容量，底层数组不会改变
	s = append(s, 1)
	fmt.Printf("s切片容量%d\n", cap(s))
	fmt.Printf("%p\n", s)
	fmt.Println(s)

	// 超过切片原始容量，底层数组已经改变
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Printf("s切片容量%d\n", cap(s))
	fmt.Printf("%p\n", s)
	fmt.Println(s)

}
