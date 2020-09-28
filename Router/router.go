package Router

import (
	"GinDemo/Controllers"
	"GinDemo/Middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	router := gin.New()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(Middlewares.Cors())
	userRouter := router.Group("/api/v1")
	{
		userRouter.GET("/user", Controllers.UserListController)
	}

	appRouter := router.Group("api/v1")
	{
		appRouter.GET("/app", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Apps",
			})
		})
	}

	projectRouter := router.Group("api/v1")
	{
		projectRouter.GET("/project", Controllers.ProjectListController)
	}

	router.Run(":8080")
}
