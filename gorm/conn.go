package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var db *gorm.DB

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(139.199.102.156:3306)/goTest?charset=utf8&parseTime=true")
	if err != nil {
		panic("连接数据库失败")
	} else {
		fmt.Println("连接数据库成功")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)
}

func main() {
	if db == nil {
		Init()
	}
	var product Product
	db.First(&product, 1)
	fmt.Println(product.Price)

}
