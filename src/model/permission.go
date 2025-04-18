package model

import (
	"time"

	"github.com/google/uuid"
)

// Permission represents an access control right in the system.
type Permission struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Code      string    `gorm:"not null;unique"`
	CreatedAt time.Time

	// Roles is the many‑to‑many join with roles.
	Roles []Role `gorm:"many2many:role_permission;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
