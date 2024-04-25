package minCoins

import (
	"fmt"
	"testing"
	"time"
)

var testCases2 = []coinsTest{
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

func TestMinCoins2(t *testing.T) {
	for _, testCase := range testCases2 {
		if res := minCoins2(testCase.value, testCase.arr); !cmprSlices(res, testCase.exp_output) {
			t.Fatalf("Expected: %v\ngot: %v\n", testCase.exp_output, res)
		}
	}
}

func TestMinCoinsEmpty2(t *testing.T) {
	var empty []int
	if res := minCoins2(13, empty); !cmprSlices(res, empty) {
		t.Fatalf("Expected: %v\ngot: %v\n", empty, res)
	}
}

func TestMinCoinsNegative2(t *testing.T) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	var err error
	go func() {
		if res := minCoins2(13, []int{-1, -2, -3}); !cmprSlices(res, []int{}) {
			err = fmt.Errorf("Expected: %v\ngot: %v\n", []int{}, res)
		}
		err = nil
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Forever loop")
	case <-done:
		if err != nil {
			t.Fatalf(err.Error())
		}
		return
	}
}

func TestMinCoinsZeroes2(t *testing.T) {
	timeout := time.After(3 * time.Second)
	done := make(chan bool)
	var err error
	go func() {
		if res := minCoins2(13, []int{0, 0, 0}); !cmprSlices(res, []int{}) {
			err = fmt.Errorf("Expected: %v\ngot: %v\n", []int{}, res)
		}
		err = nil
		done <- true
	}()
	select {
	case <-timeout:
		t.Fatal("Forever loop")
	case <-done:
		if err != nil {
			t.Fatalf(err.Error())
		}
		return
	}
}
