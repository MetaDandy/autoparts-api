package permission

import (
	"net/http"

	"github.com/MetaDandy/autoparts-api/helper"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	svc *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	router.Get("/permissions", h.FindAll)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	permission, err := h.svc.FindAll(opts)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo permisos", err.Error())
	}

	return c.JSON(permission)
}
