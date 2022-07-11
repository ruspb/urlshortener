package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ruspb/urlshortener/internal/app"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/", url.CreateUrl)
	router.GET("/:id", url.ReadUrl)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusBadRequest, "Bad Request")
	})

	router.Run(":8080")
}
