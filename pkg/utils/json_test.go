package utils

import (
	"testing"
)

func TestPrettyPrintJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string // We might just check if it's different or contains newlines
	}{
		{
			name:  "Valid JSON",
			input: `{"key":"value"}`,
		},
		{
			name:  "Invalid JSON",
			input: `not json`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := PrettyPrintJSON([]byte(tt.input))
			if tt.name == "Invalid JSON" && res != tt.input {
				t.Errorf("expected original string for invalid json, got %s", res)
			}
			if tt.name == "Valid JSON" && res == tt.input {
				t.Errorf("expected formatted json, got original")
			}
		})
	}
}
