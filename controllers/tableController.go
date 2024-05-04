package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shubhamku044/restaurant-management-system/database"
	"go.mongodb.org/mongo-driver/mongo"
)

var tableCollection *mongo.Collection = database.OpenCollection(database.Client, "table")

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getTables",
		})
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getTable",
		})
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "createTable",
		})
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "updateTable",
		})
	}
}
