package controllers

import (
	"github.com/gin-gonic/gin"
)

func ProjectListCtrl(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "",
	})
}

func ProjectCreCtrl(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Projects",
	})
}
