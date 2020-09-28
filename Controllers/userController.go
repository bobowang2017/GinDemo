package Controllers

import (
	"github.com/gin-gonic/gin"
)

func UserListController(c *gin.Context) {
	//user := Models.User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	//	//db.NewRecord(user) // => 主键为空返回`true`
	//	//db.Create(&user)
	//	//db.NewRecord(user) //
	c.JSON(200, gin.H{
		"message": "Apps",
	})
}
