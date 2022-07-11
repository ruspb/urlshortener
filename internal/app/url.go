package url

import (
	"errors"
	"math/rand"
)

var urls = map[string]string{}

func Create(longUrl string) (string, error) {
	shortUrl, _ := find(longUrl)
	if len([]rune(shortUrl)) > 0 {
		return shortUrl, nil
	}

	shortUrl = randomString(4)
	urls[shortUrl] = longUrl
	return shortUrl, nil
}

func Read(shortUrl string) (string, error) {
	longUrl, exists := urls[shortUrl]
	if !exists {
		return "", errors.New("url не найден")
	}

	return longUrl, nil
}

func randomString(n int) string {
	var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0987654321")
	str := make([]rune, n)
	for i := range str {
		str[i] = chars[rand.Intn(len(chars))]
	}
	return string(str)
}

func find(longUrl string) (string, string) {
	for key, value := range urls {
		if value == longUrl {
			return key, value
		}
	}
	return "", ""
}
