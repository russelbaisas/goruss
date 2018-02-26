package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	router = gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFile("golang.png", "golang.png")
	//nitializeRoutes()
	router.Run()
}
