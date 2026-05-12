package validation

import "testing"

func TestNameValidator(t *testing.T) {
	v := &NameValidator{}

	tests := []struct {
		name     string
		input    string
		wantErr  error
	}{
		{"vazio", "", ErrorNameRequired},
		{"valido", "José Silva", nil},
		{"apenas espacos", "   ", nil},
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
