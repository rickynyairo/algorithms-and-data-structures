package sorts

func BubbleSort(numbers []int) []int {
	n := len(numbers)
	for i, _ := range numbers {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}
