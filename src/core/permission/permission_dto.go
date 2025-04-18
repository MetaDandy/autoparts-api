package permission

import (
	"time"

	"github.com/MetaDandy/autoparts-api/src/model"
)

type PermissionResponse struct {
	ID        string    `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func permissionToDTO(p *model.Permission) PermissionResponse {
	return PermissionResponse{
		ID:        p.ID.String(),
		Name:      p.Name,
		Code:      p.Code,
		CreatedAt: p.CreatedAt,
	}
}

func PermissionToListDTO(list []model.Permission) []PermissionResponse {
	out := make([]PermissionResponse, len(list))
	for i := range list {
		out[i] = permissionToDTO(&list[i])
	}
	return out
}
