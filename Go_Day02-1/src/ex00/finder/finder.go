package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type FlagData struct {
	sl, d, f  bool
	ext, path string
}

func ParseInput() (FlagData, error) {
	var fld FlagData
	if len(os.Args) < 2 || len(os.Args) > 5 {
		return fld, fmt.Errorf("usage: ./myFind [-f -sl -d] [-ext extention] *path*")
	}

	var extfound bool
	for i := 1; i < len(os.Args)-1; i++ {
		// fmt.Println("arg:", os.Args[i])
		if extfound {
			fld.ext = os.Args[i]
		}
		switch os.Args[i] {
		case "-sl":
			fld.sl = true
		case "-d":
			fld.f = false
			fld.d = true
		case "-f":
			fld.d = false
			fld.f = true
		case "-ext":
			extfound = true
		}
	}
	fld.path = os.Args[len(os.Args)-1]

	// if fld.f && fld.d {
		// return fld, fmt.Errorf("flags f and d cannot be specified at the same time")
	// }

	return fld, nil
}

func ReadDir(fld FlagData) error {
	// files, err := os.ReadDir(fld.path)
	// if err != nil {
	// 	return fmt.Errorf("error during file reading: %s", err)
	// }
	// for _, file := range files {
	// 	// fmt.Println(file)
	// 	if fld.d && file.IsDir() {
	// 		fmt.Println(file)
	// 	} else if fld.f && file.Type().IsRegular() {
	// 		if (fld.ext != "" && strings.HasSuffix(file.Name(), fld.ext)) || fld.ext == "" {
	// 			fmt.Println(file)
	// 		}
	// 	}
	// }

	filepath.Walk(fld.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fld.d && info.IsDir() {
				fmt.Println(info.Name())
			} else if fld.f && !info.IsDir() {
				if (fld.ext != "" && strings.HasSuffix(info.Name(), fld.ext)) || fld.ext == "" {
					fmt.Println(info.Name())
				}
			}
			return nil
		})
	return nil
}
