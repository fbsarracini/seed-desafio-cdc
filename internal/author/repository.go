package author

import (
	"context"
	"database/sql"
	"time"
)

type AuthorRepository struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(ctx context.Context, req CreateRequest) error {
	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO authors (name, email, description, created_at) VALUES (?, ?, ?, ?)`,
		req.Name,
		req.Email,
		req.Description,
		time.Now().UTC(),
	)
	return err
}
