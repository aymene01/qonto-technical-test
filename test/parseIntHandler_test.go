package test

import (
	"testing"

	f "github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		input    string
		expected int
		hasError bool
	}{
		{"123", 123, false},
		{"0", 0, false},
		{"-456", -456, false},
		{"42abc", 0, true},
		{"", 0, true},
		{"99999999999999999999", 0, true},
	}

	for _, test := range tests {
		result, err := f.ParseInt(test.input)

		if test.hasError {
			if err == nil {
				t.Errorf("Expected an error for input %q, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect an error for input %q, but got %v", test.input, err)
			}
			if result != test.expected {
				t.Errorf("Expected %d for input %q, but got %d", test.expected, test.input, result)
			}
		}
	}
}