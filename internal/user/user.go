package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
}

func createUser(c *gin.Context) {
	// Parse request data
	var newUser UserData
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// user creation logic here
	// Insert newUser into database

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUser(c *gin.Context) {
	userID := c.Param("id")

	// Add user retrieval logic here
	// Query database for user by userID

	c.JSON(http.StatusOK, gin.H{"user": userID})
}
