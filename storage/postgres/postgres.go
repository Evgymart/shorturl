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
