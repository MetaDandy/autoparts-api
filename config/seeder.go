package config

import (
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {
	if err := SeedPermissions(db); err != nil {
		log.Fatalf("Error al seedear permisos: %v", err)
	}
}
