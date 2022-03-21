package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vibhordubey333/url_shortner/pkg/repository"
	"github.com/vibhordubey333/url_shortner/pkg/service"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding required"`
	UserId  string `json:"user_id" binding:"required"`
}

func CreateShortURL(c *gin.Context) {
	var creationRequest UrlCreationRequest
	if err := c.ShouldBindJSON(&creationRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := service.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	repository.SaveUrlMapping(shortURL, creationRequest.LongUrl, creationRequest.UserId)
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})
}

func HandleShortURLRedirect(c *gin.Context) {
	shortURL := c.Param("shortUrl")
	initialURL := repository.RetrieveInitialUrl(shortURL)
	c.Redirect(302, initialURL)
}
