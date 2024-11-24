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

func (u *urlShortnerRepo) SaveAlias(ctx context.Context, link models.Link) error {
	query := `
		INSERT INTO url (alias, url, expire_date)
		VALUES (?, ?, ?)
	`

	// Prepare the statement for execution
	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}
	// Ensure that the prepared statement is closed after execution
	defer stmt.Close()

	// Execute the prepared statement with the provided values
	_, err = stmt.ExecContext(ctx, link.Alias, link.Url, link.ExpireTime)
	if err != nil {
		return fmt.Errorf("failed to insert short URL: %w", err)
	}

	return nil
}

func NewUrlShortenerRepo(db *sql.DB) drivers.UrlShortenerRepo {
	return &urlShortnerRepo{db: db}
}
