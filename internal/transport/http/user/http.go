package http

import (
	"DDD-fiberv2/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *user.Service
}

func NewHandler(service *user.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoute(app *fiber.App) {
	app.Post("/users", h.Create)
	app.Get("/users/:id", h.GetByID)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}
	user, err := h.service.Create(c.Context(), req.Name, req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	user, err := h.service.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not Found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *Handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	user, err := h.service.Update(c.Context(), id, req.Name, req.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(user)
}

func (h *Handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.service.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	users, err := h.service.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
