package migrate

import (
	"github.com/OmprakashD20/go-todo-api/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{})

	return err
}
