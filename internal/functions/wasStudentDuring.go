package functions

func WasStudentDuring(student Student, targetCollege string, startYear, endYear int) bool {
	return student.College == targetCollege && student.Year >= startYear && student.Year <= endYear 
}