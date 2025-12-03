package part

import "testing"

func TestNormalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1", "first"},
		{"first", "first"},
		{"2", "second"},
		{"second", "second"},
		{"", ""},
		{"3", ""},
		{"invalid", ""},
		{"First", ""}, // Case sensitive
		{"FIRST", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Normalize(tt.input)
			if result != tt.expected {
				t.Errorf("Normalize(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"first", "1"},
		{"second", "2"},
		{"", "2"},        // Default case
		{"invalid", "2"}, // Non-first returns "2"
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Label(tt.input)
			if result != tt.expected {
				t.Errorf("Label(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
