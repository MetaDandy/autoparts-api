package api

import (
	"github.com/MetaDandy/autoparts-api/src"
	"github.com/gofiber/fiber/v2"
)

// SetupApi mounts all module routes under /api/v1.
func SetupApi(app *fiber.App, c *src.Container) {
	api := app.Group("/api/v1")

	handlers := []func(fiber.Router){
		c.PermissionHandler.RegisterRoutes,
		c.RoleHandler.RegisterRoutes,
	}

	for _, register := range handlers {
		register(api)
	}
}
