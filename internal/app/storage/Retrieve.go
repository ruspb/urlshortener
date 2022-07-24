package storage

import "errors"

func Retrieve(shortURL string) (string, error) {
	longURL, exists := urls[shortURL]
	if !exists {
		return "", errors.New("url не найден")
	}

	return longURL, nil
}
