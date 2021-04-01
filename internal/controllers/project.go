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

// @Title 项目列表
// @Author wangxiangbo
// @Description 查询项目列表
// @Tags Project
// @Param description formData string false "操作描述"
// @Param start_time formData string false "开始时间"
// @Param end_time formData string false "结束时间"
// @Param page formData string true "页数"
// @Param size formData string true "数据条数"
// @Success 200
// @Router	/projects [get]
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
