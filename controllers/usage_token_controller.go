package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateUsageToken(c *gin.Context) {
	var token models.UsageToken
	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var message models.Message
	if err := config.DB.First(&message, token.MessageID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mensagem não encontrada"})
		return
	}

	if message.Role != "assistant" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tokens só podem ser criados para mensagens do tipo 'assistant'"})
		return
	}

	if err := config.DB.Create(&token).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar token de uso"})
		return
	}

	c.JSON(http.StatusCreated, token)
}

func GetUsageTokens(c *gin.Context) {
	var tokens []models.UsageToken
	config.DB.Preload("Message").Find(&tokens)
	c.JSON(http.StatusOK, tokens)
}

func GetUsageTokenByID(c *gin.Context) {
	id := c.Param("id")
	var token models.UsageToken

	if err := config.DB.Preload("Message").First(&token, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token de uso não encontrado"})
		return
	}

	c.JSON(http.StatusOK, token)
}

func UpdateUsageToken(c *gin.Context) {
	id := c.Param("id")
	var token models.UsageToken

	if err := config.DB.First(&token, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Token de uso não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&token)
	c.JSON(http.StatusOK, token)
}

func DeleteUsageToken(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.UsageToken{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Token de uso deletado"})
}
