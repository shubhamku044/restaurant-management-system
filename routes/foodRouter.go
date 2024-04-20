package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/controllers"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods", controllers.GetFoods())
	incomingRoutes.GET("/foods/:id", controllers.GetFood())
	incomingRoutes.POST("/foods", controllers.CreateFood())
	incomingRoutes.PATCH("/foods/:id", controllers.UpdateFood())
}
