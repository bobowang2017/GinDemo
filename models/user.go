package models

import (
	"GinDemo/utils"
)

type User struct {
	BaseModel
	Username string          `gorm:"size:64;not null" json:"username" `
	Password string          `gorm:"size:64;not null" json:"password" `
	Name     string          `gorm:"size:16;default:'hello world'" json:"name"`
	Sex      *int            `gorm:"default:0" json:"sex"`
	Birthday *utils.JSONTime `json:"birthday"`
}
