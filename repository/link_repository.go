package repository

import "github.com/umtcn/link-shortener/domain"

// InMemoryLinkRepository, linklerin bellekte saklanacağı repository
type InMemoryLinkRepository struct {
	links []domain.Link
}

// Save, linki bellekte saklar ve kısa linki döner
func (r *InMemoryLinkRepository) Save(link domain.Link) string {
	r.links = append(r.links, link)
	return link.ShortURL
}

// FindByShortURL, kısa linki kullanarak ilgili linki döner
func (r *InMemoryLinkRepository) FindByShortURL(shortURL string) *domain.Link {
	for _, link := range r.links {
		if link.ShortURL == shortURL {
			return &link
		}
	}
	return nil
}
