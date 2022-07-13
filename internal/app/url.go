package url

import (
	"errors"
	"math/rand"
)

var urls = map[string]string{}

func Create(longURL string) (string, error) {
	shortURL, _ := find(longURL)
	if len([]rune(shortURL)) > 0 {
		return shortURL, nil
	}

	shortURL = randomString(4)
	urls[shortURL] = longURL
	return shortURL, nil
}

func Read(shortURL string) (string, error) {
	longURL, exists := urls[shortURL]
	if !exists {
		return "", errors.New("url не найден")
	}

	return longURL, nil
}

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
