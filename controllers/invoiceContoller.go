package controllers

import "github.com/gin-gonic/gin"

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getInvoices",
		})
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getInvoice",
		})
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "createInvoice",
		})
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "updateInvoice",
		})
	}
}
