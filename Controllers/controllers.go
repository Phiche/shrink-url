package Controllers

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"ru.phiche.shrink.url/api/Storage"
	"time"
)

func CheckHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ready"})
}

type UrlInput struct {
	OriginalUrl string `json:"originalUrl" binding:"required"`
	Name        string `json:"name"`
}

func CreateUrl(c *gin.Context) {
	var input UrlInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var shortUrl = GetShortenMD5Hash(input.OriginalUrl)

	url := Storage.Url{Hash: shortUrl,
		Original_url:  input.OriginalUrl,
		Name:          input.Name,
		Creation_date: time.Now(),
	}
	Storage.CreateUrl(url)

	c.JSON(http.StatusOK, gin.H{"short": shortUrl})
}

func Redirect(c *gin.Context) {
	hash := c.Param("hash")
	originalUrl := Storage.GetOriginalUrl(hash)
	c.Redirect(http.StatusFound, originalUrl)
}

func GetShortenMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[0:5])
}
