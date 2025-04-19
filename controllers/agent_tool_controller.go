package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateAgentTool(c *gin.Context) {
	var relation models.AgentTool
	if err := c.ShouldBindJSON(&relation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// verifica se agent existe
	var agent models.Agent
	if err := config.DB.First(&agent, relation.AgentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agente não encontrado"})
		return
	}

	// verifica se tool existe
	var tool models.Tool
	if err := config.DB.First(&tool, relation.ToolID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ferramenta não encontrada"})
		return
	}

	config.DB.Create(&relation)
	c.JSON(http.StatusCreated, relation)
}

func GetAgentTools(c *gin.Context) {
	var relations []models.AgentTool
	config.DB.Find(&relations)
	c.JSON(http.StatusOK, relations)
}

func DeleteAgentTool(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.AgentTool{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Relação agente-ferramenta deletada"})
}
