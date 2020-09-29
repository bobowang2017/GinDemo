package goMysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	DB, err := gorm.Open("mysql", "root:root@(127.0.0.1:3333)/study?charset=utf8mb4&parseTime=True&loc=Local")
	defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
}
