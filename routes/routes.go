package routes

import (
	"github.com/gin-gonic/gin"
	"sales-tracking/controllers"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/users", controllers.GetUsers)

		api.POST("/location", controllers.AddLocation)
		api.GET("/locations/:userId", controllers.GetLocations)

		api.POST("/visit/start", controllers.StartVisit)
		api.POST("/visit/end/:id", controllers.EndVisit)
		api.GET("/visits/:userId", controllers.GetVisits)
	}
}
