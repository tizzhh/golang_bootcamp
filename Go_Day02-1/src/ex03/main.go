package main

import (
	"fmt"
	"myRotate/rotate"
	"os"
)

func main() {
	savingPath, paths, err := rotate.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during input parsing: %s\n", err.Error())
		os.Exit(1)
	}

	ch := make(chan error)
	for _, path := range paths {
		go rotate.Rotate(path, savingPath, ch)
		err := <-ch
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during archiving: %s\n", err.Error())
		}
	}
}
