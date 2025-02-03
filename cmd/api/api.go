package api

import (
	"log"

	"github.com/OmprakashD20/go-todo-api/services/user"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type APIServer struct {
	addr string
	db   *gorm.DB
}

func NewAPIServer(addr string, db *gorm.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	app := fiber.New()
	api := app.Group("/api/v1")

	userService := user.NewService()
	userService.SetupUserRoutes(api)

	log.Printf("Server is running on %s", s.addr)

	return app.Listen(s.addr)
}
