package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateAgentOutputStructure(c *gin.Context) {
	var rel models.AgentOutputStructure
	if err := c.ShouldBindJSON(&rel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var agent models.Agent
	if err := config.DB.First(&agent, rel.AgentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Agente não encontrado"})
		return
	}

	var output models.OutputStructure
	if err := config.DB.First(&output, rel.OutputStructureID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Estrutura de saída não encontrada"})
		return
	}

	config.DB.Create(&rel)
	c.JSON(http.StatusCreated, rel)
}

func GetAgentOutputStructures(c *gin.Context) {
	var rels []models.AgentOutputStructure
	config.DB.Find(&rels)
	c.JSON(http.StatusOK, rels)
}

func DeleteAgentOutputStructure(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.AgentOutputStructure{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Relação agente-estrutura de saída deletada"})
}
