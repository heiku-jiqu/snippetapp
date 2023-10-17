package mocks

import (
	"time"

	"github.com/heiku-jiqu/snippetapp/internal/models"
)

type UserModel struct{}

func (u *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	if email == "bob@example.com" && password == "pa$$word" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

func (u *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (u *UserModel) Get(id int) (*models.User, error) {
	switch id {
	case 1:
		return &models.User{ID: 1, Name: "Bob", Email: "bob@example.com", Created: time.Now()}, nil
	default:
		return nil, models.ErrNoRecord
	}
}
