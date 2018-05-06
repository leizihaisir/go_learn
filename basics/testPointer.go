package main

import "fmt"

func change(b, a *int) {
	*a = *a + 1
	*b = *b + *a
}

func main() {
	a := 2
	b := 4
	change(&b, &a)
	fmt.Println(a)
	fmt.Println(b)

}
