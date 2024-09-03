package functions

type ObjArg struct {
	College string
	Year    int
}

func IsStudentOf(obj ObjArg, college string) bool {
	return obj.College == college
}