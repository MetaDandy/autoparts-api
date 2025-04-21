package user

import (
	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/MetaDandy/autoparts-api/src/core/auth"
	"github.com/MetaDandy/autoparts-api/src/core/role"
	"github.com/MetaDandy/autoparts-api/src/model"
	"github.com/google/uuid"
)

type Service struct {
	repo     *Repository
	roleRepo *role.Repository
	provider *auth.CognitoProvider
}

func NewService(r *Repository, p *auth.CognitoProvider) *Service {
	return &Service{repo: r, provider: p}
}

func (s *Service) SignUp(req *CreateUserRequest) (*model.User, error) {
	role, err := s.roleRepo.FindById(req.RoleID)
	if err != nil {
		return nil, err
	}
	sub, err := s.provider.SignUpUser(req.Email, req.Password, req.Name, req.Phone)
	if err != nil {
		return nil, err
	}

	u := &model.User{
		CognitoUserID: sub,
		Name:          req.Name,
		Email:         req.Email,
		Phone:         req.Phone,
		Address:       req.Address,
		RoleID:        uuid.MustParse(req.RoleID),
		Role:          *role,
	}
	if err := s.repo.Create(u); err != nil {
		_ = s.provider.DeleteUser(req.Email)
		return nil, err
	}
	return u, nil
}

func (s *Service) Login(email, password string) (auth.AuthTokens, error) {
	return s.provider.SignInUser(email, password)
}

func (s *Service) Refresh(token string) (auth.AuthTokens, error) {
	return s.provider.RefreshToken(token)
}

func (s *Service) FindAll(opts *helper.FindAllOptions) (*helper.PaginatedResponse[UserResponse], error) {
	users, total, err := s.repo.FindAll(opts)
	if err != nil {
		return nil, err
	}
	dtos := UsersToListDTO(users)
	pages := uint((total + int64(opts.Limit) - 1) / int64(opts.Limit))

	return &helper.PaginatedResponse[UserResponse]{
		Data:   dtos,
		Total:  total,
		Limit:  opts.Limit,
		Offset: opts.Offset,
		Pages:  pages,
	}, nil
}

func (s *Service) FindByID(id string) (*UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	dto := UserToDTO(user)
	return &dto, nil
}

func (s *Service) Update(id string, input *UpdateUserRequest) (*UserResponse, error) {
	user, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	if input.RoleID != nil {
		role, err := s.roleRepo.FindById(*input.RoleID)
		if err != nil {
			return nil, err
		}
		user.RoleID = uuid.MustParse(*input.RoleID)
		user.Role = *role
	}
	if input.Name != nil {
		user.Name = *input.Name
	}
	if input.Phone != nil {
		user.Phone = *input.Phone
	}
	if input.Address != nil {
		user.Address = *input.Address
	}

	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	reloaded, _ := s.repo.FindById(id)
	dto := UserToDTO(reloaded)
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

func (s *Service) Restore(id string) (*UserResponse, error) {
	user, err := s.repo.FindByIdUnscoped(id)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	if err := s.repo.Restore(id); err != nil {
		return nil, err
	}
	dto := UserToDTO(user)
	return &dto, nil
}
