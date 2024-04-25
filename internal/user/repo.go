package user

import (
	"context"

	"github.com/YarnOrg/yarn-app/pkg/db"
)

type User struct {
	ID           int
	Username     string
	Email        string
	PasswordHash string
}

func CreateUser(ctx context.Context, u User) (int, error) {
	const query = `INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id`
	id := 0
	err := db.Pool.QueryRow(ctx, query, u.Username, u.Email, u.PasswordHash).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func GetUserByID(ctx context.Context, id int) (User, error) {
	const query = `SELECT id, username, email, password_hash FROM users WHERE id = $1`
	var u User
	err := db.Pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
