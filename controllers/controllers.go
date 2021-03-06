package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}
