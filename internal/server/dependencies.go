package server

import (
	"gorm.io/gorm"
	auth2 "restaurant-order-system/internal/modules/auth"
	menu2 "restaurant-order-system/internal/modules/menu"
	order2 "restaurant-order-system/internal/modules/order"
)

type Dependencies struct {
	authRepo  auth2.Repository
	menuRepo  menu2.Repository
	orderRepo order2.Repository

	authService  auth2.Service
	menuService  menu2.Service
	orderService order2.Service

	authHandler  *auth2.Handler
	menuHandler  *menu2.Handler
	orderHandler *order2.Handler
}

func NewDependencies(db *gorm.DB) *Dependencies {
	authRepo := auth2.NewPostgresRepository(db)
	menuRepo := menu2.NewPostgresRepository(db)
	orderRepo := order2.NewPostgresRepository(db)

	authService := auth2.NewServiceImpl(authRepo)
	menuService := menu2.NewServiceImpl(menuRepo)
	orderService := order2.NewServiceImpl(orderRepo, menuService)

	authHandler := auth2.NewHandler(authService)
	menuHandler := menu2.NewMenuHandler(menuService)
	orderHandler := order2.NewHandler(orderService)

	return &Dependencies{
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
