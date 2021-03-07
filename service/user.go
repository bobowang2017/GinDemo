package service

import (
	"GinDemo/dao"
	"GinDemo/models"
)

//用户模块接口定义
type IUserService interface {
	CreateUser(o *models.User) error
	ListUser(PageNum, PageSize int, param map[string]interface{}) ([]*models.User, error)
	DelByUserId(userId int) error
	UpdateByUserId(userId int, params map[string]interface{}) error
	SearchUser(params map[string]interface{}) ([]*models.User, error)
	Detail(userId int) (*models.User, error)
}

//用户模块接口实现
type userService struct {
	userDao *dao.UserDao
}

func NewUserService() IUserService {
	return &userService{
		dao.NewUserDao(),
	}
}

func (u *userService) CreateUser(o *models.User) error {
	if err := u.userDao.AddUser(o); err != nil {
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

func (u *userService) Detail(userId int) (*models.User, error) {
	return u.userDao.DetailUser(userId)
}

//用户收件地址模块接口定义
type IUserAddrService interface {
	ListByUserId(userId int) ([]*models.UserAddress, error)
	CreateUserAddr(u *models.UserAddress) error
}

//用户收件地址模块接口实现
type userAddrService struct {
	userAddressDao *dao.UserAddrDao
}

func NewUserAddressService() IUserAddrService {
	return &userAddrService{
		dao.NewUserAddressDao(),
	}
}

func (u *userAddrService) ListByUserId(userId int) ([]*models.UserAddress, error) {
	param := make(map[string]interface{})
	param["user_id"] = userId
	return u.userAddressDao.SearchUserAddress(param)
}

func (u *userAddrService) CreateUserAddr(m *models.UserAddress) error {
	return u.userAddressDao.AddUserAddr(m)
}
