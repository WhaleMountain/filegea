package controller

import (
	"filegea/config"
	"filegea/internal/ui"
	"filegea/internal/util"
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

	c.Writer.WriteString(ui.Index(searchPATH, files))
}

//UploadFrom upload form
func (fgc *FileGeaController) UploadFrom(c *gin.Context) {
	savePath := c.Param("path")
	status := "Drag and Drop or Click in this area."
	c.Writer.WriteString(ui.Upload(savePath, status))
}

//UploadFromDir upload dir form
func (fgc *FileGeaController) UploadFromDir(c *gin.Context) {
	savePath := c.Param("path")
	status := "Drag and Drop or Click in this area."
	c.Writer.WriteString(ui.UploadDir(savePath, status))
}

//Upload upload file
func (fgc *FileGeaController) Upload(c *gin.Context) {
	savePath := c.Param("path")
	form, err := c.MultipartForm()
	if err != nil {
		log.Printf("Multipart form error %s\n", err)
		status := fmt.Sprintf("Error: %s", err)
		c.Writer.WriteString(ui.Upload(savePath, status))
	}

	files := form.File["file"]
	if len(files) <= 0 {
		c.Redirect(http.StatusMovedPermanently, "/filegea")
	}

	for _, file := range files {
		if err := util.Save(file, savePath); err != nil {
			log.Printf("file save error %s\n", err)
			status := fmt.Sprintf("Error: %s", err)
			c.Writer.WriteString(ui.Upload(savePath, status))
		}
	}

	status := "Upload Success !!"
	c.Writer.WriteString(ui.Upload(savePath, status))
}

//DeleteForm delete page
func (fgc *FileGeaController) DeleteForm(c *gin.Context) {
	searchPATH := c.Param("path")

	path := filepath.Join(fgc.Conf.DataPath, searchPATH)
	files, _ := ioutil.ReadDir(path)

	c.Writer.WriteString(ui.Delete(searchPATH, files))
}

//Delete delete file
func (fgc *FileGeaController) Delete(c *gin.Context) {
	c.Request.ParseForm()

	for _, paths := range c.Request.PostForm {
		if err := util.Delete(paths); err != nil {
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

	c.Writer.WriteString(ui.Download(searchPATH, files))
}

//Download download file
func (fgc *FileGeaController) Download(c *gin.Context) {
	c.Request.ParseForm()

	for _, paths := range c.Request.PostForm {
		filename, fpath, _ := util.Download(paths)

		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
		c.Writer.Header().Add("Content-Type", "application/zip")
		c.File(fpath)
	}

	// Nothing is checked
	c.Redirect(http.StatusMovedPermanently, "/filegea")
}

//Redirect / -> /filegea
func (fgc *FileGeaController) Redirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/filegea")
}
