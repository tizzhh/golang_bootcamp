package rotate

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func ParseInput() (string, []string, error) {
	var paths []string
	if len(os.Args) < 2 {
		return "", paths, fmt.Errorf("usage: ./myRotate [-a] /path1.log /path2.log")
	}

	var savingPath string
	if os.Args[1] == "-a" {
		savingPath = os.Args[2]
		dir, err := os.Stat(savingPath)
		if err != nil || !dir.IsDir() {
			return "", paths, fmt.Errorf("there should be a directory after -a")
		}
		paths = append(paths, os.Args[3:]...)
	} else {
		paths = append(paths, os.Args[1:]...)
	}

	return savingPath, paths, nil
}

func Rotate(path, savingPath string, ch chan error) {
	fmt.Println(path)
	if !strings.HasSuffix(path, ".log") {
		ch <- fmt.Errorf("%s not a log file", path)
		return
	}

	stats, err := os.Stat(path)
	if err != nil {
		ch <- err
		return
	}
	mtime := stats.ModTime().Unix()
	var logPath string
	if savingPath == "" {
		logPath = strings.TrimSuffix(path, ".log") + "_" + strconv.Itoa(int(mtime)) + ".tar.gz"
	} else {
		pathElems := strings.Split(path, "/")
		logPath = savingPath + "/" + strings.TrimSuffix(pathElems[len(pathElems)-1], ".log") + "_" + strconv.Itoa(int(mtime)) + ".tar.gz"
	}
	archive, err := os.Create(logPath)
	if err != nil {
		ch <- err
		return
	}
	defer archive.Close()

	gw := gzip.NewWriter(archive)
	tw := tar.NewWriter(archive)
	defer tw.Close()
	defer gw.Close()

	addToArchive(tw, path, ch)
	ch <- nil
}

func addToArchive(tw *tar.Writer, path string, ch chan error) {
	file, err := os.Open(path)
	if err != nil {
		ch <- err
		return
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		ch <- err
		return
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		ch <- err
		return
	}

	err = tw.WriteHeader(header)
	if err != nil {
		ch <- err
		return
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		ch <- err
		return
	}
}
