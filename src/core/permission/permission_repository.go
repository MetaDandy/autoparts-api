package permission

import (
	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll(opts *helper.FindAllOptions) ([]model.Permission, int64, error)
}

type repo struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repo{db: db}
}

func (c *repo) FindAll(opts *helper.FindAllOptions) ([]model.Permission, int64, error) {
	var types []model.Permission
	query := c.db.Model(model.Permission{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&types).Error
	return types, total, err
}
