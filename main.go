package main

import (
	"GinDemo/config"
	"GinDemo/models"
	"GinDemo/router"
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	logger.Setup()
	models.Setup()
}

func main() {
	SetUp()
	routers := router.InitRouter()
	routers.Run()
}
