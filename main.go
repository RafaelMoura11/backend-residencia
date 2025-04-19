package main

import (
	"backend-residencia/config"
	"backend-residencia/models"
	"backend-residencia/routes"

	"github.com/joho/godotenv"
	"log"
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
	); err != nil {
		log.Fatal("❌ Erro ao executar AutoMigrate:", err)
	}
	log.Println("✅ AutoMigrate concluído")

	r := routes.SetupRoutes()
	r.Run(":8080")
}
