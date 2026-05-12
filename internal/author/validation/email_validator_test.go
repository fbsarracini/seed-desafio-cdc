package validation

import "testing"

func TestEmailValidator(t *testing.T) {
	v := NewEmailValidator()

	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{"vazio", "", ErrorEmailRequired},
		{"sem arroba", "emailinvalido.com", ErrorEmailInvalid},
		{"sem dominio", "email@", ErrorEmailInvalid},
		{"sem extensao", "email@dominio", ErrorEmailInvalid},
		{"valido", "autor@casadocodigo.com.br", nil},
		{"valido simples", "a@b.co", nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.input)
			if err != tt.wantErr {
				t.Errorf("Validate(%q) = %v, want %v", tt.input, err, tt.wantErr)
			}
		})
	}
}
