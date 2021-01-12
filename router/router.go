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

	userRouter := router.Group("/api/v1/users")
	{
		userRouter.Use(middlewares.Auth())
		userRouter.GET("/", e.Wrapper(c.UserListCtrl))
		userRouter.POST("/", e.Wrapper(c.UserAddCtrl))
		userRouter.DELETE("/:userId", e.Wrapper(c.UserDelCtrl))
		userRouter.PUT("/:userId", e.Wrapper(c.UserUpdateCtrl))
		userRouter.GET("/test", e.Wrapper(c.UserTestCtrl))
	}

	appRouter := router.Group("api/v1/apps")
	{
		appRouter.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Apps",
			})
		})
	}

	projectRouter := router.Group("api/v1/projects")
	{
		projectRouter.GET("/", c.ProjectListCtrl)
		projectRouter.POST("/", c.ProjectCreCtrl)
	}

	return router
}
