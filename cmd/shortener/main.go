package main

import (
	"github.com/ruspb/urlshortener/internal/app/handlers"
)

func main() {
	router := handlers.GetRouter()
	router.POST("/", handlers.CreateShortUrl)
	router.GET("/:id", handlers.GetLongUrl)
	router.NoRoute(handlers.NoRoute)

	router.Run(":8080")
}
