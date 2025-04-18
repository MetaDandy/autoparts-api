package permission

import (
	"github.com/MetaDandy/autoparts-api/helper"
)

// Service orchestrates business logic for permissions.
type Service struct {
	repo Repository
}

// NewService constructs a Permission service with the given repository.
func NewService(r Repository) *Service {
	return &Service{repo: r}
}

// FindAll returns a paginated set of PermissionResponse DTOs.
func (s *Service) FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[PermissionResponse], error) {
	roles, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := PermissionToListDTO(roles)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[PermissionResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}
