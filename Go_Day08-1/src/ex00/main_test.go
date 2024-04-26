package main

import (
	"fmt"
	"testing"
)

type expectedRet struct {
	num int
	err error
}

type testAttrs struct {
	idx      int
	arr      []int
	expected expectedRet
}

var testCases = []testAttrs{
	{
		idx: 0,
		arr: []int{1, 2, 3},
		expected: expectedRet{
			num: 1,
			err: nil,
		},
	},
	{
		idx: 1,
		arr: []int{1, 2, 3},
		expected: expectedRet{
			num: 2,
			err: nil,
		},
	},
	{
		idx: 2,
		arr: []int{1, 2, 3},
		expected: expectedRet{
			num: 3,
			err: nil,
		},
	},
	{
		idx: -100,
		arr: []int{1, 2, 3},
		expected: expectedRet{
			num: 0,
			err: fmt.Errorf("idx out of bounds: -100"),
		},
	},
	{
		idx: 100,
		arr: []int{1, 2, 3},
		expected: expectedRet{
			num: 0,
			err: fmt.Errorf("idx out of bounds: 100"),
		},
	},
}

func TestMain(t *testing.T) {
	for i, testCase := range testCases {
		res, err := getElement(testCase.arr, testCase.idx)
		if err == nil {
			if res != testCase.expected.num {
				t.Fatalf("Test case %d failed. Int expected %d, got %d.\n", i, testCase.expected.num, res)
			}
		} else {
			if err.Error() != testCase.expected.err.Error() {
				t.Fatalf("Test case %d failed. Err expected %v, got %v\n", i, testCase.expected.err, err)
			}
		}
	}
}
