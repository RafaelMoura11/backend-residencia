package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"your_project/config"
	"your_project/models"
)

func CreateAgent(c *gin.Context) {
	var agent models.Agent
	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&agent)
	c.JSON(http.StatusCreated, agent)
}

func GetAgents(c *gin.Context) {
	var agents []models.Agent
	config.DB.Find(&agents)
	c.JSON(http.StatusOK, agents)
}

func GetAgentByID(c *gin.Context) {
	var agent models.Agent
	id := c.Param("id")

	if err := config.DB.First(&agent, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agente não encontrado"})
		return
	}
	c.JSON(http.StatusOK, agent)
}

func UpdateAgent(c *gin.Context) {
	var agent models.Agent
	id := c.Param("id")

	if err := config.DB.First(&agent, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Agente não encontrado"})
		return
	}

	if err := c.ShouldBindJSON(&agent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&agent)
	c.JSON(http.StatusOK, agent)
}

func DeleteAgent(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Agent{}, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Agente deletado"})
}
