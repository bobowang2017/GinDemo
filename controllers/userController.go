package controllers

import (
	"GinDemo/utils/redis"
	"github.com/gin-gonic/gin"
)

func UserListController(c *gin.Context) {
	//user := Models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//	//db.NewRecord(user) // => 主键为空返回`true`
	//	//db.Create(&user)
	//	//db.NewRecord(user) //
	redis.Set("hello", "wangxiangbo", 100)
	res, _ := redis.Get("bobo")
	redis.Lpush("list", []string{"1", "2", "3", "4"})
	c.JSON(200, gin.H{
		"message": res,
	})
}
