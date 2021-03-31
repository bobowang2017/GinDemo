package tests

import (
	"GinDemo/config"
	"GinDemo/dao"
	"GinDemo/models"
	"GinDemo/utils"
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

func TestHello(t *testing.T) {
	param := map[string]interface{}{
		"hello": "123",
		"world": 123.56,
	}
	res := utils.GetInt64FromMap(param, "world", 456)
	fmt.Println(res)
}
