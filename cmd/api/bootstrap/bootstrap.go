package bootstrap

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"restaurant-order-system/internal/auth"
	"restaurant-order-system/internal/menu"
	"restaurant-order-system/internal/middleware"
	"restaurant-order-system/migrations"
)

type RegisterDependencies struct {
	authRepo    auth.Repository
	menuRepo    menu.Repository
	authService auth.Service
	menuService menu.Service
	authHandler *auth.Handler
	menuHandler *menu.Handler
}

func NewRegisterDependencies(db *gorm.DB) *RegisterDependencies {
	authRepo := auth.NewPostgresRepository(db)
	menuRepo := menu.NewPostgresRepository(db)

	authService := auth.NewAuthService(authRepo)
	menuService := menu.NewMenuService(menuRepo)

	authHandler := auth.NewHandler(authService)
	menuHandler := menu.NewMenuHandler(menuService)

	return &RegisterDependencies{
		authRepo:    authRepo,
		menuRepo:    menuRepo,
		authService: authService,
		menuService: menuService,
		authHandler: authHandler,
		menuHandler: menuHandler,
	}
}

func Run(app *fiber.App) error {
	db, err := initiateDatabase()
	if err != nil {
		return fmt.Errorf("failed to connect to the database: %w", err)
	}
	if err := runMigrations(db); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	nd := NewRegisterDependencies(db)

	setupRoutes(app, nd)

	return nil
}

func setupRoutes(app *fiber.App, nd *RegisterDependencies) {
	api := app.Group("/api")

	authRoute := api.Group("/auth")
	auth.Routes(authRoute, nd.authHandler)

	menuRoute := api.Group("/menu")
	menuRoute.Use(middleware.AuthMiddleware("secret"))
	menu.Routes(menuRoute, nd.menuHandler)
}

func initiateDatabase() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=password dbname=restaurant_db port=5432 sslmode=disable TimeZone=UTC"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed connection")
	}

	log.Println("Database connected successfully")
	return db, nil
}

func runMigrations(db *gorm.DB) error {
	if err := migrations.Migrate(db); err != nil {
		return errors.New("failed to run migrations")
	}

	log.Println("Migrations run successfully")
	return nil
}
