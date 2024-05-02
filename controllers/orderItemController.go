package controllers

import "github.com/gin-gonic/gin"

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getOrderItems",
		})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getOrderItem",
		})
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getOrderItemsByOrder",
		})
	}
}

func ItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "itemsByOrder",
		})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "createOrderItem",
		})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "updateOrderItem",
		})
	}
}
