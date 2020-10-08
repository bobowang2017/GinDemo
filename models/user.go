package models

import (
	"time"
)

type User struct {
	ID          int    `gorm:"AUTO_INCREMENT"`   // 自增
	Username    string `gorm:"size:64;not null"` // string默认长度为255, 使用这种tag重设。
	Password    string `gorm:"size:64;not null"`
	Name        string `gorm:"size:16;not null"`
	Sex         int    `gorm:"default:0"`
	Birthday    time.Time
	CreatedTime *time.Time `gorm:"default:null"`
	UpdatedTime time.Time
}

func AddUser(data map[string]interface{}) error {
	article := User{
		Username: data["Username"].(string),
		Password: data["Password"].(string),
		Name:     data["Name"].(string),
		Sex:      data["Sex"].(int),
		Birthday: data["Birthday"].(time.Time),
	}
	if err := db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func ListUser(PageNum, PageSize int, params map[string]string) ([]*User, error) {
	var users []*User
	err := db.Find(&users).Offset((PageNum - 1) * PageSize).Limit(PageNum).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteById(userId int) error {
	err := db.Delete(&User{ID: userId}).Error
	if err != nil {
		return err
	}
	return nil
}
