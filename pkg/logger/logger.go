package logger

import (
	"github.com/lmittmann/tint"
	"log/slog"
	"os"
	"time"
)

const (
	envLocal = "local"
	envDev   = "development"
	envProd  = "production"
)

func NewLogger(env string, level string) *slog.Logger {
	var log *slog.Logger

	var logLevel slog.Level

	switch level {

	case "debug":
		logLevel = slog.LevelDebug
	case "info":
		logLevel = slog.LevelInfo
	case "warn":
		logLevel = slog.LevelWarn
	case "error":
		logLevel = slog.LevelError

	}

	switch env {

	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: logLevel}))

	case envDev:
		log = slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: logLevel, TimeFormat: time.Kitchen}))

	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: logLevel}))
	}

	return log
}
