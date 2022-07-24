package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateShortURL(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
	}{
		{
			name: "Создание короткого Url",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetRouter()
			r.POST("/", CreateShortURL)
			req, _ := http.NewRequest("POST", "/", strings.NewReader("http://ya.ru"))
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusCreated, w.Code)
		})
	}
}
