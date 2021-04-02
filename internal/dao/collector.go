package dao

import m "GinDemo/internal/models"

type CollectorDao struct {
}

func NewCollectorDao() *CollectorDao {
	return &CollectorDao{}
}

func (c *CollectorDao) Add(collector *m.Collector) error {
	if err := m.DB.Create(collector).Error; err != nil {
		return err
	}
	return nil
}

func (c *CollectorDao) CheckName(name string) bool {

	var total int
	if err := m.DB.Model(&m.Collector{}).Select("name").Where("name=?", name).Count(&total).Error; err != nil {
		return true
	}
	return total > 0
}

func (c *CollectorDao) List(PageNum, PageSize int, params map[string]interface{}) ([]*m.Collector, error) {
	var collectors []*m.Collector
	err := m.DB.Offset((PageNum - 1) * PageSize).Limit(PageSize).Where(params).Find(&collectors).Error
	if err != nil {
		return nil, err
	}
	return collectors, nil
}

func (c *CollectorDao) DeleteById(id int) error {
	collector := m.Collector{}
	collector.BaseModel.ID = id
	err := m.DB.Delete(&collector).Error
	if err != nil {
		return err
	}
	return nil
}
