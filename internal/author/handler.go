package author

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"time"
)

// todo autor tem um nome email e descrição
// nome, email e descrição são obrigatórios
// e a descrição é obrigatória e não pode passar de 400 caracteres
//
// resultado esperado
// um novo autor cadastrado no banco de dados e um status 200 retornado
// se o usuario já estiver criado (mesmo email), retornar 200

// 1
var emailRegex = regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`)

// 1
type createRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

func CreateHandler(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var request createRequest
		// 1
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// 1
		if errs := validate(request); len(errs) > 0 {
			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusBadRequest)

			json.NewEncoder(w).Encode(map[string][]string{"errors": errs})
			return
		}

		_, err := db.ExecContext(r.Context(),
			`INSERT INTO authors (name, email, description, created_at) VALUES (?, ?, ?, ?)`,
			request.Name, request.Email, request.Description, time.Now().UTC(),
		)

		// 1
		if err != nil {
			http.Error(w, "Erro interno no servidor", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
	}
}

func validate(request createRequest) []string {
	var errs []string

	// 1
	if request.Name == "" {
		errs = append(errs, "nome obrigatório")
	}

	// 1
	// 1
	if request.Email == "" {
		errs = append(errs, "email obrigatório")
	} else if !emailRegex.MatchString(request.Email) {
		errs = append(errs, "email inválido")
	}

	// 1
	// 1
	if request.Description == "" {
		errs = append(errs, "descrição obrigatória")
	} else if (len([]rune(request.Description))) > 400 {
		errs = append(errs, "descrição deve ter no máximo 400 caracteres")
	}

	return errs
}
