package url

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUrl(c *gin.Context) {
	c.String(http.StatusOK, "POST /")
}

func ReadUrl(c *gin.Context) {
	short := c.Param("id")
	c.String(http.StatusOK, "GET "+short)
}
