package finder

import (
	"fmt"
	"os"
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

	if fld.ext != "" && !fld.f {
		return fld, fmt.Errorf("ext only works with -f specified")
	}

	fld.path = os.Args[len(os.Args)-1]

	if !fld.d && !fld.f && !fld.sl {
		fld.d, fld.f, fld.sl = true, true, true
	}

	return fld, nil
}

func MyFind(fld FlagData) error {
	var paths []string
	var completeData []string
	err := ReadSubdirs(&paths, &completeData, fld, fld.path)
	if err != nil {
		return err
	}
	for _, file := range completeData {
		fmt.Println(file)
	}
	return nil
}

func ReadSubdirs(paths, completeData *[]string, fld FlagData, path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		if os.IsPermission(err) {
			return nil
		}
		return err
	}
	cur_dir, err := os.Getwd()
	if err != nil {
		return err
	}
	for _, file := range files {
		new_dir := cur_dir + "/" + path + "/" + file.Name()
		*paths = append(*paths, new_dir)
		if file.IsDir() {
			err = ReadSubdirs(paths, completeData, fld, path+"/"+file.Name())
			if err != nil {
				return err
			}
		}
		if file.IsDir() && fld.d {
			*completeData = append(*completeData, new_dir)
		} else if !file.IsDir() {
			link, err := os.Readlink(new_dir)
			if err != nil && fld.f {
				if fld.f && (fld.ext != "" && strings.HasSuffix(file.Name(), fld.ext)) || fld.ext == "" {
					*completeData = append(*completeData, new_dir)
				}
			} else if fld.sl {
				_, err = os.Stat(new_dir)
				if err != nil {
					*completeData = append(*completeData, new_dir+" -> [broken]")
				} else if link != "" {
					*completeData = append(*completeData, new_dir+" -> "+link)
				}
			}
		}
	}
	return nil
}
