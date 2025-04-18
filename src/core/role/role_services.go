package role

import (
	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/model"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) GetAll(opts *helper.FindAllOptions) ([]RoleResponse, int64, error) {
	roles, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, 0, err
	}
	dtos := RoleToListDTO(roles)
	return dtos, total, nil
}

func (s *Service) GetByID(id string) (*RoleResponse, error) {
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

func (s *Service) Create(input *CreateRoleRequest) (*RoleResponse, error) {
	m := &model.Role{
		Name:        input.Name,
		Description: input.Description,
		// Permissions: input.Permissions,
	}
	if err := s.repo.Create(m); err != nil {
		return nil, err
	}
	dto := RoleToDTO(m)
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

	if err := s.repo.Update(role); err != nil {
		return nil, err
	}
	dto := RoleToDTO(role)
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

func (s *Service) HardDelete(id string) (bool, error) {
	existed, err := s.repo.FindByIdUnscoped(id)
	if err != nil {
		return false, err
	}
	if existed == nil {
		return false, nil
	}
	if err := s.repo.HardDelete(id); err != nil {
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
