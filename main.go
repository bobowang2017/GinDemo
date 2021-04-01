package main

import (
	"GinDemo/common/utils/logger"
	"GinDemo/common/utils/redis"
	"GinDemo/config"
	_ "GinDemo/docs"
	"GinDemo/internal/models"
	"GinDemo/router"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	logger.Setup()
	models.Setup()
}

// @title GinDemo
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://127.0.0.1:8080

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host petstore.swagger.io
// @BasePath /api/v1
func main() {
	SetUp()
	//c := cron.New()
	//c.AddFunc("* * * * * *", func() {
	//	log.Println("Run models.CleanAllTag...")
	//})
	//c.AddFunc("* * * * * *", func() {
	//	log.Println("Run models.CleanAllArticle...")
	//})

	//c.Start()
	routers := router.InitRouter()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	routers.Run()
}
