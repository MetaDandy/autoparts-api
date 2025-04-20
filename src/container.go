package src

import (
	"github.com/MetaDandy/autoparts-api/config"
	"github.com/MetaDandy/autoparts-api/src/core/permission"
	"github.com/MetaDandy/autoparts-api/src/core/role"
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

	return &Container{
		// Permission
		PermissionRepo:    &permisionRepo,
		PermissionSvc:     permissionSvc,
		PermissionHandler: permissionHandler,

		// Role
		RoleRepo:    roleRepo,
		RoleSvc:     roleSvc,
		RoleHandler: roleHandler,
	}
}
