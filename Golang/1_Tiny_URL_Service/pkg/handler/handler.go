package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vibhordubey333/url_shortner/pkg/repository"
	"github.com/vibhordubey333/url_shortner/pkg/service"
)

type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}

func (u *UrlCreationRequest) CreateShortURL(c *gin.Context, redisClient repository.Database) {

	if err := c.ShouldBindJSON(u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortURL := service.GenerateShortLink(u.LongUrl, u.UserId)
	redisClient.SaveUrlMapping(shortURL, u.LongUrl, u.UserId)
	host := "http://localhost:9808/"
	c.JSON(200, gin.H{
		"message":   "short url created successfully",
		"short_url": host + shortURL,
	})
}

func (u *UrlCreationRequest) HandleShortURLRedirect(c *gin.Context, redisClient repository.Database) {
	shortURL := c.Param("shortURL")
	initialURL := redisClient.RetrieveInitialUrl(shortURL)
	c.Redirect(302, initialURL)
}
