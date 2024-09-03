package functions

func GenerateEvenIntList(n int) []int {
	result := []int{}

	for num := range n + 1 {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}
