package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/controllers"
)

func OrderRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orders", controllers.GetOrders())
	incomingRoutes.GET("/orders/:id", controllers.GetOrder())
	incomingRoutes.POST("/orders", controllers.CreateOrder())
	incomingRoutes.PATCH("/orders/:id", controllers.UpdateOrder())
}
