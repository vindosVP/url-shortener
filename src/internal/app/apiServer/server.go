package apiServer

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/vindosVP/url-shortener/src/internal/config"
	"github.com/vindosVP/url-shortener/src/internal/controller/http"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"golang.org/x/exp/slog"
)

func Setup(config *config.Config, l *slog.Logger, shortener usecase.Shortener) {

	const op = "apiServer.Setup"

	log := l.With(
		slog.String("op", op),
	)

	log.Info(fmt.Sprintf("starting api server on port %s", config.Server.APIPort))

	handler := fiber.New()
	setupMiddlewares(handler)

	http.SetupRoutes(handler, shortener, l)
	if err := handler.Listen(":" + config.Server.APIPort); err != nil {
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
