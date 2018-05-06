package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		sum := 0
		for i := 0; i < 100; i++ {
			sum += i + 1
		}
		ch <- sum
	}()
	fmt.Print(<-ch)
}
