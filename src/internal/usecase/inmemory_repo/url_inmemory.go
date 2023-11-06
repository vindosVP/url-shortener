package inmemory_repo

import (
	"fmt"
	"github.com/vindosVP/url-shortener/src/internal/cerrors"
	"golang.org/x/exp/slog"
)

type InmemoryRepo struct {
	originalToAlias map[string]string
	aliasToOriginal map[string]string
	l               *slog.Logger
}

func New(l *slog.Logger) *InmemoryRepo {
	return &InmemoryRepo{
		originalToAlias: make(map[string]string),
		aliasToOriginal: make(map[string]string),
		l:               l,
	}
}

func (r *InmemoryRepo) GetAlias(originalUrl string) (string, error) {
	const op = "inmemory_repo.GetAlias"

	log := r.l.With(
		slog.String("op", op),
		slog.String("originalUrl", originalUrl),
	)

	alias, exists := r.originalToAlias[originalUrl]
	if !exists {
		log.Error("alias for provided url does not exist")
		return "", cerrors.ErrAliasForURLDoesNotExist
	}

	return alias, nil
}

func (r *InmemoryRepo) AliasExists(alias string) (bool, error) {
	_, exists := r.aliasToOriginal[alias]
	return exists, nil
}

func (r *InmemoryRepo) AliasForURLExists(originalUrl string) (bool, error) {
	_, exists := r.originalToAlias[originalUrl]
	return exists, nil
}

func (r *InmemoryRepo) Save(originalUrl string, alias string) (string, error) {
	const op = "inmemory_repo.Save"

	log := r.l.With(
		slog.String("op", op),
		slog.String("originalUrl", originalUrl),
		slog.String("alias", alias),
	)

	_, exists := r.originalToAlias[originalUrl]
	if exists {
		log.Error(cerrors.ErrAliasAlreadySaved.Error())
		return "", cerrors.ErrAliasAlreadySaved
	}

	r.originalToAlias[originalUrl] = alias
	r.aliasToOriginal[alias] = originalUrl
	log.Debug("alias saved")

	return r.originalToAlias[originalUrl], nil
}

func (r *InmemoryRepo) GetOriginal(alias string) (string, error) {
	const op = "inmemory_repo.GetOriginal"

	log := r.l.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	originalUrl, exists := r.aliasToOriginal[alias]
	if !exists {
		log.Error(cerrors.ErrAliasDoesNotExist.Error())
		return "", cerrors.ErrAliasDoesNotExist
	}

	log.Debug(fmt.Sprintf("got original url %s", originalUrl))
	return originalUrl, nil
}
