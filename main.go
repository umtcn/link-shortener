package main

import (
	"fmt"
	"net/http"

	"github.com/umtcn/link-shortener/repository"
	"github.com/umtcn/link-shortener/service"
)

func main() {
	// In-memory repository ve link service örneğini oluşturuyoruz
	linkRepo := &repository.InMemoryLinkRepository{}
	linkService := service.NewLinkService(linkRepo)

	// Test için kısa bir link oluşturuyoruz
	linkService.ShortenLink("https://www.example.com", "abc123")

	// HTTP işleyicilerini ayarlayıp servisi başlatıyoruz
	http.HandleFunc("/shorten", ShortenLinkHandler(linkService))
	http.HandleFunc("/get", GetOriginalURLHandler(linkService))
	http.ListenAndServe(":8080", nil)
}

func ShortenLinkHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		originalURL := r.FormValue("original_url")
		shortURL := r.FormValue("short_url")

		if originalURL == "" || shortURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Eksik parametreler")
			return
		}

		result := linkService.ShortenLink(originalURL, shortURL)
		fmt.Fprintf(w, "Kısa link: %s", result)
	}
}

// GetOriginalURLHandler, orijinal linki almak için HTTP işleyici
func GetOriginalURLHandler(linkService *service.LinkService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		shortURL := r.URL.Query().Get("short_url")

		if shortURL == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Eksik parametre")
			return
		}

		originalURL := linkService.GetOriginalURL(shortURL)

		if originalURL == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Link bulunamadı")
			return
		}

		http.Redirect(w, r, *originalURL, http.StatusSeeOther)
	}
}
