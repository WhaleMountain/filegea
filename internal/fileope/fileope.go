package fileope

import (
	"filegea/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

//Save file save
func Save(upfile *multipart.FileHeader, savePath string) error {
	fileName := upfile.Filename

	conf := config.GetConfig()
	saveDir := filepath.Join(conf.DataPath, savePath)

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

//Delete file save
func Delete(filePaths []string) error {
	conf := config.GetConfig()
	
	for _, fp := range filePaths {
		path := filepath.Join(conf.DataPath, fp)

		if err := os.RemoveAll(path); err != nil {
			return err
		}
	}

	return nil
}
