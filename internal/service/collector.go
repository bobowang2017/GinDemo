package service

import (
	"GinDemo/common/e"
	"GinDemo/internal/dao"
	m "GinDemo/internal/models"
	"errors"
)

type ICollectorService interface {
	AddCollector(c *m.Collector) error
	ListCollector(PageNum, PageSize int, param map[string]interface{}) ([]*m.Collector, error)
	DeleteById(id int) error
}

//用户模块接口实现
type collectorService struct {
	collectorDao *dao.CollectorDao
}

func (c *collectorService) ListCollector(PageNum, PageSize int, param map[string]interface{}) ([]*m.Collector, error) {
	return c.collectorDao.List(PageNum, PageSize, param)
}

func (c *collectorService) AddCollector(m *m.Collector) error {
	//需要对name属性进行唯一性校验
	if c.collectorDao.CheckName(m.Name) {
		return errors.New(e.GetMsg(e.INVALID_PARAMS))
	}
	return c.collectorDao.Add(m)
}

func (c *collectorService) DeleteById(id int) error {
	return c.collectorDao.DeleteById(id)
}

func NewCollectorService() ICollectorService {
	return &collectorService{
		dao.NewCollectorDao(),
	}
}
