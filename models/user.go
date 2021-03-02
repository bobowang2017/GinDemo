package models

import (
	"time"
)

type User struct {
	BaseModel
	Username string     `gorm:"size:64;not null" json:"username" ` // string默认长度为255, 使用这种tag重设。
	Password string     `gorm:"size:64;not null" json:"password" `
	Name     string     `gorm:"size:16;default:'hello world'" json:"name"`
	Sex      *int       `gorm:"default:0" json:"sex"`
	Birthday *time.Time `json:"birthday"`
}
