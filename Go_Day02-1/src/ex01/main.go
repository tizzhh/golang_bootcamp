package main

import (
	"fmt"
	"myWc/wc"
	"os"
)

func main() {
	mode, paths, err := wc.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error during input parsing: %s\n", err.Error())
		os.Exit(1)
	}
	ch := make(chan wc.Wc)
	for _, path := range paths {
		go wc.WcCount(path, mode, ch)
		wcData := <-ch
		if wcData.Err != nil {
			fmt.Fprintf(os.Stderr, "Error during file reading: %s\n", wcData.Err.Error())
		} else {
			fmt.Printf("%d\t%s\n", wcData.Res, path)
		}
	}
}
