package router

import (
	"GinDemo/common/e"
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
		userRouter.Use(middlewares.Auth())
		userRouter.GET("/users", e.Wrapper(c.UserListCtrl))
		userRouter.POST("/users", e.Wrapper(c.UserAddCtrl))
		userRouter.DELETE("/users/:userId", e.Wrapper(c.UserDelCtrl))
		userRouter.PUT("/users/:userId", e.Wrapper(c.UserUpdateCtrl))
		userRouter.GET("/users/test", e.Wrapper(c.UserTestCtrl))
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
