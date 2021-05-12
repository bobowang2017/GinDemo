package tests

import (
	"GinDemo/common/utils"
	"GinDemo/common/utils/log"
	"GinDemo/common/utils/redis"
	"GinDemo/config"
	"GinDemo/internal/dao"
	"GinDemo/internal/models"
	"fmt"
	"testing"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	log.Setup()
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
