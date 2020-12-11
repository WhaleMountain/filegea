package fileope

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

const (
	basePath = "./Data"
)

//Save file save
func Save(upfile *multipart.FileHeader, savePath string) error {
	fileName := upfile.Filename

	saveDir := filepath.Join(basePath, savePath)

	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return err
	}

	saveFile := filepath.Join(saveDir, fileName)

	out, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer out.Close()

	file, err := upfile.Open()
	if err != nil {
		return err
	}
	defer file.Close()

	// Save
	if _, err := io.Copy(out, file); err != nil {
		return err
	}

	return nil
}
