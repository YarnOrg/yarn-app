package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// User represents the user model that GORM will use.
type User struct {
	gorm.Model
	Username     string `gorm:"uniqueIndex"`
	Email        string `gorm:"uniqueIndex"`
	PasswordHash string
}

// Connect initializes the database connection with GORM.
func Connect(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
