package validation

import (
	"strings"
	"testing"
)

func TestDescriptionValidator(t *testing.T) {
	v := NewDescriptionValidator(400)

	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{"vazia", "", ErrorDescriptionRequired},
		{"acima do limite", strings.Repeat("a", 401), ErrorDescriptionTooLong},
		{"exatamente no limite", strings.Repeat("a", 400), nil},
		{"valida curta", "Uma boa descrição.", nil},
		{"unicode no limite", strings.Repeat("ã", 400), nil},
		{"unicode acima do limite", strings.Repeat("ã", 401), ErrorDescriptionTooLong},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := v.Validate(tt.input)
			if err != tt.wantErr {
				t.Errorf("Validate(%q...) = %v, want %v", tt.input[:min(20, len(tt.input))], err, tt.wantErr)
			}
		})
	}
}
