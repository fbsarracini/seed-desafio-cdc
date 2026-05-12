package validation

type NameValidator struct{}

func (v *NameValidator) Validate(name string) error {
	if name == "" {
		return ErrorNameRequired
	}
	return nil
}
