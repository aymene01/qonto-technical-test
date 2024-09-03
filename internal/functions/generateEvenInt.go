package functions

func GenerateEvenIntList(n int) []int {
	result := []int{}

	for num := range n {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}
