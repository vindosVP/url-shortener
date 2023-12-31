package database

import (
	"fmt"
	"github.com/vindosVP/url-shortener/src/internal/config"
	"github.com/vindosVP/url-shortener/src/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGorm(cfg config.DB) (*gorm.DB, error) {
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

	if err := db.AutoMigrate(&entity.Url{}); err != nil {
		return err
	}

	return nil
}

func generateGormDNS(cfg config.DB) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.Port, cfg.User, cfg.Pwd, cfg.Name, cfg.SSLMode, cfg.TimeZone)
}
