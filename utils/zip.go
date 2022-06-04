package utils

import (
	"fmt"
	"io/ioutil"

	"github.com/alexmullins/zip"
)

func AddFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Encrypt(baseInZip+file.Name(), "In5BXGYxfWo/a3E0THZ9L3JDSjdeL3YxKGFjaVYiIjpARU5yampxZHd4fSxjPk9qWllSIg")
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			AddFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}
