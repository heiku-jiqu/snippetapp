package models

import (
	"database/sql"
	"errors"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type SnippetModel struct {
	DB *sql.DB
}

// Inserts snippet into database and returns id if successful
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	sql := `INSERT INTO snippets (title, content, created, expires) 
VALUES($1, $2, NOW(), NOW() + make_interval(days=> $3) )
RETURNING id`
	var id int
	err := m.DB.QueryRow(sql, title, content, expires).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, err
}

// Returns Snippet based on id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets 
WHERE expires > NOW() AND id = $1`
	s := &Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// Returns the 10 most recently created Snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
