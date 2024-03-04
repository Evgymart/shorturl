package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"main/internal/config"
	"main/internal/http-server/handlers/url/save"
	"main/storage/postgres"
	"net/http"
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
	fullUrl, err := psql.GetURL("sdfre")
	fmt.Println(err, fullUrl)
	err = psql.DeleteURL("sdfre")
	fmt.Println(err)

	mux := http.NewServeMux()
	mux.HandleFunc("POST /save", save.New(logger, psql))
	err = http.ListenAndServe("localhost:8090", mux)
	if err != nil {
		panic(err)
	}
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
