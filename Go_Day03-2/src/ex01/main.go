package main

import (
	"fmt"
	"net/http"
	"os"
	"searchRest/db"
	"searchRest/renderer"
)

func main() {
	err := db.IncreaseMaxEntries()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during max entry increasing: %s\n", err.Error())
		os.Exit(1)
	}
	http.HandleFunc("/", renderer.RenderPage)

	err = http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		panic(err)
	}
}
