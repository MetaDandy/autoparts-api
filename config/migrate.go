package config

import (
	"log"

	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&model.Permission{},
		&model.Role{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}
}
