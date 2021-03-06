package main

import (
	//"fmt"
	//"log"
	"github.com/gin-gonic/gin"
	"ru.phiche.shrink.url/api/controllers"
)

func main() {
	r := gin.Default()

	r.GET("/liveness", controllers.CheckHealth)
	r.GET("/readiness", controllers.CheckHealth)

	r.Run()
}
