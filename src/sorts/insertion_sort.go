package sorts

func InsertionSort(numbers []int) []int {
	for i := 1; i < len(numbers); i++ {
		temp := numbers[i]
		j := i
		for j > 0 && numbers[j-1] > temp {
			numbers[j] = numbers[j-1]
			j--
		}
		numbers[j] = temp
	}
	return numbers
}
