package fileope

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Souce https://stackoverflow.com/questions/37869793/how-do-i-zip-a-directory-containing-sub-directories-or-files-in-golang

func zipWriter(targets []string, zipPath string) {
	// Get a Buffer to Write To
	zipFile, err := os.Create(zipPath)
	if err != nil {
		fmt.Println(err)
	}
	defer zipFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(zipFile)

	// Add some files to the archive.
	for _, target := range targets {
		addFiles(w, target, "")
	}

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		dat, err := ioutil.ReadFile(basePath)
		if err != nil {
			fmt.Println(err)
		}

		// Add some files to the archive.
		f, err := w.Create(filepath.Join(baseInZip, filepath.Base(basePath)))
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(dat)
		if err != nil {
			fmt.Println(err)
		}

		return
	}

	for _, file := range files {
		if file.IsDir() {
			newBase := filepath.Join(basePath, file.Name(), "/")
			addFiles(w, newBase, filepath.Join(baseInZip, file.Name(), "/"))

		} else {
			dat, err := ioutil.ReadFile(filepath.Join(basePath, file.Name()))
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(filepath.Join(baseInZip, file.Name()))
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
