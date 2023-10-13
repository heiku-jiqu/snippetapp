package models

import (
	"context"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newTestDB(t *testing.T) *pgxpool.Pool {
	db, err := pgxpool.New(
		context.Background(),
		"postgres://test_web:pass@localhost:5432/test_snippetapp?sslmode=disable",
	)
	if err != nil {
		t.Fatal(err)
	}

	script, err := os.ReadFile("../../sql/testdata/setup_tables.sql")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.Exec(context.Background(), string(script))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		script, err := os.ReadFile("../../sql/testdata/teardown.sql")
		if err != nil {
			t.Fatal(err)
		}
		_, err = db.Exec(context.Background(), string(script))
		if err != nil {
			t.Fatal(err)
		}
		db.Close()
	})

	return db
}
