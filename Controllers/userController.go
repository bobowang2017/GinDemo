package Controllers

import "github.com/gin-gonic/gin"

func UserListController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Apps",
	})
}
