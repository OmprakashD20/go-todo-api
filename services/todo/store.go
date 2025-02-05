package todo

import (
	"github.com/OmprakashD20/go-todo-api/models"

	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	return &Store{db}
}

func (s *Store) CreateTodo(todo *models.Todo) error {
	err := s.db.Table("todos").Create(todo).Error

	return err
}

// func (s *Store) GetTodoById(id uint) (*models.Todo, error) {}

// func (s *Store) GetTodosByUserId(userId uint) ([]*models.Todo, error) {}

// func (s *Store) UpdateTodoById(id uint, data *models.Todo) error {}

// func (s *Store) DeleteTodoById(id uint) error {}
