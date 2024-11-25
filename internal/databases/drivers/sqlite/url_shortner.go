package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/muratovdias/url-shortner/internal/databases/drivers"
	"github.com/muratovdias/url-shortner/internal/models"
)

type urlShortnerRepo struct {
	db *sql.DB
}

func (u *urlShortnerRepo) GetUrlsList(ctx context.Context, userID string) ([]models.Link, error) {
	//TODO implement me
	panic("implement me")
}

func (u *urlShortnerRepo) Save(ctx context.Context, userID string, link models.Link) error {
	query := `
		INSERT INTO url (alias, url, user_id, expire_date)
		VALUES (?, ?, ?, ?)
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, link.Alias, link.Url, userID, link.ExpireTime)
	if err != nil {
		return fmt.Errorf("failed to insert short URL: %w", err)
	}

	return nil
}

func NewUrlShortenerRepo(db *sql.DB) drivers.UrlShortenerRepo {
	return &urlShortnerRepo{db: db}
}
