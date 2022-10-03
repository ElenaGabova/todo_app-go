package main

import (
	"github.com/ElenaGabova/todo_app-go.git"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("Server occuring while running http server: %s", err.Error())
	}
}
