package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/database"
	"github.com/shubhamku044/restaurant-management-system/middleware"
	"github.com/shubhamku044/restaurant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"

	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}
