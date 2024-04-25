package models

import (
	"gorm.io/gorm"
)

// User represents a user in the database.
type User struct {
	gorm.Model
	Username     string `json:"username"`
	Email        string `json:"email" gorm:"uniqueIndex"`
	PasswordHash string `json:"-"` // Do not serialize password hash
}

// MigrateUsers applies database migrations for the User model.
func MigrateUsers(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
