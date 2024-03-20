package main

import (
	"compareDB/comparer"
	"compareDB/reader"
	"fmt"
	"os"
)

func main() {
	mode, err := reader.ParseInput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error during input parsing: %s\n", err)
		os.Exit(1)
	}

	switch mode {
	case "reader":
		filePath := os.Args[2]
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
	case "compare2to1":
		dbcomparer1, dbcomparer2 := comparer.MakeComparer(os.Args[2], os.Args[4])
		comparer.CompareDB(dbcomparer1, dbcomparer2)
	case "compare1to2":
		dbcomparer1, dbcomparer2 := comparer.MakeComparer(os.Args[4], os.Args[2])
		comparer.CompareDB(dbcomparer1, dbcomparer2)
	}
}
