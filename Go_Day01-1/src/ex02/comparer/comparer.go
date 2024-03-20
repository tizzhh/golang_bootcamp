package comparer

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ParseInput() (string, error) {
	var mode string
	if len(os.Args) != 5 || !((os.Args[1] == "--old" && os.Args[3] == "--new") || (os.Args[1] == "--new" && os.Args[3] == "--old")) || (os.Args[1] == os.Args[3]) {
		return "", fmt.Errorf(`usage: ./compareDB -f [filename xml or json]
		or ./compareDB --[old/new] original_database.[xml/json] --[new/old] stolen_database.[json/xml]`)
	}

	if len(os.Args) == 5 {
		if os.Args[1] == "--old" {
			mode = "compare2to1"
		} else {
			mode = "compare1to2"
		}
	}
	return mode, nil
}

func CompareSnapshots(path1, path2 string) error {
	file1, err := os.Open(path1)
	if err != nil {
		return fmt.Errorf("error during file %s opening: %v", path1, err)
	}
	defer file1.Close()

	file2, err := os.Open(path2)
	if err != nil {
		return fmt.Errorf("error during file %s opening: %v", path2, err)
	}
	defer file2.Close()

	file1Contents := make(map[string]interface{})
	reader := bufio.NewReader(file1)
	for {
		file1Line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error during file %s reading %v", path2, err)
		}
		file1Contents[file1Line] = nil
	}

	reader2 := bufio.NewReader(file2)
	for {
		file2Line, err := reader2.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("error during file %s reading %v", path2, err)
		}
		if _, ok := file1Contents[file2Line]; ok {
			delete(file1Contents, file2Line)
		} else {
			fmt.Printf("ADDED %s", file2Line)
		}
	}

	for key := range file1Contents {
		fmt.Printf("REMOVED %s", key)
	}
	return nil
}
