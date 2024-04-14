package main

import (
	"fmt"
	"loadElastic/loader"
	"loadElastic/reader"
	"os"
	"strconv"
)

const CSV_PATH string = "../../materials/data.csv"
const MAX_ID int = 16358

// const CSV_PATH string = "small_data.csv"

func main() {
	indexName := "places"
	mapping := `{
		"mappings": {
			"properties": {
				"name": {
					"type":  "text"
				},
				"address": {
					"type":  "text"
				},
				"phone": {
					"type":  "text"
				},
				"location": {
				  "type": "geo_point"
				}
			}
		}
	}`
	err := loader.CreateIndex(indexName, mapping)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during index creation: %s\n", err.Error())
		os.Exit(1)
	}

	err = loader.LoadDataInDb(CSV_PATH, indexName, mapping)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during index population: %s\n", err.Error())
		os.Exit(1)
	}

	switch len(os.Args) {
	case 1:
		for i := 0; i < MAX_ID; i++ {
			go reader.GetData(strconv.Itoa(i), indexName)
		}
	case 2:
		_, err := strconv.Atoi(os.Args[1])
		if err == nil {
			err = reader.GetData(os.Args[1], indexName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error during getting data: %s\n", err.Error())
				os.Exit(1)
			}
		} else {
			err = reader.GetInfo(os.Args[1], indexName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error during getting data: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}
}
