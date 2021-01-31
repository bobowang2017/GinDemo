package controllers

import (
	"github.com/gin-gonic/gin"
)

type ProjectController struct {
}

func ProjectRouterRegister(router *gin.RouterGroup) {
	project := ProjectController{}
	router.POST("/", project.ProjectCre)
	router.GET("/", project.ProjectList)
}

func (p *ProjectController) ProjectList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Project List",
	})
}

func (p *ProjectController) ProjectCre(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Projects",
	})
}
