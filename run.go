package main

import (
	"GinDemo/apis/users"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// Set example variable
		c.Set("example", "12345")
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Print(latency)
		// access the status we are sending
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())
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
