// src/core/user/dto.go
package user

import (
	"time"

	"github.com/MetaDandy/autoparts-api/src/core/role"
	"github.com/MetaDandy/autoparts-api/src/model"
)

type CreateUserRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Name     string `json:"name"     validate:"required"`
	Phone    string `json:"phone" validate:"required,e164,regex=^\\+591\\d{7,8}$"`
	Address  string `json:"address"  validate:"required"`
	RoleID   string `json:"role_id"  validate:"required,uuid"`
}

type UpdateUserRequest struct {
	Name    *string `json:"name,omitempty"    validate:"omitempty,min=1"`
	Phone   *string `json:"phone,omitempty"   validate:"omitempty,e164"`
	Address *string `json:"address,omitempty" validate:"omitempty"`
	RoleID  *string `json:"role_id,omitempty" validate:"omitempty,uuid"`
}

type UserResponse struct {
	ID            string            `json:"id"`
	CognitoUserID string            `json:"cognito_user_id"`
	Email         string            `json:"email"`
	Name          string            `json:"name"`
	Phone         string            `json:"phone"`
	Address       string            `json:"address"`
	CreatedAt     time.Time         `json:"created_at"`
	UpdatedAt     time.Time         `json:"updated_at"`
	DeletedAt     *time.Time        `json:"deleted_at,omitempty"`
	Role          role.RoleResponse `json:"role"`
}

func UserToDTO(u *model.User) UserResponse {
	var deletedAt *time.Time
	if u.DeletedAt.Valid {
		t := u.DeletedAt.Time
		deletedAt = &t
	}

	return UserResponse{
		ID:            u.ID.String(),
		CognitoUserID: u.CognitoUserID,
		Email:         u.Email,
		Name:          u.Name,
		Phone:         u.Phone,
		Address:       u.Address,
		CreatedAt:     u.CreatedAt,
		UpdatedAt:     u.UpdatedAt,
		DeletedAt:     deletedAt,
		Role:          role.RoleToDTO(&u.Role),
	}
}

func UsersToListDTO(list []model.User) []UserResponse {
	out := make([]UserResponse, len(list))
	for i := range list {
		out[i] = UserToDTO(&list[i])
	}
	return out
}

type SignInRequest struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	ExpiresIn    int32  `json:"expires_in"`
}
