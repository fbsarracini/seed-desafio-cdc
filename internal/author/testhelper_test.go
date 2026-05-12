package author

import (
	"database/sql"
	"testing"

	_ "modernc.org/sqlite"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()
	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("abrir db: %v", err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS authors (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,
			name        TEXT    NOT NULL,
			email       TEXT    NOT NULL UNIQUE,
			description TEXT    NOT NULL,
			created_at  DATETIME NOT NULL
		)
	`)
	if err != nil {
		t.Fatalf("migrar db: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}
