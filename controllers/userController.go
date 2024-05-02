package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getUsers",
		})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "getUser",
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "createUser",
		})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "signUp",
		})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "login",
		})
	}
}
