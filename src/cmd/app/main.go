package main

import (
	"golang.org/x/exp/slog"
	stdlog "log"
	"os"
	"url-shortener/src/internal/app"
	"url-shortener/src/internal/config"
	"url-shortener/src/internal/pkg/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		stdlog.Fatalf("Config setup cerrors: %s", err)
	}

	log := setupLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
	)
	log.Debug("debug mode enabled")

	app.Run(cfg, log)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := logger.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
