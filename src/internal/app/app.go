package app

import (
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"golang.org/x/exp/slog"
	stdlog "log"
	"url-shortener/src/internal/config"
	"url-shortener/src/internal/controller/http"
	"url-shortener/src/internal/database"
	"url-shortener/src/internal/usecase"
	"url-shortener/src/internal/usecase/inmemory_repo"
	"url-shortener/src/internal/usecase/postgres_repo"
)

const (
	inmemory = "inmemory"
	postgres = "postgres"
)

func Run(config *config.Config, l *slog.Logger) {
	const op = "app.Run"
	handler := fiber.New()
	setupMiddlewares(handler)

	log := l.With(
		slog.String("op", op),
	)

	var shortener usecase.Shortener
	switch config.StorageType {
	case inmemory:
		inmemoryRepo := inmemory_repo.New(log)
		shortener = usecase.NewShortenerUseCase(inmemoryRepo, config.Domain, config.DomainProtocol, log)
	case postgres:
		db, err := database.NewGorm(config.DB)
		if err != nil {
			stdlog.Fatal("failed to setup database")
		}
		postgresRepo := postgres_repo.New(db, log)
		shortener = usecase.NewShortenerUseCase(postgresRepo, config.Domain, config.DomainProtocol, log)
	default:
		inmemoryRepo := inmemory_repo.New(log)
		shortener = usecase.NewShortenerUseCase(inmemoryRepo, config.Domain, config.DomainProtocol, log)
	}

	http.SetupRoutes(handler, shortener, l)
	if err := handler.Listen(":" + config.Server.Port); err != nil {
		log.Error("failed to start api server")
	}
}

func setupMiddlewares(handler *fiber.App) {
	handler.Use(fiberlog.New(fiberlog.Config{
		TimeZone: "Europe/Moscow",
		Format:   "[${time}] ${locals:request-id} ${status} - ${latency} ${method} ${path}​\n",
	}))

	handler.Use(requestid.New(requestid.Config{
		Header:     "X-Request-ID",
		ContextKey: "request-id",
	}))
}
