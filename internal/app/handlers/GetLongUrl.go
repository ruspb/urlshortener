package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ruspb/urlshortener/internal/app/storage"
	"net/http"
)

func GetLongURL(c *gin.Context) {
	result, err := storage.Retrieve(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, "Ошибка чтения Url")
	}

	c.Header("Location", result)
	c.String(http.StatusTemporaryRedirect, "")
	//c.String(http.StatusOK, result)
}
