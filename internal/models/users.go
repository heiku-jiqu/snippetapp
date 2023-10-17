package models

import (
	// "context"
	// "database/sql"
	// "errors"
	"context"
	"errors"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

const bcryptCost int = 12

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

type UserModelInterface interface {
	Authenticate(email string, password string) (int, error)
	Exists(id int) (bool, error)
	Insert(name string, email string, password string) error
	Get(id int) (*User, error)
	ChangePassword(id int, oldPassword, newPassword string) error
}

func (u *UserModel) Insert(name, email, password string) error {
	hashed_password, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
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

func (u *UserModel) Authenticate(email, password string) (int, error) {
	sql := `SELECT id, hashed_password FROM users WHERE email = $1`
	var id int
	var dbHashPassword []byte
	err := u.DB.QueryRow(context.Background(), sql, email).Scan(&id, &dbHashPassword)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	err = bcrypt.CompareHashAndPassword(dbHashPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}

	return id, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	var exists bool
	sql := `SELECT EXISTS(SELECT true FROM users WHERE id = $1)`
	err := u.DB.QueryRow(context.Background(), sql, id).Scan(&exists)
	return exists, err
}

func (u *UserModel) Get(id int) (*User, error) {
	sql := `SELECT name, email, created FROM users WHERE id = $1`
	user := User{ID: id}
	err := u.DB.QueryRow(context.Background(), sql, id).Scan(&user.Name, &user.Email, &user.Created)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &User{}, ErrNoRecord
		} else {
			return &User{}, err
		}
	}
	return &user, nil
}

func (u *UserModel) ChangePassword(id int, oldPassword, newPassword string) error {
	oldHash, err := u.GetPasswordHash(id)
	if err != nil {
		return ErrNoRecord
	}

	err = bcrypt.CompareHashAndPassword(oldHash, []byte(oldPassword))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidCredentials
		} else {
			return err
		}
	}

	newHash, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcryptCost)
	if err != nil {
		return err
	}

	sql := `UPDATE users SET hashed_password = $1 WHERE id = $2`
	_, err = u.DB.Exec(context.Background(), sql, newHash, id)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) GetPasswordHash(id int) ([]byte, error) {
	sql := `SELECT hashed_password FROM users WHERE id = $1`
	var hash []byte
	err := u.DB.QueryRow(context.Background(), sql, id).Scan(&hash)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrNoRecord
		}
		return nil, err
	}

	return hash, nil
}
