// Package role contains the domain logic, repository,
// services and HTTP handlers for the Role entity.
package role

import (
	"time"

	"github.com/MetaDandy/autoparts-api/src/core/permission"
	"github.com/MetaDandy/autoparts-api/src/model"
)

type CreateRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Permissions []string `json:"permissions" validate:"required,dive,uuid"`
}

type UpdateRoleRequest struct {
	Name        *string   `json:"name,omitempty"        validate:"omitempty,min=1"`
	Description *string   `json:"description,omitempty" validate:"omitempty,min=1"`
	Permissions *[]string `json:"permissions,omitempty" validate:"omitempty,dive,uuid"`
}

// RoleResponse defines the fields returned in a role response.
type RoleResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	Permissions []permission.PermissionResponse `json:"permissions"`
	// Users       []UserResponse       `json:"users,omitempty"`
}

// RoleToDTO maps a model.Role to a RoleResponse DTO.
func RoleToDTO(r *model.Role) RoleResponse {
	var deleted_at *time.Time
	if r.DeletedAt.Valid {
		t := r.DeletedAt.Time
		deleted_at = &t
	}

	perms := make([]permission.PermissionResponse, len(r.Permissions))
	for i := range r.Permissions {
		perms[i] = permission.PermissionToDTO(&r.Permissions[i])
	}

	return RoleResponse{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   deleted_at,
		Permissions: perms,
		// Users:       usersDTO,  // si incluyes usuarios
	}
}

// RoleToListDTO converts a slice of model.Role into a slice
// of RoleResponse DTOs.
func RoleToListDTO(list []model.Role) []RoleResponse {
	out := make([]RoleResponse, len(list))
	for i, c := range list {
		out[i] = RoleToDTO(&c)
	}
	return out
}
