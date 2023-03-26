package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const (
	base62     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	base62_len = len(base62)
)

var (
	shortURLs = make(map[string]string)
)

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	var shortURL strings.Builder
	for i := 0; i < 6; i++ {
		idx := rand.Intn(base62_len)
		shortURL.WriteByte(base62[idx])
	}
	return shortURL.String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	longURL := r.URL.Query().Get("url")
	if longURL == "" {
		http.Error(w, "Please provide a URL to shorten", http.StatusBadRequest)
		return
	}
	for {
		shortURL := generateShortURL()
		if _, exists := shortURLs[shortURL]; !exists {
			shortURLs[shortURL] = longURL
			fmt.Fprintf(w, "Short URL: http://localhost:8080/%s", shortURL)
			return
		}
	}
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	longURL, exists := shortURLs[path]
	if !exists {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, longURL, http.StatusFound)
}

func main() {
	http.HandleFunc("/", redirectHandler)
	http.HandleFunc("/shorten", handler)
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
