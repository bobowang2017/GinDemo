package models

import (
	"time"
)

type User struct {
	BaseModel
	Username string `gorm:"size:64;not null" json:"username" binding:"required"` // string默认长度为255, 使用这种tag重设。
	Password string `gorm:"size:64;not null" json:"password" binding:"required,min=6,max=12"`
	Name     string `gorm:"size:16;default:'hello world'" json:"name"`
	Sex      int    `gorm:"default:0"`
	Birthday time.Time
}

func AddUser(user *User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func ListUser(PageNum, PageSize int, params map[string]interface{}) ([]*User, error) {
	var users []*User
	err := db.Offset((PageNum - 1) * PageSize).Limit(PageSize).Where(params).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func DeleteById(userId int) error {
	user := User{}
	user.BaseModel.ID = userId
	err := db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateById(userId int, params map[string]interface{}) error {
	user := User{}
	user.BaseModel.ID = userId
	err := db.Model(&user).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}

func SearchUser(params map[string]interface{}) ([]*User, error) {
	var users []*User
	err := db.Where(params).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
