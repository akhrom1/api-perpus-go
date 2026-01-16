package main

import (
	"log"

	"api-perpus-go/config"
	"api-perpus-go/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config.ConnectDB()
	config.RunMigration(config.DB)

	r := gin.Default()

	routes.RegisterAPIRoutes(r)

	r.Run(":8080")
}
