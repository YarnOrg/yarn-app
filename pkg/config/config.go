// package config

// import (
// 	"github.com/spf13/viper"
// )

// type Config struct {
// 	DSN string `mapstructure:"DSN"`
// }

// func LoadConfig() (*Config, error) {
// 	viper.AddConfigPath("configs")
// 	viper.SetConfigName("dev")
// 	viper.SetConfigType("env")

// 	viper.AutomaticEnv()

// 	var config Config
// 	if err := viper.ReadInConfig(); err != nil {
// 		return nil, err
// 	}
// 	if err := viper.Unmarshal(&config); err != nil {
// 		return nil, err
// 	}
// 	return &config, nil
// }

package storage

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Password string
	User     string
	DBName   string
	SSLMode  string
}

func NewConnection(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, err
	}
	return db, nil
}
