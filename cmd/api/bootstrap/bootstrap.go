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
	"restaurant-order-system/internal/order"
	"restaurant-order-system/migrations"
)

type RegisterDependencies struct {
	authRepo  auth.Repository
	menuRepo  menu.Repository
	orderRepo order.Repository

	authService  auth.Service
	menuService  menu.Service
	orderService order.Service

	authHandler  *auth.Handler
	menuHandler  *menu.Handler
	orderHandler *order.Handler
}

func NewRegisterDependencies(db *gorm.DB) *RegisterDependencies {
	authRepo := auth.NewPostgresRepository(db)
	menuRepo := menu.NewPostgresRepository(db)
	orderRepo := order.NewPostgresRepository(db)

	authService := auth.NewServiceImpl(authRepo)
	menuService := menu.NewServiceImpl(menuRepo)
	orderService := order.NewServiceImpl(orderRepo, menuService)

	authHandler := auth.NewHandler(authService)
	menuHandler := menu.NewMenuHandler(menuService)
	orderHandler := order.NewHandler(orderService)

	return &RegisterDependencies{
		authRepo:     authRepo,
		menuRepo:     menuRepo,
		orderRepo:    orderRepo,
		authService:  authService,
		menuService:  menuService,
		orderService: orderService,
		authHandler:  authHandler,
		menuHandler:  menuHandler,
		orderHandler: orderHandler,
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

	orderRoute := api.Group("/order")
	orderRoute.Use(middleware.AuthMiddleware("secret"))
	order.Routes(orderRoute, nd.orderHandler)
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
