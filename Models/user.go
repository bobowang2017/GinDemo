package Models

import "time"

type User struct {
	Id          int    `gorm:"AUTO_INCREMENT"` // 自增
	Username    string `gorm:"size:64"`        // string默认长度为255, 使用这种tag重设。
	Password    string `gorm:"size:64"`
	Name        string `gorm:"size:16"`
	Sex         int
	Birthday    time.Time
	CreatedTime time.Time
	UpdatedTime time.Time
}
