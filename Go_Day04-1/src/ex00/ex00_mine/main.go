package main

import (
	"candyShop/api"
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/buy_candy", api.CandyShop)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		panic(err)
	}
}
