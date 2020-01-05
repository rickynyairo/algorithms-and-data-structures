package greedy

import (
	"fmt"
	"sorts"
)

func CoinChange(coins []int, amount int) {
	sorts.SelectionSort(coins)
	index := len(coins) - 1
	for {
		coin_value := coins[index]
		index--
		maxAmount := (amount / coin_value) * coin_value
		if maxAmount > 0 {
			fmt.Println("Coin Value: ", coin_value, " taken count: ", (amount / coin_value))
			amount -= maxAmount
		}
		if amount == 0 {
			break
		}

	}
}
