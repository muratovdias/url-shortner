package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muratovdias/url-shortner/internal/config"
	"github.com/muratovdias/url-shortner/internal/databases/drivers"
	"time"
)

const (
	connectionTimeout = 3 * time.Second
)

type sqliteDB struct {
	db                *sql.DB
	path              string
	dbName            string
	connectionTimeout time.Duration

	urlShortenerRepo drivers.UrlShortenerRepo
}

func (s *sqliteDB) UrlShortenerRepo() drivers.UrlShortenerRepo {
	if s.urlShortenerRepo == nil {
		s.urlShortenerRepo = NewUrlShortenerRepo(s.db)
	}

	return s.urlShortenerRepo
}

func New(conf config.DataStore) drivers.DataStore {
	return &sqliteDB{
		path:              conf.Path,
		dbName:            conf.DbName,
		connectionTimeout: connectionTimeout,
	}
}

func (s *sqliteDB) Name() string {
	return "SQLite"
}

func (s *sqliteDB) Ping() error {
	if s.db == nil {
		return fmt.Errorf("database is not initialized")
	}
	return s.db.Ping()
}

func (s *sqliteDB) Connect() error {
	if s.db != nil {
		return fmt.Errorf("database connection already established")
	}

	db, err := sql.Open("sqlite3", s.path)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %w", err)
	}

	s.db = db

	if err = s.createTable(); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func (s *sqliteDB) Close(ctx context.Context) error {
	if s.db == nil {
		return fmt.Errorf("database is not initialized")
	}
	return s.db.Close()
}

func (s *sqliteDB) createTable() error {
	stmt, err := s.db.Prepare(`
	CREATE TABLE IF NOT EXISTS url(
		alias TEXT PRIMARY KEY,
        url TEXT NOT NULL,
        user_id TEXT NOT NULL,
        clicks INTEGER DEFAULT 0,
        last_access_time DATETIME,
        expirу_date DATETIME,
    	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);
	`)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	return nil
}
