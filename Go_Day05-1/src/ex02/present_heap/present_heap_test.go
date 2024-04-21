package presentHeap_test

import (
	"errors"
	"presentHeap/present_heap"
	"testing"
)

type expec struct {
	expected []presentHeap.Present
	err      error
}

type prHeapTest struct {
	arr      []presentHeap.Present
	expected expec
	n        int
}

var testCases = []prHeapTest{
	{
		expected: expec{err: nil, expected: []presentHeap.Present{{Value: 5, Size: 1}, {Value: 5, Size: 2}}},
		arr:      []presentHeap.Present{{Value: 5, Size: 1}, {Value: 4, Size: 5}, {Value: 3, Size: 1}, {Value: 5, Size: 2}},
		n:        2,
	},
	{
		expected: expec{err: nil, expected: []presentHeap.Present{{Value: 3, Size: 69}, {Value: 3, Size: 420}, {Value: 3, Size: 69420}}},
		arr:      []presentHeap.Present{{Value: 3, Size: 69}, {Value: -1, Size: -1}, {Value: 3, Size: 69420}, {Value: 3, Size: 420}},
		n:        3,
	},
	{
		expected: expec{err: errors.New("n should be in [0; len(arr)], got: 69")},
		arr:      []presentHeap.Present{{Value: 5, Size: 1}, {Value: 4, Size: 5}, {Value: 3, Size: 1}, {Value: 5, Size: 2}},
		n:        69,
	},
}

func compareSlices(ar1, ar2 []presentHeap.Present) bool {
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

func TestGrabCoolset(t *testing.T) {
	for _, testCase := range testCases {
		res, err := presentHeap.GetNCoolestPresents(testCase.arr, testCase.n)
		if err == nil {
			if !compareSlices(res, testCase.expected.expected) {
				t.Errorf("Output \n%v not equal to expected \n%v\n", res, testCase.expected.expected)
			}
		}
		if err != nil {
			if err.Error() != testCase.expected.err.Error() {
				t.Errorf("Error not equal to expected: got: \n'%v',\n expected err \n'%v'\n", err, testCase.expected.err)
			}
		}
	}
}
