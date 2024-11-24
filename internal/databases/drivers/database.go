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
	SaveAlias(ctx context.Context, link models.Link) error
}
