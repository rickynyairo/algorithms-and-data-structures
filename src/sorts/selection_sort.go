package sorts

func SelectionSort(numbers []int) []int {
	n := len(numbers)
	for j,_ := range numbers{
		minIndex := j
		for i := j+1; i < n; i++{
			if numbers[i] < numbers[minIndex]{
				minIndex = i
			}
		}
		if minIndex != j{
			numbers[j], numbers[minIndex] = numbers[minIndex], numbers[j]
		}
	}
	return numbers
}
