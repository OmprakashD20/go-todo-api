package types

import (
	"time"

	"github.com/OmprakashD20/go-todo-api/models"
)

// Store
type UserStore interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id uint) (*models.User, error)
}

type TodoStore interface {
	CreateTodo(todo *models.Todo) error
	GetTodoById(id uint) (*models.Todo, error)
	GetTodosByUserId(userId uint) ([]*models.Todo, error)
	//UpdateTodoById(id uint, data *models.Todo) error
	//DeleteTodoById(id uint) error // Implement Batch Delete Option
}

// Payloads
type RegisterUserPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginUserPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateTodoPayload struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    string    `json:"priority"`
	DueDate     time.Time `json:"dueDate"`
}
