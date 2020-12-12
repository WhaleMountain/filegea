package controllers

import (
	"filegea/config"
	"filegea/internal/fileope"
	"filegea/internal/view"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//FileGeaController filegea controller
type FileGeaController struct {
	Conf config.Config
}

//NewFileGeaController filegea controller
func NewFileGeaController() *FileGeaController {
	return &FileGeaController{
		Conf: config.GetConfig(),
	}
}

//Index index page
func (fgc *FileGeaController) Index(c *gin.Context) {
	searchPATH := c.Param("path")

	path := filepath.Join(fgc.Conf.DataPath, searchPATH)
	files, _ := ioutil.ReadDir(path)

	c.Writer.WriteString(view.Index(searchPATH, files))
}

//UploadFrom upload form
func (fgc *FileGeaController) UploadFrom(c *gin.Context) {
	savePath := c.Param("path")
	status := "Drag and Drop or Click in this area."
	c.Writer.WriteString(view.Upload(savePath, status))
}

//UploadFromDir upload dir form
func (fgc *FileGeaController) UploadFromDir(c *gin.Context) {
	savePath := c.Param("path")
	status := "Drag and Drop or Click in this area."
	c.Writer.WriteString(view.UploadDir(savePath, status))
}

//Upload upload file
func (fgc *FileGeaController) Upload(c *gin.Context) {
	savePath := c.Param("path")
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Multipart form error %s\n", err)
		status := fmt.Sprintf("Error: %s", err)
		c.Writer.WriteString(view.Upload(savePath, status))
	}

	files := form.File["file"]

	for _, file := range files {
		if err := fileope.Save(file, savePath); err != nil {
			log.Printf("file save error %s\n", err)
			status := fmt.Sprintf("Error: %s", err)
			c.Writer.WriteString(view.Upload(savePath, status))
		}
	}

	status := "Upload Success !!"
	c.Writer.WriteString(view.Upload(savePath, status))
}

//DeleteForm delete page
func (fgc *FileGeaController) DeleteForm(c *gin.Context) {
	searchPATH := c.Param("path")

	path := filepath.Join(fgc.Conf.DataPath, searchPATH)
	files, _ := ioutil.ReadDir(path)

	c.Writer.WriteString(view.Delete(searchPATH, files))
}

//Delete delete file
func (fgc *FileGeaController) Delete(c *gin.Context) {
	c.Request.ParseForm()

	for _, paths := range c.Request.PostForm {
		if err := fileope.Delete(paths); err != nil {
			log.Printf("file delete error %s\n", err)
		}
	}

	c.Redirect(http.StatusMovedPermanently, "/filegea")
}

//DownloadForm download page
func (fgc *FileGeaController) DownloadForm(c *gin.Context) {
	searchPATH := c.Param("path")

	path := filepath.Join(fgc.Conf.DataPath, searchPATH)
	files, _ := ioutil.ReadDir(path)

	c.Writer.WriteString(view.Download(searchPATH, files))
}

//Download download file
func (fgc *FileGeaController) Download(c *gin.Context) {
	c.Request.ParseForm()

	for _, paths := range c.Request.PostForm {
		filename, fpath, _ := fileope.Download(paths)

		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.Writer.Header().Add("Content-Type", "application/zip")
		c.File(fpath)
	}
}

//Redirect / -> /filegea
func (fgc *FileGeaController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/filegea")
}
