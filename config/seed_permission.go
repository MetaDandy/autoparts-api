package config

import (
	"fmt"
	"strings"

	"github.com/MetaDandy/autoparts-api/src/model"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

var defaultPermissionCodes = []string{
	// Usuarios
	"user.create",
	"user.read",
	"user.update",
	"user.soft_delete",
	"user.hard_delete",
	"user.restore",
	// Roles
	"role.create",
	"role.read",
	"role.update",
	"role.soft_delete",
	"role.hard_delete",
	"role.restore",
	// Nota de Venta
	"sale_note.create",
	"sale_note.read",
	"sale_note.update",
	"sale_note.soft_delete",
	"sale_note.hard_delete",
	"sale_note.restore",
	// Detalle de Venta
	"sale_detail.create",
	"sale_detail.read",
	"sale_detail.update",
	"sale_detail.soft_delete",
	"sale_detail.hard_delete",
	"sale_detail.restore",
	// Nota de Ingreso
	"income_note.create",
	"income_note.read",
	"income_note.update",
	"income_note.soft_delete",
	"income_note.hard_delete",
	"income_note.restore",
	// Detalle de Ingreso
	"income_detail.create",
	"income_detail.read",
	"income_detail.update",
	"income_detail.soft_delete",
	"income_detail.hard_delete",
	"income_detail.restore",
	// Nota de Egreso
	"expense_note.create",
	"expense_note.read",
	"expense_note.update",
	"expense_note.soft_delete",
	"expense_note.hard_delete",
	"expense_note.restore",
	// Detalle de Egreso
	"expense_detail.create",
	"expense_detail.read",
	"expense_detail.update",
	"expense_detail.soft_delete",
	"expense_detail.hard_delete",
	"expense_detail.restore",
	// Depósito
	"warehouse.create",
	"warehouse.read",
	"warehouse.update",
	"warehouse.soft_delete",
	"warehouse.hard_delete",
	"warehouse.restore",
	// Producto-Depósito
	"product_warehouse.create",
	"product_warehouse.read",
	"product_warehouse.update",
	"product_warehouse.soft_delete",
	"product_warehouse.hard_delete",
	"product_warehouse.restore",
	// Producto
	"product.create",
	"product.read",
	"product.update",
	"product.soft_delete",
	"product.hard_delete",
	"product.restore",
	// Imágenes
	"image.create",
	"image.read",
	"image.update",
	"image.soft_delete",
	"image.hard_delete",
	"image.restore",
	// Tipo de Producto
	"product_type.create",
	"product_type.read",
	"product_type.update",
	"product_type.soft_delete",
	"product_type.hard_delete",
	"product_type.restore",
	// Tipo de Categoría
	"category_type.create",
	"category_type.read",
	"category_type.update",
	"category_type.soft_delete",
	"category_type.hard_delete",
	"category_type.restore",
	// Categoría
	"category.create",
	"category.read",
	"category.update",
	"category.soft_delete",
	"category.hard_delete",
	"category.restore",
	// Compatibilidad
	"compatibility.create",
	"compatibility.read",
	"compatibility.update",
	"compatibility.soft_delete",
	"compatibility.hard_delete",
	"compatibility.restore",
	// Modelo
	"model.create",
	"model.read",
	"model.update",
	"model.soft_delete",
	"model.hard_delete",
	"model.restore",
	// Marca
	"brand.create",
	"brand.read",
	"brand.update",
	"brand.soft_delete",
	"brand.hard_delete",
	"brand.restore",
	// Métricas
	"metrics.create",
	"metrics.read",
	"metrics.update",
	"metrics.soft_delete",
	"metrics.hard_delete",
	"metrics.restore",
	// Características
	"characteristics.create",
	"characteristics.read",
	"characteristics.update",
	"characteristics.soft_delete",
	"characteristics.hard_delete",
	"characteristics.restore",
}

var titleCaser = cases.Title(language.English)

func SeedPermissions(db *gorm.DB) error {
	for _, code := range defaultPermissionCodes {

		parts := strings.Split(code, ".")
		var name string

		if len(parts) == 2 {
			action := titleCaser.String(parts[1])

			entity := titleCaser.String(strings.ReplaceAll(parts[0], "_", " "))
			name = fmt.Sprintf("%s %s", action, entity)
		} else {
			name = titleCaser.String(strings.ReplaceAll(code, ".", " "))
		}

		perm := model.Permission{
			Code: code,
			Name: name,
		}
		if err := db.
			Where("code = ?", code).
			FirstOrCreate(&perm).Error; err != nil {
			return err
		}
	}
	return nil
}
