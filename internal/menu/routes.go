package menu

import "github.com/gofiber/fiber/v2"

func Routes(menuRoute fiber.Router, menuHandler *Handler) {
	menuRoute.Post("/", menuHandler.CreateMenuItem)
	menuRoute.Get("/:id", menuHandler.GetMenuItem)
	menuRoute.Put("/:id", menuHandler.UpdateMenuItem)
	menuRoute.Delete("/:id", menuHandler.DeleteMenuItem)
	menuRoute.Get("/", menuHandler.ListMenuItems)
}
