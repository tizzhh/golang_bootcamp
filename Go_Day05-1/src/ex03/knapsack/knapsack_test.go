package knapsack_test

import (
	"knapSack/knapsack"
	"testing"
)

type sackTest struct {
	expected, arr []knapsack.Present
	capacity      int
}

var testCases = []sackTest{
	{
		expected: []knapsack.Present{{Value: 3, Size: 1}, {Value: 4, Size: 1}},
		arr:      []knapsack.Present{{Value: 3, Size: 1}, {Value: 6, Size: 4}, {Value: 4, Size: 3}, {Value: 4, Size: 1}},
		capacity: 4,
	},
	{
		expected: []knapsack.Present{},
		arr:      []knapsack.Present{{Value: 3, Size: 1}, {Value: 6, Size: 4}, {Value: 4, Size: 3}, {Value: 4, Size: 1}},
		capacity: 0,
	},
	{
		expected: []knapsack.Present{},
		arr:      []knapsack.Present{},
		capacity: 4,
	},
}

func compareSlices(ar1, ar2 []knapsack.Present) bool {
	if len(ar1) != len(ar2) {
		return false
	}
	for i, val := range ar1 {
		if val != ar2[i] {
			return false
		}
	}
	return true
}

func TestGrabPresents(t *testing.T) {
	for _, testCase := range testCases {
		if res := knapsack.GrabPresents(testCase.arr, testCase.capacity); !compareSlices(res, testCase.expected) {
			t.Errorf("Output %v not equal to expected %v", res, testCase.expected)
		}
	}
}
