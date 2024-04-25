package transaction

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func createTransaction(c *gin.Context) {
	var transaction TransactionData
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// transaction logic here
	// Insert transaction record into database

	c.JSON(http.StatusCreated, gin.H{"message": "Transaction processed"})
}
