package main

import (
	"search/config"
	"search/model"
	"search/router"

	"github.com/gin-gonic/gin"
)

func main() {
	config.ConfigInit()
	model.InitConnection()
	r := gin.Default()
	router.LoadRouters(r)
	r.Run(":80")
}