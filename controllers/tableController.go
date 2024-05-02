package controllers

import "github.com/gin-gonic/gin"

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
