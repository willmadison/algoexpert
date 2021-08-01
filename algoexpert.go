package algoexpert

func IsValidSubsequence(array, sequence []int) bool {
	if len(sequence) > len(array) {
		return false
	}

	var currentIndex int

	for _, v := range array {
		if v == sequence[currentIndex] {
			currentIndex++
		}

		if currentIndex >= len(sequence) {
			return true
		}
	}

	return false
}
