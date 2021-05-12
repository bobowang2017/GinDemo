package main

import (
	"GinDemo/common/utils/log"
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
	log.Setup()
	models.Setup()
}

// @title gin-demo
// @version 1.0
// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	SetUp()
	routers := router.InitRouter()
	// The url pointing to API definition
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	routers.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	routers.Run()
}
