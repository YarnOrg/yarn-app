package marketplace

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listProduct(c *gin.Context) {
	var product ProductData
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add product listing logic here
	// Insert product into database

	c.JSON(http.StatusCreated, gin.H{"message": "Product listed"})
}
