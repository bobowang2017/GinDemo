package controllers

import (
	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	//res, _ := redis.Get("bobo")
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
