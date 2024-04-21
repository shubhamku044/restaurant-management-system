package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/controllers"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/orderItems", controllers.GetOrderItems())
	incomingRoutes.GET("/orderItems/:id", controllers.GetOrderItem())
	incomingRoutes.GET("/orderItems-order/:id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/orderItems", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/orderItems/:id", controllers.UpdateOrderItem())
}
