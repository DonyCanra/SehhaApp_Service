package config

import (
	"gorm.io/gorm"
	"sehha/domain"
)

func Migration(db *gorm.DB) {
	err := db.AutoMigrate(
		&domain.Agency{},
		&domain.User{})
	if err != nil {
		return
	}
}
