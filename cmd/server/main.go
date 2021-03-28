package main

import (
	"log"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/server"
	"todo/pkg/service"

	_ "github.com/lib/pq"

	"github.com/spf13/viper"
)

func main() {
	db, err := repository.NewPostgresDB(
		repository.Config{
			Host:     viperEnvVariable("POSTGRES_HOST"),
			Port:     viperEnvVariable("POSTGRES_PORT"),
			Username: viperEnvVariable("POSTGRES_USER"),
			Password: viperEnvVariable("POSTGRES_PASSWORD"),
			DBName:   viperEnvVariable("POSTGRES_USER"),
			SSLMode:  viperConfig("db.sslmode"),
		})

	if err != nil {
		log.Fatalf("Failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viperEnvVariable("SERVER_PORT"), handlers.InitRoutes()); err != nil {
		log.Fatalf("server error: %s", err.Error())
	}
}

func viperConfig(key string) string {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err.Error())
	}

	return viper.GetString(key)
}

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err.Error())
	}

	return viper.GetString(key)
}
