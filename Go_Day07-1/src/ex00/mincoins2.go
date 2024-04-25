package minCoins

import (
	"slices"
)

func minCoins2(val int, coins []int) []int {
	slices.Sort(coins)
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		if coins[i] < 1 {
			return []int{}
		}
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	if val != 0 {
		return []int{}
	}
	return res
}
