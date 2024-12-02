package logger

import (
	"log/slog"
	"medods-test/internal/env"
	"os"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func SetupLogger() *slog.Logger {
	var log *slog.Logger

	switch env.LOG_LEVEL {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	case envProd:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}
