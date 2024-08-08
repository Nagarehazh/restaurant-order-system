package server

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	httpAddr     string
	app          *fiber.App
	dependencies *Dependencies
}

func New(host string, port uint, db *gorm.DB) *Server {
	return &Server{
		httpAddr:     fmt.Sprintf("%s:%d", host, port),
		app:          fiber.New(),
		dependencies: NewDependencies(db),
	}
}

func (s *Server) Run() error {
	SetupRoutes(s.app, s.dependencies)
	return s.app.Listen(s.httpAddr)
}
