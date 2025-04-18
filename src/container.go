package src

import (
	"github.com/MetaDandy/autoparts-api/config"
	"github.com/MetaDandy/autoparts-api/src/core/permission"
)

// Container holds all module dependencies.
type Container struct {
	// Permissions
	PermissionRepo    *permission.Repository
	PermissionSvc     *permission.Service
	PermissionHandler *permission.Handler
}

// SetUpContainer wires repositories, services and handlers.
func SetUpContainer() *Container {
	permisionRepo := permission.NewRepository(config.DB)
	permissionSvc := permission.NewService(permisionRepo)
	permissionHandler := permission.NewHandler(permissionSvc)

	return &Container{
		PermissionRepo:    &permisionRepo,
		PermissionSvc:     permissionSvc,
		PermissionHandler: permissionHandler,
	}
}
