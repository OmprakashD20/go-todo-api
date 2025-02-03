package user

import (
	"fmt"
	"net/http"

	"github.com/OmprakashD20/go-todo-api/models"
	"github.com/OmprakashD20/go-todo-api/types"
	"github.com/OmprakashD20/go-todo-api/utils"
	"github.com/OmprakashD20/go-todo-api/validations"

	"github.com/gofiber/fiber/v2"
)

type Service struct {
	store types.UserStore
}

func NewService(store types.UserStore) *Service {
	return &Service{store}
}

func (s *Service) SetupUserRoutes(api fiber.Router) {
	api.Post("/login", validator.ValidateSchema[types.LoginUserPayload](*validator.LoginUserSchema), s.LoginUserHandler)
	api.Post("/register", validator.ValidateSchema[types.RegisterUserPayload](*validator.RegisterUserSchema), s.RegisterUserHandler)
}

func (s *Service) LoginUserHandler(ctx *fiber.Ctx) error {
	// Get validated user data from Fiber context locals
	user := ctx.Locals("validatedData").(*types.LoginUserPayload)

	// Check if user exists with the email
	existingUser, err := s.store.GetUserByEmail(user.Email)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}
	if existingUser == nil {
		return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Credentials")
	}

	// Compare the password
	if !utils.ComparePassword(existingUser.HashedPassword, user.Password, existingUser.PasswordSalt) {
		return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Invalid Credentials")
	}

	// Generate JWT
	token, err := utils.CreateJWT(int(existingUser.ID))
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Internal Server Error")
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "You have logged in!!",
		"token":   token,
	})

}

func (s *Service) RegisterUserHandler(ctx *fiber.Ctx) error {
	// Get validated user data from Fiber context locals
	user := ctx.Locals("validatedData").(*types.RegisterUserPayload)

	// Check if user already exists
	existingUser, err := s.store.GetUserByEmail(user.Email)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}
	if existingUser != nil {
		return utils.SendErrorResponse(ctx, http.StatusConflict, fmt.Sprintf("User with email %s already exists", user.Email))
	}

	// Hash password
	hashedPassword, passwordSalt, err := utils.HashPassword(user.Password)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}

	data := models.User{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Email:          user.Email,
		HashedPassword: hashedPassword,
		PasswordSalt:   passwordSalt,
	}

	// Create user
	err = s.store.CreateUser(&data)
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
	}

	// Generate JWT
	token, err := utils.CreateJWT(int(data.ID))
	if err != nil {
		return utils.SendErrorResponse(ctx, http.StatusUnauthorized, "Internal Server Error")
	}

	return ctx.Status(http.StatusCreated).JSON(&fiber.Map{
		"message": "You have successfully registered!!",
		"token":   token,
	})
}
