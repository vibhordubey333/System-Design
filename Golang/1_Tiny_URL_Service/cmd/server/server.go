package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vibhordubey333/url_shortner/pkg/handler"
	"github.com/vibhordubey333/url_shortner/pkg/repository"
)

func main() {
	log.Printf("Starting URL shortner service")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey Go URL Shortener !",
		})
	})

	redisDBResponse, errorRespose := repository.LookUpDatabase("redis")
	if errorRespose != nil {
		fmt.Println("ErrorResponse: ", errorRespose)
	}
	handlerObject := new(handler.UrlCreationRequest)
	r.POST("/create-short-url", func(c *gin.Context) {
		handlerObject.CreateShortURL(c, redisDBResponse)
	})

	r.GET("/:shortURL", func(c *gin.Context) {
		fmt.Println("Context In ShortURL:", c.Keys)
		handlerObject.HandleShortURLRedirect(c, redisDBResponse)
	})

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
