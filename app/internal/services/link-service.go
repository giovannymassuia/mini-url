package services

import (
	"fmt"

	"github.com/giovannymassuia/mini-url/internal/domain"
)

type LinkServiceInterface interface {
	CreateShortLink(longUrl string) (string, error)
	GetLongUrl(shortUrl string) (string, error)
	GetAllLinks() ([]domain.Link, error)
}

type LinkService struct {
	links map[string]domain.Link
}

func NewLinkService() *LinkService {
	return &LinkService{
		links: make(map[string]domain.Link),
	}
}

func (s *LinkService) CreateShortLink(longUrl string) (string, error) {
	newId := fmt.Sprintf("%d", len(s.links))
	newLink := domain.Link{
		Id:       newId,
		LongUrl:  longUrl,
		ShortUrl: newId,
	}

	s.links[newLink.Id] = newLink

	return newLink.ShortUrl, nil
}

func (s *LinkService) GetAllLinks() ([]domain.Link, error) {
	links := make([]domain.Link, 0, len(s.links))
	for _, link := range s.links {
		links = append(links, link)
	}

	return links, nil
}

func (s *LinkService) GetLongUrl(shortUrl string) (string, error) {
	link, ok := s.links[shortUrl]
	if !ok {
		return "", fmt.Errorf("link not found")
	}

	return link.LongUrl, nil
}
