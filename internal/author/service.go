package author

// 6 ICP Points

import (
	"context"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"
)

type AuthorService struct {
	// 1
	validator *CreateRequestValidator
	// 1
	repository *AuthorRepository
}

// construtor
func NewAuthorService(repository *AuthorRepository) *AuthorService {
	return &AuthorService{
		validator:  NewCreateRequestValidator(),
		repository: repository,
	}
}

func (s *AuthorService) Create(ctx context.Context, req CreateRequest) ([]string, error) {

	// 1
	if errs := s.validator.Validate(req); len(errs) > 0 {
		return errs, nil
	}

	exists, err := s.repository.EmailExists(ctx, req.Email)
	// 1
	if err != nil {
		return nil, err
	}

	// 1
	if exists {
		return []string{validation.ErrorEmailAlreadyExists.Error()}, nil
	}

	// 1
	if err := s.repository.Create(ctx, req); err != nil {
		return nil, err
	}

	return nil, nil
}
