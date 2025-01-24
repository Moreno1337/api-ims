package main

import (
	"log"

	"github.com/Moreno1337/api-ims/src/client/controller"
	"github.com/Moreno1337/api-ims/src/client/controller/routes"
	"github.com/Moreno1337/api-ims/src/client/model/service"
	sqlserver "github.com/Moreno1337/api-ims/src/configuration/database"
	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting API")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlserver.InitConnection()

	clientService := service.NewClientDomainService()
	clientController := controller.NewClientControllerInterface(clientService)

	router := gin.Default()

	clientsGroup := router.Group("/clients")

	clientroutes.InitRoutes(clientsGroup, clientController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
