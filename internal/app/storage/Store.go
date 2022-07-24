package storage

func Store(longURL string) (string, error) {
	shortURL, _ := find(longURL)
	if len([]rune(shortURL)) > 0 {
		return shortURL, nil
	}

	shortURL = randomString(4)
	urls[shortURL] = longURL
	return shortURL, nil
}
