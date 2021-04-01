package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method
		//origin := c.Request.Header.Get("Origin")
		token := c.Request.Header.Get("token")
		fmt.Println("Hello World", token)
		//if token != "nil" {
		//	fmt.Println("Hello world")
		//} else {
		//	c.JSON(200, gin.H{
		//		"status": 401,
		//		"msg":    "未认证",
		//		"data":   "failure",
		//	})
		//	c.Abort()
		//}
		c.Next()
	}
}
