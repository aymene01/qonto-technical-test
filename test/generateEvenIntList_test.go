package test

import (
	"reflect"
	"testing"

	"github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestGenerateEvenIntList(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int
	}{
		{"Test with 0", 0, []int{0}},
		{"Test with 1", 1, []int{0}},
		{"Test with 2", 2, []int{0, 2}},
		{"Test with 5", 5, []int{0, 2, 4}},
		{"Test with 10", 10, []int{0, 2, 4, 6, 8, 10}},
		{"Test with negative", -5, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := functions.GenerateEvenIntList(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
