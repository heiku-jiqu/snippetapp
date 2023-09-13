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

// Inserts snippet into database
func (m *SnippetModel) Insert(title, content string, expires int) (int, error) {
	return 0, nil
}

// Returns Snippet based on id
func (m *SnippetModel) Get(id int) (*Snippet, error) {
	return &Snippet{}, nil
}

// Returns the 10 most recently created Snippets.
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
