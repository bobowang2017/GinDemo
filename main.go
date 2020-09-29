package main

import (
	"GinDemo/router"
)

func main() {
	router := router.InitRouter()
	router.Run()
}
