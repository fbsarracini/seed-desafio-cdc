package author

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func newTestHandler(t *testing.T) http.Handler {
	t.Helper()
	return CreateHandler(newTestDB(t))
}

func TestHandler_ServeHTTP(t *testing.T) {
	tests := []struct {
		name       string
		body       string
		wantStatus int
		wantErrs   []string
	}{
		{
			name:       "requisicao valida",
			body:       `{"name":"Fulano","email":"fulano@email.com","description":"Autor experiente."}`,
			wantStatus: http.StatusOK,
		},
		{
			name:       "json invalido",
			body:       `{invalid}`,
			wantStatus: http.StatusBadRequest,
		},
		{
			name:       "campos obrigatorios ausentes",
			body:       `{}`,
			wantStatus: http.StatusBadRequest,
			wantErrs:   []string{"nome obrigatório", "email obrigatório", "descrição obrigatória"},
		},
		{
			name:       "email invalido",
			body:       `{"name":"Fulano","email":"nao-e-email","description":"Desc."}`,
			wantStatus: http.StatusBadRequest,
			wantErrs:   []string{"email inválido"},
		},
		{
			name:       "descricao muito longa",
			body:       `{"name":"Fulano","email":"fulano@email.com","description":"` + strings.Repeat("x", 401) + `"}`,
			wantStatus: http.StatusBadRequest,
			wantErrs:   []string{"descrição deve ter no máximo 400 caracteres"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := newTestHandler(t)
			req := httptest.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			h.ServeHTTP(w, req)

			if w.Code != tt.wantStatus {
				t.Errorf("status = %d, want %d", w.Code, tt.wantStatus)
			}

			if len(tt.wantErrs) > 0 {
				var resp ErrorResponse
				if err := json.NewDecoder(w.Body).Decode(&resp); err != nil {
					t.Fatalf("decodificar resposta de erro: %v", err)
				}
				if len(resp.Errors) != len(tt.wantErrs) {
					t.Fatalf("erros = %v, want %v", resp.Errors, tt.wantErrs)
				}
				for i, e := range resp.Errors {
					if e != tt.wantErrs[i] {
						t.Errorf("erro[%d] = %q, want %q", i, e, tt.wantErrs[i])
					}
				}
			}
		})
	}
}

func TestHandler_EmailDuplicado(t *testing.T) {
	h := newTestHandler(t)

	body := `{"name":"Autor","email":"unico@email.com","description":"Descrição."}`

	first := httptest.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(body))
	w := httptest.NewRecorder()
	h.ServeHTTP(w, first)
	if w.Code != http.StatusOK {
		t.Fatalf("primeira criacao: status = %d, want 200", w.Code)
	}

	second := httptest.NewRequest(http.MethodPost, "/authors", bytes.NewBufferString(body))
	w2 := httptest.NewRecorder()
	h.ServeHTTP(w2, second)

	if w2.Code != http.StatusConflict {
		t.Errorf("segunda criacao: status = %d, want 409", w2.Code)
	}

	var resp ErrorResponse
	if err := json.NewDecoder(w2.Body).Decode(&resp); err != nil {
		t.Fatalf("decodificar resposta: %v", err)
	}
	if len(resp.Errors) == 0 || resp.Errors[0] != "email já cadastrado" {
		t.Errorf("erros = %v, want [\"email já cadastrado\"]", resp.Errors)
	}
}
