package dao

//
//import (
//	"GinDemo/models"
//	"GinDemo/utils/logger"
//	"errors"
//	"fmt"
//	"reflect"
//)
//
//type BaseDao struct {
//}
//
//func (b *BaseDao) checkImplTableName(v models.TableNameAble) error {
//	_, ok := v.(models.TableNameAble)
//	if !ok {
//		return errors.New("value doesn't implement TableNameAble")
//	}
//	return nil
//}
//
//func (b *BaseDao) checkId(v models.TableNameAble) bool {
//	return reflect.ValueOf(v).Elem().FieldByName("ID").IsValid()
//}
//
//func (b *BaseDao) Add(v models.TableNameAble) error {
//	if err := b.checkImplTableName(v); err != nil {
//		return err
//	}
//	if err := models.DB.Save(v).Error; err != nil {
//		logger.Error("mysql", err, fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
//		return err
//	}
//	return nil
//}
//
//func (b *BaseDao) DeleteByIds(v models.TableNameAble) error {
//	if err := b.checkImplTableName(v); err != nil {
//		return err
//	}
//	if !b.checkId(v) {
//		return errors.New(fmt.Sprintf("No ID Exist in %s", v.TableName()))
//	}
//	err := models.DB.Table(v.TableName()).Where(v).Delete(v).Error
//	if err != nil {
//		return errors.New(fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
//	}
//	return nil
//}
//
//func (b *BaseDao) DeleteByCondition(v models.TableNameAble) error {
//	if err := b.checkImplTableName(v); err != nil {
//		return err
//	}
//	err := models.DB.Table(v.TableName()).Where(v).Delete(v).Error
//	if err != nil {
//		return errors.New(fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
//	}
//	return nil
//}
//
//func (b *BaseDao) UpdateById(v models.TableNameAble, param map[string]interface{}) error {
//	if err := b.checkImplTableName(v); err != nil {
//		return err
//	}
//	if !b.checkId(v) {
//		return errors.New(fmt.Sprintf("No ID Exist in %s", v.TableName()))
//	}
//	err := models.DB.Table(v.TableName()).Where(v).Update(param).Error
//	if err != nil {
//		return errors.New(fmt.Sprintf("Failed to save %s, the value is %+v", v.TableName(), v))
//	}
//	return nil
//}
