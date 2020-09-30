package router

import (
	"GinDemo/controllers"
	"GinDemo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())
	userRouter := router.Group("/api/v1")
	{
		userRouter.GET("/user", controllers.UserListController)
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
		projectRouter.GET("/projects", controllers.ProjectListController)
		projectRouter.POST("/projects", controllers.ProjectCreController)
	}

	return router
}
