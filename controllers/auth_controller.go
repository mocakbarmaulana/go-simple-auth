package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-simple-auth/models"
	"github.com/go-simple-auth/utils"
)

func Login(c *gin.Context) {
	var user models.User

	// Bind JSON request to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if user exists
	if user.Username != "admin" || user.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate JWT token
	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func Register(c *gin.Context) {
	var user models.User

	// Bind JSON request to user model
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create user
	user.ID = 1
	c.JSON(http.StatusCreated, gin.H{"data": user, "message": "User created successfully"})
}

func GetUser(c *gin.Context) {
	userID := c.MustGet("userID")

	c.JSON(http.StatusOK, gin.H{"message": "User found", "user_id": userID})
}
