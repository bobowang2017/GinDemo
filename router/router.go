package router

import (
	c "GinDemo/controllers"
	"GinDemo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	// 要在路由组之前全局使用「跨域中间件」, 否则OPTIONS会返回404
	router.Use(middlewares.Cors())

	testRouter := router.Group("/api/v1")
	{
		testRouter.GET("/test", c.TestCtrl)
	}

	userRouter := router.Group("/api/v1")
	{
		userRouter.GET("/users", c.UserListCtrl)
		userRouter.POST("/users", c.UserAddCtrl)
		userRouter.DELETE("/users/:userId", c.UserDelCtrl)
		userRouter.PUT("/users/:userId", c.UserUpdateCtrl)
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
		projectRouter.GET("/projects", c.ProjectListCtrl)
		projectRouter.POST("/projects", c.ProjectCreCtrl)
	}

	return router
}
