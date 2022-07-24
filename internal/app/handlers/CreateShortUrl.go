package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruspb/urlshortener/internal/app/storage"
	"net/http"
)

func CreateShortUrl(c *gin.Context) {
	longURL, err := c.GetRawData()
	if err != nil {
		c.String(http.StatusBadRequest, "Не передан Url")
	}

	result, err := storage.Store(string(longURL))
	if err != nil {
		c.String(http.StatusBadRequest, "Ошибка сохранения Url")
	}

	c.String(http.StatusCreated, "http://localhost:8080/"+result)
}
