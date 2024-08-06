package menu

import "github.com/gofiber/fiber/v2"

func Routes(menuRoute fiber.Router, menuHandler *Handler) {
	menuRoute.Post("/", menuHandler.CreateItem)
	menuRoute.Get("/:id", menuHandler.GetItem)
	menuRoute.Put("/:id", menuHandler.UpdateItem)
	menuRoute.Delete("/:id", menuHandler.DeleteItem)
	menuRoute.Get("/", menuHandler.ListItems)
}
