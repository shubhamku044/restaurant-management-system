package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/controllers"
)

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/menus", controllers.GetMenus())
	incomingRoutes.GET("/menus/:id", controllers.GetMenu())
	incomingRoutes.POST("/menus", controllers.CreateMenu())
	incomingRoutes.PATCH("/menus/:id", controllers.UpdateMenu())
}
