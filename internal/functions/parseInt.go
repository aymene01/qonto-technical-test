package functions

import "strconv"

func ParseInt(str string) (int, error) {
	return strconv.Atoi(str)
}