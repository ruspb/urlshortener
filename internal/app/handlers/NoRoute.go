package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NoRoute(c *gin.Context) {
	c.String(http.StatusBadRequest, "Несуществуюший Route")
}
