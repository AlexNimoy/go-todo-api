package main

import (
	"log"
	"todo/pkg/handler"
	"todo/pkg/server"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8086", handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}
