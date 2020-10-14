package service

import (
	"GinDemo/models"
	"time"
)

type UserDto struct {
	Username string
	Password string
	Name     string
	Sex      int
	Birthday time.Time

	PageNum  int
	PageSize int
}

func (u *UserDto) CreateUser() error {
	article := map[string]interface{}{
		"Username": u.Username,
		"Password": u.Password,
		"Name":     u.Name,
		"Sex":      u.Sex,
		"Birthday": u.Birthday,
	}
	if err := models.AddUser(article); err != nil {
		return err
	}
	return nil
}

func (u *UserDto) ListUser(PageNum, PageSize int, param map[string]string) ([]*models.User, error) {
	return models.ListUser(PageNum, PageSize, param)
}

func (u *UserDto) DelByUserId(userId int) error {
	return models.DeleteById(userId)
}
