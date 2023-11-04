package usecase

import (
	"fmt"
	"golang.org/x/exp/slog"
	"url-shortener/src/internal/cerrors"
	"url-shortener/src/internal/pkg/randString"
)

type ShortenerUseCase struct {
	UrlRepo        UrlRepo
	Domain         string
	DomainProtocol string
	l              *slog.Logger
}

func NewShortenerUseCase(r UrlRepo, domain string, domainProtocol string, l *slog.Logger) *ShortenerUseCase {
	return &ShortenerUseCase{
		UrlRepo:        r,
		Domain:         domain,
		DomainProtocol: domainProtocol,
		l:              l,
	}
}

func (s *ShortenerUseCase) GetLinkPattern() string {
	return fmt.Sprintf("%s:\\/\\/%s\\/\\b([-a-zA-Z0-9_]{10})$", s.DomainProtocol, s.Domain)
}

func (s *ShortenerUseCase) Save(originalUrl string) (string, error) {
	const op = "urlShortener.Save"

	log := s.l.With(
		slog.String("op", op),
		slog.String("originalUrl", originalUrl),
	)

	aliasForURLExists, err := s.UrlRepo.AliasForURLExists(originalUrl)
	if err != nil {
		return "", err
	}
	if aliasForURLExists {
		log.Debug("alias for url already exists, returning")
		alias, err := s.UrlRepo.GetAlias(originalUrl)
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("%s://%s/%s", s.DomainProtocol, s.Domain, alias)
		return url, nil
	}

	newAlias := randString.GenerateString()
	log.Debug("generated alias")
	exists, err := s.UrlRepo.AliasExists(newAlias)
	log.Debug("checking if alias exists")
	if err != nil {
		return "", err
	}
	for exists {
		log.Debug("alias already exists, generating new one")
		newAlias = randString.GenerateString()
		exists, err = s.UrlRepo.AliasExists(newAlias)
		if err != nil {
			return "", err
		}
	}

	log.Debug("alias is OK, saving...")
	savedAlias, err := s.UrlRepo.Save(originalUrl, newAlias)
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("%s://%s/%s", s.DomainProtocol, s.Domain, savedAlias)

	return url, nil
}

func (s *ShortenerUseCase) GetOriginal(alias string) (string, error) {
	const op = "urlShortener.GetOriginal"

	log := s.l.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	aliasExists, err := s.UrlRepo.AliasExists(alias)
	log.Debug("checking if alias exists")
	if err != nil {
		log.Error("failed to check if alias exist")
		return "", err
	}

	if !aliasExists {
		log.Debug("url with this alias does not exist")
		return "", cerrors.ErrAliasForURLDoesNotExist
	}

	originalURL, err := s.UrlRepo.GetOriginal(alias)
	log.Debug("getting original url")
	if err != nil {
		log.Error("failed to get original url")
		return "", err
	}

	return originalURL, nil
}
