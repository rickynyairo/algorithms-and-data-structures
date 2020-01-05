package dividenconc

// returns the maximum of the provided values
func Max(values ...int) int {
	maximum := values[0]
	for _, val := range values {
		if val > maximum {
			maximum = val
		}
	}
	return maximum
}

func HouseThief(worth []int, current int) int {
	if current >= len(worth) {
		return 0
	}
	steal_current_house := worth[current] + HouseThief(worth, current+2)
	skip_current_house := HouseThief(worth, current+1)
	return Max(steal_current_house, skip_current_house)
}
