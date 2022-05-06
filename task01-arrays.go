package homework

func average(input [15]float32) (result float32) {
	var length float32

	for i := 0; i < len(input); i++ {
		if input[i] == 0 {
			continue
		}

		result += input[i]
		length++
	}

	return result / length
}
