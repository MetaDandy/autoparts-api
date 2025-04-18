package model

import (
	"time"

	"github.com/google/uuid"
)

type Permission struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"not null"`
	Code      string    `gorm:"not null;unique"`
	CreatedAt time.Time

	Roles []Role `gorm:"many2many:role_permission;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
