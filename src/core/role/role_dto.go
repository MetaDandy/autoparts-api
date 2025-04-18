package role

import (
	"time"

	"github.com/MetaDandy/autoparts-api/src/model"
)

type CreateRoleRequest struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Permissions []string `json:"permissions" validate:"required,dive,uuid4"`
}

type UpdateRoleRequest struct {
	Name        *string   `json:"name,omitempty"        validate:"omitempty,min=1"`
	Description *string   `json:"description,omitempty" validate:"omitempty,min=1"`
	Permissions *[]string `json:"permissions,omitempty" validate:"omitempty,dive,uuid4"`
}

type RoleResponse struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`

	// Permissions []PermissionResponse  `json:"permissions"`
	// Users       []UserResponse       `json:"users,omitempty"`
}

func RoleToDTO(r *model.Role) RoleResponse {
	var deleted_at *time.Time
	if r.DeletedAt.Valid {
		t := r.DeletedAt.Time
		deleted_at = &t
	}

	// perms := make([]PermissionResponse, len(r.Permissions))
	// for i := range r.Permissions {
	//     perms[i] = permissionToDTO(&r.Permissions[i])
	// }

	return RoleResponse{
		ID:          r.ID.String(),
		Name:        r.Name,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		DeletedAt:   deleted_at,
		// Permissions: perms,
		// Users:       usersDTO,  // si incluyes usuarios
	}
}

func RoleToListDTO(list []model.Role) []RoleResponse {
	out := make([]RoleResponse, len(list))
	for i, c := range list {
		out[i] = RoleToDTO(&c)
	}
	return out
}
