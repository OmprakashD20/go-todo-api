package api

import (
	"fmt"
	"log"

	"github.com/OmprakashD20/go-todo-api/services/todo"
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

	userStore := user.NewStore(s.db)
	userService := user.NewService(userStore)
	userService.SetupUserRoutes(api.Group("/user"))

	todoStore := todo.NewStore(s.db)
	todoService := todo.NewService(todoStore, userStore)
	todoService.SetupTodoRoutes(api.Group("/todo"))

	log.Printf("Server is running on %s", s.addr)

	return app.Listen(fmt.Sprintf(":%s", s.addr))
}
