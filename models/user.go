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
