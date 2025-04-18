package role

import (
	"errors"

	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(role *model.Role) error {
	return r.db.Create(role).Error
}

func (r *Repository) Update(role *model.Role) error {
	return r.db.Save(role).Error
}

func (c *Repository) FindAll(opts *helper.FindAllOptions) ([]model.Role, int64, error) {
	var types []model.Role
	query := c.db.Model(model.Role{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&types).Error
	return types, total, err
}

func (c *Repository) FindById(id string) (*model.Role, error) {
	var Role model.Role
	err := c.db.First(&Role, "id=?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Role, err
}

func (c *Repository) FindByIdUnscoped(id string) (*model.Role, error) {
	var Role model.Role
	err := c.db.Unscoped().First(&Role, "id=?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &Role, err
}

func (c *Repository) HardDelete(id string) error {
	return c.db.Unscoped().Delete(&model.Role{}, "id=?", id).Error
}

func (c *Repository) Restore(id string) error {
	return c.db.Unscoped().
		Model(&model.Role{}).
		Where("id=?", id).
		Update("deleted_at", nil).Error
}

func (c *Repository) SoftDelete(id string) error {
	return c.db.Delete(&model.Role{}, "id = ?", id).Error
}
