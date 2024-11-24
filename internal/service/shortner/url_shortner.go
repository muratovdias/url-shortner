package shortner

import (
	"context"
	"github.com/muratovdias/url-shortner/internal/databases/drivers"
	"github.com/muratovdias/url-shortner/internal/models"
	"log/slog"
	"regexp"
	"time"
)

type UrlShortener interface {
	Save(ctx context.Context, url string) (models.Link, error)
}

type urlShortenerImpl struct {
	repo   drivers.UrlShortenerRepo
	regexp *regexp.Regexp
	log    *slog.Logger
}

func NewUrlShortener(repo drivers.UrlShortenerRepo, log *slog.Logger) UrlShortener {
	return &urlShortenerImpl{
		repo:   repo,
		log:    log,
		regexp: regexp.MustCompile(`^(https?://)?([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}(:[0-9]+)?(/.*)?$`),
	}
}

func (u *urlShortenerImpl) Save(ctx context.Context, url string) (models.Link, error) {
	if err := u.validateURL(url); err != nil {
		u.log.Error("invalid URL", err)
		return models.Link{}, err
	}

	alias, err := generateShortURL(url)
	if err != nil {
		u.log.Error("failed to generate short url", err)
		return models.Link{}, err
	}

	if err = u.repo.SaveAlias(ctx, models.Link{
		Url:        url,
		Alias:      alias,
		ExpireTime: time.Now().Add(time.Hour * 24 * 30), // 30 дней
	}); err != nil {
		u.log.Error("failed to save short url", err)
		return models.Link{}, err
	}

	return models.Link{
		Alias:      alias,
		ExpireTime: time.Now().Add(time.Hour * 24 * 30),
	}, nil
}
