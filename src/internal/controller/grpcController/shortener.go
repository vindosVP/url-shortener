package grpcController

import (
	"context"
	"github.com/asaskevich/govalidator"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"github.com/vindosVP/url-shortener/src/internal/usecase"
	"golang.org/x/exp/slog"
	"regexp"
)

type GrpcShortener struct {
	UnimplementedUrlShortenerServer
	s usecase.Shortener
	l *slog.Logger
}

func NewShortener(s usecase.Shortener, l *slog.Logger) *GrpcShortener {
	return &GrpcShortener{
		s: s,
		l: l,
	}
}

func (s GrpcShortener) Save(ctx context.Context, req *SaveRequest) (*SaveResponse, error) {
	const op = "grpcController.Save"

	log := s.l.With(
		slog.String("op", op),
	)

	isValid := govalidator.IsURL(req.Url)
	if !isValid {
		log.Info("invalid request")
		return &SaveResponse{}, cerrors.ErrInvalidUrl
	}

	url, err := s.s.Save(req.Url)
	if err != nil {
		log.Error("failed to save alias")
		return &SaveResponse{}, err
	}

	return &SaveResponse{ShortenedUrl: url}, nil
}

func (s GrpcShortener) Get(ctx context.Context, req *GetRequest) (*GetResponse, error) {
	const op = "grpcController.Get"

	log := s.l.With(
		slog.String("op", op),
	)

	isValid := govalidator.IsURL(req.ShortenedUrl)
	if !isValid {
		log.Info("invalid request")
		return &GetResponse{}, cerrors.ErrInvalidUrl
	}

	pattern := s.s.GetLinkPattern()
	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Error("failed to compile regexp")
		return &GetResponse{}, err
	}
	urlValid := reg.MatchString(req.ShortenedUrl)
	if !urlValid {
		log.Info("invalid request body")
		return &GetResponse{}, cerrors.ErrInvalidUrl
	}

	alias := req.ShortenedUrl[len(req.ShortenedUrl)-10:]
	originalURL, err := s.s.GetOriginal(alias)
	if err != nil {
		if err != cerrors.ErrAliasForURLDoesNotExist {
			log.Error("failed to get original url")
		}
		return &GetResponse{}, cerrors.ErrAliasForURLDoesNotExist
	}

	return &GetResponse{Url: originalURL}, nil
}
