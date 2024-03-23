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
	ch := make(chan int)
	for _, path := range paths {
		go wc.WcCount(path, mode, ch)
		fmt.Printf("%d\t%s\n", <-ch, path)
	}
}
