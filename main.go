package main

import (
	"GinDemo/config"
	"GinDemo/router"
	"GinDemo/utils/logger"
)

func SetUp() {
	config.Setup()
	//redis.SetUp()
	logger.Setup()
	//models.Setup()
}

func main() {
	SetUp()
	router := router.InitRouter()
	router.Run()
}
