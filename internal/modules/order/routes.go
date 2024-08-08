package order

import "github.com/gofiber/fiber/v2"

func Routes(orderRoute fiber.Router, orderHandler *Handler) {
	orderRoute.Post("/", orderHandler.CreateOrder)
	orderRoute.Get("/:id", orderHandler.GetOrder)
	orderRoute.Get("/", orderHandler.ListOrders)
	orderRoute.Put("/:id", orderHandler.UpdateOrderStatus)
}
