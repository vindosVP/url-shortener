package app

import (
	"github.com/vindosVP/url-shortener/src/internal/config"
	"github.com/vindosVP/url-shortener/src/internal/database"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"github.com/vindosVP/url-shortener/src/internal/usecase/inmemory_repo"
	"github.com/vindosVP/url-shortener/src/internal/usecase/postgres_repo"
	"golang.org/x/exp/slog"
	"log"
)

const (
	inmemory = "inmemory"
	postgres = "postgres"
)

func NewShortener(config *config.Config, l *slog.Logger) usecase.Shortener {

	var shortener usecase.Shortener
	switch config.StorageType {
	case inmemory:
		inmemoryRepo := inmemory_repo.New(l)
		shortener = usecase.NewShortenerUseCase(inmemoryRepo, config.Domain, config.DomainProtocol, l)
	case postgres:
		db, err := database.NewGorm(config.DB)
		if err != nil {
			log.Fatal("failed to setup database")
		}
		postgresRepo := postgres_repo.New(db, l)
		shortener = usecase.NewShortenerUseCase(postgresRepo, config.Domain, config.DomainProtocol, l)
	default:
		inmemoryRepo := inmemory_repo.New(l)
		shortener = usecase.NewShortenerUseCase(inmemoryRepo, config.Domain, config.DomainProtocol, l)
	}

	return shortener

}
