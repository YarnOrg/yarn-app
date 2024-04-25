package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
}

func createUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not encrypt password"})
		return
	}
	u.PasswordHash = string(hashedPassword)

	if err := CreateUser(c, u); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func getUser(c *gin.Context) {
	userID := c.Param("id")

	// Add user retrieval logic here
	// Query database for user by userID

	c.JSON(http.StatusOK, gin.H{"user": userID})
}
