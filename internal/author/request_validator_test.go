package author

import (
	"strings"
	"testing"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"
)

func TestCreateRequestValidator(t *testing.T) {
	v := NewCreateRequestValidator()

	tests := []struct {
		name     string
		req      CreateRequest
		wantErrs []string
	}{
		{
			name:     "todos os campos validos",
			req:      CreateRequest{Name: "Fulano", Email: "fulano@email.com", Description: "Uma descrição válida."},
			wantErrs: nil,
		},
		{
			name:     "nome vazio",
			req:      CreateRequest{Name: "", Email: "fulano@email.com", Description: "Descrição."},
			wantErrs: []string{validation.ErrorNameRequired.Error()},
		},
		{
			name:     "email vazio",
			req:      CreateRequest{Name: "Fulano", Email: "", Description: "Descrição."},
			wantErrs: []string{validation.ErrorEmailRequired.Error()},
		},
		{
			name:     "email invalido",
			req:      CreateRequest{Name: "Fulano", Email: "nao-e-email", Description: "Descrição."},
			wantErrs: []string{validation.ErrorEmailInvalid.Error()},
		},
		{
			name:     "descricao vazia",
			req:      CreateRequest{Name: "Fulano", Email: "fulano@email.com", Description: ""},
			wantErrs: []string{validation.ErrorDescriptionRequired.Error()},
		},
		{
			name:     "descricao muito longa",
			req:      CreateRequest{Name: "Fulano", Email: "fulano@email.com", Description: strings.Repeat("x", 401)},
			wantErrs: []string{validation.ErrorDescriptionTooLong.Error()},
		},
		{
			name: "todos os campos invalidos",
			req:  CreateRequest{Name: "", Email: "", Description: ""},
			wantErrs: []string{
				validation.ErrorNameRequired.Error(),
				validation.ErrorEmailRequired.Error(),
				validation.ErrorDescriptionRequired.Error(),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errs := v.Validate(tt.req)

			if len(errs) != len(tt.wantErrs) {
				t.Fatalf("Validate() = %v erros, want %v: got %v", len(errs), len(tt.wantErrs), errs)
			}

			for i, e := range errs {
				if e != tt.wantErrs[i] {
					t.Errorf("erro[%d] = %q, want %q", i, e, tt.wantErrs[i])
				}
			}
		})
	}
}
