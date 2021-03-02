package dao

import m "GinDemo/models"

type UserDao struct {
}

func NewUserDao() *UserDao {
	return &UserDao{}
}

func (d *UserDao) AddUser(user *m.User) error {
	if err := m.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (d *UserDao) ListUser(PageNum, PageSize int, params map[string]interface{}) ([]*m.User, error) {
	var users []*m.User
	err := m.DB.Offset((PageNum - 1) * PageSize).Limit(PageSize).Where(params).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDao) DeleteById(userId int) error {
	user := m.User{}
	user.BaseModel.ID = userId
	err := m.DB.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) UpdateById(userId int, params map[string]interface{}) error {
	user := m.User{}
	user.BaseModel.ID = userId
	err := m.DB.Model(&user).Updates(params).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) SearchUser(params map[string]interface{}) ([]*m.User, error) {
	var users []*m.User
	err := m.DB.Where(params).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDao) DetailUser(userId int) (*m.User, error) {
	user := &m.User{}
	user.BaseModel.ID = userId
	err := m.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
