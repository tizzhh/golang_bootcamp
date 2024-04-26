package minCoins

import "slices"

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

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

func minCoins2Optimized(val int, coins []int) []int {
	slices.Sort(coins)
	res := make([]int, 0)
	for i := len(coins) - 1; i >= 0; i-- {
		if coins[i] < 1 {
			return []int{}
		}
		nubmerOfSubs := val / coins[i]
		for j := 0; j < nubmerOfSubs; j++ {
			res = append(res, coins[i])
		}
		val %= coins[i]
	}
	if val != 0 {
		return []int{}
	}
	return res
}
