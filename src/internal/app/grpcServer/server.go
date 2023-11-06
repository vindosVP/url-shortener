package grpcServer

import (
	"fmt"
	"github.com/vindosVP/url-shortener/src/internal/config"
	"github.com/vindosVP/url-shortener/src/internal/controller/grpcController"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	stdlog "log"
	"net"
)

func Setup(config *config.Config, l *slog.Logger, shortener usecase.Shortener) {
	const op = "grpcServer.Setup"

	log := l.With(
		slog.String("op", op),
	)

	log.Info(fmt.Sprintf("starting grpc server on port %s", config.Server.GRPCPort))
	lis, err := net.Listen("tcp", ":"+config.Server.GRPCPort)
	if err != nil {
		stdlog.Fatal("failed to setup listener")
	}
	s := grpc.NewServer()

	service := grpcController.NewShortener(shortener, l)
	grpcController.RegisterUrlShortenerServer(s, service)

	err = s.Serve(lis)
	if err != nil {
		stdlog.Fatal("failed to serve grpc")
	}
}
