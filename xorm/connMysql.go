package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"github.com/go-xorm/core"
)

var engine *xorm.Engine

type Test struct {
	testId   string `PK`
	testName string
}

func main() {
	var err error
	engine, err = xorm.NewEngine("mysql", "root:123456@(139.199.102.156:3306)/goTest?charset=utf8")
	if err != nil {
		fmt.Println("数据库连接错误:", err.Error())
	} else {
		fmt.Println("数据库连接成功！")
	}
	engine.ShowSQL(true)
	engine.Logger().SetLevel(core.LOG_DEBUG)
	//engine.SetMapper(core.SameMapper{})

	tableErr := engine.CreateTables(&Test{})
	if tableErr != nil {
		fmt.Println(tableErr.Error())
	}
	exist, e := engine.IsTableExist("test")
	if e == nil {
		fmt.Println(exist)
	} else {
		fmt.Println(e.Error())
	}
}
