package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shubhamku044/restaurant-management-system/database"
	"github.com/shubhamku044/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		c.JSON(200, gin.H{
			"message": "getFoods",
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		foodId := c.Param("id")

		var food models.Food

		err := foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&food)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the food"})
			return
		}

		c.JSON(200, food)
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var menu models.Menu
		var food models.Food
		defer cancel()

		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(food)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

		if err != nil {
			msg := "menu not found"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		food.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		food.ID = primitive.NewObjectID()
		food.Food_id = food.ID.Hex()

		num := toFixed(*food.Price, 2)
		food.Price = &num

		result, err := foodCollection.InsertOne(ctx, food)

		if err != nil {
			msg := "food not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		defer cancel()

		c.JSON(http.StatusOK, result)
	}
}

func round(num float64) int {
	return 3
}

func toFixed(num float64, precision int) float64 {
	return 3.14159
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "updateFood",
		})
	}
}
