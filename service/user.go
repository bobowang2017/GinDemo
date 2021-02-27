package service

import (
	"GinDemo/dao"
	"GinDemo/models"
)

type UserService interface {
	ListUser(PageNum, PageSize int, param map[string]interface{}) ([]*models.User, error)
	DelByUserId(userId int) error
	UpdateByUserId(userId int, params map[string]interface{}) error
	SearchUser(params map[string]interface{}) ([]*models.User, error)
}

type userService struct {
	userDao *dao.UserDao
}

func NewUserService() UserService {
	return &userService{
		dao.NewUserDao(),
	}
}

func (u *userService) CreateUser(o models.User) error {
	user := &models.User{
		Username: o.Username,
		Password: o.Password,
		Name:     o.Name,
		Birthday: o.Birthday,
	}
	if err := u.userDao.AddUser(user); err != nil {
		return err
	}
	return nil
}

func (u *userService) ListUser(PageNum, PageSize int, param map[string]interface{}) ([]*models.User, error) {
	return u.userDao.ListUser(PageNum, PageSize, param)
}

func (u *userService) DelByUserId(userId int) error {
	//return nil
	return u.userDao.DeleteById(userId)
}

func (u *userService) UpdateByUserId(userId int, params map[string]interface{}) error {
	return u.userDao.UpdateById(userId, params)
}

func (u *userService) SearchUser(params map[string]interface{}) ([]*models.User, error) {
	return u.userDao.SearchUser(params)
}
