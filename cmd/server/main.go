package main

import (
	"log"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/server"
	"todo/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8086", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
