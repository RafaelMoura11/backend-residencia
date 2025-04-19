package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateTool(c *gin.Context) {
	var tool models.Tool
	if err := c.ShouldBindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&tool).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar ferramenta"})
		return
	}

	c.JSON(http.StatusCreated, tool)
}

func GetTools(c *gin.Context) {
	var tools []models.Tool
	config.DB.Find(&tools)
	c.JSON(http.StatusOK, tools)
}

func GetToolByID(c *gin.Context) {
	id := c.Param("id")
	var tool models.Tool

	if err := config.DB.First(&tool, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ferramenta não encontrada"})
		return
	}

	c.JSON(http.StatusOK, tool)
}

func UpdateTool(c *gin.Context) {
	id := c.Param("id")
	var tool models.Tool

	if err := config.DB.First(&tool, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ferramenta não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&tool); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&tool)
	c.JSON(http.StatusOK, tool)
}

func DeleteTool(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Tool{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ferramenta deletada"})
}
