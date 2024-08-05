package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"restaurant-order-system/cmd/api/bootstrap"
)

func main() {
	app := fiber.New()

	if err := bootstrap.Run(app); err != nil {
		log.Fatalf("Error setting up app: %v", err)
	}

	log.Fatal(app.Listen(":8080"))
}
