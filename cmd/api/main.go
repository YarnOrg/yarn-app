package main

import (
	"github.com/YarnOrg/yarn-app/internal/user/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// User module routes
	userGroup := router.Group("/users")
	{
		userGroup.POST("/", api.CreateUser)
		userGroup.GET("/:id", api.GetUser)
	}

	// Start server
	router.Run(":8080")
}
