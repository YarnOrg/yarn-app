package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"context"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUserRoutes(r *gin.Engine) {
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

	ctx := context.Background()
	id, err := CreateUser(ctx, u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func getUser(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	ctx := context.Background()
	u, err := GetUserByID(ctx, userID)
	if err != nil {
		if err == sql.ErrNoRows { // You need to import the database/sql package for this
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not retrieve user: " + err.Error()})
		return
	}

	if u.ID == 0 { // Additional check if user ID returned is 0, indicating no data found
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, u)
}
