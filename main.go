package main

import (
	"log"

	"github.com/Moreno1337/api-ims/src/configuration/logger"
	"github.com/Moreno1337/api-ims/src/client/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("Starting API")

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	clientsGroup := router.Group("/clients")

	clientroutes.InitRoutes(clientsGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
