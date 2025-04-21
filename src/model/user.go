package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CognitoUserID string    `gorm:"uniqueIndex;not null"`
	Name          string    `gorm:"not null"`
	Email         string    `gorm:"uniqueIndex;not null"`
	Phone         string    `gorm:"uniqueIndex;not null"`
	Address       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`

	RoleID uuid.UUID `gorm:"type:uuid;not null;index"`
	Role   Role      `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
