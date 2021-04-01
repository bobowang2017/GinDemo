package controllers

import (
	"GinDemo/common/e"
	"GinDemo/internal/dto"
	"GinDemo/internal/models"
	"GinDemo/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CollectorController struct {
	collectorService service.ICollectorService
}

func CollectorRouterRegister(router *gin.RouterGroup) {
	collector := CollectorController{
		service.NewCollectorService(),
	}
	router.GET("/", e.Wrapper(collector.CollectorList))
	router.POST("/", e.Wrapper(collector.CollectorAdd))
	router.DELETE("/:id", e.Wrapper(collector.CollectorDelete))
	//router.PUT("/:userId", e.Wrapper(user.UserUpdate))
	//router.GET("/:userId", e.Wrapper(user.UserDetail))
}

// @Title 采集器
// @Description 查询采集器列表
// @Tags Collector
// @Param PageNum query int false "页数"
// @Param PageSize query int false "页大小"
// @Success 200
// @Router	/api/v1/collectors [get]
func (p *CollectorController) CollectorList(c *gin.Context) interface{} {
	pageNum := c.DefaultQuery("PageNum", "1")
	pageSize := c.DefaultQuery("PageSize", "10")
	num, _ := strconv.Atoi(pageNum)
	size, _ := strconv.Atoi(pageSize)
	params := make(map[string]interface{})
	res, err := p.collectorService.ListCollector(num, size, params)
	if err != nil {
		return e.ServerError(err.Error())
	}
	return res
}

// @Title 采集器
// @Description 创建采集器
// @Tags Collector
// @Param body body dto.CollectorDto true "创建参数"
// @Success 200
// @Router	/api/v1/collectors [post]
func (p *CollectorController) CollectorAdd(c *gin.Context) interface{} {
	cDto := &dto.CollectorDto{}
	if err := c.BindJSON(cDto); err != nil {
		return e.ParameterError(err.Error())
	}
	collector := &models.Collector{
		Name:                 cDto.Name,
		Description:          cDto.Description,
		PlanId:               &cDto.PlanId,
		ProjectId:            &cDto.ProjectId,
		ProjectCode:          cDto.ProjectCode,
		EnvId:                &cDto.EnvId,
		EnvName:              cDto.EnvName,
		Cluster:              cDto.Cluster,
		TokenUrl:             cDto.TokenUrl,
		AppId:                &cDto.AppId,
		TokenRefreshInterval: &cDto.TokenRefreshInterval,
		MetricUrl:            cDto.MetricUrl,
		AlertUrl:             cDto.AlertUrl,
		Enabled:              new(int),
		HeartBeatUrl:         cDto.HeartBeatUrl,
		HeartBeatInterval:    &cDto.HeartBeatInterval,
	}
	err := p.collectorService.AddCollector(collector)
	if err != nil {
		return e.ParameterError(err.Error())
	}
	return "success"
}

func (p *CollectorController) CollectorDelete(c *gin.Context) interface{} {
	strId := c.Param("id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		return e.ParameterError(err.Error())
	}
	err = p.collectorService.DeleteById(id)
	if err != nil {
		return e.ServerError(err.Error())
	}
	return "success"
}
