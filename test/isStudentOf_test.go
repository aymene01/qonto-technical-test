package test

import (
	"testing"

	f "github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestIsStudentOf(t *testing.T) {
	tests := []struct {
		name    string
		student f.Student
		college string
		want    bool
	}{
    {"Correct college", f.Student{College: "Harvard", Year: 2022}, "Harvard", true},
    {"Incorrect college", f.Student{College: "MIT", Year: 2022}, "Harvard", false},
    {"Empty college", f.Student{College: "", Year: 2022}, "Harvard", false},
    {"Matching empty college", f.Student{College: "", Year: 2022}, "", true},
}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := f.IsStudentOf(tt.student, tt.college); got != tt.want {
				t.Errorf("isStudentOf() = %v, want %v", got, tt.want)
			}
		})
	}
}