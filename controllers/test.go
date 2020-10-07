package controllers

import (
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
	"github.com/gin-gonic/gin"
)

func TestCtrl(c *gin.Context) {
	res, _ := redis.Get("bobo")
	logger.Info(res)
	logger.Info("bobwang")
	c.JSON(200, gin.H{
		"message": "hello",
	})
}
