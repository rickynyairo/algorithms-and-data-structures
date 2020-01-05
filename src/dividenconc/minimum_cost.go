package dividenconc

import "math"

// calculates the minimum cost to traverse
// a 2D array from first cell to last
func MinCost() int {
	cost := [][]int{
		{4, 7, 8, 6, 4},
		{6, 7, 3, 9, 2},
		{3, 8, 1, 2, 4},
		{7, 1, 7, 3, 7},
		{2, 9, 8, 9, 3}}
	return minCost(cost, len(cost)-1, len(cost[0])-1)
}

func minCost(cost [][]int, row int, col int) int {
	if row == -1 || col == -1 {
		return math.MaxInt32
	}
	if row == 0 && col == 0 {
		// we are at the fist cell, stop recurring
		return cost[row][col]
	}
	minCost1 := minCost(cost, row-1, col)
	minCost2 := minCost(cost, row, col-1)
	actualMin := Min(minCost1, minCost2)
	currentCost := cost[row][col]
	return actualMin + currentCost
}
