package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sales-tracking/models"
	"time"
)

func StartVisit(c *gin.Context) {
	var visit models.Visit
	if err := c.ShouldBindJSON(&visit); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := c.MustGet("db").(*gorm.DB)
	visit.StartTime = time.Now()

	if err := db.Create(&visit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Visit started successfully"})
}

func EndVisit(c *gin.Context) {
	var visit models.Visit
	visitID := c.Param("id")
	db := c.MustGet("db").(*gorm.DB)

	if err := db.First(&visit, visitID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Visit not found"})
		return
	}

	visit.EndTime = time.Now()
	if err := db.Save(&visit).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Visit ended successfully"})
}

func GetVisits(c *gin.Context) {
	userID := c.Param("userId")
	db := c.MustGet("db").(*gorm.DB)

	var visits []models.Visit
	if err := db.Where("visitor_id = ?", userID).Find(&visits).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"visits": visits})
}
