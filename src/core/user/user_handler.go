package user

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
	grp := router.Group("/users")
	grp.Get("", h.FindAll)
	grp.Get("/:id", h.FindOne)
	grp.Post("/signup", h.Create)
	grp.Post("/signin", h.SignIn)
	grp.Post("/refresh", h.Refresh)
	grp.Patch("/:id", h.Update)
	grp.Delete("/:id", h.SoftDelete)
	grp.Post("/restore/:id", h.Restore)
}

func (h *Handler) SignIn(c *fiber.Ctx) error {
	var req SignInRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.JSONError(c, http.StatusBadRequest, "Input inválido", err.Error())
	}
	tokens, err := h.svc.Login(req.Email, req.Password)
	if err != nil {
		return helper.JSONError(c, http.StatusUnauthorized, "Credenciales inválidas", err.Error())
	}
	resp := AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		IDToken:      tokens.IDToken,
		ExpiresIn:    tokens.ExpiresIn,
	}
	return c.JSON(helper.Response{
		Data:    resp,
		Message: "Login exitoso",
	})
}

func (h *Handler) Refresh(c *fiber.Ctx) error {
	var req RefreshRequest
	if err := c.BodyParser(&req); err != nil {
		return helper.JSONError(c, http.StatusBadRequest, "Input inválido", err.Error())
	}
	tokens, err := h.svc.Refresh(req.RefreshToken)
	if err != nil {
		return helper.JSONError(c, http.StatusUnauthorized, "Refresh token inválido", err.Error())
	}
	resp := AuthResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
		IDToken:      tokens.IDToken,
		ExpiresIn:    tokens.ExpiresIn,
	}
	return c.JSON(helper.Response{
		Data:    resp,
		Message: "Token renovado",
	})
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	users, err := h.svc.FindAll(opts)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo usuarios", err.Error())
	}

	return c.JSON(users)
}

func (h *Handler) FindOne(c *fiber.Ctx) error {
	dto, err := h.svc.FindByID(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error obteniendo usuario", err.Error())
	}
	if dto == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Usuario no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    dto,
		Message: "Usuario encontrado",
	})
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var input CreateUserRequest
	if err := c.BodyParser(&input); err != nil {
		return helper.JSONError(c, http.StatusBadRequest,
			"Input inválido", err.Error())
	}
	user, err := h.svc.SignUp(&input)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error creando usuario", err.Error())
	}
	return c.Status(http.StatusCreated).JSON(helper.Response{
		Data:    user,
		Message: "Usuario creado",
	})
}

func (h *Handler) Update(c *fiber.Ctx) error {
	var input UpdateUserRequest
	if err := c.BodyParser(&input); err != nil {
		return helper.JSONError(c, http.StatusBadRequest,
			"Cuerpo inválido", err.Error())
	}
	user, err := h.svc.Update(c.Params("id"), &input)
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error actualizando usuario", err.Error())
	}
	if user == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Usuario no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    user,
		Message: "Usuario actualizado",
	})
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	ok, err := h.svc.SoftDelete(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error eliminando usuario", err.Error())
	}
	if !ok {
		return helper.JSONError(c, http.StatusNotFound, "Usuario no encontrado")
	}
	return c.SendStatus(http.StatusNoContent)
}

func (h *Handler) Restore(c *fiber.Ctx) error {
	user, err := h.svc.Restore(c.Params("id"))
	if err != nil {
		return helper.JSONError(c, http.StatusInternalServerError,
			"Error restaurando usuario", err.Error())
	}
	if user == nil {
		return helper.JSONError(c, http.StatusNotFound,
			"Usuario no encontrado")
	}
	return c.JSON(helper.Response{
		Data:    user,
		Message: "Usuario restaurado",
	})
}
