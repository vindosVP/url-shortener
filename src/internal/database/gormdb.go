package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"url-shortener/src/internal/config"
	"url-shortener/src/internal/entity"
)

func NewGorm(cfg config.DB) (*gorm.DB, error) {
	const op = "database.gormdb.NewGorm"
	dns := generateGormDNS(cfg)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = autoMigrate(db)
	if err != nil {
		return nil, err
	}

	return db, err
}

func autoMigrate(db *gorm.DB) error {
	const op = "database.gormdb.autoMigrate"

	if err := db.AutoMigrate(&entity.Url{}); err != nil {
		return err
	}

	return nil
}

func generateGormDNS(cfg config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name, cfg.SSLMode, cfg.TimeZone)
}
