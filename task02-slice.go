package homework

func reverse(input []int64) (result []int64) {
	for i := len(input)/2 - 1; i >= 0; i-- {
		j := len(input) - 1 - i
		input[i], input[j] = input[j], input[i]
	}
	result = make([]int64, 4)
	copy(result, input)
	return result
}
