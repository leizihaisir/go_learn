package main

import "fmt"

func main() {
	fmt.Println(getGrade(69))
}

func getGrade(marks int) string {

	var grade string
	switch {
	case marks > 90:
		grade = "A"
	case marks > 80:
		grade = "B"
	case marks > 60:
		grade = "C"
	default:
		grade = "D"
	}
	return grade
}
