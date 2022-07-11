package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ruspb/urlshortener/internal/app"
	"net/http"
)

func main() {
	router := gin.Default()
	router.POST("/", func(c *gin.Context) {
		longUrl, err := c.GetRawData()
		if err != nil {
			c.String(http.StatusBadRequest, "Не передан Url")
		}

		result, err := url.Create(string(longUrl))
		if err != nil {
			c.String(http.StatusBadRequest, "Ошибка сохранения Url")
		}

		c.String(http.StatusCreated, "http://localhost:8080/"+result)
	})
	router.GET("/:id", func(c *gin.Context) {
		result, err := url.Read(c.Param("id"))
		if err != nil {
			c.String(http.StatusBadRequest, "Ошибка чтения Url")
		}

		c.Header("Location", result)
		c.String(http.StatusTemporaryRedirect, "")
		//c.String(http.StatusOK, result)
	})

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusBadRequest, "Несуществуюший Route")
	})

	router.Run(":8080")
}
