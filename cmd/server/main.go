package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/dkar-dev/hitpipe/internal/config"
	"github.com/dkar-dev/hitpipe/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func main() {

	// Loading configuration from config.yaml and .env file
	// Implemented in loader.go
	cfg, err := config.Load("./config/")
	if err != nil {
		log.Fatalf("ERROR: failed to load config.yaml file: %v", err)
	}

	log := logger.NewLogger(cfg.App.Env, cfg.Logger.Level)
	slog.SetDefault(log)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DB, cfg.Postgres.SSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Error("failed to connect to the database", "error", err)
	}

	defer func(db *sqlx.DB) {
		err := db.Close()
		if err != nil {
			log.Error("failed to close database connection", "error", err)
		}
	}(db)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err = e.Start("8080:8080")
	if err != nil {
		log.Error("failed to start the server", "error", err)
		return
	}
}
