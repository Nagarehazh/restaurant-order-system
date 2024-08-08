package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"restaurant-order-system/internal/modules/auth"
)

var ErrMenuItemNotFound = "menu item not found"
var ErrOrderNotFound = "order not found"
var ErrOnlyPendingOrdersCanBeCancelled = "only pending orders can be cancelled"

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var request struct {
		Items []struct {
			MenuItemID uint `json:"menu_item_id"`
			Quantity   int  `json:"quantity"`
		} `json:"items"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request",
		})
	}

	user := c.Locals("user").(*auth.JWTClaim)
	var orderItems []Item
	for _, item := range request.Items {
		orderItems = append(orderItems, Item{
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
		})
	}

	order, err := h.service.CreateOrder(user.ID, orderItems)
	if err != nil {
		if err.Error() == ErrMenuItemNotFound {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid menu item ID",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

func (h *Handler) GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	order, err := h.service.GetOrder(uint(id))
	if err != nil {
		if err.Error() == ErrOrderNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(order)
}

func (h *Handler) UpdateOrderStatus(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		if err.Error() == ErrOrderNotFound {
			log.Errorf("Order not found: %v", err)
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Order not found",
			})
		}
		log.Errorf("Invalid order ID: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	statusQuery := c.Query("status")
	if statusQuery == "" {
		log.Error("Invalid status")
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid status",
		})
	}

	if err := h.service.UpdateOrderStatus(uint(id), Status(statusQuery)); err != nil {
		if err.Error() == ErrOnlyPendingOrdersCanBeCancelled {
			log.Errorf("Only pending orders can be cancelled: %v", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Only pending orders can be cancelled",
			})
		}
		log.Errorf("Failed to update order status: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order status updated",
	})
}

func (h *Handler) ListOrders(c *fiber.Ctx) error {
	user := c.Locals("user").(*auth.JWTClaim)

	orders, err := h.service.ListOrders(user.ID)
	if err != nil {
		log.Errorf("Failed to list orders: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to list orders"})
	}

	return c.JSON(orders)
}
