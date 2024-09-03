package functions

func CountFrequency(arr []string) map[string]int{
	count := make(map[string]int)

	for _, str := range arr {
		count[str] += 1
	}

	return count
}