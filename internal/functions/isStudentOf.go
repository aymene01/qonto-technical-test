package functions

type Student struct {
	College string
	Year    int
}

func IsStudentOf(student Student, college string) bool {
	return student.College == college
}
