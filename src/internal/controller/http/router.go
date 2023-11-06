package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"golang.org/x/exp/slog"
)

// SetupRoutes -.
// Swagger spec:
// @title       url-shortener API
// @description Url shortening api
// @version     1.0
// @host        localhost:8080
func SetupRoutes(handler *fiber.App, s usecase.Shortener, l *slog.Logger) {

	handler.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	SetupShortenRoutes(handler, s, l)
}
