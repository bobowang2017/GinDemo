package main

import (
	"GinDemo/config"
	"GinDemo/models"
	"GinDemo/router"
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
	"github.com/robfig/cron"
	"log"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	logger.Setup()
	models.Setup()
}

func main() {
	SetUp()
	c := cron.New()
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllTag...")
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models.CleanAllArticle...")
	})

	c.Start()
	routers := router.InitRouter()
	routers.Run()
}
