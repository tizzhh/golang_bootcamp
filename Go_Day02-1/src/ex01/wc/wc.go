package wc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ParseInput() (string, []string, error) {
	if len(os.Args) < 2 {
		return "", []string{}, fmt.Errorf("usage: ./myWc [-w -l -m] *file1* *file2* etc")
	}
	var mode string
	var paths []string

	if os.Args[1] != "-w" && os.Args[1] != "-l" && os.Args[1] != "-m" {
		mode = "-w"
		paths = append(paths, os.Args[1:]...)
	} else {
		mode = os.Args[1]
		paths = append(paths, os.Args[2:]...)
	}
	return mode, paths, nil
}

func WcCount(path, mode string, ch chan int) {
	file, _ := os.Open(path)
	var result int
	reader := bufio.NewReader(file)

	var fn func(line string) int

	switch mode {
	case "-l":
		fn = func(line string) int {
			if len(line) != 0 {
				return 1
			}
			return 0
		}
	case "-w":
		fn = func(line string) int {
			return len(strings.Fields(line))
		}
	case "-m":
		fn = func(line string) int {
			var res int
			for i := 0; i < len(line); i++ {
				res++
			}
			return res
		}
	}

	for {
		line, err := reader.ReadString('\n')
		result += fn(line)

		if err == io.EOF {
			break
		}
	}

	ch <- result
}
