package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controllers.GetTables())
	incomingRoutes.GET("/tables/:id", controllers.GetTable())
	incomingRoutes.POST("/tables", controllers.CreateTable())
	incomingRoutes.PATCH("/tables/:id", controllers.UpdateTable())
}
