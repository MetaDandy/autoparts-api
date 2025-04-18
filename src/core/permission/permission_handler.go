package permission

import (
	"net/http"

	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/gofiber/fiber/v2"
)

// Handler exposes HTTP endpoints for permissions.
type Handler struct {
	svc *Service
}

// NewHandler returns a new HTTP handler for the Permission service.
func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

// RegisterRoutes registers the /permissions route on the given router.
func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/permissions", h.FindAll)
}

// FindAll handles GET /permissions, returns paginated permissions.
func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	permission, err := h.svc.FindAll(opts)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo permisos", err.Error())
	}

	return c.JSON(permission)
}
