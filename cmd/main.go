package main

import (
	
	"github.com/ElenaGabova/todo_app-go.git"
	"github.com/ElenaGabova/todo_app-go.git/pkg/handler"
	"log"
)

func main() {
	handlers:= new (handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server occuring while running http server: %s", err.Error())
	}
}
