package models

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
	DB.DB().SetMaxIdleConns(30)
	DB.DB().SetMaxOpenConns(30)
	/**
	      禁用表名复数>
	     !!!如不禁用则会出现表 y结尾边ies的问题
	      !!!如果只是部分表需要使用源表名，请在实体类中声明TableName的构造函数
	  ```
	      func (实体名) TableName() string {
	          return "数据库表名"
	      }
	  ```
	*/
	DB.SingularTable(true)
}

func Create(o interface{}) {
	DB.Create(o)
}
