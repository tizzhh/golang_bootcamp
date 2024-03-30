package main

import (
	"fmt"
	"myFind/finder"
	"os"
)

func main() {
	inputData, err := finder.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during input parsing: %s\n", err.Error())
		os.Exit(1)
	}
	err = finder.MyFind(inputData)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during find: %s\n", err.Error())
		os.Exit(1)
	}
}
