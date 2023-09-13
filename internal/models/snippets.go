package models

import (
	"database/sql"
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
	return &Snippet{}, nil
}

// Returns the 10 most recently created Snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
