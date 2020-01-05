package dividenconc

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
func HouseThief(worth []int, current int) int {
	if current >= len(worth) {
		return 0
	}
	steal_current_house := worth[current] + HouseThief(worth, current+2)
	skip_current_house := HouseThief(worth, current+1)
	return max(steal_current_house, skip_current_house)
}
