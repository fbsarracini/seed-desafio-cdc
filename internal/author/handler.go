package author

// 6 ICP Points

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

// Handler implementa o handler HTTP para criar autores
type Handler struct {
	// 1
	service *AuthorService
}

// CreateHandler é a factory function que cria um handler completo com suas dependências
func CreateHandler(db *sql.DB) http.Handler {
	// 1
	repository := NewAuthorRepository(db)
	service := NewAuthorService(repository)
	return &Handler{service: service}
}

// ServeHTTP processa a requisição HTTP
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// 1
	var req CreateRequest
	// 1
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	validationErrs, err := h.service.Create(r.Context(), req)

	// 1
	if len(validationErrs) > 0 {
		h.respondWithErrors(w, validationErrs)
		return
	}

	// 1
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respondWithErrors(w http.ResponseWriter, errs []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
}
