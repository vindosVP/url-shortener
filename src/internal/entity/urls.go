package entity

import "gorm.io/gorm"

type Url struct {
	gorm.Model
	ID    int    `gorm:"primary_key;unique;not_null;autoIncrement:true"`
	URL   string `gorm:"type:varchar(255);not_null;unique"`
	Alias string `gorm:"type:varchar(255);not_null;unique"`
}
