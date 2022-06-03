package utils

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

func CopyFile(src, dst string) (int64, error) {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func CopyFolderFiles(src, dst string) {

	if _, err := os.Stat(dst); os.IsNotExist(err) {
		err := os.Mkdir(dst, 0644)
		if err != nil {
			return
		}
	}
	files, e := os.ReadDir(src)
	if e != nil {
		return
	}
	for _, file := range files {
		if strings.Contains(file.Name(), "cache") {
			continue
		}
		if file.IsDir() {
			err := os.Mkdir(path.Join(dst, file.Name()), 0644)
			if err != nil {
				continue
			}

			CopyFolderFiles(path.Join(src, file.Name()), path.Join(dst, file.Name()))
			continue
		}

		_, err := CopyFile(path.Join(src, file.Name()), path.Join(dst, file.Name()))
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func ContainsString(slice []os.DirEntry, value string) bool {

	for _, s := range slice {
		if s.Name() == value {
			return true
		}
	}

	return false
}
