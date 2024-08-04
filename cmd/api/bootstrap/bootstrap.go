package bootstrap

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"log"
	"restaurant-order-system/internal/auth"
)

func SetupApp(app *fiber.App) error {
	// Initialize database
	db, err := sql.Open("postgres", "postgres://postgres:password@postgres/restaurant_db?sslmode=disable")
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize repositories
	authRepo := auth.NewPostgresRepository(db)

	// Initialize services
	authService := auth.NewAuthService(authRepo)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)

	// Setup routes and middleware
	setupRoutes(app, authHandler)

	return nil
}

func setupRoutes(app *fiber.App, authHandler *auth.Handler) {
	api := app.Group("/api")

	// Auth routes
	authRoute := api.Group("/auth")
	authRoute.Post("/register", authHandler.Register)
	authRoute.Post("/login", authHandler.Login)

	// Other routes will be added here later
}
