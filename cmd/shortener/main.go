package main

import (
	"github.com/ruspb/urlshortener/internal/app/handlers"
)

func main() {
	router := handlers.GetRouter()
	router.POST("/", handlers.CreateShortURL)
	router.GET("/:id", handlers.GetLongURL)
	router.NoRoute(handlers.NoRoute)

	router.Run(":8080")
}
