package drivers

import (
	"context"
	"github.com/muratovdias/url-shortner/src/models"
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
	Save(ctx context.Context, link models.Link) error
	GetUrlsList(ctx context.Context) ([]models.Link, error)
	Delete(ctx context.Context, url string) error
	Stats(ctx context.Context, alias string) (models.UrlStats, error)
	GetOriginalUrl(ctx context.Context, alias string) (models.Link, error)
	UpdateStats(ctx context.Context, link models.Link) error
}
