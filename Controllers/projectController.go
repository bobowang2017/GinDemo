package Controllers

import (
	"GinDemo/Models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
)

func ProjectListController(c *gin.Context) {
	user := Models.User{Name: "WangBoBo", Sex: 18, Birthday: time.Now()}
	DB, err := gorm.Open("mysql", "root:root@(127.0.0.1:3333)/study?charset=utf8mb4&parseTime=True&loc=Local")
	defer DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	DB.Create(&user)
	c.JSON(200, gin.H{
		"message": "Projects",
	})
}
