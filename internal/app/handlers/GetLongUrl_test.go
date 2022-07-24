package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLongURL(t *testing.T) {
	type args struct {
		c    *gin.Context
		long string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Создание короткого Url",
			args: args{long: "http://ya.ru"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := GetRouter()

			r.GET("/:id", GetLongURL)
			req, _ := http.NewRequest("GET", "/short", nil)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
		})
	}
}
