package shortner

import (
	"context"
	"errors"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	"github.com/muratovdias/url-shortner/src/models"
	"log/slog"
	"regexp"
	"time"
)

type UrlShortener interface {
	Save(ctx context.Context, url string) (models.Link, error)
	GetUrlsList(ctx context.Context) ([]models.Link, error)
	Delete(ctx context.Context, alias string) error
	Stats(ctx context.Context, alias string) (models.UrlStats, error)
	Redirect(ctx context.Context, alias string) (string, error)
}

type urlShortenerImpl struct {
	repo   drivers.UrlShortenerRepo
	regexp *regexp.Regexp
	log    *slog.Logger
}

func (u *urlShortenerImpl) Redirect(ctx context.Context, alias string) (string, error) {
	link, err := u.repo.GetOriginalUrl(ctx, alias)
	if err != nil {
		u.log.Error("failed to get original url", "alias", alias, "error", err)
		return "", err
	}

	if link.ExpireTime.Before(time.Now()) {
		return "", models.ErrExpired
	}

	link.UrlStats.Clicks += 1
	link.UrlStats.LastAccessTime = time.Now()
	link.Alias = alias

	if err = u.repo.UpdateStats(ctx, link); err != nil {
		u.log.Error("failed to update stats", "alias", alias, "error", err)
		return "", err
	}

	return link.Url, nil
}

func (u *urlShortenerImpl) Stats(ctx context.Context, alias string) (models.UrlStats, error) {
	stats, err := u.repo.Stats(ctx, alias)
	if err != nil {
		u.log.Error("failed to delete url", "error", err)
		return models.UrlStats{}, err
	}

	return stats, nil
}

func (u *urlShortenerImpl) Delete(ctx context.Context, alias string) error {
	if err := u.repo.Delete(ctx, alias); err != nil {
		u.log.Error("failed to delete url", "error", err)
		return err
	}
	return nil
}

func (u *urlShortenerImpl) GetUrlsList(ctx context.Context) ([]models.Link, error) {
	links, err := u.repo.GetUrlsList(ctx)
	if err != nil {
		u.log.Error("failed to get url list", "error", err)
		return nil, err
	}

	return links, nil
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
		u.log.Error("invalid URL", err.Error())
		return models.Link{}, ErrInvalidUrl
	}

	alias, err := generateShortURL()
	if err != nil {
		u.log.Error("failed to generate short url", err.Error())
		return models.Link{}, err
	}

	expireTime := time.Now().Add(time.Hour * 24 * 30) // 30 дней

	if err = u.repo.Save(ctx, models.Link{
		Url:        url,
		Alias:      alias,
		ExpireTime: expireTime,
	}); err != nil {
		u.log.Error("failed to save short url", err.Error())
		if errors.Is(errors.Unwrap(err), models.ErrAlreadyExists) {
			return models.Link{}, err
		}
		return models.Link{}, err
	}

	return models.Link{
		Alias:      alias,
		ExpireTime: expireTime,
	}, nil
}
