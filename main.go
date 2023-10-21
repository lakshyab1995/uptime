package url_short

import (
	"golang.org/x/exp/slog"
	"log"
	"net/http"
)

func main() {
	shortener := &UrlShortener{
		urls: make(map[string]string),
	}

	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

	slog.Info("URL Shortener is running on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("URL Shortener service stopped:%s", err)
	}
}
