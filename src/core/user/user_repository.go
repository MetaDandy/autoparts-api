package user

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

func (r *Repository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *Repository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (c *Repository) FindAll(opts *helper.FindAllOptions) ([]model.User, int64, error) {
	var user []model.User
	query := c.db.Model(model.User{})
	var total int64
	query, total = helper.ApplyFindAllOptions(query, opts)

	err := query.Find(&user).Error
	return user, total, err
}

func (c *Repository) FindById(id string) (*model.User, error) {
	var user model.User
	err := c.db.Preload("Permissions").First(&user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (c *Repository) FindByIdUnscoped(id string) (*model.User, error) {
	var user model.User
	err := c.db.Unscoped().Preload("Permissions").First(&user, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (c *Repository) SoftDelete(id string) error {
	return c.db.Delete(&model.User{}, "id = ?", id).Error
}

func (c *Repository) Restore(id string) error {
	return c.db.Unscoped().
		Model(&model.User{}).
		Where("id=?", id).
		Update("deleted_at", nil).Error
}
