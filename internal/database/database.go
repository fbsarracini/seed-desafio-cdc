package database

import (
	"database/sql"

	// _ import: carrega o driver sem usar diretamente.
	_ "modernc.org/sqlite"
)

func Open(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}

	if err := migrate(db); err != nil {
		return nil, err
	}

	return db, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS authors (
			id          INTEGER PRIMARY KEY AUTOINCREMENT,  -- ID único, auto-incrementa
			name        TEXT    NOT NULL,                   -- Nome obrigatório
			email       TEXT    NOT NULL UNIQUE,            -- Email obrigatório e único
			description TEXT    NOT NULL,                   -- Descrição obrigatória
			created_at  DATETIME NOT NULL                   -- Timestamp de criação
		)
	`)
	return err
}
