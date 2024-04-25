package minCoins

import (
	"fmt"
	"testing"
	"time"
)

type coinsTest struct {
	value           int
	arr, exp_output []int
}

var testCases = []coinsTest{
	{
		value:      13,
		arr:        []int{1, 5, 10},
		exp_output: []int{10, 1, 1, 1},
	},
	{
		value:      13,
		arr:        []int{1, 1, 5, 5, 10, 10},
		exp_output: []int{10, 1, 1, 1},
	},
	{
		value:      0,
		arr:        []int{1, 5, 10},
		exp_output: []int{},
	},
	{
		value:      18,
		arr:        []int{1, 3, 7},
		exp_output: []int{7, 7, 3, 1},
	},
	{
		value:      13,
		arr:        []int{2, 5, 10},
		exp_output: []int{},
	},
	{
		value:      13,
		arr:        []int{10, 5, 1},
		exp_output: []int{10, 1, 1, 1},
	},
}

func cmprSlices(ar1, ar2 []int) bool {
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

func TestMinCoins(t *testing.T) {
	for _, testCase := range testCases {
		if res := minCoins(testCase.value, testCase.arr); !cmprSlices(res, testCase.exp_output) {
			t.Fatalf("Expected: %v\ngot: %v\n", testCase.exp_output, res)
		}
	}
}

func TestMinCoinsEmpty(t *testing.T) {
	var empty []int
	if res := minCoins(13, empty); !cmprSlices(res, empty) {
		t.Fatalf("Expected: %v\ngot: %v\n", empty, res)
	}
}

func TestMinCoinsNegative(t *testing.T) {
	timeout := time.After(1 * time.Second)
	done := make(chan bool)
	err := make(chan error)
	go func() {
		if res := minCoins(13, []int{-1, -2, -3}); !cmprSlices(res, []int{}) {
			err <- fmt.Errorf("Expected: %v\ngot: %v\n", []int{}, res)
		}
		err <- nil
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Forever loop")
	case <-done:
		errs := <-err
		if err != nil {
			t.Fatalf(errs.Error())
		}
		return
	}
}

func TestMinCoinsZeroes(t *testing.T) {
	timeout := time.After(1 * time.Second)
	done := make(chan bool)
	err := make(chan error)
	go func() {
		if res := minCoins(13, []int{0, 0, 0}); !cmprSlices(res, []int{}) {
			err <- fmt.Errorf("Expected: %v\ngot: %v\n", []int{}, res)
		}
		err <- nil
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Forever loop")
	case <-done:
		errs := <-err
		if err != nil {
			t.Fatalf(errs.Error())
		}
		return
	}
}
