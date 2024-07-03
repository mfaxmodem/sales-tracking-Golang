package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"sales-tracking/models"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	db := c.MustGet("db").(*gorm.DB)

	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, users)
}
