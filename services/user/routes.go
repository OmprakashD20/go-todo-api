package user

import "github.com/gofiber/fiber/v2"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SetupUserRoutes(api fiber.Router) {
	api.Post("/login", s.LoginHandler)
	api.Post("/register", s.RegisterHandler)
}

func (s *Service) LoginHandler(ctx *fiber.Ctx) error {
	return nil
}

func (s *Service) RegisterHandler(ctx *fiber.Ctx) error {
	return nil
}
