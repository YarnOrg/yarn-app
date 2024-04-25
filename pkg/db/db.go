package db

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

var Pool *pgxpool.Pool

func ConnectDB(dsn string) error {
	var err error
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}

	config.MaxConnLifetime = time.Hour
	config.MaxConns = 20
	config.MinConns = 5

	Pool, err = pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return err
	}
	log.Println("Database connection established")
	return nil
}
