package main

import (
	"GinDemo/config"
	"GinDemo/router"
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
)

func SetUp() {
	config.Setup()
	redis.SetUpRedis()
	logger.SetupLog()
}

func main() {
	SetUp()
	router := router.InitRouter()
	router.Run()
}
