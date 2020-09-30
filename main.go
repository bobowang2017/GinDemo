package main

import (
	"GinDemo/router"
	"GinDemo/utils"
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

func SetUp() {
	utils.InitConfigIni("./config/setting.ini")
	utils.SetUpRedis()
}

func main() {
	SetUp()
	router := router.InitRouter()
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	router.Run()
}
