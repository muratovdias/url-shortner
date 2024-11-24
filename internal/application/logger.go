package application

import (
	"log/slog"
	"os"
)

const (
	dev = "dev"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case dev:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
