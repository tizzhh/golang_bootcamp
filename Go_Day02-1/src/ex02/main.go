package main

import (
	"fmt"
	"myXargs/xargs"
	"os"
)

func main() {
	params, err := xargs.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err.Error())
		os.Exit(1)
	}

	xargs.MyXargs(params)
}
