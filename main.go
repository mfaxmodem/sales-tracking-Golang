package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"sales-tracking/config"
	"sales-tracking/models"
	"sales-tracking/routes"
)

func main() {
	// اتصال به پایگاه داده
	config.ConnectDatabase()

	// مهاجرت مدل‌ها
	err := config.DB.AutoMigrate(&models.User{}, &models.Route{}, &models.Checkpoint{}, &models.Store{}, &models.Visit{})
	if err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	router := gin.Default()

	// اضافه کردن اتصال پایگاه داده به کانتکست Gin
	router.Use(func(c *gin.Context) {
		c.Set("db", config.DB)
		c.Next()
	})

	routes.SetupRoutes(router)

	err = router.Run(":8080")
	if err != nil {
		return
	}
}
