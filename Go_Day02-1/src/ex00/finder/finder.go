package finder

import (
	"fmt"
	"os"
	// "path/filepath"
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
			fld.d = true
		case "-f":
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

	// filepath.Walk(fld.path,
	// 	func(path string, info os.FileInfo, err error) error {
	// 		if err != nil {
	// 			return err
	// 		}
	// 		root_path, err := os.Getwd()
	// 		if err != nil {
	// 			return err
	// 		}
	// 		if fld.d && info.IsDir() {
	// 			fmt.Println(root_path + "/" + info.Name())
	// 		}
	// 		if fld.f && !info.IsDir() {
	// 			if (fld.ext != "" && strings.HasSuffix(info.Name(), fld.ext)) || fld.ext == "" {
	// 				fmt.Println(root_path + "/" + info.Name())
	// 			}
	// 		}
	// 		return nil
	// 	})
	var paths []string
	err := ReadSubdirs(&paths, fld, fld.path)
	if err != nil {
		return err
	}
	for _, file := range paths {
		fmt.Println(file)
	}
	// files, err := os.ReadDir(fld.path)
	// if err != nil {
	// 	return err
	// }
	// for _, file := range files {
	// 	// info, _ := os.Readlink("/home/tizzhh/Desktop/golang_bootcamp/Go_Day02-1/src/ex00/test_dir/" + file.Name())
	// 	fmt.Println(file)
	// }
	return nil
}

func ReadSubdirs(paths *[]string, fld FlagData, path string) error {
	// fmt.Println(paths)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println(path)
		return err
	}
	// var count_dirs = 0
	cur_dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, file := range files {
		// info, _ := os.Readlink("/home/tizzhh/Desktop/golang_bootcamp/Go_Day02-1/src/ex00/test_dir/" + file.Name())
		// fmt.Println(file)
		// fmt.Println(file)
		if file.IsDir() {
			// count_dirs++
			new_dir := cur_dir + "/" + path + "/" + file.Name()
			*paths = append(*paths, new_dir)
			err = ReadSubdirs(paths, fld, path+"/"+file.Name())
			if err != nil {
				return err
			}
		} else if !file.IsDir() { // тут нужно допилить что если у нас не указан fld.d, то все равно прозход по всех дирам
			if (fld.ext != "" && strings.HasSuffix(file.Name(), fld.ext)) || fld.ext == "" {
				*paths = append(*paths, cur_dir+"/"+file.Name())
			}
		}

	}
	return nil
}
