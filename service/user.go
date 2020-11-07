package service

import (
	"GinDemo/models"
)

type UserDto struct {
	models.User
	PageNum  int
	PageSize int
}

func (u *UserDto) CreateUser() error {
	user := &models.User{
		Username: u.Username,
		Password: u.Password,
		Name:     u.Name,
		Birthday: u.Birthday,
	}
	if err := models.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserDto) ListUser(PageNum, PageSize int, param map[string]string) ([]*models.User, error) {
	return models.ListUser(PageNum, PageSize, param)
}

func (u *UserDto) DelByUserId(userId int) error {
	//return nil
	return models.DeleteById(userId)
}

func (u *UserDto) UpdateByUserId(userId int, params map[string]interface{}) error {
	return models.UpdateById(userId, params)
}
