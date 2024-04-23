package main

import (
	"fmt"
	"logoGen/frog_draw"
	"os"
)

const (
	FILENAME string = "amazing_logo.png"
)

func main() {
	err := frogdraw.CreateLogoFrog(FILENAME)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during image creation: %s\n", err.Error())
		os.Exit(1)
	}
}
