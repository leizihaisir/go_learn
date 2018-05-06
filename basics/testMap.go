package main

import "fmt"

func main() {

	// 1：make函数创建map
	var dict = make(map[string]string)
	dict["uasename"] = "zihailei"
	dict["password"] = "123456"

	for k, v := range dict {
		fmt.Printf("key:%s,valve:%s\n", k, v)
	}

	// 2：map字面量的方式创map
	dict0 := map[string]int{"ileir": 90, "obama": 10}

	for k, v := range dict0 {
		fmt.Printf("map_key:%s,map_valve:%d\n", k, v)
	}

	//创建一个nil的Map
	var dict1 map[string]int
	//使用nil的Map必须先要初始化
	dict1 = make(map[string]int)
	dict1["张三"] = 43
	fmt.Println(dict1)

	// 两个map交换key和value
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

}
