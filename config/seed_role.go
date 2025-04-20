package config

import (
	"log"

	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

// SeedRoles asegura que exista un rol "Admin" ligado a todos los permisos.
func SeedRoles(db *gorm.DB) error {
	const adminRoleName = "Admin"

	var existing model.Role
	err := db.Preload("Permissions").
		Where("name = ?", adminRoleName).
		First(&existing).Error

	if err == nil {
		log.Printf("⚠️ Rol %q ya existe con %d permisos; skip.", adminRoleName, len(existing.Permissions))
		return err
	}
	if err != gorm.ErrRecordNotFound {
		log.Fatalf("❌ Error buscando rol %q: %v", adminRoleName, err)
	}

	var allPerms []model.Permission
	if err := db.Find(&allPerms).Error; err != nil {
		log.Fatalf("❌ Error cargando permisos para seed de rol: %v", err)
	}

	role := model.Role{
		Name:        adminRoleName,
		Description: "Administrador con todos los permisos",
	}

	if err := db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&role).Error; err != nil {
			return err
		}
		if err := tx.Model(&role).Association("Permissions").
			Replace(allPerms); err != nil {
			return err
		}
		return nil
	}); err != nil {
		log.Fatalf("❌ Error creando rol %q: %v", adminRoleName, err)
	}

	log.Printf("✅ Rol %q creado con %d permisos.", adminRoleName, len(allPerms))
	return nil
}
