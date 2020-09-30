package models

import "time"

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
