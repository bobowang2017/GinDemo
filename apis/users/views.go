package users

import "github.com/gin-gonic/gin"

func UserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Users",
	})
}
