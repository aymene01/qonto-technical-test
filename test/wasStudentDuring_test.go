package test

import (
	"testing"

	"github.com/aymene01/qonto-technical-test/internal/functions"
)

func TestWasStudentDuring(t *testing.T) {
	tests := []struct {
		student       functions.Student
		targetCollege string
		startYear     int
		endYear       int
		expected      bool
	}{
		{
			student:       functions.Student{College: "cambridge", Year: 2008},
			targetCollege: "cambridge",
			startYear:     2004,
			endYear:       2009,
			expected:      true,
		},
		{
			student:       functions.Student{College: "oxford", Year: 2005},
			targetCollege: "cambridge",
			startYear:     2004,
			endYear:       2007,
			expected:      false,
		},
		{
			student:       functions.Student{College: "cambridge", Year: 2002},
			targetCollege: "cambridge",
			startYear:     2000,
			endYear:       2004,
			expected:      true,
		},
		{
			student:       functions.Student{College: "cambridge", Year: 2004},
			targetCollege: "cambridge",
			startYear:     2004,
			endYear:       2004,
			expected:      true,
		},
		{
			student:       functions.Student{College: "cambridge", Year: 2010},
			targetCollege: "cambridge",
			startYear:     2005,
			endYear:       2009,
			expected:      false,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := functions.WasStudentDuring(tt.student, tt.targetCollege, tt.startYear, tt.endYear)
			if actual != tt.expected {
				t.Errorf("WasStudentDuring(%v, %s, %d, %d) = %v; want %v",
					tt.student, tt.targetCollege, tt.startYear, tt.endYear, actual, tt.expected)
			}
		})
	}
}