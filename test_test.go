package main

import (
	"testing"

	f "github.com/aymene01/qonto-technical-test/functions"
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
			if got := f.IsTruthy(tt.arr); got != tt.want {
				t.Errorf("isTruthy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFrequency(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		want map[string]int
	}{
		{"Empty array", []string{}, map[string]int{}},
		{"Single element", []string{"a"}, map[string]int{"a": 1}},
		{"Multiple elements with one occurrence", []string{"a", "b", "c"}, map[string]int{"a": 1, "b": 1, "c": 1}},
		{"Multiple elements with multiple occurrences", []string{"a", "b", "a", "b", "c"}, map[string]int{"a": 2, "b": 2, "c": 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := f.CountFrequency(tt.arr)
			for key, value := range tt.want {
				if got[key] != value {
					t.Errorf("countFrequency()[%v] = %v, want %v", key, got[key], value)
				}
			}
		})
	}
}

func TestIsStudentOf(t *testing.T) {
	tests := []struct {
		name    string
		obj     f.ObjArg
		college string
		want    bool
	}{
    {"Correct college", f.ObjArg{College: "Harvard", Year: 2022}, "Harvard", true},
    {"Incorrect college", f.ObjArg{College: "MIT", Year: 2022}, "Harvard", false},
    {"Empty college", f.ObjArg{College: "", Year: 2022}, "Harvard", false},
    {"Matching empty college", f.ObjArg{College: "", Year: 2022}, "", true},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.IsStudentOf(tt.obj, tt.college); got != tt.want {
				t.Errorf("isStudentOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
		{"99999999999999999999", 0, true}, // out of range
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