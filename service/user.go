package service

import (
	"GinDemo/models"
)

type UserService struct {
}

func (u *UserService) CreateUser(o models.User) error {
	user := &models.User{
		Username: o.Username,
		Password: o.Password,
		Name:     o.Name,
		Birthday: o.Birthday,
	}
	if err := models.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (u *UserService) ListUser(PageNum, PageSize int, param map[string]interface{}) ([]*models.User, error) {
	return models.ListUser(PageNum, PageSize, param)
}

func (u *UserService) DelByUserId(userId int) error {
	//return nil
	return models.DeleteById(userId)
}

func (u *UserService) UpdateByUserId(userId int, params map[string]interface{}) error {
	return models.UpdateById(userId, params)
}

func (u *UserService) SearchUser(params map[string]interface{}) ([]*models.User, error) {
	return models.SearchUser(params)
}
