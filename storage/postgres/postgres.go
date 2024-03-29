package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"main/internal/config"
	"main/storage"
)

const (
	PSQL_UNIQUE_CONSTRAINT_VIOLATION = "23505"
)

type Storage struct {
	DB *sql.DB
}

func New(cfg config.Config) (*Storage, error) {
	const operation = "storage.postgres.New"
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbServer.Ip,
		cfg.DbServer.Port,
		cfg.DbServer.User,
		cfg.DbServer.Password,
		cfg.DbServer.Database,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", operation, err)
	}

	return &Storage{
		DB: db,
	}, nil
}

func (s *Storage) SaveURL(fullUrl string, alias string) error {
	const operation = "storage.postgres.SaveURL"
	stmt, err := s.DB.Prepare("INSERT INTO urls (url, alias) VALUES ($1, $2)")
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	_, err = stmt.Exec(fullUrl, alias)
	if err != nil {
		var psqlError *pq.Error
		if errors.As(err, &psqlError) {
			if psqlError.Code == PSQL_UNIQUE_CONSTRAINT_VIOLATION {
				return fmt.Errorf("%s: %w", operation, storage.ErrUrlAlreadyExists)
			}
		}

		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}

func (s *Storage) GetURL(alias string) (string, error) {
	const operation = "storage.postgres.GetURL"
	stmt, err := s.DB.Prepare("SELECT url FROM urls WHERE alias = $1 LIMIT 1")
	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	var fullUrl string
	err = stmt.QueryRow(alias).Scan(&fullUrl)
	if errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("%s: %w", operation, storage.ErrUrlNotFound)
	}

	if err != nil {
		return "", fmt.Errorf("%s: %w", operation, err)
	}

	return fullUrl, nil
}

func (s *Storage) DeleteURL(alias string) error {
	const operation = "storage.postgres.DeleteURL"
	stmt, err := s.DB.Prepare("DELETE FROM urls WHERE alias = $1")
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	_, err = stmt.Exec(alias)
	return err
}
