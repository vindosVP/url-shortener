package http

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"url-shortener/src/internal/usecase"
)

func SetupRoutes(handler *fiber.App, s usecase.Shortener, l *slog.Logger) {

	handler.Get("/ping", func(ctx *fiber.Ctx) error {
		return ctx.SendStatus(fiber.StatusOK)
	})
	SetupShortenRoutes(handler, s, l)
}
