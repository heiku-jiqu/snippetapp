package models

import (
	// "context"
	// "database/sql"
	// "errors"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
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
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	sql := `INSERT INTO users (name, email, hashed_password, created)
	   VALUES ($1, $2, $3, NOW())
	   RETURNING id`
	_, err = u.DB.Exec(
		context.Background(),
		sql,
		name, email, string(hashed_password),
	)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" && strings.Contains(pgErr.Message, "unique") {
				return ErrDuplicateEmail
			}
		}
		return err
	}
	return nil
}

func (u *UserModel) Authenticate(email, password string) (id int, err error) {
	return 0, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
