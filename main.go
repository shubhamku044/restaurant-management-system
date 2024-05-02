package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/shubhamku044/restaurant-management-system/database"
	"github.com/shubhamku044/restaurant-management-system/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.New()
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
		router.Use(gin.Recovery())
	} else {
		router.Use(gin.Logger())
	}
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is up and running on port " + port,
		})
	})
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.MenuRoutes(router)
	routes.InvoiceRoutes(router)
	routes.TableRoutes(router)

	err = router.Run(":" + port)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
	fmt.Println("Hello, World!")
}
