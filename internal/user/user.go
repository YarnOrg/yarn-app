package user

import (
	"net/http"

	"github.com/YarnOrg/yarn-app/pkg/db" // Adjust the import path based on your module name

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Repository holds DB connection logic
type Repository struct {
	DB *gorm.DB
}

// NewRepository creates a new repository with the given DB connection
func NewRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

// CreateUser handler
func (repo *Repository) CreateUser(c *gin.Context) {
	var user db.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if result := repo.DB.Create(&user); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// GetUser handler
func (repo *Repository) GetUser(c *gin.Context) {
	id := c.Param("id")
	var user db.User
	if result := repo.DB.First(&user, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// RegisterUserRoutes initializes routes
func RegisterUserRoutes(router *gin.Engine, repo *Repository) {
	router.POST("/users", repo.CreateUser)
	router.GET("/users/:id", repo.GetUser)
}
