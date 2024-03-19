package main

import (
	"Anscombe/process"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Println("Input sequence of numbers separeted by newlines. Input 'end' to end the input")
	nums := []int{}
	outputTypes := flag.String("type", "not specified", `Input desired output type. 
	Available variants are: 'Mean', 'Median', 'Mode', 'SD'`)
	flag.Parse()

	err := process.ParseInput(&nums)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during parsing: %s\n", err.Error())
		os.Exit(1)
	}
	if len(nums) == 0 {
		fmt.Fprintf(os.Stderr, "Empty input")
		os.Exit(1)
	}

	sort.Ints(nums)
	var (
		mean, median, sd float64
		mode             int
	)

	mean = process.CalcMean(nums)
	median = process.CalcMedian(nums)
	mode = process.CalcMode(nums)
	sd = process.CalcSd(nums, mean)
	switch *outputTypes {
	case "Mean":
		fmt.Printf("Mean: %.2f\n", mean)
	case "Median":
		fmt.Printf("Median: %.2f\n", median)
	case "Mode":
		fmt.Println("Mode:", mode)
	case "SD":
		fmt.Printf("SD: %.2f\n", sd)
	default:
		fmt.Printf("Mean: %.2f\nMedian: %.2f\nMode: %d\nSD: %.2f\n", mean, median, mode, sd)
	}
}
