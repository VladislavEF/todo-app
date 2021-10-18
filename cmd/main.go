package main

import (
	"log"

	"github.com/VladislavEF/todo-app"
	"github.com/VladislavEF/todo-app/pkg/handlers"
	"github.com/VladislavEF/todo-app/pkg/repository"
	"github.com/VladislavEF/todo-app/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	service := service.NewService(repos)
	handlers := handlers.NewHandler(service)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running server:", err.Error())
	}
}
