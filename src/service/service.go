package service

import "github.com/muratovdias/url-shortner/src/service/shortner"

type Service struct {
	UrlShortener shortner.UrlShortener
}

func NewService(urlShortener shortner.UrlShortener) *Service {
	return &Service{UrlShortener: urlShortener}
}
