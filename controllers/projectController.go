package controllers

import (
	"GinDemo/models"
	"github.com/gin-gonic/gin"
	"time"
)

func ProjectListController(c *gin.Context) {
	users := models.GetDB().Find(&models.User{})
	defer models.GetDB().Close()
	c.JSON(200, gin.H{
		"message": users,
	})
}

func ProjectCreController(c *gin.Context) {
	user := models.User{Name: "WangBoBo", Sex: 18, Birthday: time.Now()}
	models.GetDB().Create(&user)
	defer models.GetDB().Close()
	c.JSON(200, gin.H{
		"message": "Projects",
	})
}
