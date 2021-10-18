package main

import (
	"log"

	"github.com/VladislavEF/todo-app"
	"github.com/VladislavEF/todo-app/pkg/handlers"
)

func main() {
	handlers := new(handlers.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running server: %s", err.Error())
	}
}
