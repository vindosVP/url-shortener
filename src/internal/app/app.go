package app

import (
	"github.com/vindosVP/url-shortener/src/internal/app/apiServer"
	"github.com/vindosVP/url-shortener/src/internal/app/grpcServer"
	"github.com/vindosVP/url-shortener/src/internal/config"
	"golang.org/x/exp/slog"
)

const (
	api  = "api"
	grpc = "grpc"
)

func Run(config *config.Config, l *slog.Logger) {
	shortener := NewShortener(config, l)

	switch config.Server.Type {
	case api:
		apiServer.Setup(config, l, shortener)
	case grpc:
		grpcServer.Setup(config, l, shortener)
	default:
		apiServer.Setup(config, l, shortener)
	}

}
