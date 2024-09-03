package functions

func IsTruthy(arr []any) bool {
	for _, e := range arr {
		if e == true {
			return true
		}
	}
	return false
}
