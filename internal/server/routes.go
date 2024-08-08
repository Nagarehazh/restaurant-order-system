package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	auth2 "restaurant-order-system/internal/modules/auth"
	menu2 "restaurant-order-system/internal/modules/menu"
	order2 "restaurant-order-system/internal/modules/order"
	"restaurant-order-system/internal/server/middleware"
)

func SetupRoutes(app *fiber.App, d *Dependencies) {
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(cors.New())
	app.Use(recover.New())
	app.Use(favicon.New())

	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint: "/health-check",
	}))

	api := app.Group("/api")

	authRoute := api.Group("/auth")
	auth2.Routes(authRoute, d.authHandler)

	menuRoute := api.Group("/menu")
	menuRoute.Use(middleware.AuthMiddleware("secret"))
	menu2.Routes(menuRoute, d.menuHandler)

	orderRoute := api.Group("/order")
	orderRoute.Use(middleware.AuthMiddleware("secret"))
	order2.Routes(orderRoute, d.orderHandler)
}
