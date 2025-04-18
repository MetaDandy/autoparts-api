package permission

import (
	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

// Repository defines data access methods for permissions.
type Repository interface {
	// FindAll returns a slice of Permission and the total count.
	FindAll(opts *helper.FindAllOptions) ([]model.Permission, int64, error)
}

type repo struct {
	db *gorm.DB
}

// NewRepository creates a new Permission repository backed by GORM.
func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

// FindAll applies filtering, sorting and pagination to retrieve Permissions.
func (c *repo) FindAll(opts *helper.FindAllOptions) ([]model.Permission, int64, error) {
	var types []model.Permission
	query := c.db.Model(model.Permission{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&types).Error
	return types, total, err
}
