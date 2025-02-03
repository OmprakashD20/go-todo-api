package database

import (
	"fmt"

	"github.com/OmprakashD20/go-todo-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(config *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.DBHost, config.DBPort, config.DBUser, config.DBPort, config.DBName, config.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
