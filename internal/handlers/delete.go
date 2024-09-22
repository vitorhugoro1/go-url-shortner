package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	service "github.com/vitorhugoro1/go-url-shortner/internal/services"
)

func DeleteShortenHandler(c *gin.Context) {
	shortId := c.Param("id")

	redisConnec, _ := c.Get("redis")
	redis := redisConnec.(*redis.Client)
	shorten := service.NewShorten(redis)

	err := shorten.Delete(c.Request.Context(), shortId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	}

    c.JSON(http.StatusOK, gin.H{"message": "URL deleted successfully"})
}
