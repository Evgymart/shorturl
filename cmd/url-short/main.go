package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"main/internal/config"
	"main/storage/postgres"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	logger := setupLogger(cfg.Env)
	logger.Debug("Logger works", slog.String("env", cfg.Env))

	psql, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}

	err = psql.DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	err = psql.SaveURL("http://google.com", "sdfre")
	fmt.Println(err)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}),
		)
	}

	return log
}
