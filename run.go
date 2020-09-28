package main

import (
	"GinDemo/apis/users"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	userRouter := router.Group("/api/v1")
	{
		userRouter.GET("/user", users.UserList)
	}
	appRouter := router.Group("api/v1")
	{
		appRouter.GET("/app", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Apps",
			})
		})
	}
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})
	router.Run()
}
