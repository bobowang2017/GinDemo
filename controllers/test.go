package controllers

import (
	"GinDemo/utils/logger"
	"github.com/gin-gonic/gin"
)

func TestCtrl(c *gin.Context) {
	//res, _ := redis.Get("bobo")
	//logger.Info(res)
	logger.Info("Hello World")
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
