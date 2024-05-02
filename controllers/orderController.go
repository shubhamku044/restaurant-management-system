package controllers

import "github.com/gin-gonic/gin"

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getOrders",
		})
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getOrder",
		})
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "createOrder",
		})
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "updateOrder",
		})
	}
}
