package functions

import "strconv"

func ParseInt(str string) (int, error) {
	result, err := strconv.Atoi(str)
	return result, err
}