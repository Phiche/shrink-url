package main

import (
	"github.com/gin-gonic/gin"
	"ru.phiche.shrink.url/api/Controllers"
	"ru.phiche.shrink.url/api/Storage"
)

func main() {
	CassandraSession := Storage.Session
	defer CassandraSession.Close()
	r := gin.Default()

	r.GET("/probe/liveness", Controllers.CheckHealth)
	r.GET("/probe/readiness", Controllers.CheckHealth)

	r.POST("/url", Controllers.CreateUrl)
	//todo: gin can't have wildcard parent uri because of conflict
	r.GET("/proxy/:hash", Controllers.Redirect)

	r.Run()
}
