package main

import (
	"time"
	"fmt"
)

func serial() {
	sum, count, start := 0, 0, time.Now()
	for count < 8e8 {
		sum += count
		count += 1
	}
	end := time.Now()
	fmt.Println(fmt.Sprintf("result is %d, time is %s", sum, end.Sub(start)))
}
func parallel() {
	sum, count, ch, start := 0, 0, make(chan int, 4), time.Now()
	for count < 8 {
		go func(count int) {
			value := 0
			for i := count * 1e8; i < (count+1)*1e8; i++ {
				value += i
			}
			ch <- value
		}(count)
		count++
	}
	for count > 0 {
		sum += <-ch
		count--
	}
	end := time.Now()
	fmt.Println(fmt.Sprintf("result is %d, time is %s", sum, end.Sub(start)))
}

func main() {
	serial()
	parallel()
}
