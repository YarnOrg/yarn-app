package main

import (
	"log"

	"github.com/YarnOrg/yarn-app/internal/user/api"
	"github.com/YarnOrg/yarn-app/pkg/db"

	// "internal/user"
	// "pkg/db"

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

	router.Group("/users").POST("/", user.createUser)
	router.Group("/users").GET("/:id", user.getUser)
	router.Group("/transactions").POST("/", transaction.createTransaction)
	router.Group("/messages").POST("/", chat.sendMessage)
	router.Group("/products").POST("/", marketplace.listProduct)

	//
	if err := db.ConnectDB("user=postgres dbname=socialfinance sslmode=disable"); err != nil {
		log.Fatal("Could not connect to database:", err)
	}

	user.RegisterUserRoutes(router)
	// Start server
	router.Run(":8080")
}
