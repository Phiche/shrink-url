package main

import (
	"ru.phiche.shrink.url/api/Controllers"
	//"fmt"
	//"log"
	"github.com/gin-gonic/gin"
	"ru.phiche.shrink.url/api/Storage"
)

func main() {
	CassandraSession := Storage.Session
	defer CassandraSession.Close()
	r := gin.Default()

	r.GET("/liveness", Controllers.CheckHealth)
	r.GET("/readiness", Controllers.CheckHealth)

	r.POST("/url", Controllers.CreateUrl)
	r.GET("/proxy/:hash", Controllers.Redirect)

	r.Run()
}
