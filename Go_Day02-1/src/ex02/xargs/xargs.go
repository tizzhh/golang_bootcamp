package xargs

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ParseInput() (string, error) {
	var params string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input xargs arguments:")

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		params += " " + text
	}

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error during reading input: %s", err)
	}
	return params[1:], nil
}

func MyXargs(params string) {
	paramsList := strings.Fields(params)
	var flags string
	flagExists := false
	if len(os.Args) >= 3 {
		flags = strings.Join(os.Args[2:], " ")
		flagExists = true
	}

	for _, param := range paramsList {
		var res []byte
		var err error
		if flagExists {
			res, err = exec.Command(os.Args[1], flags, param).Output()
		} else {
			res, err = exec.Command(os.Args[1], param).Output()
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		} else {
			fmt.Println(string(res))
		}
	}
}
