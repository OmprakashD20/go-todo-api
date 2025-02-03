package user

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

func (s *Store) CreateUser(user *models.User) error {
	err := s.db.Create(user).Error
	return err
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := s.db.Where("email = ?", email).First(&user).Error

	// GORM doesn't return nil when no record is found
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, err
}

func (s *Store) GetUserById(id uint) (*models.User, error) {
	user := models.User{}
	err := s.db.Where("id = ?", id).First(&user).Error

	// GORM doesn't return nil when no record is found
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, err
}
