package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log/slog"
	"main/internal/config"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	conf := config.MustLoad()
	fmt.Println(conf)
	logger := setupLogger(conf.Env)
	logger.Debug("Logger works", slog.String("env", conf.Env))

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.DbServer.Ip, conf.DbServer.Port, conf.DbServer.User, conf.DbServer.Password, conf.DbServer.Database)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
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
