package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

func GetDB() *gorm.DB {
	return getConn()
}

func getConn() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@(127.0.0.1:3333)/study?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
	}
	db.DB().SetMaxIdleConns(3000)
	db.DB().SetMaxOpenConns(3000)
	db.DB().SetConnMaxLifetime(time.Second * 300)
	db.SingularTable(true)
	if err := db.DB().Ping(); err != nil {
		db.Close()
		fmt.Println(err)
	}
	return db
}
