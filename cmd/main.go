package main

import (
	"log"
	todo "restAPI"
	"restAPI/package/handler"
	"restAPI/package/repository"
	"restAPI/package/service"
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

func initConfig() error {
	return nil
}
