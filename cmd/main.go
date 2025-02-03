package main

import (
	"log"

	"github.com/OmprakashD20/go-todo-api/cmd/api"
	"github.com/OmprakashD20/go-todo-api/cmd/migrate"
	"github.com/OmprakashD20/go-todo-api/config"
	"github.com/OmprakashD20/go-todo-api/database"
)

func main() {
	db, err := database.ConnectDB(&config.Envs.DB)

	if err != nil {
		log.Fatalf("Failed to connect to the database")
	}

	if err := migrate.Migrate(db); err != nil {
		log.Fatalf("Failed to migrate the database")
	}

	server := api.NewAPIServer(config.Envs.Port, db)

	if err := server.Run(); err != nil {
		log.Fatal("Failed to run the server")
	}
}
