package author

import (
	"context"
	"testing"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"
)

func TestAuthorService_Create(t *testing.T) {
	ctx := context.Background()

	t.Run("cria com sucesso", func(t *testing.T) {
		db := newTestDB(t)
		svc := NewAuthorService(NewAuthorRepository(db))

		err := svc.Create(ctx, CreateRequest{
			Name:        "Beltrano",
			Email:       "beltrano@email.com",
			Description: "Escritora.",
		})
		if err != nil {
			t.Fatalf("Create() erro inesperado: %v", err)
		}
	})

	t.Run("retorna erro se email ja existe", func(t *testing.T) {
		db := newTestDB(t)
		repo := NewAuthorRepository(db)
		svc := NewAuthorService(repo)

		req := CreateRequest{Name: "Cicrano", Email: "cicrano@email.com", Description: "Dev."}
		if err := svc.Create(ctx, req); err != nil {
			t.Fatalf("primeira criacao falhou: %v", err)
		}

		err := svc.Create(ctx, req)
		if err != validation.ErrorEmailAlreadyExists {
			t.Errorf("Create() = %v, want ErrorEmailAlreadyExists", err)
		}
	})
}
