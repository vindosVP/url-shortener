package http

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
	"regexp"
	"url-shortener/src/internal/cerrors"
	"url-shortener/src/internal/controller/http/resp"
	"url-shortener/src/internal/usecase"
)

type shortenRequest struct {
	URL string `json:"url" binding:"required" example:"https://www.ozon.ru/category/smartfony-15502/" validate:"required,url"`
}

type getRequest struct {
	URL string `json:"url" binding:"required" example:"https://mydomain/JAHBG_068H" validate:"required,url"`
}

type ShortenRoutes struct {
	s usecase.Shortener
	l *slog.Logger
}

func SetupShortenRoutes(handler fiber.Router, s usecase.Shortener, l *slog.Logger) {
	r := &ShortenRoutes{
		s: s,
		l: l,
	}
	handler.Get("", r.get)
	handler.Post("", r.save)
}

// @Summary     Get
// @Description Get original url by alias
// @ID          get
// @Tags  	    url-shortener
// @Accept      json
// @Produce     json
// @Param       request body getRequest true "Alias"
// @Success     200 {object} resp.Response
// @Failure     400 {object} resp.Response
// @Failure     500 {object} resp.Response
// @Router      / [get]
func (r *ShortenRoutes) get(c *fiber.Ctx) error {
	const op = "controller.http.get"

	log := r.l.With(
		slog.String("op", op),
		slog.String("request_id", c.Locals("request-id").(string)),
	)

	req := &getRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Error("failed to decode request body")
		return resp.ErrorResponse(c, fiber.StatusInternalServerError, "failed to decode request body")
	}

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		log.Info("invalid request body")
		return resp.ValidationError(c, fiber.StatusBadRequest, validateErr)
	}

	url := req.URL
	pattern := r.s.GetLinkPattern()
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Error("failed to compile regexp")
		return resp.ErrorResponse(c, fiber.StatusInternalServerError, "failed to compile regexp")
	}

	urlValid := reg.MatchString(url)
	if !urlValid {
		log.Info("invalid request body")
		return resp.ErrorResponse(c, fiber.StatusBadRequest, "url is not valid")
	}

	alias := url[len(url)-10:]

	originalURL, err := r.s.GetOriginal(alias)
	if err != nil {
		if err == cerrors.ErrAliasForURLDoesNotExist {
			return resp.ErrorResponse(c, fiber.StatusNotFound, "ulr with this alias does not exist")
		}
		return resp.ErrorResponse(c, fiber.StatusInternalServerError, "internal server error")
	}

	return resp.OkResponse(c, fiber.StatusOK, originalURL)
}

// @Summary     Save
// @Description Save alias for an url
// @ID          save
// @Tags  	    url-shortener
// @Accept      json
// @Produce     json
// @Param       request body shortenRequest true "Url"
// @Success     200 {object} resp.Response
// @Failure     400 {object} resp.Response
// @Failure     500 {object} resp.Response
// @Router      / [post]
func (r *ShortenRoutes) save(c *fiber.Ctx) error {
	const op = "controller.http.save"

	log := r.l.With(
		slog.String("op", op),
		slog.String("request_id", c.Locals("request-id").(string)),
	)

	req := &shortenRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Error("failed to decode request body")
		return resp.ErrorResponse(c, fiber.StatusInternalServerError, "failed to decode request body")
	}

	if err := validator.New().Struct(req); err != nil {
		validateErr := err.(validator.ValidationErrors)
		log.Info("invalid request body")
		return resp.ValidationError(c, fiber.StatusBadRequest, validateErr)
	}

	url, err := r.s.Save(req.URL)
	if err != nil {
		log.Error("failed to save alias")
		return resp.ErrorResponse(c, fiber.StatusInternalServerError, "failed to save alias")
	}

	return resp.OkResponse(c, fiber.StatusOK, url)
}
