package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	service "github.com/vitorhugoro1/go-url-shortner/internal/services"
)

func ShortenHandler(c *gin.Context) {
	var body map[string]string

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	originalUrl := body["long_url"]

	redisConnec, _ := c.Get("redis")
	redis := redisConnec.(*redis.Client)

	shorten := service.NewShorten(redis)

	shortedKey, err := shorten.CreateShorten(c.Request.Context(), originalUrl)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to store URL",
			"trace": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"shortned": "http://localhost:8080/" + *shortedKey,
	})
}
