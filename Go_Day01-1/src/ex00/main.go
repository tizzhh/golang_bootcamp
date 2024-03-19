package main

import (
	"fmt"
	"os"
	"readDB/reader"
)

func main() {
	filePath, err := reader.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during input parsing: %s\n", err)
		os.Exit(1)
	}

	var dbreader reader.DBReader = reader.CreateReader(filePath)
	err = dbreader.ReadData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s during reading file: %s\n", err, filePath)
		os.Exit(1)
	}

	err = dbreader.OutputData()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %s during marshalling file: %s\n", err, filePath)
		os.Exit(1)
	}

	var j_x_reader reader.DBReader = &reader.JSONReader{Path: "stolen_database.json"}
	j_x_reader.ReadData()
	reader.OutputJsonXml(j_x_reader)
	j_x_reader = &reader.XMLReader{Path: "original_database.xml"}
	j_x_reader.ReadData()
	reader.OutputJsonXml(j_x_reader)

}
