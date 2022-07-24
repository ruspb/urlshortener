package storage

import (
	"math/rand"
)

var urls = map[string]string{}

func randomString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func find(longURL string) (string, string) {
	for key, value := range urls {
		if value == longURL {
			return key, value
		}
	}
	return "", ""
}
