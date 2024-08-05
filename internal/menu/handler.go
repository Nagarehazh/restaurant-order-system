package menu

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Handler struct {
	service Service
}

func NewMenuHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateMenuItem(c *fiber.Ctx) error {
	var req MenuItem
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateMenuItem(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create menu item"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu item created successfully"})
}

func (h *Handler) GetMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	item, err := h.service.GetMenuItem(uint(idParsed))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu item not found"})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *Handler) UpdateMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	item, err := h.service.GetMenuItem(uint(idParsed))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu item not found"})
	}

	var req MenuItem
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	item.Name = req.Name
	item.Price = req.Price
	item.Description = req.Description

	if err := h.service.UpdateMenuItem(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update menu item"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Menu item updated successfully"})
}

func (h *Handler) DeleteMenuItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.service.DeleteMenuItem(uint(idParsed)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete menu item"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Menu item deleted successfully"})
}

func (h *Handler) ListMenuItems(c *fiber.Ctx) error {
	items, err := h.service.ListMenuItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list menu items"})
	}

	return c.Status(fiber.StatusOK).JSON(items)
}
