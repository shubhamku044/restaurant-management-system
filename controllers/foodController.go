package controllers

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/shubhamku044/restaurant-management-system/database"
	"github.com/shubhamku044/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection *mongo.Collection = database.OpenCollection(database.Client, "food")
var validate = validator.New()

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		recordPerPage, err := strconv.Atoi(c.Query("recordPerPage"))

		if err != nil || recordPerPage < 1 {
			recordPerPage = 10
		}

		page, err := strconv.Atoi(c.Query("page"))
		if err != nil || page < 1 {
			page = 1
		}

		startIdx := (page - 1) * recordPerPage
		startIdx, err = strconv.Atoi(c.Query("startIndex"))

		matchStage := bson.D{{Key: "$match", Value: bson.D{{}}}}
		groupStage := bson.D{{
			Key: "$group",
			Value: bson.D{
				{
					Key: "_id",
					Value: bson.D{
						{
							Key:   "_id",
							Value: "null",
						},
					},
				},
				{
					Key:   "total_count",
					Value: bson.D{{Key: "$sum", Value: 1}},
				},
				{
					Key:   "data",
					Value: bson.D{{Key: "$push", Value: "$$ROOT"}},
				},
			},
		}}
		projectStage := bson.D{{
			Key: "$project",
			Value: bson.D{
				{
					Key:   "_id",
					Value: 0,
				},
				{
					Key:   "total_count",
					Value: 1,
				},
				{
					Key: "food_items",
					Value: bson.D{
						{
							Key: "$slice",
							Value: []interface{}{
								"$data",
								startIdx,
								recordPerPage,
							},
						},
					},
				},
			},
		}}

		result, err := foodCollection.Aggregate(ctx, mongo.Pipeline{
			matchStage, groupStage, projectStage,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the food"})
			return
		}
		allFoods := []bson.M{}
		if err = result.All(ctx, &allFoods); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error occurred while listing the food"})
			fmt.Println(err)
		}

		c.JSON(http.StatusOK, allFoods)
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
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var food models.Food
		var menu models.Menu
		defer cancel()

		foodId := c.Param("id")
		if err := c.BindJSON(&food); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updateObj primitive.D

		if food.Name != nil {
			updateObj = append(updateObj, primitive.E{Key: "name", Value: food.Name})
		}
		if food.Price != nil {
			updateObj = append(updateObj, primitive.E{Key: "price", Value: food.Price})
		}
		if food.Food_image != nil {
			updateObj = append(updateObj, primitive.E{Key: "food_image", Value: food.Food_image})
		}
		if food.Menu_id != nil {
			err := menuCollection.FindOne(ctx, bson.M{"menu_id": food.Menu_id}).Decode(&menu)

			if err != nil {
				msg := "menu not found"
				c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
				return
			}
			updateObj = append(updateObj, primitive.E{Key: "menu_id", Value: food.Menu_id})
		}

		food.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, primitive.E{Key: "updated_at", Value: food.Updated_at})

		upsert := true

		filter := bson.M{"food_id": foodId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_, err := foodCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			msg := "food not updated"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		updatedFood := models.Food{}
		err = foodCollection.FindOne(ctx, bson.M{"food_id": foodId}).Decode(&updatedFood)

		if err != nil {
			msg := "Error fetching updated food"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		c.JSON(http.StatusOK, updatedFood)
	}
}
