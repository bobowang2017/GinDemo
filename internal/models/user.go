package models

import (
	"GinDemo/common/utils"
)

// User   用户对象Model信息
type User struct {
	BaseModel
	Username string          `gorm:"size:64;not null" json:"username" `
	Password string          `gorm:"size:64;not null" json:"password" `
	Name     string          `gorm:"size:16;default:'hello world'" json:"name"`
	Sex      *int            `gorm:"default:0" json:"sex"`
	Birthday *utils.JSONTime `json:"birthday"`
}

func (u *User) TableName() string {
	return "user"
}

// UserAddress   用户收件地址对象Model信息
type UserAddress struct {
	BaseModel
	UserId    int    `gorm:"size:64;not null" json:"userId" `
	Consignee string `gorm:"size:64;not null" json:"consignee" `
	Province  int    `gorm:"default:0;not null" json:"province"`
	City      int    `gorm:"default:0;not null" json:"city"`
	District  int    `gorm:"not null" json:"district"`
	Street    string `gorm:"not null" json:"street"`
	Zipcode   string `json:"zipcode"`
	Mobile    string `gorm:"not null" json:"mobile"`
	Default   bool   `json:"default"`
}
