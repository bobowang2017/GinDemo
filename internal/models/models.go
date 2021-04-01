package models

import (
	"GinDemo/common/utils"
	"GinDemo/config"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

type BaseModel struct {
	ID          int            `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedTime utils.JSONTime `gorm:"default:null" json:"createdTime"`
	UpdatedTime utils.JSONTime `gorm:"default:null" json:"updatedTime"`
}

type TableNameAble interface {
	TableName() string
}

// Setup initializes the database instance
func Setup() {
	var err error
	DB, err = gorm.Open(config.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return config.DatabaseSetting.TablePrefix + defaultTableName
	//}
	DB.LogMode(true)
	DB.SingularTable(true)
	DB.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	DB.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer DB.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now()
		if createTimeField, ok := scope.FieldByName("CreatedTime"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("UpdatedTime"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("UpdatedTime", time.Now())
	}
}

// deleteCallback will set `DeletedOn` where deleting
//func deleteCallback(scope *gorm.Scope) {
//	if !scope.HasError() {
//		var extraOption string
//		if str, ok := scope.Get("gorm:delete_option"); ok {
//			extraOption = fmt.Sprint(str)
//		}
//
//		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
//
//		if !scope.Search.Unscoped && hasDeletedOnField {
//			scope.Raw(fmt.Sprintf(
//				"UPDATE %v SET %v=%v%v%v",
//				scope.QuotedTableName(),
//				scope.Quote(deletedOnField.DBName),
//				scope.AddToVars(time.Now().Unix()),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		} else {
//			scope.Raw(fmt.Sprintf(
//				"DELETE FROM %v%v%v",
//				scope.QuotedTableName(),
//				addExtraSpaceIfExist(scope.CombinedConditionSql()),
//				addExtraSpaceIfExist(extraOption),
//			)).Exec()
//		}
//	}
//}

// addExtraSpaceIfExist adds a separator
//func addExtraSpaceIfExist(str string) string {
//	if str != "" {
//		return " " + str
//	}
//	return ""
//}
