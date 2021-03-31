package main

import (
	"GinDemo/config"
	_ "GinDemo/docs"
	"GinDemo/models"
	"GinDemo/router"
	"GinDemo/utils/logger"
	"GinDemo/utils/redis"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUp() {
	config.Setup()
	redis.SetUp()
	logger.Setup()
	models.Setup()
}

// @title Hello World
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v1
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
