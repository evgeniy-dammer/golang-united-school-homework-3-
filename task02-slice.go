package homework

func reverse(input []int64) (result []int64) {
	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i])
	}

	return result
}
