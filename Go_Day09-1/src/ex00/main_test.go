package main

import "testing"

type testAttrs struct {
	inp, exp []int
}

var testCases = []testAttrs{
	{
		inp: []int{5, 4, 3, 1, 2},
		exp: []int{1, 2, 3, 4, 5},
	},
	{
		inp: []int{1, 1, 1},
		exp: []int{1, 1, 1},
	},
}

func compareSlices(sl1, sl2 []int) bool {
	if len(sl1) != len(sl2) {
		return false
	}
	for i, val := range sl1 {
		if val != sl2[i] {
			return false
		}
	}
	return true
}

func TestSleepSort(t *testing.T) {
	for i, testCase := range testCases {
		ch := sleepSort(testCase.inp)
		res := make([]int, 0, len(testCase.inp))
		for i := 0; i < len(testCase.inp); i++ {
			res = append(res, <-ch)
		}
		close(ch)
		if !compareSlices(res, testCase.exp) {
			t.Fatalf("Test case %d failed. Expected %v, got %v\n", i, testCase.exp, res)
		}
	}
}
