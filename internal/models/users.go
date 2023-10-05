package models

import (
	// "context"
	// "database/sql"
	// "errors"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	ID      int
	Name    string
	Email   string
	Hash    []byte
	Created time.Time
}

type UserModel struct {
	DB *pgxpool.Pool
}

func (u *UserModel) Insert(name, email, password string) error {
	// sql := `INSERT INTO users (name, email, hashed_password, created)
	//    VALUES ($1, $2, $3, NOW())
	//    RETURNING id`
	// u.DB.QueryRow(
	// 	context.Background(),
	// 	sql,
	// 	name, email,
	// )
	return nil
}

func (u *UserModel) Authenticate(email, password string) (id int, err error) {
	return 0, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
