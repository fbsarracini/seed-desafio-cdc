package author

// 4 ICP Points

import "context"

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

	// 1
	if err := s.repository.Create(ctx, req); err != nil {
		return nil, err
	}

	return nil, nil
}
