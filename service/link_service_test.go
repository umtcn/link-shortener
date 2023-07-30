package service

import (
	"testing"

	"github.com/umtcn/link-shortener/repository"
)

func TestLinkService(t *testing.T) {
	linkRepo := &repository.InMemoryLinkRepository{}
	linkService := NewLinkService(linkRepo)

	// Link kısaltma testi
	originalURL := "https://www.example.com"
	shortURL := "abc123"
	result := linkService.ShortenLink(originalURL, shortURL)

	if result != shortURL {
		t.Errorf("Link kısaltma işlemi başarısız. Beklenen: %s, Alınan: %s", shortURL, result)
	}

	// Orijinal URL'yi almak için kısa linkle sorgulama testi
	original := linkService.GetOriginalURL(shortURL)

	if original == nil || *original != originalURL {
		t.Errorf("Orijinal URL alma işlemi başarısız. Beklenen: %s, Alınan: %v", originalURL, original)
	}
}
