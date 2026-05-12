package author

type CreateRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}
