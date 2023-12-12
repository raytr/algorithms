package main

import (
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

var urlMap = make(map[string]string)

func shortenURL(longURL string) string {
	hash := sha1.Sum([]byte(longURL))
	shortURL := base64.URLEncoding.EncodeToString(hash[:])[:8]
	urlMap[shortURL] = longURL
	return shortURL
}

func getLongURL(shortURL string) string {
	return urlMap[shortURL]
}

func main() {
	longURL := "https://www.example.com/this-is-a-very-long-url"

	shortURL := shortenURL(longURL)
	fmt.Println("Short URL:", shortURL)

	retrievedLongURL := getLongURL(shortURL)
	fmt.Println("Retrieved Long URL:", retrievedLongURL)
}
