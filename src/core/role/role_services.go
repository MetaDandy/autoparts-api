package role

import (
	"errors"

	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/core/permission"
	"github.com/MetaDandy/autoparts-api/src/model"
	"gorm.io/gorm"
)

type Service struct {
	repo     *Repository
	permRepo permission.Repository
}

func NewService(r *Repository, p permission.Repository) *Service {
	return &Service{repo: r, permRepo: p}
}

func (s *Service) FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[RoleResponse], error) {
	roles, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := RoleToListDTO(roles)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[RoleResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindByID(id string) (*RoleResponse, error) {
	role, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, nil
	}
	dto := RoleToDTO(role)
	return &dto, nil
}

func (s *Service) findPermissions(ids []string) ([]model.Permission, error) {
	perms, err := s.permRepo.FindByIDs(ids)
	if err != nil {
		return nil, err
	}
	if len(perms) != len(ids) {
		return nil, errors.New("one or more permissions not found")
	}

	return perms, nil
}

func (s *Service) Create(input *CreateRoleRequest) (*RoleResponse, error) {
	perms, err := s.findPermissions(input.Permissions)
	if err != nil {
		return nil, err
	}

	var role model.Role
	err = s.repo.db.Transaction(func(tx *gorm.DB) error {
		role = model.Role{
			Name:        input.Name,
			Description: input.Description,
		}
		if err := tx.Create(&role).Error; err != nil {
			return err
		}
		if err := tx.Model(&role).Association("Permissions").
			Replace(perms); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	loaded, _ := s.repo.FindById(role.ID.String())
	dto := RoleToDTO(loaded)
	return &dto, nil
}

func (s *Service) Update(id string, input *UpdateRoleRequest) (*RoleResponse, error) {
	role, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, nil
	}
	if input.Name != nil {
		role.Name = *input.Name
	}
	if input.Description != nil {
		role.Description = *input.Description
	}

	err = s.repo.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Save(&role).Error; err != nil {
			return err
		}

		if input.Permissions != nil {
			perms, err := s.findPermissions(*input.Permissions)
			if err != nil {
				return nil
			}
			if err := tx.Model(&role).Association("Permissions").
				Replace(perms); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	reloaded, _ := s.repo.FindById(id)
	dto := RoleToDTO(reloaded)
	return &dto, nil
}

func (s *Service) SoftDelete(id string) (bool, error) {
	existed, err := s.repo.FindById(id)
	if err != nil {
		return false, err
	}
	if existed == nil {
		return false, nil
	}
	if err := s.repo.SoftDelete(id); err != nil {
		return false, err
	}
	return true, nil
}

func (s *Service) Restore(id string) (*RoleResponse, error) {
	role, err := s.repo.FindByIdUnscoped(id)
	if err != nil {
		return nil, err
	}
	if role == nil {
		return nil, nil
	}
	if err := s.repo.Restore(id); err != nil {
		return nil, err
	}
	dto := RoleToDTO(role)
	return &dto, nil
}
