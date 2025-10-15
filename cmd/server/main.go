package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/dkar-dev/HitPipe/internal/config"
	"github.com/dkar-dev/HitPipe/pkg/logger"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {

	// Loading configuration from yaml file
	// The loader and its Load() method is implemented in loader.go
	cfg, err := config.Load[config.Config]("config/config.yaml")
	if err != nil {
		log.Fatalf("error loading config.yaml file: %v", err)
	}

	// Loading .env file
	// Checking that the program is running in production mode
	// Depends on the "ENV" variable
	// Take values: "production", "development", "preprod"
	if cfg.App.Env != "production" && cfg.App.Env != "preprod" {
		err := godotenv.Load("config/.env") // Loading .env from root directory
		if err != nil {
			log.Fatalf("error loading .env file: %v", err)
		}
	}

	log := logger.NewLogger(cfg.App.Env, cfg.Logger.Level)
	slog.SetDefault(log)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host, cfg.Postgres.Port, cfg.Postgres.User, cfg.Postgres.Password, cfg.Postgres.DBName, cfg.Postgres.SSLMode)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Error("failed to connect to database", "error", err)
		//log.Fatalf("cannot connect to the database: %v", err)
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

	err = e.Start(":8080")
	if err != nil {
		return
	}
}
