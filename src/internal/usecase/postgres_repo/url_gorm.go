package postgres_repo

import (
	"fmt"
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
	"url-shortener/src/internal/entity"
)

type PostgresRepo struct {
	db *gorm.DB
	l  *slog.Logger
}

func New(db *gorm.DB, l *slog.Logger) *PostgresRepo {
	return &PostgresRepo{
		db: db,
		l:  l,
	}
}

func (r *PostgresRepo) GetAlias(originalUrl string) (string, error) {
	const op = "postgres_repo.GetAlias"

	log := r.l.With(
		slog.String("op", op),
		slog.String("originalUrl", originalUrl),
	)

	var alias string
	err := r.db.Model(&entity.Url{}).
		Select("urls.alias").
		Where("urls.url = ?", originalUrl).
		First(&alias).Error

	if err != nil {
		log.Error("failed to get url by alias")
	}

	return alias, err
}

func (r *PostgresRepo) AliasForURLExists(originalUrl string) (bool, error) {
	const op = "postgres_repo.AliasForURLExists"

	log := r.l.With(
		slog.String("op", op),
		slog.String("originalUrl", originalUrl),
	)

	var count int64
	err := r.db.Model(&entity.Url{}).Where("url = ?", originalUrl).Count(&count).Error

	if err != nil {
		log.Error("failed to check if alias for url exists")
	}

	return count > 0, err
}

func (r *PostgresRepo) AliasExists(alias string) (bool, error) {
	const op = "postgres_repo.AliasExists"

	log := r.l.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	var count int64
	err := r.db.Model(&entity.Url{}).Where("alias = ?", alias).Count(&count).Error

	if err != nil {
		log.Error("failed to check if alias exists")
	}

	return count > 0, err
}

func (r *PostgresRepo) Save(originalUrl string, alias string) (string, error) {
	const op = "postgres_repo.Save"

	log := r.l.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	URL := &entity.Url{
		URL:   originalUrl,
		Alias: alias,
	}
	err := r.db.Create(URL).Error

	if err != nil {
		log.Error("failed to save alias")
	} else {
		log.Debug("alias saved")
	}

	return alias, err
}

func (r *PostgresRepo) GetOriginal(alias string) (string, error) {
	const op = "postgres_repo.GetOriginal"

	log := r.l.With(
		slog.String("op", op),
		slog.String("alias", alias),
	)

	var originalURL string
	err := r.db.Model(&entity.Url{}).
		Select("urls.url").
		Where("urls.alias = ?", alias).
		First(&originalURL).Error

	if err != nil {
		log.Error("failed to get original url")
	} else {
		log.Debug(fmt.Sprintf("got original url %s", originalURL))
	}

	return originalURL, err
}
