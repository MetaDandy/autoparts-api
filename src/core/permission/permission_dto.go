// Package permission contains the domain logic, repository,
// services and HTTP handlers for the Permission entity.
package permission

import (
	"time"

	"github.com/MetaDandy/autoparts-api/src/model"
)

// PermissionResponse defines the fields returned in a permission response.
type PermissionResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// permissionToDTO maps a model.Permission to a PermissionResponse DTO.
func permissionToDTO(p *model.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        p.ID.String(),
		Name:      p.Name,
		Code:      p.Code,
		CreatedAt: p.CreatedAt,
	}
}

// PermissionToListDTO converts a slice of model.Permission into a slice
// of PermissionResponse DTOs.
func PermissionToListDTO(list []model.Permission) []PermissionResponse {
	out := make([]PermissionResponse, len(list))
	for i := range list {
		out[i] = permissionToDTO(&list[i])
	}
	return out
}
