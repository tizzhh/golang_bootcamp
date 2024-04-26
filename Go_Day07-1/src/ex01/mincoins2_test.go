package minCoins

import (
	"fmt"
	"os"
	"sort"
	"testing"
	"time"
)

type coinsTest struct {
	value           int
	arr, exp_output []int
}

type testCaseAttrs struct {
	number    int
	res, res2 time.Duration
}

const (
	FILE_NAME_TOP_10  string = "top10.txt"
	FILE_NAME_COMPARE string = "compare_times.txt"
)

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
	{
		value:      13,
		arr:        []int{-1, -2, 0},
		exp_output: []int{},
	},
	{
		value:      13,
		arr:        []int{},
		exp_output: []int{},
	},
	{
		value:      42069240,
		arr:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		exp_output: []int{},
	},
	{
		value:      420692406,
		arr:        []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		exp_output: []int{},
	},
}

func TestTimes(t *testing.T) {
	var times []testCaseAttrs

	for i, testCase := range testCases {
		start := time.Now()
		minCoins2(testCase.value, testCase.arr)
		times = append(times, testCaseAttrs{number: i, res: time.Duration(time.Since(start).Nanoseconds())})
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].res > times[j].res
	})

	file, err := os.Create(FILE_NAME_TOP_10)
	if err != nil {
		t.Fatalf("Error creating a file: %s", err.Error())
	}
	defer file.Close()
	for _, testTime := range times {
		fmt.Fprintf(file, "Test case number %d finished with %d ㎲\n", testTime.number, testTime.res)
	}
}

func TestTimesCompareFuncs(t *testing.T) {
	var times []testCaseAttrs

	for i, testCase := range testCases {
		start := time.Now()
		minCoins2(testCase.value, testCase.arr)
		res1Time := time.Duration(time.Since(start).Nanoseconds())
		start = time.Now()
		minCoins2Optimized(testCase.value, testCase.arr)
		res2Time := time.Duration(time.Since(start).Nanoseconds())
		times = append(times, testCaseAttrs{number: i, res: res1Time, res2: res2Time})
	}

	sort.Slice(times, func(i, j int) bool {
		return times[i].res > times[j].res
	})

	file, err := os.Create(FILE_NAME_COMPARE)
	if err != nil {
		t.Fatalf("Error creating a file: %s", err.Error())
	}
	defer file.Close()
	for _, testTime := range times {
		fmt.Fprintf(file, "Test case number %d minCoins2 finished with %d ㎲, minCoins2Optimzed with %d ㎲. The difference is: %d\n", testTime.number, testTime.res, testTime.res2, testTime.res-testTime.res2)
	}
}
