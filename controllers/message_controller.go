package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateMessage(c *gin.Context) {
	var message models.Message
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message.Timestamp = time.Now()
	if err := config.DB.Create(&message).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar mensagem"})
		return
	}

	c.JSON(http.StatusCreated, message)
}

func GetMessages(c *gin.Context) {
	var messages []models.Message
	config.DB.Preload("Agent").Find(&messages)
	c.JSON(http.StatusOK, messages)
}

func GetMessageByID(c *gin.Context) {
	var message models.Message
	id := c.Param("id")

	if err := config.DB.Preload("Agent").First(&message, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mensagem não encontrada"})
		return
	}

	c.JSON(http.StatusOK, message)
}

func UpdateMessage(c *gin.Context) {
	var message models.Message
	id := c.Param("id")

	if err := config.DB.First(&message, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Mensagem não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&message)
	c.JSON(http.StatusOK, message)
}

func DeleteMessage(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Message{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Mensagem deletada"})
}
