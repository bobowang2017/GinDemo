package Controllers

import "github.com/gin-gonic/gin"

func ProjectListController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Projects",
	})
}
