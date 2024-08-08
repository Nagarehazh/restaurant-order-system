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

func (h *Handler) CreateItem(c *fiber.Ctx) error {
	var req Item
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	if err := h.service.CreateItem(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create menu item"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Menu item created successfully"})
}

func (h *Handler) GetItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	item, err := h.service.GetItem(uint(idParsed))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu item not found"})
	}

	return c.Status(fiber.StatusOK).JSON(item)
}

func (h *Handler) UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	item, err := h.service.GetItem(uint(idParsed))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Menu item not found"})
	}

	var req Item
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	item.Name = req.Name
	item.Price = req.Price
	item.Description = req.Description

	if err := h.service.UpdateItem(item); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update menu item"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Menu item updated successfully"})
}

func (h *Handler) DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	idParsed, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := h.service.DeleteItem(uint(idParsed)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete menu item"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Menu item deleted successfully"})
}

func (h *Handler) ListItems(c *fiber.Ctx) error {
	items, err := h.service.ListItems()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list menu items"})
	}

	return c.Status(fiber.StatusOK).JSON(items)
}
