package role

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
	grp := router.Group("/roles")
	grp.Get("", h.FindAll)
	grp.Get("/:id", h.FindOne)
	grp.Post("", h.Create)
	grp.Patch("/:id", h.Update)
	grp.Delete("/:id", h.SoftDelete)
	grp.Post("/restore/:id", h.Restore)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	roles, err := h.svc.FindAll(opts)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo roles", err.Error())
	}

	return c.JSON(roles)
}

func (h *Handler) FindOne(c *fiber.Ctx) error {
	dto, err := h.svc.FindByID(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo rol", err.Error())
	}
	if dto == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Rol no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    dto,
		Message: "Rol encontrado",
	})
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var input CreateRoleRequest
	if err := c.BodyParser(&input); err != nil {
		return helper.JSONError(c, http.StatusBadRequest,
			"Input inválido", err.Error())
	}
	role, err := h.svc.Create(&input)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error creando rol", err.Error())
	}
	return c.Status(http.StatusCreated).JSON(helper.Response{
		Data:    role,
		Message: "Rol creado",
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input UpdateRoleRequest
	if err := c.BodyParser(&input); err != nil {
		return helper.JSONError(c, http.StatusBadRequest,
			"Cuerpo inválido", err.Error())
	}
	role, err := h.svc.Update(c.Params("id"), &input)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error actualizando rol", err.Error())
	}
	if role == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Rol no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    role,
		Message: "Rol actualizado",
	})
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	ok, err := h.svc.SoftDelete(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error eliminando rol", err.Error())
	}
	if !ok {
		return helper.JSONError(c, http.StatusNotFound, "Rol no encontrado")
	}
	return c.SendStatus(http.StatusNoContent)
}

func (h *Handler) Restore(c *fiber.Ctx) error {
	role, err := h.svc.Restore(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error restaurando rol", err.Error())
	}
	if role == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Rol no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    role,
		Message: "Rol restaurado",
	})
}
