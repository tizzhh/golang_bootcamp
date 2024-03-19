package process

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func ParseInput(nums *[]int) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return fmt.Errorf("error occured: %s", err.Error())
		}
		line = line[:len(line)-1]
		if line == "end" {
			break
		}
		elems := strings.Fields(line)
		for _, val := range elems {
			num, err := strconv.Atoi(val)
			if err != nil {
				return fmt.Errorf("wrong input: %s", err.Error())
			}
			*nums = append(*nums, num)
		}
	}
	return nil
}

func CalcMean(nums []int) (mean float64) {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return float64(sum) / float64(len(nums))
}

func CalcMedian(nums []int) (median float64) {
	if len(nums)%2 != 0 {
		median = float64(nums[len(nums)/2])
	} else {
		median = (float64(nums[len(nums)/2]-1) + float64(nums[len(nums)/2])) / 2.0
	}
	return
}

func CalcMode(nums []int) (mode int) {
	numberCounts := make(map[int]int, len(nums))
	for _, num := range nums {
		numberCounts[num]++
	}
	mode, max_val := 0, 0
	for key, val := range numberCounts {
		if val >= max_val {
			max_val = val
			mode = key
		} else if val == max_val {
			if key < mode {
				mode = key
			}
		}
	}

	return mode
}

func CalcSd(nums []int, mean float64) (sd float64) {
	for _, num := range nums {
		sd += math.Pow(float64(num)-mean, 2)
	}
	return math.Sqrt(sd / float64(len(nums)-1))
}
