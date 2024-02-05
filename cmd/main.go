package main

import (
	"log"

	todo "github.com/AndreyAlekseev/todo_list"
	"github.com/AndreyAlekseev/todo_list/pkg/handler"
	"github.com/AndreyAlekseev/todo_list/pkg/repository"
	"github.com/AndreyAlekseev/todo_list/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
