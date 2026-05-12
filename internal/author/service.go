package author

// 4 ICP Points

import (
	"context"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"
)

type AuthorService struct {
	// 1
	repository *AuthorRepository
}

func NewAuthorService(repository *AuthorRepository) *AuthorService {
	return &AuthorService{repository: repository}
}

func (s *AuthorService) Create(ctx context.Context, req CreateRequest) error {
	exists, err := s.repository.EmailExists(ctx, req.Email)
	// 1
	if err != nil {
		return err
	}

	// 1
	if exists {
		return validation.ErrorEmailAlreadyExists
	}

	// 1
	if err := s.repository.Create(ctx, req); err != nil {
		return err
	}

	return nil
}
