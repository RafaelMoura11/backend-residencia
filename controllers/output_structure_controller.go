package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"backend-residencia/config"
	"backend-residencia/models"
)

func CreateOutputStructure(c *gin.Context) {
	var output models.OutputStructure
	if err := c.ShouldBindJSON(&output); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&output).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar estrutura de saída"})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func GetOutputStructures(c *gin.Context) {
	var outputs []models.OutputStructure
	config.DB.Find(&outputs)
	c.JSON(http.StatusOK, outputs)
}

func GetOutputStructureByID(c *gin.Context) {
	id := c.Param("id")
	var output models.OutputStructure

	if err := config.DB.First(&output, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estrutura de saída não encontrada"})
		return
	}

	c.JSON(http.StatusOK, output)
}

func UpdateOutputStructure(c *gin.Context) {
	id := c.Param("id")
	var output models.OutputStructure

	if err := config.DB.First(&output, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Estrutura de saída não encontrada"})
		return
	}

	if err := c.ShouldBindJSON(&output); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&output)
	c.JSON(http.StatusOK, output)
}

func DeleteOutputStructure(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.OutputStructure{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Estrutura de saída deletada"})
}
