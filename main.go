package main

import (
	"backend-residencia/config"
	"backend-residencia/models"
	"backend-residencia/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	log.Println("🔁 Carregando variáveis de ambiente...")

	config.ConnectDatabase()

	log.Println("🔧 Executando AutoMigrate...")
	if err := config.DB.AutoMigrate(&models.Message{}, &models.Agent{})	; err != nil {
		log.Fatal("❌ Erro no AutoMigrate:", err)
	}

	log.Println("🚀 Servidor iniciado em http://localhost:8080")
	r := routes.SetupRoutes()
	r.Run(":8080")
}
