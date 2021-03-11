package dao

import (
	"GinDemo/models"
	"GinDemo/utils/logger"
	"errors"
	"fmt"
)

type BaseDao struct {
}

func (b *BaseDao) checkImplTableName(v models.TableNameAble) error {
	_, ok := v.(models.TableNameAble)
	if !ok {
		return errors.New("value doesn't implement TableNameAble")
	}
	return nil
}

func (b *BaseDao) Add(v models.TableNameAble) error {
	if err := b.checkImplTableName(v); err != nil {
		return err
	}
	if err := models.DB.Save(v).Error; err != nil {
		logger.Error("mysql", err, fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
		return err
	}
	return nil
}

func (b *BaseDao) DeleteById(v models.TableNameAble) error {
	if err := b.checkImplTableName(v); err != nil {
		return err
	}
	err := models.DB.Table(v.TableName()).Where(v).Delete(v).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
	}
	return nil
}

func (b *BaseDao) DeleteByCondition(v models.TableNameAble) error {
	if err := b.checkImplTableName(v); err != nil {
		return err
	}
	err := models.DB.Table(v.TableName()).Where(v).Delete(v).Error
	if err != nil {
		return errors.New(fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
	}
	return nil
}

func (b *BaseDao) Update(v models.TableNameAble) error {
	if err := b.checkImplTableName(v); err != nil {
		return err
	}
	return nil
}
