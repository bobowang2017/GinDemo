package tests

import (
	"GinDemo/config"
	"GinDemo/dao"
	"GinDemo/models"
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
	"fmt"
	"testing"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	logger.Setup()
	models.Setup()
}

func TestUserList(t *testing.T) {
	SetUp()
	userDao := dao.NewUserDao()
	res, err := userDao.ListUser(1, 10, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
}
