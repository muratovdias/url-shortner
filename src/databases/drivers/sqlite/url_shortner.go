package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/mattn/go-sqlite3"
	"github.com/muratovdias/url-shortner/src/databases/drivers"
	"github.com/muratovdias/url-shortner/src/models"
)

type urlShortnerRepo struct {
	db *sql.DB
}

func (u *urlShortnerRepo) UpdateStats(ctx context.Context, link models.Link) error {
	query := `
		UPDATE url
		SET clicks = clicks + $1, last_access_time = $2 
		WHERE alias = $3 ;
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.ExecContext(ctx, link.UrlStats.Clicks, link.UrlStats.LastAccessTime, link.Alias); err != nil {
		return err
	}

	return nil
}

func (u *urlShortnerRepo) GetOriginalUrl(ctx context.Context, alias string) (models.Link, error) {
	query := `
		SELECT url, expire_date
		FROM url
		WHERE alias = $1
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return models.Link{}, err
	}

	var link models.Link
	if err = stmt.QueryRowContext(ctx, alias).Scan(&link.Url, &link.ExpireTime); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Link{}, sql.ErrNoRows
		}
		return models.Link{}, err
	}

	return link, nil
}

func (u *urlShortnerRepo) Stats(ctx context.Context, alias string) (models.UrlStats, error) {
	query := `
        SELECT clicks, last_access_time
        FROM url
        WHERE alias = ?
    `

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return models.UrlStats{}, err
	}
	defer stmt.Close()

	var stats models.UrlStats
	var lastAccessedAt sql.NullTime

	row := stmt.QueryRowContext(ctx, alias)
	if err = row.Scan(&stats.Clicks, &lastAccessedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return stats, sql.ErrNoRows
		}
		return stats, err
	}

	if lastAccessedAt.Valid {
		stats.LastAccessTime = lastAccessedAt.Time
	}

	return stats, nil
}

func (u *urlShortnerRepo) Delete(ctx context.Context, alias string) error {
	query := `
		DELETE FROM url
		WHERE alias = ?
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, alias)
	if err != nil {
		return fmt.Errorf("failed to delete short link: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to fetch affected rows: %w", err)
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (u *urlShortnerRepo) GetUrlsList(ctx context.Context) ([]models.Link, error) {
	query := `
	SELECT alias, url, expire_date
	FROM url
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare query: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var links []models.Link

	for rows.Next() {
		var link models.Link

		err = rows.Scan(
			&link.Alias,
			&link.Url,
			&link.ExpireTime,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return links, nil
}

func (u *urlShortnerRepo) Save(ctx context.Context, link models.Link) error {
	query := `
		INSERT INTO url (alias, url, expire_date)
		VALUES (?, ?, ?)
	`

	stmt, err := u.db.PrepareContext(ctx, query)
	if err != nil {
		return fmt.Errorf("failed to prepare query: %w", err)
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, link.Alias, link.Url, link.ExpireTime)
	if err != nil {
		if sqliteErr, ok := err.(sqlite3.Error); ok && sqliteErr.ExtendedCode == sqlite3.ErrConstraintUnique {
			return fmt.Errorf("%s: %w", err, models.ErrAlreadyExists)
		}
		return fmt.Errorf("failed to insert short URL: %w", err)
	}

	return nil
}

func NewUrlShortenerRepo(db *sql.DB) drivers.UrlShortenerRepo {
	return &urlShortnerRepo{db: db}
}
