package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestPerson() {
	// p是一个结构体的具体实现
	p := person{
		Name: "ileir",
		Age:  25,
	}

	// p1是一个结构体的具体实现的指针
	p1 := &person{
		Name: "juanjuan",
		Age:  24,
	}

	jsonStr, err := json.Marshal(p)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", string(jsonStr))

	jsonStrP1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%s\n", string(jsonStrP1))
}

type human struct {
	Sex  int
	Name string
	Age  int
}

// go语言没有继承，struct使用的是组合的方式减少重复字段和代码
type teacher struct {
	human
	Course string
}

type student struct {
	human           human
	HeadTeacher     teacher
	LanguageTeacher teacher
}

func testCombiFun() {

	t := teacher{
		human: human{
			Name: "job",
			Sex:  1,
			Age:  32,
		},
		Course: "语文",
	}
	jsonStrT, err := json.Marshal(t)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("teacher%s\n", string(jsonStrT))

	s := student{
		human: human{
			Name: "zihailei",
			Sex:  1,
			Age:  18,
		},
		HeadTeacher: teacher{
			human: human{
				Name: "job",
				Sex:  1,
				Age:  32,
			},
			Course: "数学",
		},
		LanguageTeacher: teacher{
			human: human{
				Name: "job",
				Sex:  1,
				Age:  32,
			},
			Course: "语文",
		},
	}
	jsonStrs, err := json.Marshal(s)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("student%s\n", string(jsonStrs))
}

func main() {
	TestPerson()
	testCombiFun()

}
