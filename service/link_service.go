package service

import (
	"github.com/umtcn/link-shortener/domain"
	"github.com/umtcn/link-shortener/repository"
)

// LinkService, link kısaltma işlemlerini yönetecek servis
type LinkService struct {
	linkRepository *repository.InMemoryLinkRepository
}

// NewLinkService, yeni bir LinkService örneği oluşturur
func NewLinkService(linkRepository *repository.InMemoryLinkRepository) *LinkService {
	return &LinkService{linkRepository: linkRepository}
}

// ShortenLink, uzun bir linki kısaltır ve kısa linki döner
func (s *LinkService) ShortenLink(originalURL, shortURL string) string {
	link := domain.Link{OriginalURL: originalURL, ShortURL: shortURL}
	return s.linkRepository.Save(link)
}

// GetOriginalURL, kısa linki kullanarak ilgili orijinal linki döner
func (s *LinkService) GetOriginalURL(shortURL string) *string {
	link := s.linkRepository.FindByShortURL(shortURL)
	if link != nil {
		return &link.OriginalURL
	}
	return nil
}
