package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/database"
	"github.com/shubhamku044/restaurant-management-system/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var orderCollection *mongo.Collection = database.OpenCollection(database.Client, "order")

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		result, err := orderCollection.Find(context.TODO(), bson.M{})
		if err != nil {
			c.JSON(500, gin.H{"error": "error occurred while listing the orders"})
			return
		}

		var allOrders []bson.M
		if err = result.All(ctx, &allOrders); err != nil {
			c.JSON(500, gin.H{"error": "error occurred while listing the orders"})
			return
		}

		c.JSON(http.StatusOK, allOrders)
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		orderId := c.Param("id")
		err := orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&order)
		if err != nil {
			c.JSON(500, gin.H{"error": "error occurred while listing the order"})
			return
		}

		c.JSON(200, gin.H{
			"message": order,
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var order models.Order
		var table models.Table

		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validateErr := validate.Struct(order)

		if validateErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validateErr.Error()})
			return
		}

		err := tableCollection.FindOne(context.Background(), bson.M{"table_id": order.Table_id}).Decode(&table)
		if err != nil {
			c.JSON(500, gin.H{"error": "error occurred while creating the order"})
			return
		}

		order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		order.Order_id = primitive.NewObjectID().Hex()

		result, insertErr := orderCollection.InsertOne(ctx, order)

		if insertErr != nil {
			c.JSON(500, gin.H{"error": "error occurred while creating the order"})
			return
		}

		c.JSON(http.StatusCreated, result)
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var table models.Table
		var order models.Order
		var updateObj primitive.D
		orderId := c.Param("id")
		if err := c.BindJSON(&order); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if order.Table_id != nil {
			err := tableCollection.FindOne(context.Background(), bson.M{"table_id": order.Table_id}).Decode(&table)
			if err != nil {
				c.JSON(500, gin.H{"error": "error occurred while updating the order"})
				return
			}
			updateObj = append(updateObj, bson.E{Key: "table_id", Value: order.Table_id})
		}

		order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		updateObj = append(updateObj, bson.E{Key: "updated_at", Value: order.Updated_at})

		upsert := true

		filter := bson.M{"order_id": orderId}

		opt := options.UpdateOptions{
			Upsert: &upsert,
		}

		_, err := orderCollection.UpdateOne(
			ctx,
			filter,
			bson.D{
				{Key: "$set", Value: updateObj},
			},
			&opt,
		)

		if err != nil {
			c.JSON(500, gin.H{"error": "error occurred while updating the order"})
			return
		}

		updatedOrder := models.Order{}
		err = orderCollection.FindOne(ctx, bson.M{"order_id": orderId}).Decode(&updatedOrder)

		if err != nil {
			c.JSON(500, gin.H{"error": "error occurred while updating the order"})
			return
		}

		c.JSON(http.StatusOK, updatedOrder)
	}
}

func OrderItemOrderCreator(order models.Order) string {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	order.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	order.ID = primitive.NewObjectID()
	order.Order_id = order.ID.Hex()

	orderCollection.InsertOne(ctx, order)
	return order.Order_id
}
