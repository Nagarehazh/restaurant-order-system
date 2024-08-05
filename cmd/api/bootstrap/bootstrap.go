package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"restaurant-order-system/internal/auth"
	"restaurant-order-system/internal/menu"
	"restaurant-order-system/migrations"
)

func SetupApp(app *fiber.App) error {
	// Connect to the database
	dsn := "host=postgres user=postgres password=password dbname=restaurant_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := migrations.Migrate(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	} else {
		log.Println("Migrations ran successfully")
	}

	// Initialize repositories
	authRepo := auth.NewPostgresRepository(db)
	menuRepo := menu.NewPostgresRepository(db)

	// Initialize services
	authService := auth.NewAuthService(authRepo)
	menuService := menu.NewMenuService(menuRepo)

	// Initialize handlers
	authHandler := auth.NewHandler(authService)
	menuHandler := menu.NewMenuHandler(menuService)

	// Setup routes and middleware
	setupRoutes(app, authHandler, menuHandler)

	return nil
}

func setupRoutes(app *fiber.App, authHandler *auth.Handler, menuHandler *menu.Handler) {
	api := app.Group("/api")

	// Auth routes
	authRoute := api.Group("/auth")
	authRoute.Post("/register", authHandler.Register)
	authRoute.Post("/login", authHandler.Login)

	// Menu routes
	menuRoute := api.Group("/menu")
	menuRoute.Post("/", menuHandler.CreateMenuItem)
	menuRoute.Get("/:id", menuHandler.GetMenuItem)
	menuRoute.Put("/:id", menuHandler.UpdateMenuItem)
	menuRoute.Delete("/:id", menuHandler.DeleteMenuItem)
	menuRoute.Get("/", menuHandler.ListMenuItems)
}
