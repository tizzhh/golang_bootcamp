// This package provides functions for calculating the minimum amount of coins required to buy something.
package minCoins

import "slices"

// Original function for minimum coins calculation.
func MinCoins(val int, coins []int) []int {
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

// minCoins v2. Some corner cases are addressed:
// 1) If the coins slice is not sorted;
// 2) If there are values < 1 in the coins slice;
// 3) If you can't spend all your money with the slice provided.
func MinCoins2(val int, coins []int) []int {
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

// minCoins v2 Optimized.
// Added division and modulo to increase calculation speeds.
func MinCoins2Optimized(val int, coins []int) []int {
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
