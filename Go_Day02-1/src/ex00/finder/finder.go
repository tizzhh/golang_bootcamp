package finder

import (
	"fmt"
	"os"
)

type FlagData struct {
	sl, d, f bool
	ext, path string
}

func ParseInput() (FlagData, error) {
	var fld FlagData
	fmt.Println("aboba")
	if len(os.Args) < 2 || len(os.Args) > 5 {
		return fld, fmt.Errorf("sage: ./myFind [-f -sl -d] [-ext extention] *path*")
	}

	for arg := range os.Args[1:] {
		fmt.Println(arg)
	}
	return fld, nil
}
