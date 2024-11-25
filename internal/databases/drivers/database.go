package drivers

import (
	"context"
	"github.com/muratovdias/url-shortner/internal/models"
)

type DataStore interface {
	Base
	UrlShortenerRepo() UrlShortenerRepo
}

type Base interface {
	Name() string
	Ping() error
	Connect() error
	Close(ctx context.Context) error
}

type UrlShortenerRepo interface {
	Save(ctx context.Context, userID string, link models.Link) error
	GetUrlsList(ctx context.Context, userID string) ([]models.Link, error)
}
