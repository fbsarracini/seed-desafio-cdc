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

func (r *AuthorRepository) EmailExists(ctx context.Context, email string) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM authors WHERE email = ?)`,
		email,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
