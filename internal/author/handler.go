package author

// 7 ICP Points

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"
)

type Handler struct {
	// 1
	validator *CreateRequestValidator
	// 1
	service *AuthorService
}

func CreateHandler(db *sql.DB) http.Handler {
	// 1
	repository := NewAuthorRepository(db)
	service := NewAuthorService(repository)
	validator := NewCreateRequestValidator()
	return &Handler{service: service, validator: validator}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1
	var req CreateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request json", http.StatusBadRequest)
		return
	}

	// 1
	if errs := h.validator.Validate(req); len(errs) > 0 {
		h.respondWithErrors(w, errs, http.StatusBadRequest)
		return
	}

	err := h.service.Create(r.Context(), req)

	// 1
	if errors.Is(err, validation.ErrorEmailAlreadyExists) {
		h.respondWithErrors(w, []string{err.Error()}, http.StatusConflict)
		return
	}

	// 1
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) respondWithErrors(w http.ResponseWriter, errs []string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorResponse{Errors: errs})
}
