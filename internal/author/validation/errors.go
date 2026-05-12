package validation

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

var (
	ErrorNameRequired        = ValidationError{"nome obrigatório"}
	ErrorEmailRequired       = ValidationError{"email obrigatório"}
	ErrorEmailInvalid        = ValidationError{"email inválido"}
	ErrorEmailAlreadyExists  = ValidationError{"email já cadastrado"}
	ErrorDescriptionRequired = ValidationError{"descrição obrigatória"}
	ErrorDescriptionTooLong  = ValidationError{"descrição deve ter no máximo 400 caracteres"}
)
