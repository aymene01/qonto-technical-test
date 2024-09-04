package test

import (
	"reflect"
	"testing"

	"github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestCountFrequency(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string]int
	}{
		{
			input:    []string{"apple", "banana", "apple", "orange", "banana", "apple"},
			expected: map[string]int{"apple": 3, "banana": 2, "orange": 1},
		},
		{
			input:    []string{"a", "b", "a", "c", "b", "a", "c", "c"},
			expected: map[string]int{"a": 3, "b": 2, "c": 3},
		},
		{
			input:    []string{"single"},
			expected: map[string]int{"single": 1},
		},
		{
			input:    []string{},
			expected: map[string]int{},
		},
		{
			input:    []string{"repeat", "repeat", "repeat"},
			expected: map[string]int{"repeat": 3},
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := functions.CountFrequency(tt.input)
			if !reflect.DeepEqual(actual, tt.expected) {
				t.Errorf("got %v, want %v", actual, tt.expected)
			}
		})
	}
}
