package controllers

import (
	"GinDemo/common/utils/redis"
	"github.com/gin-gonic/gin"
)

func TestCtrl(c *gin.Context) {
	var slice = []int{1, 2, 3, 4, 5}
	slice[6] = 6
	_, _ = redis.Get("bobo")

	c.JSON(200, gin.H{
		"message": "hello",
	})
}
