package db

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(dsn string) error {
	var err error
	DB, err = pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		return err
	}
	return nil
}
