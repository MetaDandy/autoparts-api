package src

import (
	"github.com/MetaDandy/autoparts-api/config"
	"github.com/MetaDandy/autoparts-api/src/core/auth"
	"github.com/MetaDandy/autoparts-api/src/core/permission"
	"github.com/MetaDandy/autoparts-api/src/core/role"
	"github.com/MetaDandy/autoparts-api/src/core/user"
)

// Container holds all module dependencies.
type Container struct {
	// Permissions
	PermissionRepo    *permission.Repository
	PermissionSvc     *permission.Service
	PermissionHandler *permission.Handler

	// Role
	RoleRepo    *role.Repository
	RoleSvc     *role.Service
	RoleHandler *role.Handler

	// User
	UserRepo    *user.Repository
	UserSvc     *user.Service
	UserHandler *user.Handler
}

// SetUpContainer wires repositories, services and handlers.
func SetUpContainer() *Container {
	// Permission
	permisionRepo := permission.NewRepository(config.DB)
	permissionSvc := permission.NewService(permisionRepo)
	permissionHandler := permission.NewHandler(permissionSvc)

	// Role
	roleRepo := role.NewRepository(config.DB)
	roleSvc := role.NewService(roleRepo, permisionRepo)
	roleHandler := role.NewHandler(roleSvc)

	// User
	userRepo := user.NewRepository(config.DB)
	userSvc := user.NewService(userRepo, &auth.CognitoProvider{})
	userHandler := user.NewHandler(userSvc)

	return &Container{
		// Permission
		PermissionRepo:    &permisionRepo,
		PermissionSvc:     permissionSvc,
		PermissionHandler: permissionHandler,

		// Role
		RoleRepo:    roleRepo,
		RoleSvc:     roleSvc,
		RoleHandler: roleHandler,

		// User
		UserRepo:    userRepo,
		UserSvc:     userSvc,
		UserHandler: userHandler,
	}
}
