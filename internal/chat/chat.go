package chat

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendMessage(c *gin.Context) {
	var message MessageData
	if err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add message sending logic here
	// Insert message into database

	c.JSON(http.StatusCreated, gin.H{"message": "Message sent"})
}
