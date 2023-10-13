package models

import (
	"testing"

	"github.com/heiku-jiqu/snippetapp/internal/assert"
)

func TestUserModelExists(t *testing.T) {
	conn := newTestDB(t)
	users := UserModel{conn}
	t.Run("User exists", func(t *testing.T) {
		res, err := users.Exists(1)
		assert.Equal(t, res, true)
		assert.NilError(t, err)
	})
	t.Run("Zero ID does not exist", func(t *testing.T) {
		res, err := users.Exists(0)
		assert.Equal(t, res, false)
		assert.NilError(t, err)
	})
	t.Run("User does not exist", func(t *testing.T) {
		res, err := users.Exists(999)
		assert.Equal(t, res, false)
		assert.NilError(t, err)
	})
}
