package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/vitorhugoro1/go-url-shortner/internal/handlers"
)

func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return client
}

func main() {
	rds := initRedis()

	router := gin.Default()

	router.Use(func(ctx *gin.Context) {
		ctx.Set("redis", rds)
	})

	router.POST("/shorten", handlers.ShortenHandler)
	router.GET("/:id", handlers.RedirectHandler)
	router.DELETE("/:id", handlers.DeleteShortenHandler)

	router.Run(":8080")
}
