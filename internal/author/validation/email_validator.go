package validation

import "regexp"

type EmailValidator struct {
	regex *regexp.Regexp
}

func NewEmailValidator() *EmailValidator {
	return &EmailValidator{
		regex: regexp.MustCompile(`^[^@\s]+@[^@\s]+\.[^@\s]+$`),
	}
}

func (v *EmailValidator) Validate(email string) error {
	if email == "" {
		return ErrorEmailRequired
	}
	if !v.regex.MatchString(email) {
		return ErrorEmailInvalid
	}
	return nil
}
