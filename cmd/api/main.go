// package main

// import (
// 	"github.com/YarnOrg/yarn-app/internal/user"
// 	// "github.com/YarnOrg/yarn-app/pkg/db"

// 	// "internal/user"
// 	// "pkg/db"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	router := gin.Default()

// 	//
// 	// if err := db.ConnectDB("user=mohameda password=postgres dbname=yarn sslmode=disable"); err != nil {
// 	// 	log.Fatal("Could not connect to database:", err)
// 	// }

// 	user.RegisterUserRoutes(router)
// 	// Start server
// 	router.Run(":8080")
// }

package main

import (
	"log"
	"os"

	"github.com/YarnOrg/yarn-app/models"
	"github.com/YarnOrg/yarn-app/storage"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}

	if err := models.MigrateUsers(db); err != nil {
		log.Fatalf("could not migrate database: %v", err)
	}

	app := fiber.New()
	repo := storage.Repository{DB: db}
	repo.SetupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
