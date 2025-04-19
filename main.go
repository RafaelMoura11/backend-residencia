package main

import (
	"backend-residencia/config"
	"backend-residencia/models"
	"backend-residencia/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	config.ConnectDatabase()

	log.Println("📦 Executando AutoMigrate...")
	if err := config.DB.AutoMigrate(
		&models.Agent{},
		&models.Message{},
		&models.UsageToken{},
		&models.Tool{},
		&models.AgentTool{},
		&models.OutputStructure{},
		&models.AgentOutputStructure{},
	); err != nil {
		log.Fatal("❌ Erro ao executar AutoMigrate:", err)
	}
	log.Println("✅ AutoMigrate concluído")

	r := routes.SetupRoutes()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
