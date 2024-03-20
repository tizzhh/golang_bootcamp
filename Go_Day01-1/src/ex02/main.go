package main

import (
	"compareFS/comparer"
	"fmt"
	"os"
)

func main() {
	mode, err := comparer.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during parsing input parsing: %s\n", err.Error())
		os.Exit(1)
	}

	switch mode {
	case "compare2to1":
		err := comparer.CompareSnapshots(os.Args[2], os.Args[4])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during file parsing: %s", err.Error())
			os.Exit(1)
		}
	case "compare1to2":
		err := comparer.CompareSnapshots(os.Args[4], os.Args[2])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error during file parsing: %s", err.Error())
			os.Exit(1)
		}
	}
}
