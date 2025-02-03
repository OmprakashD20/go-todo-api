package api

import (
	"log"

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

	log.Printf("Server is running on %s", s.addr)

	return app.Listen(s.addr)
}
