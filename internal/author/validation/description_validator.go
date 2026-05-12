package validation

type DescriptionValidator struct {
	maxLength int
}

func NewDescriptionValidator(maxLength int) *DescriptionValidator {
	return &DescriptionValidator{maxLength: maxLength}
}

func (v *DescriptionValidator) Validate(description string) error {
	if description == "" {
		return ErrorDescriptionRequired
	}
	if len([]rune(description)) > v.maxLength {
		return ErrorDescriptionTooLong
	}
	return nil
}
