package main

import (
	"log"
	"os"

	"github.com/YarnOrg/yarn-app/internal/user"
	"github.com/YarnOrg/yarn-app/pkg/db"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("DATABASE_URL")
	database, err := db.Connect(dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := user.NewRepository(database)

	router := gin.Default()
	user.RegisterUserRoutes(router, repo)

	log.Fatal(router.Run(":8080"))
}
