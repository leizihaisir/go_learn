package main

import (
	"reflect"
	"fmt"
	"time"
)

type User struct {
	Name string
	Age  int
}

func testMakeFunc(count int) {
	sum := 0
	for i := 0; i < count; i++ {
		sum += 1
	}
	fmt.Println(sum)
}

func main() {
	user := User{
		"zihailei",
		25,
	}
	// 需要传入实例
	userType := reflect.TypeOf(user)
	fmt.Println(userType.Name())
	fmt.Println(userType.Kind())
	firstField := userType.Field(0)
	fmt.Println(firstField.Name, firstField.Type)
	userValue := reflect.ValueOf(&user).Elem()
	userValue.FieldByName("Name").SetString("rose")
	userValue.FieldByName("Age").SetInt(18)
	fmt.Println(user)
	userTypeNew := reflect.TypeOf(user)
	newUser := reflect.New(userTypeNew)
	newUser.Elem().FieldByName("Name").SetString("rose")
	newUser.Elem().FieldByName("Age").SetInt(10)
	fmt.Println(newUser)

	// 函数反射
	funcType := reflect.TypeOf(testMakeFunc)
	funcValue := reflect.ValueOf(testMakeFunc)
	newFunc := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()
		out := funcValue.Call(args)
		end := time.Now()
		fmt.Println(end.Sub(start))
		return out
	})
	var count int = 1e8
	newFunc.Call([]reflect.Value{reflect.ValueOf(count)})
}
