package test

import (
	"testing"

	"github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestIsTruthy(t *testing.T) {
	tests := []struct {
		name string
		arr  []any
		want bool
	}{
		{"Empty array", []any{}, false},
		{"Array with no true", []any{false, 0, "", nil}, false},
		{"Array with true", []any{false, true, 0, ""}, true},
		{"Array with only true", []any{true}, true},
		{"Array with multiple true", []any{true, true}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := functions.IsTruthy(tt.arr); got != tt.want {
				t.Errorf("isTruthy() = %v, want %v", got, tt.want)
			}
		})
	}
}
