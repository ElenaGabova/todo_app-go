package main

import (
	"log"

	"github.com/ElenaGabova/todo_app-go.git"
	"github.com/ElenaGabova/todo_app-go.git/pkg/handler"
	"github.com/ElenaGabova/todo_app-go.git/pkg/repository"
	"github.com/ElenaGabova/todo_app-go.git/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	servives := service.NewService(repos)
	handlers := handler.NewHandler(servives)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Server occuring while running http server: %s", err.Error())
	}
}
