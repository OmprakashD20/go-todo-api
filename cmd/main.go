package main

import (
	"log"

	"github.com/OmprakashD20/go-todo-api/cmd/api"
	"github.com/OmprakashD20/go-todo-api/config"
	"github.com/OmprakashD20/go-todo-api/database"
)

func main() {
	_, err := database.ConnectDB(&config.Envs.DB)

	if err != nil {
		log.Fatalf("Failed to connect to the database")
	}

	server := api.NewAPIServer(config.Envs.Port, nil)

	if err := server.Run(); err != nil {
		log.Fatal("Failed to run the server")
	}
}
