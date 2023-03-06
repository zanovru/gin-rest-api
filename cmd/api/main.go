package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/zanovru/gin-rest-api/internal/api"
	"github.com/zanovru/gin-rest-api/internal/api/handlers"
	"github.com/zanovru/gin-rest-api/internal/config"
	"github.com/zanovru/gin-rest-api/internal/repositories"
	"github.com/zanovru/gin-rest-api/internal/repositories/postgres"

	services2 "github.com/zanovru/gin-rest-api/internal/services"
	"os"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

func main() {

	configs, err := config.Init("apiserv")

	if err != nil {
		log.Fatalf("Error occured when initaliazing config: %s", err.Error())
	}

	logFile := configs.OutputFile

	f, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + logFile)
		panic(err)
	}
	defer f.Close()

	log.SetOutput(f)

	db, err := postgres.NewPostgresDB(configs)
	if err != nil {
		log.Fatalf("Error occured when connecting to PostgreSQL Database: %s", err.Error())
	}

	repos := repositories.NewRepositories(db)
	services := services2.NewServices(repos)
	routing := handlers.NewRouting(services)

	server := api.NewServer()

	if err := server.Start(routing.InitRoutes(), configs); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
