package author

import (
	"context"
	"testing"
)

func TestAuthorRepository_EmailExists(t *testing.T) {
	db := newTestDB(t)
	repo := NewAuthorRepository(db)
	ctx := context.Background()

	t.Run("nao existe", func(t *testing.T) {
		exists, err := repo.EmailExists(ctx, "novo@email.com")
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		if exists {
			t.Error("esperava false, mas obteve true")
		}
	})

	t.Run("existe apos criacao", func(t *testing.T) {
		req := CreateRequest{Name: "Beltrano", Email: "beltrano@email.com", Description: "Descrição"}
		if err := repo.Create(ctx, req); err != nil {
			t.Fatalf("Create() erro: %v", err)
		}

		exists, err := repo.EmailExists(ctx, req.Email)
		if err != nil {
			t.Fatalf("erro inesperado: %v", err)
		}
		if !exists {
			t.Error("esperava true, mas obteve false")
		}
	})
}

func TestAuthorRepository_Create(t *testing.T) {
	db := newTestDB(t)
	repo := NewAuthorRepository(db)
	ctx := context.Background()

	t.Run("cria com sucesso", func(t *testing.T) {
		req := CreateRequest{Name: "Fulano", Email: "fulano@github.com", Description: "Especialista em Go."}
		if err := repo.Create(ctx, req); err != nil {
			t.Fatalf("Create() erro: %v", err)
		}
	})

	t.Run("falha com email duplicado", func(t *testing.T) {
		req := CreateRequest{Name: "Ciclano", Email: "ciclano@email.com", Description: "Dev backend."}
		if err := repo.Create(ctx, req); err != nil {
			t.Fatalf("primeira criacao falhou: %v", err)
		}
		if err := repo.Create(ctx, req); err == nil {
			t.Error("esperava erro de email duplicado, mas obteve nil")
		}
	})
}
