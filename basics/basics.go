package main

import "fmt"

var a = "菜鸟教程"
var b = "runoob.com"

const (
	weigth  = 100
	Unknown = 0
	Female  = 1
	Male    = 2
)

func main() {
	c := false
	d := true
	fmt.Println(a, b, c, d)
	fmt.Printf("weigth %v\n", weigth)
	fmt.Printf("Unknown value: %d,Female value: %d,Male value: %d\n", Unknown, Female, Male)

	var count = 100
	var ptr = &count

	fmt.Println(&count)
	fmt.Println(*ptr)

	var m = 2
	var n = 16

	var x = m << 2
	fmt.Println(x)
	fmt.Println(n >> 2)
	fmt.Println(judgeSex(Female))

}

func judgeSex(sex int) string {

	var sexChinese string
	if sex == 0 {
		sexChinese = "中性"
	} else if sex == 1 {
		sexChinese = "男"
	} else {
		sexChinese = "女"
	}

	return sexChinese
}
