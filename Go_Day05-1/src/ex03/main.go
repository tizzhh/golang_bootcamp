package main

import (
	"fmt"
	"knapSack/knapsack"
)

func main() {
	presents := []knapsack.Present{{Value: 3, Size: 1}, {Value: 6, Size: 4}, {Value: 4, Size: 3}, {Value: 4, Size: 1}}
	fmt.Println(knapsack.GrabPresents(presents, 4))
}
