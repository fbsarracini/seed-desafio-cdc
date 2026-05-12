package author

import "github.com/fbsarracini/seed-desafio-cdc/internal/author/validation"

type CreateRequestValidator struct {
	nameValidator        *validation.NameValidator
	emailValidator       *validation.EmailValidator
	descriptionValidator *validation.DescriptionValidator
}

func NewCreateRequestValidator() *CreateRequestValidator {
	return &CreateRequestValidator{
		nameValidator:        &validation.NameValidator{},
		emailValidator:       validation.NewEmailValidator(),
		descriptionValidator: validation.NewDescriptionValidator(400),
	}
}

func (v *CreateRequestValidator) Validate(req CreateRequest) []string {
	var errs []string

	if err := v.nameValidator.Validate(req.Name); err != nil {
		errs = append(errs, err.Error())
	}

	if err := v.emailValidator.Validate(req.Email); err != nil {
		errs = append(errs, err.Error())
	}

	if err := v.descriptionValidator.Validate(req.Description); err != nil {
		errs = append(errs, err.Error())
	}

	return errs
}
