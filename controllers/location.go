package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sales-tracking/models"
)

func AddLocation(c *gin.Context) {
	var location models.Checkpoint
	if err := c.ShouldBindJSON(&location); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&location).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Location added successfully"})
}

func GetLocations(c *gin.Context) {
	userId := c.Param("userId")

	db := c.MustGet("db").(*gorm.DB)
	var checkpoints []models.Checkpoint
	if err := db.Where("user_id = ?", userId).Find(&checkpoints).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"locations": checkpoints})
}
