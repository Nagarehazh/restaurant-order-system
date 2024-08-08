package auth

import "github.com/gofiber/fiber/v2"

func Routes(authRoute fiber.Router, authHandler *Handler) {
	authRoute.Post("/register", authHandler.Register)
	authRoute.Post("/login", authHandler.Login)
}
