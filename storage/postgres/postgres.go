package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"main/internal/config"
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

	stmt, err := s.DB.Prepare("INSERT INTO urls(`url`, `alias`) VALUES (?, ?)")
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	_, err = stmt.Exec(fullUrl, alias)
	if err != nil {
		return fmt.Errorf("%s: %w", operation, err)
	}

	return nil
}
