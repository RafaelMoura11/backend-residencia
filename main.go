package main

import (
	"your_project/config"
	"your_project/models"
	"your_project/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config.ConnectDatabase()

	// auto migrate
	config.DB.AutoMigrate(&models.Agent{})

	r := routes.SetupRoutes()
	r.Run(":8080")
}
