package types

import (
	"github.com/OmprakashD20/go-todo-api/models"
)

// Store
type UserStore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
}

// Payloads
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayload struct{
	Email string `json:"email"`
	Password string `json:"password"`
}