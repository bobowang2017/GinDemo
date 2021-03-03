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

	userRouter := router.Group("/api/v1/users")
	{
		//userRouter.Use(middlewares.Auth())
		c.UserRouterRegister(userRouter)
	}

	projectRouter := router.Group("api/v1/projects")
	{
		c.ProjectRouterRegister(projectRouter)
	}

	return router
}
