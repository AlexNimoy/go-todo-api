package main

import (
	"context"
	"os"
	"os/signal"
	"todo/pkg/handler"
	"todo/pkg/repository"
	"todo/pkg/server"
	"todo/pkg/service"

	_ "github.com/lib/pq"

	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

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
		logrus.Fatalf("Failed to init db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viperEnvVariable("SERVER_PORT"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("server error: %s", err.Error())
		}
	}()

	logrus.Printf("Server started at port: %s", viperEnvVariable("SERVER_PORT"))

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Server error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("DB error: %s", err.Error())
	}

	logrus.Printf("Server stoped")
}

func viperConfig(key string) string {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	err := viper.ReadInConfig()

	if err != nil {
		logrus.Fatalf("Error while reading config file %s", err.Error())
	}

	return viper.GetString(key)
}

func viperEnvVariable(key string) string {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		logrus.Fatalf("Error while reading config file %s", err.Error())
	}

	return viper.GetString(key)
}
