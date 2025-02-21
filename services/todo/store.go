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
	err := s.db.Model(&models.Todo{}).Create(todo).Error

	return err
}

func (s *Store) GetTodoById(id uint) (*models.Todo, error) {
	todo := models.Todo{}

	err := s.db.Model(&models.Todo{}).Where("id = ?", id).First(&todo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &todo, err
}

func (s *Store) GetTodosByUserId(userId uint) ([]*models.Todo, error) {
	todos := []*models.Todo{}

	err := s.db.Model(&models.Todo{}).Where("user_id = ?", userId).Order("due_date ASC").Find(&todos).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return todos, err
}

func (s *Store) UpdateTodoById(id uint, data *models.Todo) error {
	err := s.db.Model(&models.Todo{}).Where("id = ?", id).Updates(data).Error

	return err
}

// func (s *Store) DeleteTodoById(id uint) error {}
