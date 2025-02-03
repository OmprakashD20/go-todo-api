package main

import (
	"log"

	"github.com/OmprakashD20/go-todo-api/cmd/api"
)

func main() {
	server := api.NewAPIServer(":8000", nil)

	if err := server.Run(); err != nil {
		log.Fatal("Failed to run the server")
	}
}
